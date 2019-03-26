package common

import (
	"testing"

	v2 "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
	apiclient "github.com/openservicebrokerapi/osb-checker/client"
	"github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateInstance(
	t *testing.T,
	instanceID string,
	req *v2.ServiceInstanceUpdateRequest,
	async bool,
) {
	Convey("UPDATE - request syntax", t, func() {

		So(testAPIVersionHeader(config.GenerateInstanceURL(instanceID), "PATCH"), ShouldBeNil)
		So(testAuthentication(config.GenerateInstanceURL(instanceID), "PATCH"), ShouldBeNil)
		if async {
			So(testAsyncParameters(config.GenerateInstanceURL(instanceID), "PATCH"), ShouldBeNil)
		}

		var emptyValue, fakeValue = "", "xxxx-xxxx"
		Convey("should reject if missing service_id", func() {
			tempBody := new(v2.ServiceInstanceUpdateRequest)
			deepCopy(req, tempBody)
			tempBody.ServiceID = &emptyValue
			code, _, err := apiclient.Default.UpdateInstance(instanceID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if service_id is invalid", func() {
			tempBody := new(v2.ServiceInstanceUpdateRequest)
			deepCopy(req, tempBody)
			tempBody.ServiceID = &fakeValue
			code, _, err := apiclient.Default.UpdateInstance(instanceID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if parameters are not following schema", func() {
			tempBody := new(v2.ServiceInstanceUpdateRequest)
			deepCopy(req, tempBody)
			tempBody.Parameters = map[string]interface{}{
				"can-not": "be-good",
			}
			if err := testCatalogSchema(&SchemaOpts{
				ServiceID:  *tempBody.ServiceID,
				PlanID:     tempBody.PlanID,
				Parameters: tempBody.Parameters,
				SchemaType: config.TypeServiceInstance,
				Action:     config.ActionUpdate,
			}); err == nil {
				return
			}
			code, _, err := apiclient.Default.UpdateInstance(instanceID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})
	})

	Convey("UPDATE - new", t, func() {
		Convey("should accept a valid update request", func() {
			tempBody := new(v2.ServiceInstanceUpdateRequest)
			deepCopy(req, tempBody)
			code, asyncBody, err := apiclient.Default.UpdateInstance(instanceID, tempBody, async)

			So(err, ShouldBeNil)
			if async {
				So(code, ShouldEqual, 202)
				So(testJSONSchema(asyncBody), ShouldBeNil)
			} else {
				So(code, ShouldEqual, 200)
			}
		})
	})

	if async {
		Convey("UPDATE - poll", t, func(c C) {
			testPollInstanceLastOperation(instanceID)

			So(pollInstanceLastOperationStatus(instanceID), ShouldBeNil)
		})
	}
}
