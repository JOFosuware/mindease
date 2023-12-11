package handlers

import (
	"net/http"
	"strconv"

	"github.com/jofosuware/mindease/internal/config"
	"github.com/jofosuware/mindease/internal/driver"
	"github.com/jofosuware/mindease/internal/forms"
	"github.com/jofosuware/mindease/internal/helpers"
	"github.com/jofosuware/mindease/internal/models"
	"github.com/jofosuware/mindease/internal/render"
	"github.com/jofosuware/mindease/internal/repository"
	"github.com/jofosuware/mindease/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// GetHome is the home page handler
func (m *Repository) GetHome(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "home",
	}

	data["meta"] = fmd

	render.Template(w, r, "home.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

// GetBookSession handles session page handler
func (m *Repository) GetBookSession(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "book",
	}

	data["meta"] = fmd

	render.Template(w, r, "book.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

// GetBipolarExpert authenticates user and lists bipolar experts
func (m *Repository) GetBipolarExperts(w http.ResponseWriter, r *http.Request) {
	client, _ := m.App.Session.Get(r.Context(), "client").(models.Client)
	// if !ok {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	m.App.ErrorLog.Println("Failed to get client data from session")
	// 	return
	// }
	if !helpers.IsAuthenticated(r) {
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		return
	}

	pdrs, err := m.DB.FetchProviders()
	if err != nil {
		m.App.ErrorLog.Println(err)
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		return
	}

	meta := models.FormMetaData{
		Section: "book",
	}

	data := make(map[string]interface{})
	data["providers"] = pdrs
	data["meta"] = meta
	data["client"] = client

	render.Template(w, r, "providers.page.html", &models.TemplateData{
		Data: data,
	})
}

// GetClientForm handles request for client's form
func (m *Repository) GetClientForm(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "book",
	}

	data["meta"] = fmd

	render.Template(w, r, "register.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostClientInformation handles the processing of client's information
func (m *Repository) PostClientInformation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		m.App.ErrorLog.Println("Form Parse error", err)
		return
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone, _ := strconv.Atoi(r.Form.Get("phone"))

	c := models.Client{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "phone")

	if !form.Valid() {
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		m.App.ErrorLog.Println("empty form field", err)
		return
	}

	cl, err := m.DB.FetchClientByEmail(c.Email)

	if err != nil && err.Error() != "sql: no rows in result set" {
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	if cl.Email == email {
		m.App.InfoLog.Println("Username already exist!, choose another one")
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		return
	}

	id, err := m.DB.InsertClient(c)
	if err != nil {
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	m.App.Session.Put(r.Context(), "client_id", id)
	m.App.Session.Put(r.Context(), "client", c)
	http.Redirect(w, r, "/bipolar", http.StatusSeeOther)

}

// Providers
// GetProvidersPage handles request for providers welcoming page
func (m *Repository) GetProviderPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "provider",
	}

	data["meta"] = fmd

	render.Template(w, r, "starting.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// GetProviderForm handles request from providers' form
func (m *Repository) GetProviderForm(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "provider",
	}

	data["meta"] = fmd

	render.Template(w, r, "signup.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostClientProfile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		m.App.ErrorLog.Println("Form Parse error", err)
		return
	}

	name := r.Form.Get("name")
	description := r.Form.Get("description")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	photo, _, _ := r.FormFile("photo")
	defer photo.Close()

	providerImage, err := helpers.ProcessImage(photo)
	if err != nil {
		http.Redirect(w, r, "/admin/signup", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	p := models.Provider{
		Name:        name,
		Username:    username,
		Description: description,
		Password:    password,
		Photo:       providerImage,
	}

	form := forms.New(r.PostForm)
	form.Required("name", "description")

	if !form.Valid() {
		http.Redirect(w, r, "/admin/signup", http.StatusSeeOther)
		m.App.ErrorLog.Println("empty form field", err)
		return
	}

	pdr, err := m.DB.FetchProvider(username)

	if err != nil && err.Error() != "sql: no rows in result set" {
		m.App.ErrorLog.Println(err)
		http.Redirect(w, r, "/admin/signup", http.StatusSeeOther)
		return
	}

	if pdr.Username == username {
		m.App.ErrorLog.Println("Username already taken!")
		http.Redirect(w, r, "/admin/signup", http.StatusSeeOther)
		return
	}

	err = m.DB.InsertProvider(p)
	if err != nil {
		m.App.ErrorLog.Println(err)
		http.Redirect(w, r, "/admin/signup", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})

	data["provider"] = p

	//m.App.Session.Put(context.Background(), "provider", p)
	//http.Redirect(w, r, fmt.Sprintf("/profile/%s", username), http.StatusSeeOther)
	render.Template(w, r, "profile.page.html", &models.TemplateData{
		Data: data,
	})
}

// GetProviderLoginForm handles request for providers' login form
func (m *Repository) GetProviderLoginForm(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	//errMsg := m.App.Session.Pop(r.Context(), "error").(string)

	meta := models.FormMetaData{
		Section: "login",
	}

	data["meta"] = meta
	//data["error"] = errMsg
	render.Template(w, r, "login.page.html", &models.TemplateData{
		Data: data,
	})
}

// PostLoginProvider handles the request to allow access to a provider
func (m *Repository) PostLoginProvider(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println(err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.Required("username", "password")

	if !form.Valid() {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		m.App.ErrorLog.Println("empty form field", err)
		return
	}

	p := models.Provider{
		Username: username,
		Password: password,
	}

	pdr, err := m.DB.FetchProvider(p.Username)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	if pdr.Password != p.Password {
		//m.App.Session.Put(r.Context(), "error", "Password incorrect!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	notifs, err := m.DB.FetchNotifications()
	if err != nil {
		m.App.ErrorLog.Println(err)
	}

	data := make(map[string]interface{})
	data["provider"] = pdr
	data["notifications"] = notifs

	m.App.Session.Put(r.Context(), "provider", pdr)
	render.Template(w, r, "provider.page.html", &models.TemplateData{
		Data: data,
	})

}

// GetProviderProfile handles request for provider's profile page
func (m *Repository) GetProviderProfile(w http.ResponseWriter, r *http.Request) {
	pdrs, err := m.DB.FetchProviders()
	if err != nil {
		m.App.ErrorLog.Println(err)
		http.Redirect(w, r, "/create-account", http.StatusSeeOther)
		return
	}

	meta := models.FormMetaData{
		Section: "experts",
	}

	data := make(map[string]interface{})
	data["providers"] = pdrs
	data["meta"] = meta

	render.Template(w, r, "experts.page.html", &models.TemplateData{
		Data: data,
	})
}

// Notification
// PostToProvider
func (m *Repository) PostToProvider(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println(err)
		return
	}

	m.App.ErrorLog.Println("hits..")

	name := r.Form.Get("name")
	phone := r.Form.Get("phone")
	condition := r.Form.Get("condition")

	form := forms.New(r.PostForm)
	form.Required("name", "phone", "condition")

	if !form.Valid() {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		m.App.ErrorLog.Println("empty form field", err)
		return
	}

	n := models.Notification{
		Name:      name,
		Phone:     phone,
		Condition: condition,
	}

	err = m.DB.InsertNotification(n)
	if err != nil {
		m.App.ErrorLog.Println(err)
		return
	}

	render.Template(w, r, "toclient.page.html", &models.TemplateData{})
}

// Pharmacy
// GetDrugForm handles request for drug request form
func (m *Repository) GetDrugForm(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	meta := models.FormMetaData{
		Section: "drugs",
	}

	data["meta"] = meta

	render.Template(w, r, "request.page.html", &models.TemplateData{
		Data: data,
	})
}

// PostDrugRequest handles request to process drug request
func (m *Repository) PostDrugRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Redirect(w, r, "/request-drug", http.StatusSeeOther)
		m.App.ErrorLog.Println("Form Parse error", err)
		return
	}

	name := r.Form.Get("name")
	institution := r.Form.Get("institution")
	physician := r.Form.Get("physician")
	location := r.Form.Get("location")
	photo, _, _ := r.FormFile("prescription")
	defer photo.Close()

	formImage, err := helpers.ProcessImageLarge(photo)
	if err != nil {
		http.Redirect(w, r, "/request-drug", http.StatusSeeOther)
		m.App.ErrorLog.Println(err)
		return
	}

	p := models.Prescription{
		Name:        name,
		Institution: institution,
		Physician:   physician,
		Location:    location,
		Image:       formImage,
	}

	form := forms.New(r.PostForm)
	form.Required("name", "institution", "physician")

	if !form.Valid() {
		http.Redirect(w, r, "/request-drug", http.StatusSeeOther)
		m.App.ErrorLog.Println("empty form field", err)
		return
	}

	id, err := m.DB.InsertPrescription(p)
	if err != nil {
		m.App.ErrorLog.Println(err)
		http.Redirect(w, r, "/request-drug", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})

	data["form_id"] = id

	m.App.Session.Put(r.Context(), "form_id", id)
	render.Template(w, r, "processing.page.html", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "about",
	}

	data["meta"] = fmd

	render.Template(w, r, "about.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) GetServicesPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "services",
	}

	data["meta"] = fmd

	render.Template(w, r, "services.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// GetPortfolioPage
func (m *Repository) GetPortfolioPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	fmd := models.FormMetaData{
		Section: "portfolio",
	}

	data["meta"] = fmd

	render.Template(w, r, "portfolio.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
