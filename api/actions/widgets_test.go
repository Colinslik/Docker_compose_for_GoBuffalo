package actions

import (
        "github.com/gobuffalo/authrecipe/models"
)

func (as *ActionSuite) Test_Widgets_New() {
        res := as.HTML("/widgets/new").Get()
        as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_WidgetsResource_Create() {
  // clear out the uploads directory
  os.RemoveAll("./uploads")

  // setup a new Widget
  w := &models.Widget{Name: "Foo"}

  // find the file we want to upload
  r, err := os.Open("./logo.svg")
  as.NoError(err)
  // setup a new httptest.File to hold the file information
  f := httptest.File{
    // ParamName is the name of the form parameter
    ParamName: "someFile",
    // FileName is the name of the file being uploaded
    FileName: r.Name(),
    // Reader is the file that is to be uploaded, any io.Reader works
    Reader: r,
  }

  // Post the Widget and the File(s) to /widgets
  res, err := as.HTML("/widgets").MultiPartPost(w, f)
  as.NoError(err)
  as.Equal(302, res.Code)

  // assert the file exists on disk
  _, err = os.Stat("./uploads/logo.svg")
  as.NoError(err)

  // assert the Widget was saved to the DB correctly
  as.NoError(as.DB.First(w))
  as.Equal("Foo", w.Name)
  as.NotZero(w.ID)
}

