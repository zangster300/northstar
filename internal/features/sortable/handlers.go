package sortable

import (
	"net/http"

	"northstar/internal/features/sortable/pages"
)

func HandleSortablePage(w http.ResponseWriter, r *http.Request) {
	if err := pages.SortableInitial().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
