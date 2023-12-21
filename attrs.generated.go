package goml

import (
	"io"
)

type IDAttribute struct{
	Value string
}

func (a IDAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" id='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func IDAttr(value string) IDAttribute {
	return IDAttribute{Value: value}
}

func (a IDAttribute) RenderElement(w io.Writer) error { return nil }

type ClassAttribute struct{
	Value string
}

func (a ClassAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" class='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func ClassAttr(value string) ClassAttribute {
	return ClassAttribute{Value: value}
}

func (a ClassAttribute) RenderElement(w io.Writer) error { return nil }

type HrefAttribute struct{
	Value string
}

func (a HrefAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" href='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func HrefAttr(value string) HrefAttribute {
	return HrefAttribute{Value: value}
}

func (a HrefAttribute) RenderElement(w io.Writer) error { return nil }

type SrcAttribute struct{
	Value string
}

func (a SrcAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" src='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func SrcAttr(value string) SrcAttribute {
	return SrcAttribute{Value: value}
}

func (a SrcAttribute) RenderElement(w io.Writer) error { return nil }

type AltAttribute struct{
	Value string
}

func (a AltAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" alt='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func AltAttr(value string) AltAttribute {
	return AltAttribute{Value: value}
}

func (a AltAttribute) RenderElement(w io.Writer) error { return nil }

type TitleAttribute struct{
	Value string
}

func (a TitleAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" title='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func TitleAttr(value string) TitleAttribute {
	return TitleAttribute{Value: value}
}

func (a TitleAttribute) RenderElement(w io.Writer) error { return nil }

type WidthAttribute struct{
	Value string
}

func (a WidthAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" width='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func WidthAttr(value string) WidthAttribute {
	return WidthAttribute{Value: value}
}

func (a WidthAttribute) RenderElement(w io.Writer) error { return nil }

type HeightAttribute struct{
	Value string
}

func (a HeightAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" height='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func HeightAttr(value string) HeightAttribute {
	return HeightAttribute{Value: value}
}

func (a HeightAttribute) RenderElement(w io.Writer) error { return nil }

type PlaceholderAttribute struct{
	Value string
}

func (a PlaceholderAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" placeholder='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func PlaceholderAttr(value string) PlaceholderAttribute {
	return PlaceholderAttribute{Value: value}
}

func (a PlaceholderAttribute) RenderElement(w io.Writer) error { return nil }

type ValueAttribute struct{
	Value string
}

func (a ValueAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" value='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func ValueAttr(value string) ValueAttribute {
	return ValueAttribute{Value: value}
}

func (a ValueAttribute) RenderElement(w io.Writer) error { return nil }

type NameAttribute struct{
	Value string
}

func (a NameAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" name='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func NameAttr(value string) NameAttribute {
	return NameAttribute{Value: value}
}

func (a NameAttribute) RenderElement(w io.Writer) error { return nil }

type TypeAttribute struct{
	Value string
}

func (a TypeAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" type='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func TypeAttr(value string) TypeAttribute {
	return TypeAttribute{Value: value}
}

func (a TypeAttribute) RenderElement(w io.Writer) error { return nil }

type StyleAttribute struct{
	Value string
}

func (a StyleAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" style='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func StyleAttr(value string) StyleAttribute {
	return StyleAttribute{Value: value}
}

func (a StyleAttribute) RenderElement(w io.Writer) error { return nil }

type TargetAttribute struct{
	Value string
}

func (a TargetAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" target='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func TargetAttr(value string) TargetAttribute {
	return TargetAttribute{Value: value}
}

func (a TargetAttribute) RenderElement(w io.Writer) error { return nil }

type RelAttribute struct{
	Value string
}

func (a RelAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" rel='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func RelAttr(value string) RelAttribute {
	return RelAttribute{Value: value}
}

func (a RelAttribute) RenderElement(w io.Writer) error { return nil }

type RoleAttribute struct{
	Value string
}

