package dbrepo

import (
	"context"
	"time"

	"github.com/jofosuware/mindease/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// // Authenticate authenticates a user
// func (m *postgresDBRepo) Authenticate(username, password string) (models.User, error) {

// 	u, err := m.FetchUser(username)
// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// 	if err == bcrypt.ErrMismatchedHashAndPassword {
// 		return models.User{}, errors.New("incorrect password")
// 	} else if err != nil {
// 		return models.User{}, err
// 	}

// 	return u, nil
// }

// InsertClient add client to the database
func (m *postgresDBRepo) InsertClient(c models.Client) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	query := `
		insert into clients (name, email, phone, created_at, updated_at) 
		values ($1, $2, $3, $4, $5) returning id
	`
	err := m.DB.QueryRowContext(ctx, query,
		c.Name,
		c.Email,
		c.Phone,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// FetchClientByEmail retrieves a client by email
func (m *postgresDBRepo) FetchClientByEmail(email string) (models.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c := models.Client{}
	quary := `select 
				id, name, email, phone, created_at, updated_at
		      from clients where email = $1`

	row := m.DB.QueryRowContext(ctx, quary, email)
	err := row.Scan(
		&c.Name,
		&c.Email,
		&c.Phone,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return models.Client{}, err
	}

	return c, nil
}

// InsertProvider save the provider profile into the database
func (m *postgresDBRepo) InsertProvider(p models.Provider) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into providers 
				(name, username, description, password, image, created_at, updated_at) 
			  values 
			  	($1, $2, $3, $4, $5, $6, $7) 
	`
	_, err := m.DB.ExecContext(ctx, query,
		p.Name,
		p.Username,
		p.Description,
		p.Password,
		p.Photo,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// FetchProvider retrieves a provider's profile data
func (m *postgresDBRepo) FetchProvider(username string) (models.Provider, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := models.Provider{}
	quary := `select 
				id, name, username, description, password, image, created_at, updated_at
		      from providers where username = $1`

	row := m.DB.QueryRowContext(ctx, quary, username)
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Username,
		&p.Description,
		&p.Password,
		&p.Photo,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return models.Provider{}, err
	}

	return p, nil
}

// FetchProviders retrieves a provider's profile data
func (m *postgresDBRepo) FetchProviders() ([]models.Provider, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var p []models.Provider
	quary := `select 
				id, name, username, description, image, created_at, updated_at
		      from providers`

	rows, err := m.DB.QueryContext(ctx, quary)
	if err != nil {
		return p, err
	}

	defer rows.Close()

	for rows.Next() {
		pdr := models.Provider{}

		err = rows.Scan(
			&pdr.ID,
			&pdr.Name,
			&pdr.Username,
			&pdr.Description,
			&pdr.Photo,
			&pdr.CreatedAt,
			&pdr.UpdatedAt,
		)

		if err != nil {
			return p, err
		}

		p = append(p, pdr)

		if err != nil {
			return p, err
		}

		if err = rows.Err(); err != nil {
			return p, err
		}
	}

	return p, nil
}

// InsertNotification
func (m *postgresDBRepo) InsertNotification(n models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into notifications 
				(name, phone, condition, created_at, updated_at) 
			  values 
			  	($1, $2, $3, $4, $5) 
	`
	_, err := m.DB.ExecContext(ctx, query,
		n.Name,
		n.Phone,
		n.Condition,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// FetchNotifications
func (m *postgresDBRepo) FetchNotifications() ([]models.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var n []models.Notification
	quary := `select 
				id, name, phone, condition, created_at, updated_at
		      from notifications`

	rows, err := m.DB.QueryContext(ctx, quary)
	if err != nil {
		return n, err
	}

	defer rows.Close()

	for rows.Next() {
		notif := models.Notification{}

		err = rows.Scan(
			&notif.ID,
			&notif.Name,
			&notif.Phone,
			&notif.Condition,
			&notif.CreatedAt,
			&notif.UpdatedAt,
		)

		if err != nil {
			return n, err
		}

		n = append(n, notif)

		if err != nil {
			return n, err
		}

		if err = rows.Err(); err != nil {
			return n, err
		}
	}

	return n, nil
}

// Pharmacy Database
// InsertPrescription
func (m *postgresDBRepo) InsertPrescription(p models.Prescription) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	query := `insert into prescriptions 
				(form_id, name, institution, physician, location, image, created_at, updated_at) 
			  values 
			  	($1, $2, $3, $4, $5, $6, $7, $8) returning id
	`
	err := m.DB.QueryRowContext(ctx, query,
		p.FormID,
		p.Name,
		p.Institution,
		p.Physician,
		p.Location,
		p.Image,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return newID, err
	}

	return newID, nil
}

// FetchPrescriptions
func (m *postgresDBRepo) FetchPrescription(formId string) (models.Prescription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := models.Prescription{}
	quary := `select 
				id, form_id, name, institution, physician, location, image, created_at, updated_at
		      from prescription where form_id = $1`

	row := m.DB.QueryRowContext(ctx, quary, formId)
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Institution,
		&p.Physician,
		&p.Location,
		&p.Image,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return models.Prescription{}, err
	}

	return p, nil
}
