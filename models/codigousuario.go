// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// CodigoUsuario represents a row from 'v1wq1ics1m037sn6.codigo_usuario'.
type CodigoUsuario struct {
	ID        int            `json:"id"`         // id
	Codigo    string         `json:"codigo"`     // codigo
	Nombres   string         `json:"nombres"`    // nombres
	Apellidos string         `json:"apellidos"`  // apellidos
	CreatedBy sql.NullString `json:"created_by"` // created_by
	UpdatedBy sql.NullString `json:"updated_by"` // updated_by

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CodigoUsuario exists in the database.
func (cu *CodigoUsuario) Exists() bool {
	return cu._exists
}

// Deleted provides information if the CodigoUsuario has been deleted from the database.
func (cu *CodigoUsuario) Deleted() bool {
	return cu._deleted
}

// Insert inserts the CodigoUsuario to the database.
func (cu *CodigoUsuario) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if cu._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.codigo_usuario (` +
		`codigo, nombres, apellidos, created_by, updated_by` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, cu.Codigo, cu.Nombres, cu.Apellidos, cu.CreatedBy, cu.UpdatedBy)
	res, err := db.Exec(sqlstr, cu.Codigo, cu.Nombres, cu.Apellidos, cu.CreatedBy, cu.UpdatedBy)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	cu.ID = int(id)
	cu._exists = true

	return nil
}

// Update updates the CodigoUsuario in the database.
func (cu *CodigoUsuario) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cu._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if cu._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE v1wq1ics1m037sn6.codigo_usuario SET ` +
		`codigo = ?, nombres = ?, apellidos = ?, created_by = ?, updated_by = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, cu.Codigo, cu.Nombres, cu.Apellidos, cu.CreatedBy, cu.UpdatedBy, cu.ID)
	_, err = db.Exec(sqlstr, cu.Codigo, cu.Nombres, cu.Apellidos, cu.CreatedBy, cu.UpdatedBy, cu.ID)
	return err
}

// Save saves the CodigoUsuario to the database.
func (cu *CodigoUsuario) Save(db XODB) error {
	if cu.Exists() {
		return cu.Update(db)
	}

	return cu.Insert(db)
}

// Delete deletes the CodigoUsuario from the database.
func (cu *CodigoUsuario) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cu._exists {
		return nil
	}

	// if deleted, bail
	if cu._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.codigo_usuario WHERE id = ?`

	// run query
	XOLog(sqlstr, cu.ID)
	_, err = db.Exec(sqlstr, cu.ID)
	if err != nil {
		return err
	}

	// set deleted
	cu._deleted = true

	return nil
}

// CodigoUsuarioByCodigo retrieves a row from 'v1wq1ics1m037sn6.codigo_usuario' as a CodigoUsuario.
//
// Generated from index 'UNIQ_8BA7FCA120332D99'.
func CodigoUsuarioByCodigo(db XODB, codigo string) (*CodigoUsuario, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, codigo, nombres, apellidos, created_by, updated_by ` +
		`FROM v1wq1ics1m037sn6.codigo_usuario ` +
		`WHERE codigo = ?`

	// run query
	XOLog(sqlstr, codigo)
	cu := CodigoUsuario{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, codigo).Scan(&cu.ID, &cu.Codigo, &cu.Nombres, &cu.Apellidos, &cu.CreatedBy, &cu.UpdatedBy)
	if err != nil {
		return nil, err
	}

	return &cu, nil
}

// CodigoUsuarioByID retrieves a row from 'v1wq1ics1m037sn6.codigo_usuario' as a CodigoUsuario.
//
// Generated from index 'codigo_usuario_id_pkey'.
func CodigoUsuarioByID(db XODB, id int) (*CodigoUsuario, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, codigo, nombres, apellidos, created_by, updated_by ` +
		`FROM v1wq1ics1m037sn6.codigo_usuario ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	cu := CodigoUsuario{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&cu.ID, &cu.Codigo, &cu.Nombres, &cu.Apellidos, &cu.CreatedBy, &cu.UpdatedBy)
	if err != nil {
		return nil, err
	}

	return &cu, nil
}