func (a RoleAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" role='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func RoleAttr(value string) RoleAttribute {
	return RoleAttribute{Value: value}
}

func (a RoleAttribute) RenderElement(w io.Writer) error { return nil }

type FormAttribute struct{
	Value string
}

func (a FormAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" form='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func FormAttr(value string) FormAttribute {
	return FormAttribute{Value: value}
}

func (a FormAttribute) RenderElement(w io.Writer) error { return nil }

type MaxAttribute struct{
	Value string
}

func (a MaxAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" max='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func MaxAttr(value string) MaxAttribute {
	return MaxAttribute{Value: value}
}

func (a MaxAttribute) RenderElement(w io.Writer) error { return nil }

type MinAttribute struct{
	Value string
}

func (a MinAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" min='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func MinAttr(value string) MinAttribute {
	return MinAttribute{Value: value}
}

func (a MinAttribute) RenderElement(w io.Writer) error { return nil }

type StepAttribute struct{
	Value string
}

func (a StepAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" step='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func StepAttr(value string) StepAttribute {
	return StepAttribute{Value: value}
}

func (a StepAttribute) RenderElement(w io.Writer) error { return nil }

type AutoCompleteAttribute struct{
	Value string
}

func (a AutoCompleteAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" autocomplete='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func AutoCompleteAttr(value string) AutoCompleteAttribute {
	return AutoCompleteAttribute{Value: value}
}

func (a AutoCompleteAttribute) RenderElement(w io.Writer) error { return nil }

type PatternAttribute struct{
	Value string
}

func (a PatternAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" pattern='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func PatternAttr(value string) PatternAttribute {
	return PatternAttribute{Value: value}
}

func (a PatternAttribute) RenderElement(w io.Writer) error { return nil }

type ForAttribute struct{
	Value string
}

func (a ForAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" for='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func ForAttr(value string) ForAttribute {
	return ForAttribute{Value: value}
}

func (a ForAttribute) RenderElement(w io.Writer) error { return nil }

type LabelAttribute struct{
	Value string
}

func (a LabelAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" label='" + a.Value + "'")) 
	if err != nil {
		return err
	}

	return nil
}

func LabelAttr(value string) LabelAttribute {
	return LabelAttribute{Value: value}
}

func (a LabelAttribute) RenderElement(w io.Writer) error { return nil }

type DisabledAttribute struct{}

func (a DisabledAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" disabled")) 
	if err != nil {
		return err
	}

	return nil
}

func DisabledAttr() DisabledAttribute {
	return DisabledAttribute{}
}

func (a DisabledAttribute) RenderElement(w io.Writer) error { return nil }

type ReadOnlyAttribute struct{}

func (a ReadOnlyAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" readonly")) 
	if err != nil {
		return err
	}

	return nil
}

func ReadOnlyAttr() ReadOnlyAttribute {
	return ReadOnlyAttribute{}
}

func (a ReadOnlyAttribute) RenderElement(w io.Writer) error { return nil }

type CheckedAttribute struct{}

func (a CheckedAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" checked")) 
	if err != nil {
		return err
	}

	return nil
}

func CheckedAttr() CheckedAttribute {
	return CheckedAttribute{}
}

func (a CheckedAttribute) RenderElement(w io.Writer) error { return nil }

type RequiredAttribute struct{}

func (a RequiredAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" required")) 
	if err != nil {
		return err
	}

	return nil
}

func RequiredAttr() RequiredAttribute {
	return RequiredAttribute{}
}

func (a RequiredAttribute) RenderElement(w io.Writer) error { return nil }

type MultipleAttribute struct{}

func (a MultipleAttribute) RenderAttr(w io.Writer) error {
	_, err := w.Write([]byte(" multiple")) 
	if err != nil {
		return err
	}

	return nil
}

func MultipleAttr() MultipleAttribute {
	return MultipleAttribute{}
}

func (a MultipleAttribute) RenderElement(w io.Writer) error { return nil }
