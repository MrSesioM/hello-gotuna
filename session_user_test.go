package gotdd_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcalbg/gotdd"
	"github.com/alcalbg/gotdd/test/assert"
	"github.com/alcalbg/gotdd/test/doubles"
)

func TestReadingUserIDFromEmptyStore(t *testing.T) {

	r := &http.Request{}
	sessionStoreSpy := doubles.NewGorillaSessionStoreSpy("")
	ses := gotdd.NewSession(sessionStoreSpy)

	assert.Equal(t, true, ses.IsGuest(r))

	id, err := ses.GetUserID(r)
	assert.Error(t, err)
	assert.Equal(t, "", id)

	assert.Equal(t, 0, sessionStoreSpy.SaveCalls)
}

func TestSaveUserIDAndRetrieve(t *testing.T) {

	r := &http.Request{}
	w := httptest.NewRecorder()
	sessionStoreSpy := doubles.NewGorillaSessionStoreSpy("")
	ses := gotdd.NewSession(sessionStoreSpy)

	err := ses.SetUserID(w, r, doubles.FakeUser1.GetID())
	assert.NoError(t, err)

	id, err := ses.GetUserID(r)
	assert.NoError(t, err)
	assert.Equal(t, doubles.FakeUser1.GetID(), id)
	assert.Equal(t, false, ses.IsGuest(r))
	assert.Equal(t, 1, sessionStoreSpy.SaveCalls)
}

func TestSettingUserLocale(t *testing.T) {

	r := &http.Request{}
	w := httptest.NewRecorder()
	sessionStoreSpy := doubles.NewGorillaSessionStoreSpy("")
	ses := gotdd.NewSession(sessionStoreSpy)

	assert.Equal(t, "", ses.GetUserLocale(r))

	err := ses.SetUserLocale(w, r, "fr-FR")
	assert.NoError(t, err)

	assert.Equal(t, "fr-FR", ses.GetUserLocale(r))
	assert.Equal(t, 1, sessionStoreSpy.SaveCalls)
}
