// Code generated by reactGen. DO NOT EDIT.

package downloadbutton

import "myitcv.io/react"

type DownloadItemsBtnElem struct {
	react.Element
}

func (d DownloadItemsBtnDef) ShouldComponentUpdateIntf(nextProps react.Props, prevState, nextState react.State) bool {
	res := false

	{
		res = d.Props() != nextProps.(DownloadItemsBtnProps) || res
	}
	return res
}

func buildDownloadItemsBtn(cd react.ComponentDef) react.Component {
	return DownloadItemsBtnDef{ComponentDef: cd}
}

func buildDownloadItemsBtnElem(props DownloadItemsBtnProps, children ...react.Element) *DownloadItemsBtnElem {
	return &DownloadItemsBtnElem{
		Element: react.CreateElement(buildDownloadItemsBtn, props, children...),
	}
}

// IsProps is an auto-generated definition so that DownloadItemsBtnProps implements
// the myitcv.io/react.Props interface.
func (d DownloadItemsBtnProps) IsProps() {}

// Props is an auto-generated proxy to the current props of DownloadItemsBtn
func (d DownloadItemsBtnDef) Props() DownloadItemsBtnProps {
	uprops := d.ComponentDef.Props()
	return uprops.(DownloadItemsBtnProps)
}

func (d DownloadItemsBtnProps) EqualsIntf(val react.Props) bool {
	return d == val.(DownloadItemsBtnProps)
}

var _ react.Props = DownloadItemsBtnProps{}
