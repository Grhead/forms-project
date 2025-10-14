package internal

import "google.golang.org/api/forms/v1"

func createForms(Title string, FormTitle string) (forms.Form, error) {

	return forms.Form{}, nil
}
func getForm() {

}
func setItems(items []*forms.Item) {
	var requests []*forms.Request
	for i, item := range items {
		requests = append(requests, &forms.Request{
			CreateItem: &forms.CreateItemRequest{
				Item: item,
				Location: &forms.Location{
					Index:           int64(i),
					ForceSendFields: []string{"Index"},
				},
			},
		})
	}
}
func getResponses() {

}
func detailResponses() {

}
