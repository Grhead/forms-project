package services

import (
	"context"
	"testing"
	"tusur-forms/internal/domain"
)

func TestGoogleForms_NewService(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filename string
		want     FormService
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var g GoogleForms
			got, gotErr := g.NewService(context.Background(), tt.filename)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewService() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewService() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_googleFormsAdapter_NewForm(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		title         string
		documentTitle string
		want          domain.Form
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var g googleFormsAdapter
			got, gotErr := g.NewForm(tt.title, tt.documentTitle)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewForm() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewForm() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_googleFormsAdapter_GetForm(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		formID  string
		want    domain.Form
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var g googleFormsAdapter
			got, gotErr := g.GetForm(tt.formID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetForm() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetForm() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_googleFormsAdapter_SetQuestions(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		form      domain.Form
		questions []*domain.Question
		want      domain.Form
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var g googleFormsAdapter
			got, gotErr := g.SetQuestions(tt.form, tt.questions)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SetQuestions() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SetQuestions() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SetQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
