package api_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	testapi "github.com/cloudfoundry/cli/cf/api/fakes"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/net"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testnet "github.com/cloudfoundry/cli/testhelpers/net"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/cf/api"
	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("route repository", func() {
	var (
		ts         *httptest.Server
		handler    *testnet.TestHandler
		configRepo core_config.Repository
		repo       CloudControllerRouteRepository
	)

	BeforeEach(func() {
		configRepo = testconfig.NewRepositoryWithDefaults()
		configRepo.SetSpaceFields(models.SpaceFields{
			Guid: "the-space-guid",
			Name: "the-space-name",
		})
		gateway := net.NewCloudControllerGateway(configRepo, time.Now, &testterm.FakeUI{})
		repo = NewCloudControllerRouteRepository(configRepo, gateway)
	})

	AfterEach(func() {
		if ts != nil {
			ts.Close()
		}
	})

	Describe("List routes", func() {
		It("lists routes in the current space", func() {
			ts, handler = testnet.NewServer([]testnet.TestRequest{
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "GET",
					Path:     "/v2/spaces/the-space-guid/routes?inline-relations-depth=1",
					Response: firstPageRoutesResponse,
				}),
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "GET",
					Path:     "/v2/spaces/the-space-guid/routes?inline-relations-depth=1&page=2",
					Response: secondPageRoutesResponse,
				}),
			})
			configRepo.SetApiEndpoint(ts.URL)

			routes := []models.Route{}
			apiErr := repo.ListRoutes(func(route models.Route) bool {
				routes = append(routes, route)
				return true
			})

			Expect(len(routes)).To(Equal(2))
			Expect(routes[0].Guid).To(Equal("route-1-guid"))
			Expect(routes[0].Path).To(Equal(""))
			Expect(routes[0].ServiceInstance.Guid).To(Equal("service-guid"))
			Expect(routes[0].ServiceInstance.Name).To(Equal("test-service"))
			Expect(routes[1].Guid).To(Equal("route-2-guid"))
			Expect(routes[1].Path).To(Equal("/path-2"))
			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})

		It("lists routes from all the spaces of current org", func() {
			ts, handler = testnet.NewServer([]testnet.TestRequest{
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "GET",
					Path:     "/v2/routes?q=organization_guid:my-org-guid&inline-relations-depth=1",
					Response: firstPageRoutesOrgLvlResponse,
				}),
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "GET",
					Path:     "/v2/routes?q=organization_guid:my-org-guid&inline-relations-depth=1&page=2",
					Response: secondPageRoutesResponse,
				}),
			})
			configRepo.SetApiEndpoint(ts.URL)

			routes := []models.Route{}
			apiErr := repo.ListAllRoutes(func(route models.Route) bool {
				routes = append(routes, route)
				return true
			})

			Expect(len(routes)).To(Equal(2))
			Expect(routes[0].Guid).To(Equal("route-1-guid"))
			Expect(routes[0].Space.Guid).To(Equal("space-1-guid"))
			Expect(routes[0].ServiceInstance.Guid).To(Equal("service-guid"))
			Expect(routes[0].ServiceInstance.Name).To(Equal("test-service"))
			Expect(routes[1].Guid).To(Equal("route-2-guid"))
			Expect(routes[1].Space.Guid).To(Equal("space-2-guid"))

			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})

		Describe("Find", func() {
			var ccServer *ghttp.Server
			BeforeEach(func() {
				ccServer = ghttp.NewServer()
				configRepo.SetApiEndpoint(ccServer.URL())
			})

			AfterEach(func() {
				ccServer.Close()
			})

			Context("when the path is empty", func() {
				BeforeEach(func() {
					v := url.Values{}
					v.Set("inline-relations-depth", "1")
					v.Set("q", "host:my-cool-app;domain_guid:my-domain-guid")

					ccServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", "/v2/routes", v.Encode()),
							ghttp.VerifyHeader(http.Header{
								"accept": []string{"application/json"},
							}),
							ghttp.RespondWith(http.StatusCreated, findResponseBodyWithoutPath),
						),
					)
				})

				It("returns the route", func() {
					domain := models.DomainFields{}
					domain.Guid = "my-domain-guid"

					route, apiErr := repo.Find("my-cool-app", domain, "", 0)

					Expect(apiErr).NotTo(HaveOccurred())
					Expect(route.Host).To(Equal("my-cool-app"))
					Expect(route.Guid).To(Equal("my-route-guid"))
					Expect(route.Path).To(Equal(""))
					Expect(route.Domain.Guid).To(Equal(domain.Guid))
				})
			})

			Context("when the route is found", func() {
				BeforeEach(func() {
					v := url.Values{}
					v.Set("inline-relations-depth", "1")
					v.Set("q", "host:my-cool-app;domain_guid:my-domain-guid;path:/somepath;port:8148")

					ccServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", "/v2/routes", v.Encode()),
							ghttp.VerifyHeader(http.Header{
								"accept": []string{"application/json"},
							}),
							ghttp.RespondWith(http.StatusCreated, findResponseBody),
						),
					)
				})

				It("returns the route", func() {
					domain := models.DomainFields{}
					domain.Guid = "my-domain-guid"

					route, apiErr := repo.Find("my-cool-app", domain, "somepath", 8148)

					Expect(apiErr).NotTo(HaveOccurred())
					Expect(route.Host).To(Equal("my-cool-app"))
					Expect(route.Guid).To(Equal("my-route-guid"))
					Expect(route.Path).To(Equal("/somepath"))
					Expect(route.Domain.Guid).To(Equal(domain.Guid))
				})
			})

			Context("when the route is not found", func() {
				BeforeEach(func() {
					v := url.Values{}
					v.Set("inline-relations-depth", "1")
					v.Set("q", "host:my-cool-app;domain_guid:my-domain-guid;path:/somepath;port:1478")

					ccServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", "/v2/routes", v.Encode()),
							ghttp.VerifyHeader(http.Header{
								"accept": []string{"application/json"},
							}),
							ghttp.RespondWith(http.StatusOK, `{ "resources": [] }`),
						),
					)
				})

				It("returns 'not found'", func() {
					ts, handler = testnet.NewServer([]testnet.TestRequest{
						testapi.NewCloudControllerTestRequest(testnet.TestRequest{
							Method:   "GET",
							Path:     "/v2/routes?q=host%3Amy-cool-app%3Bdomain_guid%3Amy-domain-guid",
							Response: testnet.TestResponse{Status: http.StatusOK, Body: `{ "resources": [ ] }`},
						}),
					})
					configRepo.SetApiEndpoint(ts.URL)

					domain := models.DomainFields{}
					domain.Guid = "my-domain-guid"

					_, apiErr := repo.Find("my-cool-app", domain, "somepath", 1478)

					Expect(handler).To(HaveAllRequestsCalled())

					Expect(apiErr.(*errors.ModelNotFoundError)).NotTo(BeNil())
				})
			})
		})
	})

	Describe("CreateInSpace", func() {
		var ccServer *ghttp.Server
		BeforeEach(func() {
			ccServer = ghttp.NewServer()
			configRepo.SetApiEndpoint(ccServer.URL())
		})

		AfterEach(func() {
			if ccServer != nil {
				ccServer.Close()
			}
		})

		Context("when no host, path, or port are given", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/v2/routes", "inline-relations-depth=1&async=true"),
						ghttp.VerifyJSON(`
							{
								"domain_guid":"my-domain-guid",
								"space_guid":"my-space-guid",
								"generate_port":false
							}
						`),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
					),
				)
			})

			It("tries to create a route", func() {
				repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 0, false)
				Expect(ccServer.ReceivedRequests()).To(HaveLen(1))
			})

			Context("when creating the route succeeds", func() {
				BeforeEach(func() {
					h := ccServer.GetHandler(0)
					ccServer.SetHandler(0, ghttp.CombineHandlers(
						h,
						ghttp.RespondWith(http.StatusCreated, `
								{
									"metadata": { "guid": "my-route-guid" },
									"entity": { "host": "my-cool-app" }
								}
							`),
					))
				})

				It("returns the created route", func() {
					createdRoute, err := repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 0, false)
					Expect(err).NotTo(HaveOccurred())
					Expect(createdRoute.Guid).To(Equal("my-route-guid"))
				})
			})
		})

		Context("when a host is given", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/v2/routes", "inline-relations-depth=1&async=true"),
						ghttp.VerifyJSON(`
							{
								"host":"the-host",
								"domain_guid":"my-domain-guid",
								"space_guid":"my-space-guid",
								"generate_port":false
							}
						`),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
					),
				)
			})

			It("tries to create a route", func() {
				repo.CreateInSpace("the-host", "", "my-domain-guid", "my-space-guid", 0, false)
				Expect(ccServer.ReceivedRequests()).To(HaveLen(1))
			})

			Context("when creating the route succeeds", func() {
				BeforeEach(func() {
					h := ccServer.GetHandler(0)
					ccServer.SetHandler(0, ghttp.CombineHandlers(
						h,
						ghttp.RespondWith(http.StatusCreated, `
								{
									"metadata": { "guid": "my-route-guid" },
									"entity": { "host": "the-host" }
								}
							`),
					))
				})

				It("returns the created route", func() {
					createdRoute, err := repo.CreateInSpace("the-host", "", "my-domain-guid", "my-space-guid", 0, false)
					Expect(err).NotTo(HaveOccurred())
					Expect(createdRoute.Host).To(Equal("the-host"))
				})
			})
		})

		Context("when a path is given", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/v2/routes", "inline-relations-depth=1&async=true"),
						ghttp.VerifyJSON(`
							{
								"domain_guid":"my-domain-guid",
								"space_guid":"my-space-guid",
								"path":"/the-path",
								"generate_port":false
							}
						`),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
					),
				)
			})

			It("tries to create a route", func() {
				repo.CreateInSpace("", "the-path", "my-domain-guid", "my-space-guid", 0, false)
				Expect(ccServer.ReceivedRequests()).To(HaveLen(1))
			})

			Context("when creating the route succeeds", func() {
				BeforeEach(func() {
					h := ccServer.GetHandler(0)
					ccServer.SetHandler(0, ghttp.CombineHandlers(
						h,
						ghttp.RespondWith(http.StatusCreated, `
								{
									"metadata": { "guid": "my-route-guid" },
									"entity": { "path": "the-path" }
								}
							`),
					))
				})

				It("returns the created route", func() {
					createdRoute, err := repo.CreateInSpace("", "the-path", "my-domain-guid", "my-space-guid", 0, false)
					Expect(err).NotTo(HaveOccurred())
					Expect(createdRoute.Path).To(Equal("the-path"))
				})
			})

			Context("when creating the route fails", func() {
				BeforeEach(func() {
					ccServer.Close()
					ccServer = nil
				})

				It("returns an error", func() {
					_, err := repo.CreateInSpace("", "the-path", "my-domain-guid", "my-space-guid", 0, false)
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("when a port is given", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/v2/routes", "inline-relations-depth=1&async=true"),
						ghttp.VerifyJSON(`
							{
								"port":9090,
								"domain_guid":"my-domain-guid",
								"space_guid":"my-space-guid",
								"generate_port":false
							}
						`),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
					),
				)
			})

			It("tries to create a route", func() {
				repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 9090, false)
				Expect(ccServer.ReceivedRequests()).To(HaveLen(1))
			})

			Context("when creating the route succeeds", func() {
				BeforeEach(func() {
					h := ccServer.GetHandler(0)
					ccServer.SetHandler(0, ghttp.CombineHandlers(
						h,
						ghttp.RespondWith(http.StatusCreated, `
							{
								"metadata": { "guid": "my-route-guid" },
								"entity": { "port": 9090 }
							}
						`),
					))
				})

				It("returns the created route", func() {
					createdRoute, err := repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 9090, false)
					Expect(err).NotTo(HaveOccurred())
					Expect(createdRoute.Port).To(Equal(9090))
				})
			})
		})

		Context("when random-port is true", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/v2/routes", "inline-relations-depth=1&async=true"),
						ghttp.VerifyJSON(`
							{
								"domain_guid":"my-domain-guid",
								"space_guid":"my-space-guid",
								"generate_port":true
							}
						`),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
					),
				)
			})

			It("tries to create a route", func() {
				repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 0, true)
				Expect(ccServer.ReceivedRequests()).To(HaveLen(1))
			})

			Context("when creating the route succeeds", func() {
				BeforeEach(func() {
					h := ccServer.GetHandler(0)
					ccServer.SetHandler(0, ghttp.CombineHandlers(
						h,
						ghttp.RespondWith(http.StatusCreated, `
							{
								"metadata": { "guid": "my-route-guid" },
								"entity": { "port": 50321 }
							}
						`),
					))
				})

				It("returns the created route", func() {
					createdRoute, err := repo.CreateInSpace("", "", "my-domain-guid", "my-space-guid", 0, true)
					Expect(err).NotTo(HaveOccurred())
					Expect(createdRoute.Port).To(Equal(50321))
				})
			})
		})
	})

	Describe("Check routes", func() {
		var (
			ccServer *ghttp.Server
			domain   models.DomainFields
		)

		BeforeEach(func() {
			domain = models.DomainFields{
				Guid: "domain-guid",
			}
			ccServer = ghttp.NewServer()
			configRepo.SetApiEndpoint(ccServer.URL())
		})

		AfterEach(func() {
			ccServer.Close()
		})

		Context("when the route is found", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/routes/reserved/domain/domain-guid/host/my-host", "path=/some-path"),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
						ghttp.RespondWith(http.StatusNoContent, nil),
					),
				)
			})

			It("returns true", func() {
				found, err := repo.CheckIfExists("my-host", domain, "some-path")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeTrue())
			})
		})

		Context("when the route is not found", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/routes/reserved/domain/domain-guid/host/my-host", "path=/some-path"),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
						ghttp.RespondWith(http.StatusNotFound, nil),
					),
				)
			})

			It("returns false", func() {
				found, err := repo.CheckIfExists("my-host", domain, "some-path")
				Expect(err).NotTo(HaveOccurred())
				Expect(found).To(BeFalse())
			})
		})

		Context("when finding the route fails", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/routes/reserved/domain/domain-guid/host/my-host", "path=/some-path"),
						ghttp.VerifyHeader(http.Header{
							"accept": []string{"application/json"},
						}),
						ghttp.RespondWith(http.StatusForbidden, nil),
					),
				)
			})

			It("returns an error", func() {
				_, err := repo.CheckIfExists("my-host", domain, "some-path")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when the path is empty", func() {
			BeforeEach(func() {
				ccServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.RespondWith(http.StatusNoContent, nil),
					),
				)
			})

			It("should not add a path query param", func() {
				_, err := repo.CheckIfExists("my-host", domain, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(len(ccServer.ReceivedRequests())).To(Equal(1))
				req := ccServer.ReceivedRequests()[0]
				vals := req.URL.Query()
				_, ok := vals["path"]
				Expect(ok).To(BeFalse())
			})
		})
	})

	Describe("Bind routes", func() {
		It("binds routes", func() {
			ts, handler = testnet.NewServer([]testnet.TestRequest{
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "PUT",
					Path:     "/v2/apps/my-cool-app-guid/routes/my-cool-route-guid",
					Response: testnet.TestResponse{Status: http.StatusCreated, Body: ""},
				}),
			})
			configRepo.SetApiEndpoint(ts.URL)

			apiErr := repo.Bind("my-cool-route-guid", "my-cool-app-guid")
			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})

		It("unbinds routes", func() {
			ts, handler = testnet.NewServer([]testnet.TestRequest{
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "DELETE",
					Path:     "/v2/apps/my-cool-app-guid/routes/my-cool-route-guid",
					Response: testnet.TestResponse{Status: http.StatusCreated, Body: ""},
				}),
			})
			configRepo.SetApiEndpoint(ts.URL)

			apiErr := repo.Unbind("my-cool-route-guid", "my-cool-app-guid")
			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})

	})

	Describe("Delete routes", func() {
		It("deletes routes", func() {
			ts, handler = testnet.NewServer([]testnet.TestRequest{
				testapi.NewCloudControllerTestRequest(testnet.TestRequest{
					Method:   "DELETE",
					Path:     "/v2/routes/my-cool-route-guid",
					Response: testnet.TestResponse{Status: http.StatusCreated, Body: ""},
				}),
			})
			configRepo.SetApiEndpoint(ts.URL)

			apiErr := repo.Delete("my-cool-route-guid")
			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})
	})

})

