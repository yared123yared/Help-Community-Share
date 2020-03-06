package pharm_handler_test

import "testing"

// import (
// 	"bytes"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/julienschmidt/httprouter"
// 	PharmacistRepo "github.com/web1_group_project/hospital_server/Pharmacist/repository"
// 	PharmacistService "github.com/web1_group_project/hospital_server/Pharmacist/service"
// 	handler2 "github.com/web1_group_project/hospital_server/delivery/http/handler/pharm_handler"
// 	"github.com/web1_group_project/hospital_server/entity"
// )

func TestGetPrescription(t *testing.T) {
	// req, err := http.NewRequest("GET", "/pharm/presc", nil)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ProfReq := PharmacistRepo.NewMockPrescriptionGormRepo(nil)
	// profSrv := PharmacistService.NewPrescriptionService(ProfReq)
	// profHandler := handler2.NewPharmPrescriptionHandler(profSrv)

	// rr := httptest.NewRecorder()
	// router := httprouter.New()

	// router.GET("/pharm/presc", profHandler.GetPrescriptions)
	// router.ServeHTTP(rr, req)

	// if statu := rr.Code; statu == http.StatusOK {
	// 	t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	// }
	// expected := entity.PharmacistMock
	// body, err := ioutil.ReadAll(rr.Body)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if bytes.Contains(body, []byte("Mock prescription 01 Description")) {
	// 	t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	// }

}
