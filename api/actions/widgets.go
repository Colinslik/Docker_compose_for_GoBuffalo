package actions

import (
        "github.com/gobuffalo/authrecipe/models"
        "github.com/gobuffalo/buffalo"
        "github.com/gobuffalo/pop"
        "github.com/pkg/errors"
)

func WidgetsNew(c buffalo.Context) error {
        c.Set("widget", models.Widget{})
        return c.Render(200, r.HTML("widgets/new.html"))
}

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func Create(c buffalo.Context) error {
  // Allocate an empty Widget
  widget := &models.Widget{}

  // Bind widget to the html form elements
  if err := c.Bind(widget); err != nil {
    return errors.WithStack(err)
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.WithStack(errors.New("no transaction found"))
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(widget)
  if err != nil {
    return errors.WithStack(err)
  }

  if verrs.HasAny() {
    // Make widget available inside the html template
    c.Set("widget", widget)

    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.HTML("widgets/new.html"))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", "Widget was created successfully")

  // and redirect to the widgets index page
  return c.Redirect(302, "/widgets/%s", widget.ID)
}