var firstPageRoutesResponse = testnet.TestResponse{Status: http.StatusOK, Body: `
{
  "next_url": "/v2/spaces/the-space-guid/routes?inline-relations-depth=1&page=2",
  "resources": [
    {
      "metadata": {
        "guid": "route-1-guid"
      },
      "entity": {
        "host": "route-1-host",
        "path": "",
        "domain": {
          "metadata": {
            "guid": "domain-1-guid"
          },
          "entity": {
            "name": "cfapps.io"
          }
        },
        "space": {
          "metadata": {
            "guid": "space-1-guid"
          },
          "entity": {
            "name": "space-1"
          }
        },
        "apps": [
       	  {
       	    "metadata": {
              "guid": "app-1-guid"
            },
            "entity": {
              "name": "app-1"
       	    }
       	  }
        ],
         "service_instance_url": "/v2/service_instances/service-guid",
        "service_instance": {
           "metadata": {
              "guid": "service-guid",
              "url": "/v2/service_instances/service-guid"
           },
           "entity": {
              "name": "test-service",
              "credentials": {
                 "username": "user",
                 "password": "password"
              },
              "type": "managed_service_instance",
              "route_service_url": "https://something.awesome.com",
              "space_url": "/v2/spaces/space-1-guid"
           }
        }
      }
    }
  ]
}`}

