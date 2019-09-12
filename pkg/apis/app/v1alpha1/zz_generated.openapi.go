// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Condition":                      schema_pkg_apis_app_v1alpha1_Condition(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Deployments":                    schema_pkg_apis_app_v1alpha1_Deployments(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Env":                            schema_pkg_apis_app_v1alpha1_Env(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.GitSource":                      schema_pkg_apis_app_v1alpha1_GitSource(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.InfinispanConnectionProperties": schema_pkg_apis_app_v1alpha1_InfinispanConnectionProperties(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoApp":                      schema_pkg_apis_app_v1alpha1_KogitoApp(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppBuildObject":           schema_pkg_apis_app_v1alpha1_KogitoAppBuildObject(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppServiceObject":         schema_pkg_apis_app_v1alpha1_KogitoAppServiceObject(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppSpec":                  schema_pkg_apis_app_v1alpha1_KogitoAppSpec(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppStatus":                schema_pkg_apis_app_v1alpha1_KogitoAppStatus(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndex":                schema_pkg_apis_app_v1alpha1_KogitoDataIndex(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexSpec":            schema_pkg_apis_app_v1alpha1_KogitoDataIndexSpec(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexStatus":          schema_pkg_apis_app_v1alpha1_KogitoDataIndexStatus(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.ResourceMap":                    schema_pkg_apis_app_v1alpha1_ResourceMap(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Resources":                      schema_pkg_apis_app_v1alpha1_Resources(ref),
		"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.WebhookSecret":                  schema_pkg_apis_app_v1alpha1_WebhookSecret(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition - The condition for the kogito-cloud-operator",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_app_v1alpha1_Deployments(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Deployments ...",
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are ready to serve requests",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"starting": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are starting, may or may not succeed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"stopped": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments are not starting, unclear what next step will be",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"failed": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployments failed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_Env(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Env Data to define environment variables in key/value pair fashion",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of an environment variable",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Description: "Value for that environment variable",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_GitSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitSource Git coordinates to locate the source code to build",
				Properties: map[string]spec.Schema{
					"uri": {
						SchemaProps: spec.SchemaProps{
							Description: "Git URI for the s2i source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reference": {
						SchemaProps: spec.SchemaProps{
							Description: "Branch to use in the git repository",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"contextDir": {
						SchemaProps: spec.SchemaProps{
							Description: "Context/subdirectory where the code is located, relatively to repo root",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"uri"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_InfinispanConnectionProperties(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfinispanConnectionProperties is the confuguration needed for authenticating on Infinispan cluster. More information can be found at https://docs.jboss.org/infinispan/10.0/apidocs/org/infinispan/client/hotrod/configuration/package-summary.html#package.description",
				Properties: map[string]spec.Schema{
					"credentials": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.SecretCredentialsType"),
						},
					},
					"useAuth": {
						SchemaProps: spec.SchemaProps{
							Description: "UseAuth will be set to true if the credentials are set. Will set the property infinispan.client.hotrod.use_auth",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"authRealm": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the Infinispan authentication realm. Will set the property infinispan.client.hotrod.auth_realm",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"saslMechanism": {
						SchemaProps: spec.SchemaProps{
							Description: "SaslMechanism defined for the authentication. Will set the property infinispan.client.hotrod.sasl_mechanism",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceURI": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceURI is the service URI to connect to the Infinispan cluster, e.g. myinifisan-cluster:11222",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.SecretCredentialsType"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoApp(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoApp is the Schema for the kogitoapps API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppSpec", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoAppBuildObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoAppBuildObject Data to define how to build an application from source",
				Properties: map[string]spec.Schema{
					"incremental": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Env"),
									},
								},
							},
						},
					},
					"gitSource": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.GitSource"),
						},
					},
					"webhooks": {
						SchemaProps: spec.SchemaProps{
							Description: "WebHook secrets for build configs",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.WebhookSecret"),
									},
								},
							},
						},
					},
					"imageS2I": {
						SchemaProps: spec.SchemaProps{
							Description: "ImageS2I is used by build configurations to build the image from source",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Image"),
						},
					},
					"imageRuntime": {
						SchemaProps: spec.SchemaProps{
							Description: "ImageRuntime is used build configurations to build a final runtime image based on a s2i configuration",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Image"),
						},
					},
					"native": {
						SchemaProps: spec.SchemaProps{
							Description: "Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is quarkus. See: https://www.graalvm.org/docs/reference-manual/aot-compilation/",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"gitSource"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Env", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.GitSource", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Image", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.WebhookSecret"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoAppServiceObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoAppServiceObject Data to define the service of the kogito app",
				Properties: map[string]spec.Schema{
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels for the application service",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoAppSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoAppSpec defines the desired state of KogitoApp",
				Properties: map[string]spec.Schema{
					"runtime": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the runtime used, either quarkus or springboot, defaults to quarkus",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Env"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "The resources for the deployed pods, like memory and cpu",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Resources"),
						},
					},
					"build": {
						SchemaProps: spec.SchemaProps{
							Description: "S2I Build configuration",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppBuildObject"),
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Description: "Service configuration",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppServiceObject"),
						},
					},
				},
				Required: []string{"build"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Env", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppBuildObject", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoAppServiceObject", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Resources"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoAppStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoAppStatus defines the observed state of KogitoApp",
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"route": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"deployments": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Deployments"),
						},
					},
				},
				Required: []string{"conditions", "deployments"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Condition", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.Deployments"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoDataIndex(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoDataIndex is the Schema for the kogitodataindices API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexSpec", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KogitoDataIndexStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoDataIndexSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoDataIndexSpec defines the desired state of KogitoDataIndex",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the Data Index Service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the number of pod replicas that the Data Index Service will spin",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Description: "Env is a collection of additional environment variables to add to the Data Index container",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image to use for this service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"memoryLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "MemoryLimit is the limit of Memory which will be available for the container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"memoryRequest": {
						SchemaProps: spec.SchemaProps{
							Description: "MemoryRequest is the request of Memory which will be available for the container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cpuLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "CPULimit is the limit of CPU which will be available for the container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cpuRequest": {
						SchemaProps: spec.SchemaProps{
							Description: "CPURequest is the request of CPU which will be available for the container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"infinispan": {
						SchemaProps: spec.SchemaProps{
							Description: "Infinispan has the data used by the Kogito Data Index to connect to the Infinispan cluster",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.InfinispanConnectionProperties"),
						},
					},
					"kafka": {
						SchemaProps: spec.SchemaProps{
							Description: "Kafka has the data used by the Kogito Data Index to connecto to a Kafka cluster",
							Ref:         ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KafkaConnectionProperties"),
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.InfinispanConnectionProperties", "github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.KafkaConnectionProperties"},
	}
}

func schema_pkg_apis_app_v1alpha1_KogitoDataIndexStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoDataIndexStatus defines the observed state of KogitoDataIndex",
				Properties: map[string]spec.Schema{
					"deploymentStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the Data Index Service Deployment created and managed by it",
							Ref:         ref("k8s.io/api/apps/v1.StatefulSetStatus"),
						},
					},
					"serviceStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the Database Service created and managed by it",
							Ref:         ref("k8s.io/api/core/v1.ServiceStatus"),
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "OK when all resources are created successfully",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.DataIndexCondition"),
									},
								},
							},
						},
					},
					"dependenciesStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "All dependencies OK means that everything was found within the namespace",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"route": {
						SchemaProps: spec.SchemaProps{
							Description: "Route is where the service is exposed",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"deploymentStatus", "serviceStatus", "conditions", "dependenciesStatus"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.DataIndexCondition", "k8s.io/api/apps/v1.StatefulSetStatus", "k8s.io/api/core/v1.ServiceStatus"},
	}
}

func schema_pkg_apis_app_v1alpha1_ResourceMap(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceMap Data to define a list of possible Resources",
				Properties: map[string]spec.Schema{
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource type like cpu and memory",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Description: "Value of this resource in Kubernetes format",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"resource", "value"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_Resources(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Resources Data to define Resources needed for each deployed pod",
				Properties: map[string]spec.Schema{
					"limits": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.ResourceMap"),
									},
								},
							},
						},
					},
					"requests": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.ResourceMap"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1.ResourceMap"},
	}
}

func schema_pkg_apis_app_v1alpha1_WebhookSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebhookSecret Secret to use for a given webhook",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "WebHook type, either GitHub or Generic",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret value for webhook",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
