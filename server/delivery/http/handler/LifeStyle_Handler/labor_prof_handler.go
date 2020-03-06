package labor_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/web1_group_project/hospital_server/Laboratorist"
	"github.com/web1_group_project/hospital_server/entity"
)

// LaborProfileHandler handles comment related http requests
type LaborProfileHandler struct {
	profileService Laboratorist.LabratoristProfileService
}

// NewPharmMedicineHandler returns new AdminCommentHandler object
func NewLaborProfileHandler(profcService Laboratorist.LabratoristProfileService) *LaborProfileHandler {
	return &LaborProfileHandler{profileService: profcService}
}

// GetMedicines handles GET /v1/admin/comments request
func (ach *LaborProfileHandler) GetProfiles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	medicines, errs := ach.profileService.Profiles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(medicines, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleMedicine handles GET /v1/admin/comments/:id request
func (ach *LaborProfileHandler) GetSingleProfile(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.profileService.Profile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(medicine, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostMedicine handles POST /v1/admin/comments request
func (ach *LaborProfileHandler) PostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("post  Pharmacist")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	medicine := &entity.Laboratorist{}

	err := json.Unmarshal(body, medicine)
	log.Println("Unmarshal1 done")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	log.Println("Unmarshal2 done")

	medicine, errs := ach.profileService.AddProfile(medicine)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/labor/profile/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutMedicine handles PUT /v1/admin/comments/:id request
func (ach *LaborProfileHandler) PutProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.profileService.Profile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &medicine)

	medicine, errs = ach.profileService.UpdateProfile(medicine)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(medicine, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//DeleteMedicine handles DELETE /v1/admin/comments/:id request
func (ach *LaborProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.profileService.DeleteProfile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