var secondPageRoutesResponse = testnet.TestResponse{Status: http.StatusOK, Body: `
{
  "resources": [
    {
      "metadata": {
        "guid": "route-2-guid"
      },
      "entity": {
        "host": "route-2-host",
        "path": "/path-2",
        "domain": {
          "metadata": {
            "guid": "domain-2-guid"
          },
          "entity": {
            "name": "example.com"
          }
        },
        "space": {
          "metadata": {
            "guid": "space-2-guid"
          },
          "entity": {
            "name": "space-2"
          }
        },
        "apps": [
       	  {
       	    "metadata": {
              "guid": "app-2-guid"
            },
            "entity": {
              "name": "app-2"
       	    }
       	  },
       	  {
       	    "metadata": {
              "guid": "app-3-guid"
            },
            "entity": {
              "name": "app-3"
       	    }
       	  }
        ]
      }
    }
  ]
}`}

var findResponseBody = `
{ "resources": [
	{
		"metadata": {
			"guid": "my-route-guid"
		},
		"entity": {
			"host": "my-cool-app",
			"domain": {
				"metadata": {
					"guid": "my-domain-guid"
				}
			},
			"path": "/somepath"
		}
	}
]}`

var findResponseBodyWithoutPath = `
{ "resources": [
	{
		"metadata": {
			"guid": "my-route-guid"
		},
		"entity": {
			"host": "my-cool-app",
			"domain": {
				"metadata": {
					"guid": "my-domain-guid"
				}
			}
		}
	}
]}`

var firstPageRoutesOrgLvlResponse = testnet.TestResponse{Status: http.StatusOK, Body: `
{
  "next_url": "/v2/routes?q=organization_guid:my-org-guid&inline-relations-depth=1&page=2",
  "resources": [
    {
      "metadata": {
        "guid": "route-1-guid"
      },
      "entity": {
        "host": "route-1-host",
        "domain": {
          "metadata": {
            "guid": "domain-1-guid"
          },
          "entity": {
            "name": "cfapps.io"
          }
        },
        "space": {
          "metadata": {
            "guid": "space-1-guid"
          },
          "entity": {
            "name": "space-1"
          }
        },
        "apps": [
          {
            "metadata": {
              "guid": "app-1-guid"
            },
            "entity": {
              "name": "app-1"
            }
          }
        ],
        "service_instance_url": "/v2/service_instances/service-guid",
        "service_instance": {
           "metadata": {
              "guid": "service-guid",
              "url": "/v2/service_instances/service-guid"
           },
           "entity": {
              "name": "test-service",
              "credentials": {
                 "username": "user",
                 "password": "password"
              },
              "type": "managed_service_instance",
              "route_service_url": "https://something.awesome.com",
              "space_url": "/v2/spaces/space-1-guid"
           }
        }
      }
    }
  ]
}`}
