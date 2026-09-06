// Google Sheets API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package sheets

// Adds a new banded range to the spreadsheet.
type AddBandingRequest struct {
	BandedRange BandedRange `json:"bandedRange,omitempty\"` // The banded range to add. The bandedRangeId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of a range that already exists.)

}

// The result of adding a banded range.
type AddBandingResponse struct {
	BandedRange BandedRange `json:"bandedRange,omitempty\"` // The banded range that was added.

}

// Adds a chart to a sheet in the spreadsheet.
type AddChartRequest struct {
	Chart EmbeddedChart `json:"chart,omitempty\"` // The chart that should be added to the spreadsheet, including the position where it should be placed. The chartId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of an embedded object that already exists.)

}

// The result of adding a chart to a spreadsheet.
type AddChartResponse struct {
	Chart EmbeddedChart `json:"chart,omitempty\"` // The newly added chart.

}

// Adds a new conditional format rule at the given index. All subsequent rules' indexes are incremented.
type AddConditionalFormatRuleRequest struct {
	Index int `json:"index,omitempty\"` // The zero-based index where the rule should be inserted.

	Rule ConditionalFormatRule `json:"rule,omitempty\"` // The rule to add.

}

// Adds a data source. After the data source is added successfully, an associated DATA_SOURCE sheet is created and an execution is triggered to refresh the sheet to read data from the data source. The request requires an additional `bigquery.readonly` OAuth scope if you are adding a BigQuery data source.
type AddDataSourceRequest struct {
	DataSource DataSource `json:"dataSource,omitempty\"` // The data source to add.

}

// The result of adding a data source.
type AddDataSourceResponse struct {
	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // The data execution status.

	DataSource DataSource `json:"dataSource,omitempty\"` // The data source that was created.

}

// Creates a group over the specified range. If the requested range is a superset of the range of an existing group G, then the depth of G is incremented and this new group G' has the depth of that group. For example, a group [C:D, depth 1] + [B:E] results in groups [B:E, depth 1] and [C:D, depth 2]. If the requested range is a subset of the range of an existing group G, then the depth of the new group G' becomes one greater than the depth of G. For example, a group [B:E, depth 1] + [C:D] results in groups [B:E, depth 1] and [C:D, depth 2]. If the requested range starts before and ends within, or starts within and ends after, the range of an existing group G, then the range of the existing group G becomes the union of the ranges, and the new group G' has depth one greater than the depth of G and range as the intersection of the ranges. For example, a group [B:D, depth 1] + [C:E] results in groups [B:E, depth 1] and [C:D, depth 2].
type AddDimensionGroupRequest struct {
	RangeValue DimensionRange `json:"range,omitempty\"` // The range over which to create a group.

}

// The result of adding a group.
type AddDimensionGroupResponse struct {
	DimensionGroups []DimensionGroup `json:"dimensionGroups,omitempty\"` // All groups of a dimension after adding a group to that dimension.

}

// Adds a filter view.
type AddFilterViewRequest struct {
	Filter FilterView `json:"filter,omitempty\"` // The filter to add. The filterViewId field is optional. If one is not set, an ID will be randomly generated. (It is an error to specify the ID of a filter that already exists.)

}

// The result of adding a filter view.
type AddFilterViewResponse struct {
	Filter FilterView `json:"filter,omitempty\"` // The newly added filter view.

}

// Adds a named range to the spreadsheet.
type AddNamedRangeRequest struct {
	NamedRange NamedRange `json:"namedRange,omitempty\"` // The named range to add. The namedRangeId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of a range that already exists.)

}

// The result of adding a named range.
type AddNamedRangeResponse struct {
	NamedRange NamedRange `json:"namedRange,omitempty\"` // The named range to add.

}

// Adds a new protected range.
type AddProtectedRangeRequest struct {
	ProtectedRange ProtectedRange `json:"protectedRange,omitempty\"` // The protected range to be added. The protectedRangeId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of a range that already exists.)

}

// The result of adding a new protected range.
type AddProtectedRangeResponse struct {
	ProtectedRange ProtectedRange `json:"protectedRange,omitempty\"` // The newly added protected range.

}

// Adds a new sheet. When a sheet is added at a given index, all subsequent sheets' indexes are incremented. To add an object sheet, use AddChartRequest instead and specify EmbeddedObjectPosition.sheetId or EmbeddedObjectPosition.newSheet.
type AddSheetRequest struct {
	Properties SheetProperties `json:"properties,omitempty\"` // The properties the new sheet should have. All properties are optional. The sheetId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of a sheet that already exists.)

}

// The result of adding a sheet.
type AddSheetResponse struct {
	Properties SheetProperties `json:"properties,omitempty\"` // The properties of the newly added sheet.

}

// Adds a slicer to a sheet in the spreadsheet.
type AddSlicerRequest struct {
	Slicer Slicer `json:"slicer,omitempty\"` // The slicer that should be added to the spreadsheet, including the position where it should be placed. The slicerId field is optional; if one is not set, an id will be randomly generated. (It is an error to specify the ID of a slicer that already exists.)

}

// The result of adding a slicer to a spreadsheet.
type AddSlicerResponse struct {
	Slicer Slicer `json:"slicer,omitempty\"` // The newly added slicer.

}

// Adds a new table to the spreadsheet.
type AddTableRequest struct {
	Table Table `json:"table,omitempty\"` // Required. The table to add.

}

// The result of adding a table.
type AddTableResponse struct {
	Table Table `json:"table,omitempty\"` // Output only. The table that was added.

}

// Adds new cells after the last row with data in a sheet, inserting new rows into the sheet if necessary.
type AppendCellsRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields of CellData that should be updated. At least one field must be specified. The root is the CellData; 'row.values.' should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Rows []RowData `json:"rows,omitempty\"` // The data to append.

	SheetId int `json:"sheetId,omitempty\"` // The sheet ID to append the data to.

	TableId string `json:"tableId,omitempty\"` // The ID of the table to append data to. The data will be only appended to the table body. This field also takes precedence over the `sheet_id` field.

}

// Appends rows or columns to the end of a sheet.
type AppendDimensionRequest struct {
	Dimension string `json:"dimension,omitempty\"` // Whether rows or columns should be appended.

	Length int `json:"length,omitempty\"` // The number of rows or columns to append.

	SheetId int `json:"sheetId,omitempty\"` // The sheet to append rows or columns to.

}

// The response when updating a range of values in a spreadsheet.
type AppendValuesResponse struct {
	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

	TableRange string `json:"tableRange,omitempty\"` // The range (in A1 notation) of the table that values are being appended to (before the values were appended). Empty if no table was found.

	Updates UpdateValuesResponse `json:"updates,omitempty\"` // Information about the updates that were applied.

}

// Fills in more data based on existing data.
type AutoFillRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range to autofill. This will examine the range and detect the location that has data and automatically fill that data in to the rest of the range.

	SourceAndDestination SourceAndDestination `json:"sourceAndDestination,omitempty\"` // The source and destination areas to autofill. This explicitly lists the source of the autofill and where to extend that data.

	UseAlternateSeries bool `json:"useAlternateSeries,omitempty\"` // True if we should generate data with the "alternate" series. This differs based on the type and amount of source data.

}

// Automatically resizes one or more dimensions based on the contents of the cells in that dimension.
type AutoResizeDimensionsRequest struct {
	DataSourceSheetDimensions DataSourceSheetDimensionRange `json:"dataSourceSheetDimensions,omitempty\"` // The dimensions on a data source sheet to automatically resize.

	Dimensions DimensionRange `json:"dimensions,omitempty\"` // The dimensions to automatically resize.

}

// A banded (alternating colors) range in a sheet.
type BandedRange struct {
	BandedRangeId int `json:"bandedRangeId,omitempty\"` // The ID of the banded range. If unset, refer to banded_range_reference.

	BandedRangeReference string `json:"bandedRangeReference,omitempty\"` // Output only. The reference of the banded range, used to identify the ID that is not supported by the banded_range_id.

	ColumnProperties BandingProperties `json:"columnProperties,omitempty\"` // Properties for column bands. These properties are applied on a column- by-column basis throughout all the columns in the range. At least one of row_properties or column_properties must be specified.

	RangeValue GridRange `json:"range,omitempty\"` // The range over which these properties are applied.

	RowProperties BandingProperties `json:"rowProperties,omitempty\"` // Properties for row bands. These properties are applied on a row-by-row basis throughout all the rows in the range. At least one of row_properties or column_properties must be specified.

}

// Properties referring a single dimension (either row or column). If both BandedRange.row_properties and BandedRange.column_properties are set, the fill colors are applied to cells according to the following rules: * header_color and footer_color take priority over band colors. * first_band_color takes priority over second_band_color. * row_properties takes priority over column_properties. For example, the first row color takes priority over the first column color, but the first column color takes priority over the second row color. Similarly, the row header takes priority over the column header in the top left cell, but the column header takes priority over the first row color if the row header is not set.
type BandingProperties struct {
	FirstBandColor Color `json:"firstBandColor,omitempty\"` // The first color that is alternating. (Required) Deprecated: Use first_band_color_style.

	FirstBandColorStyle ColorStyle `json:"firstBandColorStyle,omitempty\"` // The first color that is alternating. (Required) If first_band_color is also set, this field takes precedence.

	FooterColor Color `json:"footerColor,omitempty\"` // The color of the last row or column. If this field is not set, the last row or column is filled with either first_band_color or second_band_color, depending on the color of the previous row or column. Deprecated: Use footer_color_style.

	FooterColorStyle ColorStyle `json:"footerColorStyle,omitempty\"` // The color of the last row or column. If this field is not set, the last row or column is filled with either first_band_color or second_band_color, depending on the color of the previous row or column. If footer_color is also set, this field takes precedence.

	HeaderColor Color `json:"headerColor,omitempty\"` // The color of the first row or column. If this field is set, the first row or column is filled with this color and the colors alternate between first_band_color and second_band_color starting from the second row or column. Otherwise, the first row or column is filled with first_band_color and the colors proceed to alternate as they normally would. Deprecated: Use header_color_style.

	HeaderColorStyle ColorStyle `json:"headerColorStyle,omitempty\"` // The color of the first row or column. If this field is set, the first row or column is filled with this color and the colors alternate between first_band_color and second_band_color starting from the second row or column. Otherwise, the first row or column is filled with first_band_color and the colors proceed to alternate as they normally would. If header_color is also set, this field takes precedence.

	SecondBandColor Color `json:"secondBandColor,omitempty\"` // The second color that is alternating. (Required) Deprecated: Use second_band_color_style.

	SecondBandColorStyle ColorStyle `json:"secondBandColorStyle,omitempty\"` // The second color that is alternating. (Required) If second_band_color is also set, this field takes precedence.

}

// Formatting options for baseline value.
type BaselineValueFormat struct {
	ComparisonType string `json:"comparisonType,omitempty\"` // The comparison type of key value with baseline value.

	Description string `json:"description,omitempty\"` // Description which is appended after the baseline value. This field is optional.

	NegativeColor Color `json:"negativeColor,omitempty\"` // Color to be used, in case baseline value represents a negative change for key value. This field is optional. Deprecated: Use negative_color_style.

	NegativeColorStyle ColorStyle `json:"negativeColorStyle,omitempty\"` // Color to be used, in case baseline value represents a negative change for key value. This field is optional. If negative_color is also set, this field takes precedence.

	Position TextPosition `json:"position,omitempty\"` // Specifies the horizontal text positioning of baseline value. This field is optional. If not specified, default positioning is used.

	PositiveColor Color `json:"positiveColor,omitempty\"` // Color to be used, in case baseline value represents a positive change for key value. This field is optional. Deprecated: Use positive_color_style.

	PositiveColorStyle ColorStyle `json:"positiveColorStyle,omitempty\"` // Color to be used, in case baseline value represents a positive change for key value. This field is optional. If positive_color is also set, this field takes precedence.

	TextFormat TextFormat `json:"textFormat,omitempty\"` // Text formatting options for baseline value. The link field is not supported.

}

// An axis of the chart. A chart may not have more than one axis per axis position.
type BasicChartAxis struct {
	Format TextFormat `json:"format,omitempty\"` // The format of the title. Only valid if the axis is not associated with the domain. The link field is not supported.

	Position string `json:"position,omitempty\"` // The position of this axis.

	Title string `json:"title,omitempty\"` // The title of this axis. If set, this overrides any title inferred from headers of the data.

	TitleTextPosition TextPosition `json:"titleTextPosition,omitempty\"` // The axis title text position.

	ViewWindowOptions ChartAxisViewWindowOptions `json:"viewWindowOptions,omitempty\"` // The view window options for this axis.

}

// The domain of a chart. For example, if charting stock prices over time, this would be the date.
type BasicChartDomain struct {
	Domain ChartData `json:"domain,omitempty\"` // The data of the domain. For example, if charting stock prices over time, this is the data representing the dates.

	Reversed bool `json:"reversed,omitempty\"` // True to reverse the order of the domain values (horizontal axis).

}

// A single series of data in a chart. For example, if charting stock prices over time, multiple series may exist, one for the "Open Price", "High Price", "Low Price" and "Close Price".
type BasicChartSeries struct {
	Color Color `json:"color,omitempty\"` // The color for elements (such as bars, lines, and points) associated with this series. If empty, a default color is used. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // The color for elements (such as bars, lines, and points) associated with this series. If empty, a default color is used. If color is also set, this field takes precedence.

	DataLabel DataLabel `json:"dataLabel,omitempty\"` // Information about the data labels for this series.

	LineStyle LineStyle `json:"lineStyle,omitempty\"` // The line style of this series. Valid only if the chartType is AREA, LINE, or SCATTER. COMBO charts are also supported if the series chart type is AREA or LINE.

	PointStyle PointStyle `json:"pointStyle,omitempty\"` // The style for points associated with this series. Valid only if the chartType is AREA, LINE, or SCATTER. COMBO charts are also supported if the series chart type is AREA, LINE, or SCATTER. If empty, a default point style is used.

	Series ChartData `json:"series,omitempty\"` // The data being visualized in this chart series.

	StyleOverrides []BasicSeriesDataPointStyleOverride `json:"styleOverrides,omitempty\"` // Style override settings for series data points.

	TargetAxis string `json:"targetAxis,omitempty\"` // The minor axis that will specify the range of values for this series. For example, if charting stocks over time, the "Volume" series may want to be pinned to the right with the prices pinned to the left, because the scale of trading volume is different than the scale of prices. It is an error to specify an axis that isn't a valid minor axis for the chart's type.

	TypeValue string `json:"type,omitempty\"` // The type of this series. Valid only if the chartType is COMBO. Different types will change the way the series is visualized. Only LINE, AREA, and COLUMN are supported.

}

// The specification for a basic chart. See BasicChartType for the list of charts this supports.
type BasicChartSpec struct {
	Axis []BasicChartAxis `json:"axis,omitempty\"` // The axis on the chart.

	ChartType string `json:"chartType,omitempty\"` // The type of the chart.

	CompareMode string `json:"compareMode,omitempty\"` // The behavior of tooltips and data highlighting when hovering on data and chart area.

	Domains []BasicChartDomain `json:"domains,omitempty\"` // The domain of data this is charting. Only a single domain is supported.

	HeaderCount int `json:"headerCount,omitempty\"` // The number of rows or columns in the data that are "headers". If not set, Google Sheets will guess how many rows are headers based on the data. (Note that BasicChartAxis.title may override the axis title inferred from the header values.)

	InterpolateNulls bool `json:"interpolateNulls,omitempty\"` // If some values in a series are missing, gaps may appear in the chart (e.g, segments of lines in a line chart will be missing). To eliminate these gaps set this to true. Applies to Line, Area, and Combo charts.

	LegendPosition string `json:"legendPosition,omitempty\"` // The position of the chart legend.

	LineSmoothing bool `json:"lineSmoothing,omitempty\"` // Gets whether all lines should be rendered smooth or straight by default. Applies to Line charts.

	Series []BasicChartSeries `json:"series,omitempty\"` // The data this chart is visualizing.

	StackedType string `json:"stackedType,omitempty\"` // The stacked type for charts that support vertical stacking. Applies to Area, Bar, Column, Combo, and Stepped Area charts.

	ThreeDimensional bool `json:"threeDimensional,omitempty\"` // True to make the chart 3D. Applies to Bar and Column charts.

	TotalDataLabel DataLabel `json:"totalDataLabel,omitempty\"` // Controls whether to display additional data labels on stacked charts which sum the total value of all stacked values at each value along the domain axis. These data labels can only be set when chart_type is one of AREA, BAR, COLUMN, COMBO or STEPPED_AREA and stacked_type is either STACKED or PERCENT_STACKED. In addition, for COMBO, this will only be supported if there is only one type of stackable series type or one type has more series than the others and each of the other types have no more than one series. For example, if a chart has two stacked bar series and one area series, the total data labels will be supported. If it has three bar series and two area series, total data labels are not allowed. Neither CUSTOM nor placement can be set on the total_data_label.

}

// The default filter associated with a sheet. For more information, see [Manage data visibility with filters](https://developers.google.com/workspace/sheets/api/guides/filters).
type BasicFilter struct {
	Criteria map[string]interface{} `json:"criteria,omitempty\"` // The criteria for showing/hiding values per column. The map's key is the column index, and the value is the criteria for that column. This field is deprecated in favor of filter_specs.

	FilterSpecs []FilterSpec `json:"filterSpecs,omitempty\"` // The filter criteria per column. Both criteria and filter_specs are populated in responses. If both fields are specified in an update request, this field takes precedence.

	RangeValue GridRange `json:"range,omitempty\"` // The range the filter covers.

	SortSpecs []SortSpec `json:"sortSpecs,omitempty\"` // The sort order per column. Later specifications are used when values are equal in the earlier specifications.

	TableId string `json:"tableId,omitempty\"` // The table this filter is backed by, if any. When writing, only one of range or table_id may be set.

}

// Style override settings for a single series data point.
type BasicSeriesDataPointStyleOverride struct {
	Color Color `json:"color,omitempty\"` // Color of the series data point. If empty, the series default is used. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // Color of the series data point. If empty, the series default is used. If color is also set, this field takes precedence.

	Index int `json:"index,omitempty\"` // The zero-based index of the series data point.

	PointStyle PointStyle `json:"pointStyle,omitempty\"` // Point style of the series data point. Valid only if the chartType is AREA, LINE, or SCATTER. COMBO charts are also supported if the series chart type is AREA, LINE, or SCATTER. If empty, the series default is used.

}

// The request for clearing more than one range selected by a DataFilter in a spreadsheet.
type BatchClearValuesByDataFilterRequest struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The DataFilters used to determine which ranges to clear.

}

// The response when clearing a range of values selected with DataFilters in a spreadsheet.
type BatchClearValuesByDataFilterResponse struct {
	ClearedRanges []string `json:"clearedRanges,omitempty\"` // The ranges that were cleared, in [A1 notation](https://developers.google.com/workspace/sheets/api/guides/concepts#cell). If the requests are for an unbounded range or a range larger than the bounds of the sheet, this is the actual ranges that were cleared, bounded to the sheet's limits.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

}

// The request for clearing more than one range of values in a spreadsheet.
type BatchClearValuesRequest struct {
	Ranges []string `json:"ranges,omitempty\"` // The ranges to clear, in [A1 notation or R1C1 notation](https://developers.google.com/workspace/sheets/api/guides/concepts#cell).

}

// The response when clearing a range of values in a spreadsheet.
type BatchClearValuesResponse struct {
	ClearedRanges []string `json:"clearedRanges,omitempty\"` // The ranges that were cleared, in A1 notation. If the requests are for an unbounded range or a range larger than the bounds of the sheet, this is the actual ranges that were cleared, bounded to the sheet's limits.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

}

// The request for retrieving a range of values in a spreadsheet selected by a set of DataFilters.
type BatchGetValuesByDataFilterRequest struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The data filters used to match the ranges of values to retrieve. Ranges that match any of the specified data filters are included in the response.

	DateTimeRenderOption string `json:"dateTimeRenderOption,omitempty\"` // How dates, times, and durations should be represented in the output. This is ignored if value_render_option is FORMATTED_VALUE. The default dateTime render option is SERIAL_NUMBER.

	MajorDimension string `json:"majorDimension,omitempty\"` // The major dimension that results should use. For example, if the spreadsheet data is: `A1=1,B1=2,A2=3,B2=4`, then a request that selects that range and sets `majorDimension=ROWS` returns `[[1,2],[3,4]]`, whereas a request that sets `majorDimension=COLUMNS` returns `[[1,3],[2,4]]`.

	ValueRenderOption string `json:"valueRenderOption,omitempty\"` // How values should be represented in the output. The default render option is FORMATTED_VALUE.

}

// The response when retrieving more than one range of values in a spreadsheet selected by DataFilters.
type BatchGetValuesByDataFilterResponse struct {
	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the spreadsheet the data was retrieved from.

	ValueRanges []MatchedValueRange `json:"valueRanges,omitempty\"` // The requested values with the list of data filters that matched them.

}

// The response when retrieving more than one range of values in a spreadsheet.
type BatchGetValuesResponse struct {
	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the spreadsheet the data was retrieved from.

	ValueRanges []ValueRange `json:"valueRanges,omitempty\"` // The requested values. The order of the ValueRanges is the same as the order of the requested ranges.

}

// The request for updating any aspect of a spreadsheet.
type BatchUpdateSpreadsheetRequest struct {
	IncludeSpreadsheetInResponse bool `json:"includeSpreadsheetInResponse,omitempty\"` // Determines if the update response should include the spreadsheet resource.

	Requests []Request `json:"requests,omitempty\"` // A list of updates to apply to the spreadsheet. Requests will be applied in the order they are specified. If any request is not valid, no requests will be applied.

	ResponseIncludeGridData bool `json:"responseIncludeGridData,omitempty\"` // True if grid data should be returned. Meaningful only if include_spreadsheet_in_response is 'true'. This parameter is ignored if a field mask was set in the request.

	ResponseRanges []string `json:"responseRanges,omitempty\"` // Limits the ranges included in the response spreadsheet. Meaningful only if include_spreadsheet_in_response is 'true'.

}

// The reply for batch updating a spreadsheet.
type BatchUpdateSpreadsheetResponse struct {
	Replies []Response `json:"replies,omitempty\"` // The reply of the updates. This maps 1:1 with the updates, although replies to some requests may be empty.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

	UpdatedSpreadsheet Spreadsheet `json:"updatedSpreadsheet,omitempty\"` // The spreadsheet after updates were applied. This is only set if BatchUpdateSpreadsheetRequest.include_spreadsheet_in_response is `true`.

}

// The request for updating more than one range of values in a spreadsheet.
type BatchUpdateValuesByDataFilterRequest struct {
	Data []DataFilterValueRange `json:"data,omitempty\"` // The new values to apply to the spreadsheet. If more than one range is matched by the specified DataFilter the specified values are applied to all of those ranges.

	IncludeValuesInResponse bool `json:"includeValuesInResponse,omitempty\"` // Determines if the update response should include the values of the cells that were updated. By default, responses do not include the updated values. The `updatedData` field within each of the BatchUpdateValuesResponse.responses contains the updated values. If the range to write was larger than the range actually written, the response includes all values in the requested range (excluding trailing empty rows and columns).

	ResponseDateTimeRenderOption string `json:"responseDateTimeRenderOption,omitempty\"` // Determines how dates, times, and durations in the response should be rendered. This is ignored if response_value_render_option is FORMATTED_VALUE. The default dateTime render option is SERIAL_NUMBER.

	ResponseValueRenderOption string `json:"responseValueRenderOption,omitempty\"` // Determines how values in the response should be rendered. The default render option is FORMATTED_VALUE.

	ValueInputOption string `json:"valueInputOption,omitempty\"` // How the input data should be interpreted.

}

// The response when updating a range of values in a spreadsheet.
type BatchUpdateValuesByDataFilterResponse struct {
	Responses []UpdateValuesByDataFilterResponse `json:"responses,omitempty\"` // The response for each range updated.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

	TotalUpdatedCells int `json:"totalUpdatedCells,omitempty\"` // The total number of cells updated.

	TotalUpdatedColumns int `json:"totalUpdatedColumns,omitempty\"` // The total number of columns where at least one cell in the column was updated.

	TotalUpdatedRows int `json:"totalUpdatedRows,omitempty\"` // The total number of rows where at least one cell in the row was updated.

	TotalUpdatedSheets int `json:"totalUpdatedSheets,omitempty\"` // The total number of sheets where at least one cell in the sheet was updated.

}

// The request for updating more than one range of values in a spreadsheet.
type BatchUpdateValuesRequest struct {
	Data []ValueRange `json:"data,omitempty\"` // The new values to apply to the spreadsheet.

	IncludeValuesInResponse bool `json:"includeValuesInResponse,omitempty\"` // Determines if the update response should include the values of the cells that were updated. By default, responses do not include the updated values. The `updatedData` field within each of the BatchUpdateValuesResponse.responses contains the updated values. If the range to write was larger than the range actually written, the response includes all values in the requested range (excluding trailing empty rows and columns).

	ResponseDateTimeRenderOption string `json:"responseDateTimeRenderOption,omitempty\"` // Determines how dates, times, and durations in the response should be rendered. This is ignored if response_value_render_option is FORMATTED_VALUE. The default dateTime render option is SERIAL_NUMBER.

	ResponseValueRenderOption string `json:"responseValueRenderOption,omitempty\"` // Determines how values in the response should be rendered. The default render option is FORMATTED_VALUE.

	ValueInputOption string `json:"valueInputOption,omitempty\"` // How the input data should be interpreted.

}

// The response when updating a range of values in a spreadsheet.
type BatchUpdateValuesResponse struct {
	Responses []UpdateValuesResponse `json:"responses,omitempty\"` // One UpdateValuesResponse per requested range, in the same order as the requests appeared.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

	TotalUpdatedCells int `json:"totalUpdatedCells,omitempty\"` // The total number of cells updated.

	TotalUpdatedColumns int `json:"totalUpdatedColumns,omitempty\"` // The total number of columns where at least one cell in the column was updated.

	TotalUpdatedRows int `json:"totalUpdatedRows,omitempty\"` // The total number of rows where at least one cell in the row was updated.

	TotalUpdatedSheets int `json:"totalUpdatedSheets,omitempty\"` // The total number of sheets where at least one cell in the sheet was updated.

}

// The specification of a BigQuery data source that's connected to a sheet.
type BigQueryDataSourceSpec struct {
	ProjectId string `json:"projectId,omitempty\"` // The ID of a BigQuery enabled Google Cloud project with a billing account attached. For any queries executed against the data source, the project is charged.

	QuerySpec BigQueryQuerySpec `json:"querySpec,omitempty\"` // A BigQueryQuerySpec.

	TableSpec BigQueryTableSpec `json:"tableSpec,omitempty\"` // A BigQueryTableSpec.

}

// Specifies a custom BigQuery query.
type BigQueryQuerySpec struct {
	RawQuery string `json:"rawQuery,omitempty\"` // The raw query string.

}

// Specifies a BigQuery table definition. Only [native tables](https://cloud.google.com/bigquery/docs/tables-intro) are allowed.
type BigQueryTableSpec struct {
	DatasetId string `json:"datasetId,omitempty\"` // The BigQuery dataset id.

	TableId string `json:"tableId,omitempty\"` // The BigQuery table id.

	TableProjectId string `json:"tableProjectId,omitempty\"` // The ID of a BigQuery project the table belongs to. If not specified, the project_id is assumed.

}

// A condition that can evaluate to true or false. BooleanConditions are used by conditional formatting, data validation, and the criteria in filters.
type BooleanCondition struct {
	TypeValue string `json:"type,omitempty\"` // The type of condition.

	Values []ConditionValue `json:"values,omitempty\"` // The values of the condition. The number of supported values depends on the condition type. Some support zero values, others one or two values, and ConditionType.ONE_OF_LIST supports an arbitrary number of values.

}

// A rule that may or may not match, depending on the condition.
type BooleanRule struct {
	Condition BooleanCondition `json:"condition,omitempty\"` // The condition of the rule. If the condition evaluates to true, the format is applied.

	Format CellFormat `json:"format,omitempty\"` // The format to apply. Conditional formatting can only apply a subset of formatting: bold, italic, strikethrough, foreground color and, background color.

}

// A border along a cell.
type Border struct {
	Color Color `json:"color,omitempty\"` // The color of the border. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // The color of the border. If color is also set, this field takes precedence.

	Style string `json:"style,omitempty\"` // The style of the border.

	Width int `json:"width,omitempty\"` // The width of the border, in pixels. Deprecated; the width is determined by the "style" field.

}

// The borders of the cell.
type Borders struct {
	Bottom Border `json:"bottom,omitempty\"` // The bottom border of the cell.

	Left Border `json:"left,omitempty\"` // The left border of the cell.

	Right Border `json:"right,omitempty\"` // The right border of the cell.

	Top Border `json:"top,omitempty\"` // The top border of the cell.

}

// A bubble chart.
type BubbleChartSpec struct {
	BubbleBorderColor Color `json:"bubbleBorderColor,omitempty\"` // The bubble border color. Deprecated: Use bubble_border_color_style.

	BubbleBorderColorStyle ColorStyle `json:"bubbleBorderColorStyle,omitempty\"` // The bubble border color. If bubble_border_color is also set, this field takes precedence.

	BubbleLabels ChartData `json:"bubbleLabels,omitempty\"` // The data containing the bubble labels. These do not need to be unique.

	BubbleMaxRadiusSize int `json:"bubbleMaxRadiusSize,omitempty\"` // The max radius size of the bubbles, in pixels. If specified, the field must be a positive value.

	BubbleMinRadiusSize int `json:"bubbleMinRadiusSize,omitempty\"` // The minimum radius size of the bubbles, in pixels. If specific, the field must be a positive value.

	BubbleOpacity float64 `json:"bubbleOpacity,omitempty\"` // The opacity of the bubbles between 0 and 1.0. 0 is fully transparent and 1 is fully opaque.

	BubbleSizes ChartData `json:"bubbleSizes,omitempty\"` // The data containing the bubble sizes. Bubble sizes are used to draw the bubbles at different sizes relative to each other. If specified, group_ids must also be specified. This field is optional.

	BubbleTextStyle TextFormat `json:"bubbleTextStyle,omitempty\"` // The format of the text inside the bubbles. Strikethrough, underline, and link are not supported.

	Domain ChartData `json:"domain,omitempty\"` // The data containing the bubble x-values. These values locate the bubbles in the chart horizontally.

	GroupIds ChartData `json:"groupIds,omitempty\"` // The data containing the bubble group IDs. All bubbles with the same group ID are drawn in the same color. If bubble_sizes is specified then this field must also be specified but may contain blank values. This field is optional.

	LegendPosition string `json:"legendPosition,omitempty\"` // Where the legend of the chart should be drawn.

	Series ChartData `json:"series,omitempty\"` // The data containing the bubble y-values. These values locate the bubbles in the chart vertically.

}

// Cancels one or multiple refreshes of data source objects in the spreadsheet by the specified references. The request requires an additional `bigquery.readonly` OAuth scope if you are cancelling a refresh on a BigQuery data source.
type CancelDataSourceRefreshRequest struct {
	DataSourceId string `json:"dataSourceId,omitempty\"` // Reference to a DataSource. If specified, cancels all associated data source object refreshes for this data source.

	IsAll bool `json:"isAll,omitempty\"` // Cancels all existing data source object refreshes for all data sources in the spreadsheet.

	References DataSourceObjectReferences `json:"references,omitempty\"` // References to data source objects whose refreshes are to be cancelled.

}

// The response from cancelling one or multiple data source object refreshes.
type CancelDataSourceRefreshResponse struct {
	Statuses []CancelDataSourceRefreshStatus `json:"statuses,omitempty\"` // The cancellation statuses of refreshes of all data source objects specified in the request. If is_all is specified, the field contains only those in failure status. Refreshing and canceling refresh the same data source object is also not allowed in the same `batchUpdate`.

}

// The status of cancelling a single data source object refresh.
type CancelDataSourceRefreshStatus struct {
	Reference DataSourceObjectReference `json:"reference,omitempty\"` // Reference to the data source object whose refresh is being cancelled.

	RefreshCancellationStatus RefreshCancellationStatus `json:"refreshCancellationStatus,omitempty\"` // The cancellation status.

}

// A candlestick chart.
type CandlestickChartSpec struct {
	Data []CandlestickData `json:"data,omitempty\"` // The Candlestick chart data. Only one CandlestickData is supported.

	Domain CandlestickDomain `json:"domain,omitempty\"` // The domain data (horizontal axis) for the candlestick chart. String data will be treated as discrete labels, other data will be treated as continuous values.

}

// The Candlestick chart data, each containing the low, open, close, and high values for a series.
type CandlestickData struct {
	CloseSeries CandlestickSeries `json:"closeSeries,omitempty\"` // The range data (vertical axis) for the close/final value for each candle. This is the top of the candle body. If greater than the open value the candle will be filled. Otherwise the candle will be hollow.

	HighSeries CandlestickSeries `json:"highSeries,omitempty\"` // The range data (vertical axis) for the high/maximum value for each candle. This is the top of the candle's center line.

	LowSeries CandlestickSeries `json:"lowSeries,omitempty\"` // The range data (vertical axis) for the low/minimum value for each candle. This is the bottom of the candle's center line.

	OpenSeries CandlestickSeries `json:"openSeries,omitempty\"` // The range data (vertical axis) for the open/initial value for each candle. This is the bottom of the candle body. If less than the close value the candle will be filled. Otherwise the candle will be hollow.

}

// The domain of a CandlestickChart.
type CandlestickDomain struct {
	Data ChartData `json:"data,omitempty\"` // The data of the CandlestickDomain.

	Reversed bool `json:"reversed,omitempty\"` // True to reverse the order of the domain values (horizontal axis).

}

// The series of a CandlestickData.
type CandlestickSeries struct {
	Data ChartData `json:"data,omitempty\"` // The data of the CandlestickSeries.

}

// Data about a specific cell.
type CellData struct {
	ChipRuns []ChipRun `json:"chipRuns,omitempty\"` // Optional. Runs of chips applied to subsections of the cell. Properties of a run start at a specific index in the text and continue until the next run. When reading, all chipped and non-chipped runs are included. Non-chipped runs will have an empty Chip. When writing, only runs with chips are included. Runs containing chips are of length 1 and are represented in the user-entered text by an “@” placeholder symbol. New runs will overwrite any prior runs. Writing a new user_entered_value will erase previous runs.

	DataSourceFormula DataSourceFormula `json:"dataSourceFormula,omitempty\"` // Output only. Information about a data source formula on the cell. The field is set if user_entered_value is a formula referencing some DATA_SOURCE sheet, e.g. `=SUM(DataSheet!Column)`.

	DataSourceTable DataSourceTable `json:"dataSourceTable,omitempty\"` // A data source table anchored at this cell. The size of data source table itself is computed dynamically based on its configuration. Only the first cell of the data source table contains the data source table definition. The other cells will contain the display values of the data source table result in their effective_value fields.

	DataValidation DataValidationRule `json:"dataValidation,omitempty\"` // A data validation rule on the cell, if any. When writing, the new data validation rule will overwrite any prior rule.

	EffectiveFormat CellFormat `json:"effectiveFormat,omitempty\"` // The effective format being used by the cell. This includes the results of applying any conditional formatting and, if the cell contains a formula, the computed number format. If the effective format is the default format, effective format will not be written. This field is read-only.

	EffectiveValue ExtendedValue `json:"effectiveValue,omitempty\"` // The effective value of the cell. For cells with formulas, this is the calculated value. For cells with literals, this is the same as the user_entered_value. This field is read-only.

	FormattedValue string `json:"formattedValue,omitempty\"` // The formatted value of the cell. This is the value as it's shown to the user. This field is read-only.

	Hyperlink string `json:"hyperlink,omitempty\"` // A hyperlink this cell points to, if any. If the cell contains multiple hyperlinks, this field will be empty. This field is read-only. To set it, use a `=HYPERLINK` formula in the userEnteredValue.formulaValue field. A cell-level link can also be set from the userEnteredFormat.textFormat field. Alternatively, set a hyperlink in the textFormatRun.format.link field that spans the entire cell.

	Note string `json:"note,omitempty\"` // Any note on the cell.

	PivotTable PivotTable `json:"pivotTable,omitempty\"` // A pivot table anchored at this cell. The size of pivot table itself is computed dynamically based on its data, grouping, filters, values, etc. Only the top-left cell of the pivot table contains the pivot table definition. The other cells will contain the calculated values of the results of the pivot in their effective_value fields.

	TextFormatRuns []TextFormatRun `json:"textFormatRuns,omitempty\"` // Runs of rich text applied to subsections of the cell. Runs are only valid on user entered strings, not formulas, bools, or numbers. Properties of a run start at a specific index in the text and continue until the next run. Runs will inherit the properties of the cell unless explicitly changed. When writing, the new runs will overwrite any prior runs. When writing a new user_entered_value, previous runs are erased.

	UserEnteredFormat CellFormat `json:"userEnteredFormat,omitempty\"` // The format the user entered for the cell. When writing, the new format will be merged with the existing format.

	UserEnteredValue ExtendedValue `json:"userEnteredValue,omitempty\"` // The value the user entered in the cell. e.g., `1234`, `'Hello'`, or `=NOW()` Note: Dates, Times and DateTimes are represented as doubles in serial number format.

}

// The format of a cell.
type CellFormat struct {
	BackgroundColor Color `json:"backgroundColor,omitempty\"` // The background color of the cell. Deprecated: Use background_color_style.

	BackgroundColorStyle ColorStyle `json:"backgroundColorStyle,omitempty\"` // The background color of the cell. If background_color is also set, this field takes precedence.

	Borders Borders `json:"borders,omitempty\"` // The borders of the cell.

	HorizontalAlignment string `json:"horizontalAlignment,omitempty\"` // The horizontal alignment of the value in the cell.

	HyperlinkDisplayType string `json:"hyperlinkDisplayType,omitempty\"` // If one exists, how a hyperlink should be displayed in the cell.

	NumberFormat NumberFormat `json:"numberFormat,omitempty\"` // A format describing how number values should be represented to the user.

	Padding Padding `json:"padding,omitempty\"` // The padding of the cell.

	TextDirection string `json:"textDirection,omitempty\"` // The direction of the text in the cell.

	TextFormat TextFormat `json:"textFormat,omitempty\"` // The format of the text in the cell (unless overridden by a format run). Setting a cell-level link here clears the cell's existing links. Setting the link field in a TextFormatRun takes precedence over the cell-level link.

	TextRotation TextRotation `json:"textRotation,omitempty\"` // The rotation applied to text in the cell.

	VerticalAlignment string `json:"verticalAlignment,omitempty\"` // The vertical alignment of the value in the cell.

	WrapStrategy string `json:"wrapStrategy,omitempty\"` // The wrap strategy for the value in the cell.

}

// The options that define a "view window" for a chart (such as the visible values in an axis).
type ChartAxisViewWindowOptions struct {
	ViewWindowMax float64 `json:"viewWindowMax,omitempty\"` // The maximum numeric value to be shown in this view window. If unset, will automatically determine a maximum value that looks good for the data.

	ViewWindowMin float64 `json:"viewWindowMin,omitempty\"` // The minimum numeric value to be shown in this view window. If unset, will automatically determine a minimum value that looks good for the data.

	ViewWindowMode string `json:"viewWindowMode,omitempty\"` // The view window's mode.

}

// Custom number formatting options for chart attributes.
type ChartCustomNumberFormatOptions struct {
	Prefix string `json:"prefix,omitempty\"` // Custom prefix to be prepended to the chart attribute. This field is optional.

	Suffix string `json:"suffix,omitempty\"` // Custom suffix to be appended to the chart attribute. This field is optional.

}

// The data included in a domain or series.
type ChartData struct {
	AggregateType string `json:"aggregateType,omitempty\"` // The aggregation type for the series of a data source chart. Only supported for data source charts.

	ColumnReference DataSourceColumnReference `json:"columnReference,omitempty\"` // The reference to the data source column that the data reads from.

	GroupRule ChartGroupRule `json:"groupRule,omitempty\"` // The rule to group the data by if the ChartData backs the domain of a data source chart. Only supported for data source charts.

	SourceRange ChartSourceRange `json:"sourceRange,omitempty\"` // The source ranges of the data.

}

// Allows you to organize the date-time values in a source data column into buckets based on selected parts of their date or time values.
type ChartDateTimeRule struct {
	TypeValue string `json:"type,omitempty\"` // The type of date-time grouping to apply.

}

// An optional setting on the ChartData of the domain of a data source chart that defines buckets for the values in the domain rather than breaking out each individual value. For example, when plotting a data source chart, you can specify a histogram rule on the domain (it should only contain numeric values), grouping its values into buckets. Any values of a chart series that fall into the same bucket are aggregated based on the aggregate_type.
type ChartGroupRule struct {
	DateTimeRule ChartDateTimeRule `json:"dateTimeRule,omitempty\"` // A ChartDateTimeRule.

	HistogramRule ChartHistogramRule `json:"histogramRule,omitempty\"` // A ChartHistogramRule

}

// Allows you to organize numeric values in a source data column into buckets of constant size.
type ChartHistogramRule struct {
	IntervalSize float64 `json:"intervalSize,omitempty\"` // The size of the buckets that are created. Must be positive.

	MaxValue float64 `json:"maxValue,omitempty\"` // The maximum value at which items are placed into buckets. Values greater than the maximum are grouped into a single bucket. If omitted, it is determined by the maximum item value.

	MinValue float64 `json:"minValue,omitempty\"` // The minimum value at which items are placed into buckets. Values that are less than the minimum are grouped into a single bucket. If omitted, it is determined by the minimum item value.

}

// Source ranges for a chart.
type ChartSourceRange struct {
	Sources []GridRange `json:"sources,omitempty\"` // The ranges of data for a series or domain. Exactly one dimension must have a length of 1, and all sources in the list must have the same dimension with length 1. The domain (if it exists) & all series must have the same number of source ranges. If using more than one source range, then the source range at a given offset must be in order and contiguous across the domain and series. For example, these are valid configurations: domain sources: A1:A5 series1 sources: B1:B5 series2 sources: D6:D10 domain sources: A1:A5, C10:C12 series1 sources: B1:B5, D10:D12 series2 sources: C1:C5, E10:E12

}

// The specifications of a chart.
type ChartSpec struct {
	AltText string `json:"altText,omitempty\"` // The alternative text that describes the chart. This is often used for accessibility.

	BackgroundColor Color `json:"backgroundColor,omitempty\"` // The background color of the entire chart. Not applicable to Org charts. Deprecated: Use background_color_style.

	BackgroundColorStyle ColorStyle `json:"backgroundColorStyle,omitempty\"` // The background color of the entire chart. Not applicable to Org charts. If background_color is also set, this field takes precedence.

	BasicChart BasicChartSpec `json:"basicChart,omitempty\"` // A basic chart specification, can be one of many kinds of charts. See BasicChartType for the list of all charts this supports.

	BubbleChart BubbleChartSpec `json:"bubbleChart,omitempty\"` // A bubble chart specification.

	CandlestickChart CandlestickChartSpec `json:"candlestickChart,omitempty\"` // A candlestick chart specification.

	DataSourceChartProperties DataSourceChartProperties `json:"dataSourceChartProperties,omitempty\"` // If present, the field contains data source chart specific properties.

	FilterSpecs []FilterSpec `json:"filterSpecs,omitempty\"` // The filters applied to the source data of the chart. Only supported for data source charts.

	FontName string `json:"fontName,omitempty\"` // The name of the font to use by default for all chart text (e.g. title, axis labels, legend). If a font is specified for a specific part of the chart it will override this font name.

	HiddenDimensionStrategy string `json:"hiddenDimensionStrategy,omitempty\"` // Determines how the charts will use hidden rows or columns.

	HistogramChart HistogramChartSpec `json:"histogramChart,omitempty\"` // A histogram chart specification.

	Maximized bool `json:"maximized,omitempty\"` // True to make a chart fill the entire space in which it's rendered with minimum padding. False to use the default padding. (Not applicable to Geo and Org charts.)

	OrgChart OrgChartSpec `json:"orgChart,omitempty\"` // An org chart specification.

	PieChart PieChartSpec `json:"pieChart,omitempty\"` // A pie chart specification.

	ScorecardChart ScorecardChartSpec `json:"scorecardChart,omitempty\"` // A scorecard chart specification.

	SortSpecs []SortSpec `json:"sortSpecs,omitempty\"` // The order to sort the chart data by. Only a single sort spec is supported. Only supported for data source charts.

	Subtitle string `json:"subtitle,omitempty\"` // The subtitle of the chart.

	SubtitleTextFormat TextFormat `json:"subtitleTextFormat,omitempty\"` // The subtitle text format. Strikethrough, underline, and link are not supported.

	SubtitleTextPosition TextPosition `json:"subtitleTextPosition,omitempty\"` // The subtitle text position. This field is optional.

	Title string `json:"title,omitempty\"` // The title of the chart.

	TitleTextFormat TextFormat `json:"titleTextFormat,omitempty\"` // The title text format. Strikethrough, underline, and link are not supported.

	TitleTextPosition TextPosition `json:"titleTextPosition,omitempty\"` // The title text position. This field is optional.

	TreemapChart TreemapChartSpec `json:"treemapChart,omitempty\"` // A treemap chart specification.

	WaterfallChart WaterfallChartSpec `json:"waterfallChart,omitempty\"` // A waterfall chart specification.

}

// The Smart Chip.
type Chip struct {
	PersonProperties PersonProperties `json:"personProperties,omitempty\"` // Properties of a linked person.

	RichLinkProperties RichLinkProperties `json:"richLinkProperties,omitempty\"` // Properties of a rich link.

}

// The run of a chip. The chip continues until the start index of the next run.
type ChipRun struct {
	Chip Chip `json:"chip,omitempty\"` // Optional. The chip of this run.

	StartIndex int `json:"startIndex,omitempty\"` // Required. The zero-based character index where this run starts, in UTF-16 code units.

}

// Clears the basic filter, if any exists on the sheet.
type ClearBasicFilterRequest struct {
	SheetId int `json:"sheetId,omitempty\"` // The sheet ID on which the basic filter should be cleared.

}

// The request for clearing a range of values in a spreadsheet.
type ClearValuesRequest struct {
}

// The response when clearing a range of values in a spreadsheet.
type ClearValuesResponse struct {
	ClearedRange string `json:"clearedRange,omitempty\"` // The range (in A1 notation) that was cleared. (If the request was for an unbounded range or a range larger than the bounds of the sheet, this will be the actual range that was cleared, bounded to the sheet's limits.)

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

}

// Represents a color in the RGBA color space. This representation is designed for simplicity of conversion to and from color representations in various languages over compactness. For example, the fields of this representation can be trivially provided to the constructor of `java.awt.Color` in Java; it can also be trivially provided to UIColor's `+colorWithRed:green:blue:alpha` method in iOS; and, with just a little work, it can be easily formatted into a CSS `rgba()` string in JavaScript. This reference page doesn't have information about the absolute color space that should be used to interpret the RGB value—for example, sRGB, Adobe RGB, DCI-P3, and BT.2020. By default, applications should assume the sRGB color space. When color equality needs to be decided, implementations, unless documented otherwise, treat two colors as equal if all their red, green, blue, and alpha values each differ by at most `1e-5`. Example (Java): import com.google.type.Color; // ... public static java.awt.Color fromProto(Color protocolor) { float alpha = protocolor.hasAlpha() ? protocolor.getAlpha().getValue() : 1.0; return new java.awt.Color( protocolor.getRed(), protocolor.getGreen(), protocolor.getBlue(), alpha); } public static Color toProto(java.awt.Color color) { float red = (float) color.getRed(); float green = (float) color.getGreen(); float blue = (float) color.getBlue(); float denominator = 255.0; Color.Builder resultBuilder = Color .newBuilder() .setRed(red / denominator) .setGreen(green / denominator) .setBlue(blue / denominator); int alpha = color.getAlpha(); if (alpha != 255) { result.setAlpha( FloatValue .newBuilder() .setValue(((float) alpha) / denominator) .build()); } return resultBuilder.build(); } // ... Example (iOS / Obj-C): // ... static UIColor* fromProto(Color* protocolor) { float red = [protocolor red]; float green = [protocolor green]; float blue = [protocolor blue]; FloatValue* alpha_wrapper = [protocolor alpha]; float alpha = 1.0; if (alpha_wrapper != nil) { alpha = [alpha_wrapper value]; } return [UIColor colorWithRed:red green:green blue:blue alpha:alpha]; } static Color* toProto(UIColor* color) { CGFloat red, green, blue, alpha; if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) { return nil; } Color* result = [[Color alloc] init]; [result setRed:red]; [result setGreen:green]; [result setBlue:blue]; if (alpha <= 0.9999) { [result setAlpha:floatWrapperWithValue(alpha)]; } [result autorelease]; return result; } // ... Example (JavaScript): // ... var protoToCssColor = function(rgb_color) { var redFrac = rgb_color.red || 0.0; var greenFrac = rgb_color.green || 0.0; var blueFrac = rgb_color.blue || 0.0; var red = Math.floor(redFrac * 255); var green = Math.floor(greenFrac * 255); var blue = Math.floor(blueFrac * 255); if (!('alpha' in rgb_color)) { return rgbToCssColor(red, green, blue); } var alphaFrac = rgb_color.alpha.value || 0.0; var rgbParams = [red, green, blue].join(','); return ['rgba(', rgbParams, ',', alphaFrac, ')'].join(”); }; var rgbToCssColor = function(red, green, blue) { var rgbNumber = new Number((red << 16) | (green << 8) | blue); var hexString = rgbNumber.toString(16); var missingZeros = 6 - hexString.length; var resultBuilder = ['#']; for (var i = 0; i < missingZeros; i++) { resultBuilder.push('0'); } resultBuilder.push(hexString); return resultBuilder.join(”); }; // ...
type Color struct {
	Alpha float64 `json:"alpha,omitempty\"` // The fraction of this color that should be applied to the pixel. That is, the final pixel color is defined by the equation: `pixel color = alpha * (this color) + (1.0 - alpha) * (background color)` This means that a value of 1.0 corresponds to a solid color, whereas a value of 0.0 corresponds to a completely transparent color. This uses a wrapper message rather than a simple float scalar so that it is possible to distinguish between a default value and the value being unset. If omitted, this color object is rendered as a solid color (as if the alpha value had been explicitly given a value of 1.0).

	Blue float64 `json:"blue,omitempty\"` // The amount of blue in the color as a value in the interval [0, 1].

	Green float64 `json:"green,omitempty\"` // The amount of green in the color as a value in the interval [0, 1].

	Red float64 `json:"red,omitempty\"` // The amount of red in the color as a value in the interval [0, 1].

}

// A color value.
type ColorStyle struct {
	RgbColor Color `json:"rgbColor,omitempty\"` // RGB color. The [`alpha`](https://developers.google.com/workspace/sheets/api/reference/rest/v4/spreadsheets/other#Color.FIELDS.alpha) value in the [`Color`](https://developers.google.com/workspace/sheets/api/reference/rest/v4/spreadsheets/other#color) object isn't generally supported.

	ThemeColor string `json:"themeColor,omitempty\"` // Theme color.

}

// The value of the condition.
type ConditionValue struct {
	RelativeDate string `json:"relativeDate,omitempty\"` // A relative date (based on the current date). Valid only if the type is DATE_BEFORE, DATE_AFTER, DATE_ON_OR_BEFORE or DATE_ON_OR_AFTER. Relative dates are not supported in data validation. They are supported only in conditional formatting and conditional filters.

	UserEnteredValue string `json:"userEnteredValue,omitempty\"` // A value the condition is based on. The value is parsed as if the user typed into a cell. Formulas are supported (and must begin with an `=` or a '+').

}

// A rule describing a conditional format.
type ConditionalFormatRule struct {
	BooleanRule BooleanRule `json:"booleanRule,omitempty\"` // The formatting is either "on" or "off" according to the rule.

	GradientRule GradientRule `json:"gradientRule,omitempty\"` // The formatting will vary based on the gradients in the rule.

	Ranges []GridRange `json:"ranges,omitempty\"` // The ranges that are formatted if the condition is true. All the ranges must be on the same grid.

}

// Copies data from the source to the destination.
type CopyPasteRequest struct {
	Destination GridRange `json:"destination,omitempty\"` // The location to paste to. If the range covers a span that's a multiple of the source's height or width, then the data will be repeated to fill in the destination range. If the range is smaller than the source range, the entire source data will still be copied (beyond the end of the destination range).

	PasteOrientation string `json:"pasteOrientation,omitempty\"` // How that data should be oriented when pasting.

	PasteType string `json:"pasteType,omitempty\"` // What kind of data to paste.

	Source GridRange `json:"source,omitempty\"` // The source range to copy.

}

// The request to copy a sheet across spreadsheets.
type CopySheetToAnotherSpreadsheetRequest struct {
	DestinationSpreadsheetId string `json:"destinationSpreadsheetId,omitempty\"` // The ID of the spreadsheet to copy the sheet to.

}

// A request to create developer metadata.
type CreateDeveloperMetadataRequest struct {
	DeveloperMetadata DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata to create.

}

// The response from creating developer metadata.
type CreateDeveloperMetadataResponse struct {
	DeveloperMetadata DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata that was created.

}

// Moves data from the source to the destination.
type CutPasteRequest struct {
	Destination GridCoordinate `json:"destination,omitempty\"` // The top-left coordinate where the data should be pasted.

	PasteType string `json:"pasteType,omitempty\"` // What kind of data to paste. All the source data will be cut, regardless of what is pasted.

	Source GridRange `json:"source,omitempty\"` // The source data to cut.

}

// The data execution status. A data execution is created to sync a data source object with the latest data from a DataSource. It is usually scheduled to run at background, you can check its state to tell if an execution completes There are several scenarios where a data execution is triggered to run: * Adding a data source creates an associated data source sheet as well as a data execution to sync the data from the data source to the sheet. * Updating a data source creates a data execution to refresh the associated data source sheet similarly. * You can send refresh request to explicitly refresh one or multiple data source objects.
type DataExecutionStatus struct {
	ErrorCode string `json:"errorCode,omitempty\"` // The error code.

	ErrorMessage string `json:"errorMessage,omitempty\"` // The error message, which may be empty.

	LastRefreshTime string `json:"lastRefreshTime,omitempty\"` // Gets the time the data last successfully refreshed.

	State string `json:"state,omitempty\"` // The state of the data execution.

}

// Filter that describes what data should be selected or returned from a request. For more information, see [Read, write, and search metadata](https://developers.google.com/workspace/sheets/api/guides/metadata).
type DataFilter struct {
	A1Range string `json:"a1Range,omitempty\"` // Selects data that matches the specified A1 range.

	DeveloperMetadataLookup DeveloperMetadataLookup `json:"developerMetadataLookup,omitempty\"` // Selects data associated with the developer metadata matching the criteria described by this DeveloperMetadataLookup.

	GridRange GridRange `json:"gridRange,omitempty\"` // Selects data that matches the range described by the GridRange.

}

// A range of values whose location is specified by a DataFilter.
type DataFilterValueRange struct {
	DataFilter DataFilter `json:"dataFilter,omitempty\"` // The data filter describing the location of the values in the spreadsheet.

	MajorDimension string `json:"majorDimension,omitempty\"` // The major dimension of the values.

	Values [][]interface{} `json:"values,omitempty\"` // The data to be written. If the provided values exceed any of the ranges matched by the data filter then the request fails. If the provided values are less than the matched ranges only the specified values are written, existing values in the matched ranges remain unaffected.

}

// Settings for one set of data labels. Data labels are annotations that appear next to a set of data, such as the points on a line chart, and provide additional information about what the data represents, such as a text representation of the value behind that point on the graph.
type DataLabel struct {
	CustomLabelData ChartData `json:"customLabelData,omitempty\"` // Data to use for custom labels. Only used if type is set to CUSTOM. This data must be the same length as the series or other element this data label is applied to. In addition, if the series is split into multiple source ranges, this source data must come from the next column in the source data. For example, if the series is B2:B4,E6:E8 then this data must come from C2:C4,F6:F8.

	Placement string `json:"placement,omitempty\"` // The placement of the data label relative to the labeled data.

	TextFormat TextFormat `json:"textFormat,omitempty\"` // The text format used for the data label. The link field is not supported.

	TypeValue string `json:"type,omitempty\"` // The type of the data label.

}

// Information about an external data source in the spreadsheet.
type DataSource struct {
	CalculatedColumns []DataSourceColumn `json:"calculatedColumns,omitempty\"` // All calculated columns in the data source.

	DataSourceId string `json:"dataSourceId,omitempty\"` // The spreadsheet-scoped unique ID that identifies the data source. Example: 1080547365.

	SheetId int `json:"sheetId,omitempty\"` // The ID of the Sheet connected with the data source. The field cannot be changed once set. When creating a data source, an associated DATA_SOURCE sheet is also created, if the field is not specified, the ID of the created sheet will be randomly generated.

	Spec DataSourceSpec `json:"spec,omitempty\"` // The DataSourceSpec for the data source connected with this spreadsheet.

}

// Properties of a data source chart.
type DataSourceChartProperties struct {
	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // Output only. The data execution status.

	DataSourceId string `json:"dataSourceId,omitempty\"` // ID of the data source that the chart is associated with.

}

// A column in a data source.
type DataSourceColumn struct {
	Formula string `json:"formula,omitempty\"` // The formula of the calculated column.

	Reference DataSourceColumnReference `json:"reference,omitempty\"` // The column reference.

}

// An unique identifier that references a data source column.
type DataSourceColumnReference struct {
	Name string `json:"name,omitempty\"` // The display name of the column. It should be unique within a data source.

}

// A data source formula.
type DataSourceFormula struct {
	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // Output only. The data execution status.

	DataSourceId string `json:"dataSourceId,omitempty\"` // The ID of the data source the formula is associated with.

}

// Reference to a data source object.
type DataSourceObjectReference struct {
	ChartId int `json:"chartId,omitempty\"` // References to a data source chart.

	DataSourceFormulaCell GridCoordinate `json:"dataSourceFormulaCell,omitempty\"` // References to a cell containing DataSourceFormula.

	DataSourcePivotTableAnchorCell GridCoordinate `json:"dataSourcePivotTableAnchorCell,omitempty\"` // References to a data source PivotTable anchored at the cell.

	DataSourceTableAnchorCell GridCoordinate `json:"dataSourceTableAnchorCell,omitempty\"` // References to a DataSourceTable anchored at the cell.

	SheetId string `json:"sheetId,omitempty\"` // References to a DATA_SOURCE sheet.

}

// A list of references to data source objects.
type DataSourceObjectReferences struct {
	References []DataSourceObjectReference `json:"references,omitempty\"` // The references.

}

// A parameter in a data source's query. The parameter allows the user to pass in values from the spreadsheet into a query.
type DataSourceParameter struct {
	Name string `json:"name,omitempty\"` // Named parameter. Must be a legitimate identifier for the DataSource that supports it. For example, [BigQuery identifier](https://cloud.google.com/bigquery/docs/reference/standard-sql/lexical#identifiers).

	NamedRangeId string `json:"namedRangeId,omitempty\"` // ID of a NamedRange. Its size must be 1x1.

	RangeValue GridRange `json:"range,omitempty\"` // A range that contains the value of the parameter. Its size must be 1x1.

}

// A schedule for data to refresh every day in a given time interval.
type DataSourceRefreshDailySchedule struct {
	StartTime TimeOfDay `json:"startTime,omitempty\"` // The start time of a time interval in which a data source refresh is scheduled. Only `hours` part is used. The time interval size defaults to that in the Sheets editor.

}

// A monthly schedule for data to refresh on specific days in the month in a given time interval.
type DataSourceRefreshMonthlySchedule struct {
	DaysOfMonth []int `json:"daysOfMonth,omitempty\"` // Days of the month to refresh. Only 1-28 are supported, mapping to the 1st to the 28th day. At least one day must be specified.

	StartTime TimeOfDay `json:"startTime,omitempty\"` // The start time of a time interval in which a data source refresh is scheduled. Only `hours` part is used. The time interval size defaults to that in the Sheets editor.

}

// Schedule for refreshing the data source. Data sources in the spreadsheet are refreshed within a time interval. You can specify the start time by clicking the Scheduled Refresh button in the Sheets editor, but the interval is fixed at 4 hours. For example, if you specify a start time of 8 AM , the refresh will take place between 8 AM and 12 PM every day.
type DataSourceRefreshSchedule struct {
	DailySchedule DataSourceRefreshDailySchedule `json:"dailySchedule,omitempty\"` // Daily refresh schedule.

	Enabled bool `json:"enabled,omitempty\"` // True if the refresh schedule is enabled, or false otherwise.

	MonthlySchedule DataSourceRefreshMonthlySchedule `json:"monthlySchedule,omitempty\"` // Monthly refresh schedule.

	NextRun Interval `json:"nextRun,omitempty\"` // Output only. The time interval of the next run.

	RefreshScope string `json:"refreshScope,omitempty\"` // The scope of the refresh. Must be ALL_DATA_SOURCES.

	WeeklySchedule DataSourceRefreshWeeklySchedule `json:"weeklySchedule,omitempty\"` // Weekly refresh schedule.

}

// A weekly schedule for data to refresh on specific days in a given time interval.
type DataSourceRefreshWeeklySchedule struct {
	DaysOfWeek []string `json:"daysOfWeek,omitempty\"` // Days of the week to refresh. At least one day must be specified.

	StartTime TimeOfDay `json:"startTime,omitempty\"` // The start time of a time interval in which a data source refresh is scheduled. Only `hours` part is used. The time interval size defaults to that in the Sheets editor.

}

// A range along a single dimension on a DATA_SOURCE sheet.
type DataSourceSheetDimensionRange struct {
	ColumnReferences []DataSourceColumnReference `json:"columnReferences,omitempty\"` // The columns on the data source sheet.

	SheetId int `json:"sheetId,omitempty\"` // The ID of the data source sheet the range is on.

}

// Additional properties of a DATA_SOURCE sheet.
type DataSourceSheetProperties struct {
	Columns []DataSourceColumn `json:"columns,omitempty\"` // The columns displayed on the sheet, corresponding to the values in RowData.

	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // The data execution status.

	DataSourceId string `json:"dataSourceId,omitempty\"` // ID of the DataSource the sheet is connected to.

}

// This specifies the details of the data source. For example, for BigQuery, this specifies information about the BigQuery source.
type DataSourceSpec struct {
	BigQuery BigQueryDataSourceSpec `json:"bigQuery,omitempty\"` // A BigQueryDataSourceSpec.

	Looker LookerDataSourceSpec `json:"looker,omitempty\"` // A LookerDatasourceSpec.

	Parameters []DataSourceParameter `json:"parameters,omitempty\"` // The parameters of the data source, used when querying the data source.

}

// A data source table, which allows the user to import a static table of data from the DataSource into Sheets. This is also known as "Extract" in the Sheets editor.
type DataSourceTable struct {
	ColumnSelectionType string `json:"columnSelectionType,omitempty\"` // The type to select columns for the data source table. Defaults to SELECTED.

	Columns []DataSourceColumnReference `json:"columns,omitempty\"` // Columns selected for the data source table. The column_selection_type must be SELECTED.

	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // Output only. The data execution status.

	DataSourceId string `json:"dataSourceId,omitempty\"` // The ID of the data source the data source table is associated with.

	FilterSpecs []FilterSpec `json:"filterSpecs,omitempty\"` // Filter specifications in the data source table.

	RowLimit int `json:"rowLimit,omitempty\"` // The limit of rows to return. If not set, a default limit is applied. Please refer to the Sheets editor for the default and max limit.

	SortSpecs []SortSpec `json:"sortSpecs,omitempty\"` // Sort specifications in the data source table. The result of the data source table is sorted based on the sort specifications in order.

}

// A data validation rule.
type DataValidationRule struct {
	Condition BooleanCondition `json:"condition,omitempty\"` // The condition that data in the cell must match.

	InputMessage string `json:"inputMessage,omitempty\"` // A message to show the user when adding data to the cell.

	ShowCustomUi bool `json:"showCustomUi,omitempty\"` // True if the UI should be customized based on the kind of condition. If true, "List" conditions will show a dropdown.

	Strict bool `json:"strict,omitempty\"` // True if invalid data should be rejected.

}

// Allows you to organize the date-time values in a source data column into buckets based on selected parts of their date or time values. For example, consider a pivot table showing sales transactions by date: +----------+--------------+ | Date | SUM of Sales | +----------+--------------+ | 1/1/2017 | $621.14 | | 2/3/2017 | $708.84 | | 5/8/2017 | $326.84 | ... +----------+--------------+ Applying a date-time group rule with a DateTimeRuleType of YEAR_MONTH results in the following pivot table. +--------------+--------------+ | Grouped Date | SUM of Sales | +--------------+--------------+ | 2017-Jan | $53,731.78 | | 2017-Feb | $83,475.32 | | 2017-Mar | $94,385.05 | ... +--------------+--------------+
type DateTimeRule struct {
	TypeValue string `json:"type,omitempty\"` // The type of date-time grouping to apply.

}

// Removes the banded range with the given ID from the spreadsheet.
type DeleteBandingRequest struct {
	BandedRangeId int `json:"bandedRangeId,omitempty\"` // The ID of the banded range to delete.

}

// Deletes a conditional format rule at the given index. All subsequent rules' indexes are decremented.
type DeleteConditionalFormatRuleRequest struct {
	Index int `json:"index,omitempty\"` // The zero-based index of the rule to be deleted.

	SheetId int `json:"sheetId,omitempty\"` // The sheet the rule is being deleted from.

}

// The result of deleting a conditional format rule.
type DeleteConditionalFormatRuleResponse struct {
	Rule ConditionalFormatRule `json:"rule,omitempty\"` // The rule that was deleted.

}

// Deletes a data source. The request also deletes the associated data source sheet, and unlinks all associated data source objects.
type DeleteDataSourceRequest struct {
	DataSourceId string `json:"dataSourceId,omitempty\"` // The ID of the data source to delete.

}

// A request to delete developer metadata.
type DeleteDeveloperMetadataRequest struct {
	DataFilter DataFilter `json:"dataFilter,omitempty\"` // The data filter describing the criteria used to select which developer metadata entry to delete.

}

// The response from deleting developer metadata.
type DeleteDeveloperMetadataResponse struct {
	DeletedDeveloperMetadata []DeveloperMetadata `json:"deletedDeveloperMetadata,omitempty\"` // The metadata that was deleted.

}

// Deletes a group over the specified range by decrementing the depth of the dimensions in the range. For example, assume the sheet has a depth-1 group over B:E and a depth-2 group over C:D. Deleting a group over D:E leaves the sheet with a depth-1 group over B:D and a depth-2 group over C:C.
type DeleteDimensionGroupRequest struct {
	RangeValue DimensionRange `json:"range,omitempty\"` // The range of the group to be deleted.

}

// The result of deleting a group.
type DeleteDimensionGroupResponse struct {
	DimensionGroups []DimensionGroup `json:"dimensionGroups,omitempty\"` // All groups of a dimension after deleting a group from that dimension.

}

// Deletes the dimensions from the sheet.
type DeleteDimensionRequest struct {
	RangeValue DimensionRange `json:"range,omitempty\"` // The dimensions to delete from the sheet.

}

// Removes rows within this range that contain values in the specified columns that are duplicates of values in any previous row. Rows with identical values but different letter cases, formatting, or formulas are considered to be duplicates. This request also removes duplicate rows hidden from view (for example, due to a filter). When removing duplicates, the first instance of each duplicate row scanning from the top downwards is kept in the resulting range. Content outside of the specified range isn't removed, and rows considered duplicates do not have to be adjacent to each other in the range.
type DeleteDuplicatesRequest struct {
	ComparisonColumns []DimensionRange `json:"comparisonColumns,omitempty\"` // The columns in the range to analyze for duplicate values. If no columns are selected then all columns are analyzed for duplicates.

	RangeValue GridRange `json:"range,omitempty\"` // The range to remove duplicates rows from.

}

// The result of removing duplicates in a range.
type DeleteDuplicatesResponse struct {
	DuplicatesRemovedCount int `json:"duplicatesRemovedCount,omitempty\"` // The number of duplicate rows removed.

}

// Deletes the embedded object with the given ID.
type DeleteEmbeddedObjectRequest struct {
	ObjectId int `json:"objectId,omitempty\"` // The ID of the embedded object to delete.

}

// Deletes a particular filter view.
type DeleteFilterViewRequest struct {
	FilterId int `json:"filterId,omitempty\"` // The ID of the filter to delete.

}

// Removes the named range with the given ID from the spreadsheet.
type DeleteNamedRangeRequest struct {
	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the named range to delete.

}

// Deletes the protected range with the given ID.
type DeleteProtectedRangeRequest struct {
	ProtectedRangeId int `json:"protectedRangeId,omitempty\"` // The ID of the protected range to delete.

}

// Deletes a range of cells, shifting other cells into the deleted area.
type DeleteRangeRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range of cells to delete.

	ShiftDimension string `json:"shiftDimension,omitempty\"` // The dimension from which deleted cells will be replaced with. If ROWS, existing cells will be shifted upward to replace the deleted cells. If COLUMNS, existing cells will be shifted left to replace the deleted cells.

}

// Deletes the requested sheet.
type DeleteSheetRequest struct {
	SheetId int `json:"sheetId,omitempty\"` // The ID of the sheet to delete. If the sheet is of DATA_SOURCE type, the associated DataSource is also deleted.

}

// Removes the table with the given ID from the spreadsheet.
type DeleteTableRequest struct {
	TableId string `json:"tableId,omitempty\"` // The ID of the table to delete.

}

// Developer metadata associated with a location or object in a spreadsheet. For more information, see [Read, write, and search metadata](https://developers.google.com/workspace/sheets/api/guides/metadata). Developer metadata may be used to associate arbitrary data with various parts of a spreadsheet and it will remain associated at those locations as they move around and the spreadsheet is edited. For example, if developer metadata is associated with row 5 and another row is then subsequently inserted above row 5, that original metadata is still associated with the row it was first associated with (what is now row 6). If the associated object is deleted then its metadata is deleted too.
type DeveloperMetadata struct {
	Location DeveloperMetadataLocation `json:"location,omitempty\"` // The location where the metadata is associated.

	MetadataId int `json:"metadataId,omitempty\"` // The spreadsheet-scoped unique ID that identifies the metadata. IDs may be specified when metadata is created, otherwise one will be randomly generated and assigned. Must be positive.

	MetadataKey string `json:"metadataKey,omitempty\"` // The metadata key. There may be multiple metadata in a spreadsheet with the same key. Developer metadata must always have a key specified.

	MetadataValue string `json:"metadataValue,omitempty\"` // Data associated with the metadata's key.

	Visibility string `json:"visibility,omitempty\"` // The metadata visibility. Developer metadata must always have visibility specified.

}

// A location where metadata may be associated in a spreadsheet.
type DeveloperMetadataLocation struct {
	DimensionRange DimensionRange `json:"dimensionRange,omitempty\"` // Represents the row or column when metadata is associated with a dimension. The specified DimensionRange must represent a single row or column. It cannot be unbounded or span multiple rows or columns.

	LocationType string `json:"locationType,omitempty\"` // The type of location this object represents. This field is read-only.

	SheetId int `json:"sheetId,omitempty\"` // The ID of the sheet when metadata is associated with an entire sheet.

	Spreadsheet bool `json:"spreadsheet,omitempty\"` // True when metadata is associated with an entire spreadsheet.

}

// Selects DeveloperMetadata that matches all of the specified fields. For example, if only a metadata ID is specified this considers the DeveloperMetadata with that particular unique ID. If a metadata key is specified, this considers all developer metadata with that key. If a key, visibility, and location type are all specified, this considers all developer metadata with that key and visibility that are associated with a location of that type. In general, this selects all DeveloperMetadata that match the intersection of all the specified fields; any field or combination of fields may be specified.
type DeveloperMetadataLookup struct {
	LocationMatchingStrategy string `json:"locationMatchingStrategy,omitempty\"` // Determines how this lookup matches the location. If this field is specified as EXACT, only developer metadata associated on the exact location specified is matched. If this field is specified to INTERSECTING, developer metadata associated on intersecting locations is also matched. If left unspecified, this field assumes a default value of INTERSECTING. If this field is specified, a metadataLocation must also be specified.

	LocationType string `json:"locationType,omitempty\"` // Limits the selected developer metadata to those entries which are associated with locations of the specified type. For example, when this field is specified as ROW this lookup only considers developer metadata associated on rows. If the field is left unspecified, all location types are considered. This field cannot be specified as SPREADSHEET when the locationMatchingStrategy is specified as INTERSECTING or when the metadataLocation is specified as a non-spreadsheet location. Spreadsheet metadata cannot intersect any other developer metadata location. This field also must be left unspecified when the locationMatchingStrategy is specified as EXACT.

	MetadataId int `json:"metadataId,omitempty\"` // Limits the selected developer metadata to that which has a matching DeveloperMetadata.metadata_id.

	MetadataKey string `json:"metadataKey,omitempty\"` // Limits the selected developer metadata to that which has a matching DeveloperMetadata.metadata_key.

	MetadataLocation DeveloperMetadataLocation `json:"metadataLocation,omitempty\"` // Limits the selected developer metadata to those entries associated with the specified location. This field either matches exact locations or all intersecting locations according the specified locationMatchingStrategy.

	MetadataValue string `json:"metadataValue,omitempty\"` // Limits the selected developer metadata to that which has a matching DeveloperMetadata.metadata_value.

	Visibility string `json:"visibility,omitempty\"` // Limits the selected developer metadata to that which has a matching DeveloperMetadata.visibility. If left unspecified, all developer metadata visible to the requesting project is considered.

}

// A group over an interval of rows or columns on a sheet, which can contain or be contained within other groups. A group can be collapsed or expanded as a unit on the sheet.
type DimensionGroup struct {
	Collapsed bool `json:"collapsed,omitempty\"` // This field is true if this group is collapsed. A collapsed group remains collapsed if an overlapping group at a shallower depth is expanded. A true value does not imply that all dimensions within the group are hidden, since a dimension's visibility can change independently from this group property. However, when this property is updated, all dimensions within it are set to hidden if this field is true, or set to visible if this field is false.

	Depth int `json:"depth,omitempty\"` // The depth of the group, representing how many groups have a range that wholly contains the range of this group.

	RangeValue DimensionRange `json:"range,omitempty\"` // The range over which this group exists.

}

// Properties about a dimension.
type DimensionProperties struct {
	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // Output only. If set, this is a column in a data source sheet.

	DeveloperMetadata []DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata associated with a single row or column.

	HiddenByFilter bool `json:"hiddenByFilter,omitempty\"` // True if this dimension is being filtered. This field is read-only.

	HiddenByUser bool `json:"hiddenByUser,omitempty\"` // True if this dimension is explicitly hidden.

	PixelSize int `json:"pixelSize,omitempty\"` // The height (if a row) or width (if a column) of the dimension in pixels.

}

// A range along a single dimension on a sheet. All indexes are zero-based. Indexes are half open: the start index is inclusive and the end index is exclusive. Missing indexes indicate the range is unbounded on that side.
type DimensionRange struct {
	Dimension string `json:"dimension,omitempty\"` // The dimension of the span.

	EndIndex int `json:"endIndex,omitempty\"` // The end (exclusive) of the span, or not set if unbounded.

	SheetId int `json:"sheetId,omitempty\"` // The sheet this span is on.

	StartIndex int `json:"startIndex,omitempty\"` // The start (inclusive) of the span, or not set if unbounded.

}

// Duplicates a particular filter view.
type DuplicateFilterViewRequest struct {
	FilterId int `json:"filterId,omitempty\"` // The ID of the filter being duplicated.

}

// The result of a filter view being duplicated.
type DuplicateFilterViewResponse struct {
	Filter FilterView `json:"filter,omitempty\"` // The newly created filter.

}

// Duplicates the contents of a sheet.
type DuplicateSheetRequest struct {
	InsertSheetIndex int `json:"insertSheetIndex,omitempty\"` // The zero-based index where the new sheet should be inserted. The index of all sheets after this are incremented.

	NewSheetId int `json:"newSheetId,omitempty\"` // If set, the ID of the new sheet. If not set, an ID is chosen. If set, the ID must not conflict with any existing sheet ID. If set, it must be non-negative.

	NewSheetName string `json:"newSheetName,omitempty\"` // The name of the new sheet. If empty, a new name is chosen for you.

	SourceSheetId int `json:"sourceSheetId,omitempty\"` // The sheet to duplicate. If the source sheet is of DATA_SOURCE type, its backing DataSource is also duplicated and associated with the new copy of the sheet. No data execution is triggered, the grid data of this sheet is also copied over but only available after the batch request completes.

}

// The result of duplicating a sheet.
type DuplicateSheetResponse struct {
	Properties SheetProperties `json:"properties,omitempty\"` // The properties of the duplicate sheet.

}

// The editors of a protected range.
type Editors struct {
	DomainUsersCanEdit bool `json:"domainUsersCanEdit,omitempty\"` // True if anyone in the document's domain has edit access to the protected range. Domain protection is only supported on documents within a domain.

	Groups []string `json:"groups,omitempty\"` // The email addresses of groups with edit access to the protected range.

	Users []string `json:"users,omitempty\"` // The email addresses of users with edit access to the protected range.

}

// A chart embedded in a sheet.
type EmbeddedChart struct {
	Border EmbeddedObjectBorder `json:"border,omitempty\"` // The border of the chart.

	ChartId int `json:"chartId,omitempty\"` // The ID of the chart.

	Position EmbeddedObjectPosition `json:"position,omitempty\"` // The position of the chart.

	Spec ChartSpec `json:"spec,omitempty\"` // The specification of the chart.

}

// A border along an embedded object.
type EmbeddedObjectBorder struct {
	Color Color `json:"color,omitempty\"` // The color of the border. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // The color of the border. If color is also set, this field takes precedence.

}

// The position of an embedded object such as a chart.
type EmbeddedObjectPosition struct {
	NewSheet bool `json:"newSheet,omitempty\"` // If true, the embedded object is put on a new sheet whose ID is chosen for you. Used only when writing.

	OverlayPosition OverlayPosition `json:"overlayPosition,omitempty\"` // The position at which the object is overlaid on top of a grid.

	SheetId int `json:"sheetId,omitempty\"` // The sheet this is on. Set only if the embedded object is on its own sheet. Must be non-negative.

}

// An error in a cell.
type ErrorValue struct {
	Message string `json:"message,omitempty\"` // A message with more information about the error (in the spreadsheet's locale).

	TypeValue string `json:"type,omitempty\"` // The type of error.

}

// The kinds of value that a cell in a spreadsheet can have.
type ExtendedValue struct {
	BoolValue bool `json:"boolValue,omitempty\"` // Represents a boolean value.

	ErrorValue ErrorValue `json:"errorValue,omitempty\"` // Represents an error. This field is read-only.

	FormulaValue string `json:"formulaValue,omitempty\"` // Represents a formula.

	NumberValue float64 `json:"numberValue,omitempty\"` // Represents a double value. Note: Dates, Times and DateTimes are represented as doubles in SERIAL_NUMBER format.

	StringValue string `json:"stringValue,omitempty\"` // Represents a string value. Leading single quotes are not included. For example, if the user typed `'123` into the UI, this would be represented as a `stringValue` of `"123"`.

}

// Criteria for showing or hiding rows in a filter or filter view.
type FilterCriteria struct {
	Condition BooleanCondition `json:"condition,omitempty\"` // A condition that must be `true` for values to be shown. (This does not override hidden_values -- if a value is listed there, it will still be hidden.)

	HiddenValues []string `json:"hiddenValues,omitempty\"` // Values that should be hidden.

	VisibleBackgroundColor Color `json:"visibleBackgroundColor,omitempty\"` // The background fill color to filter by; only cells with this fill color are shown. Mutually exclusive with visible_foreground_color. Deprecated: Use visible_background_color_style.

	VisibleBackgroundColorStyle ColorStyle `json:"visibleBackgroundColorStyle,omitempty\"` // The background fill color to filter by; only cells with this fill color are shown. This field is mutually exclusive with visible_foreground_color, and must be set to an RGB-type color. If visible_background_color is also set, this field takes precedence.

	VisibleForegroundColor Color `json:"visibleForegroundColor,omitempty\"` // The foreground color to filter by; only cells with this foreground color are shown. Mutually exclusive with visible_background_color. Deprecated: Use visible_foreground_color_style.

	VisibleForegroundColorStyle ColorStyle `json:"visibleForegroundColorStyle,omitempty\"` // The foreground color to filter by; only cells with this foreground color are shown. This field is mutually exclusive with visible_background_color, and must be set to an RGB-type color. If visible_foreground_color is also set, this field takes precedence.

}

// The filter criteria associated with a specific column.
type FilterSpec struct {
	ColumnIndex int `json:"columnIndex,omitempty\"` // The zero-based column index.

	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // Reference to a data source column.

	FilterCriteria FilterCriteria `json:"filterCriteria,omitempty\"` // The criteria for the column.

}

// A filter view. For more information, see [Manage data visibility with filters](https://developers.google.com/workspace/sheets/api/guides/filters).
type FilterView struct {
	Criteria map[string]interface{} `json:"criteria,omitempty\"` // The criteria for showing/hiding values per column. The map's key is the column index, and the value is the criteria for that column. This field is deprecated in favor of filter_specs.

	FilterSpecs []FilterSpec `json:"filterSpecs,omitempty\"` // The filter criteria for showing or hiding values per column. Both criteria and filter_specs are populated in responses. If both fields are specified in an update request, this field takes precedence.

	FilterViewId int `json:"filterViewId,omitempty\"` // The ID of the filter view.

	NamedRangeId string `json:"namedRangeId,omitempty\"` // The named range this filter view is backed by, if any. When writing, only one of range, named_range_id, or table_id may be set.

	RangeValue GridRange `json:"range,omitempty\"` // The range this filter view covers. When writing, only one of range, named_range_id, or table_id may be set.

	SortSpecs []SortSpec `json:"sortSpecs,omitempty\"` // The sort order per column. Later specifications are used when values are equal in the earlier specifications.

	TableId string `json:"tableId,omitempty\"` // The table this filter view is backed by, if any. When writing, only one of range, named_range_id, or table_id may be set.

	Title string `json:"title,omitempty\"` // The name of the filter view.

}

// Finds and replaces data in cells over a range, sheet, or all sheets.
type FindReplaceRequest struct {
	AllSheets bool `json:"allSheets,omitempty\"` // True to find/replace over all sheets.

	Find string `json:"find,omitempty\"` // The value to search.

	IncludeFormulas bool `json:"includeFormulas,omitempty\"` // True if the search should include cells with formulas. False to skip cells with formulas.

	MatchCase bool `json:"matchCase,omitempty\"` // True if the search is case sensitive.

	MatchEntireCell bool `json:"matchEntireCell,omitempty\"` // True if the find value should match the entire cell.

	RangeValue GridRange `json:"range,omitempty\"` // The range to find/replace over.

	Replacement string `json:"replacement,omitempty\"` // The value to use as the replacement.

	SearchByRegex bool `json:"searchByRegex,omitempty\"` // True if the find value is a regex. The regular expression and replacement should follow Java regex rules at https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html. The replacement string is allowed to refer to capturing groups. For example, if one cell has the contents `"Google Sheets"` and another has `"Google Docs"`, then searching for `"o.* (.*)"` with a replacement of `"$1 Rocks"` would change the contents of the cells to `"GSheets Rocks"` and `"GDocs Rocks"` respectively.

	SheetId int `json:"sheetId,omitempty\"` // The sheet to find/replace over.

}

// The result of the find/replace.
type FindReplaceResponse struct {
	FormulasChanged int `json:"formulasChanged,omitempty\"` // The number of formula cells changed.

	OccurrencesChanged int `json:"occurrencesChanged,omitempty\"` // The number of occurrences (possibly multiple within a cell) changed. For example, if replacing `"e"` with `"o"` in `"Google Sheets"`, this would be `"3"` because `"Google Sheets"` -> `"Googlo Shoots"`.

	RowsChanged int `json:"rowsChanged,omitempty\"` // The number of rows changed.

	SheetsChanged int `json:"sheetsChanged,omitempty\"` // The number of sheets changed.

	ValuesChanged int `json:"valuesChanged,omitempty\"` // The number of non-formula cells changed.

}

// The request for retrieving a Spreadsheet.
type GetSpreadsheetByDataFilterRequest struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The DataFilters used to select which ranges to retrieve from the spreadsheet.

	ExcludeTablesInBandedRanges bool `json:"excludeTablesInBandedRanges,omitempty\"` // True if tables should be excluded in the banded ranges. False if not set.

	IncludeGridData bool `json:"includeGridData,omitempty\"` // True if grid data should be returned. This parameter is ignored if a field mask was set in the request.

}

// A rule that applies a gradient color scale format, based on the interpolation points listed. The format of a cell will vary based on its contents as compared to the values of the interpolation points.
type GradientRule struct {
	Maxpoint InterpolationPoint `json:"maxpoint,omitempty\"` // The final interpolation point.

	Midpoint InterpolationPoint `json:"midpoint,omitempty\"` // An optional midway interpolation point.

	Minpoint InterpolationPoint `json:"minpoint,omitempty\"` // The starting interpolation point.

}

// A coordinate in a sheet. All indexes are zero-based.
type GridCoordinate struct {
	ColumnIndex int `json:"columnIndex,omitempty\"` // The column index of the coordinate.

	RowIndex int `json:"rowIndex,omitempty\"` // The row index of the coordinate.

	SheetId int `json:"sheetId,omitempty\"` // The sheet this coordinate is on.

}

// Data in the grid, as well as metadata about the dimensions.
type GridData struct {
	ColumnMetadata []DimensionProperties `json:"columnMetadata,omitempty\"` // Metadata about the requested columns in the grid, starting with the column in start_column.

	RowData []RowData `json:"rowData,omitempty\"` // The data in the grid, one entry per row, starting with the row in startRow. The values in RowData will correspond to columns starting at start_column.

	RowMetadata []DimensionProperties `json:"rowMetadata,omitempty\"` // Metadata about the requested rows in the grid, starting with the row in start_row.

	StartColumn int `json:"startColumn,omitempty\"` // The first column this GridData refers to, zero-based.

	StartRow int `json:"startRow,omitempty\"` // The first row this GridData refers to, zero-based.

}

// Properties of a grid.
type GridProperties struct {
	ColumnCount int `json:"columnCount,omitempty\"` // The number of columns in the grid.

	ColumnGroupControlAfter bool `json:"columnGroupControlAfter,omitempty\"` // True if the column grouping control toggle is shown after the group.

	FrozenColumnCount int `json:"frozenColumnCount,omitempty\"` // The number of columns that are frozen in the grid.

	FrozenRowCount int `json:"frozenRowCount,omitempty\"` // The number of rows that are frozen in the grid.

	HideGridlines bool `json:"hideGridlines,omitempty\"` // True if the grid isn't showing gridlines in the UI.

	RowCount int `json:"rowCount,omitempty\"` // The number of rows in the grid.

	RowGroupControlAfter bool `json:"rowGroupControlAfter,omitempty\"` // True if the row grouping control toggle is shown after the group.

}

// A range on a sheet. All indexes are zero-based. Indexes are half open, i.e. the start index is inclusive and the end index is exclusive -- [start_index, end_index). Missing indexes indicate the range is unbounded on that side. For example, if `"Sheet1"` is sheet ID 123456, then: `Sheet1!A1:A1 == sheet_id: 123456, start_row_index: 0, end_row_index: 1, start_column_index: 0, end_column_index: 1` `Sheet1!A3:B4 == sheet_id: 123456, start_row_index: 2, end_row_index: 4, start_column_index: 0, end_column_index: 2` `Sheet1!A:B == sheet_id: 123456, start_column_index: 0, end_column_index: 2` `Sheet1!A5:B == sheet_id: 123456, start_row_index: 4, start_column_index: 0, end_column_index: 2` `Sheet1 == sheet_id: 123456` The start index must always be less than or equal to the end index. If the start index equals the end index, then the range is empty. Empty ranges are typically not meaningful and are usually rendered in the UI as `#REF!`.
type GridRange struct {
	EndColumnIndex int `json:"endColumnIndex,omitempty\"` // The end column (exclusive) of the range, or not set if unbounded.

	EndRowIndex int `json:"endRowIndex,omitempty\"` // The end row (exclusive) of the range, or not set if unbounded.

	SheetId int `json:"sheetId,omitempty\"` // The sheet this range is on.

	StartColumnIndex int `json:"startColumnIndex,omitempty\"` // The start column (inclusive) of the range, or not set if unbounded.

	StartRowIndex int `json:"startRowIndex,omitempty\"` // The start row (inclusive) of the range, or not set if unbounded.

}

// A histogram chart. A histogram chart groups data items into bins, displaying each bin as a column of stacked items. Histograms are used to display the distribution of a dataset. Each column of items represents a range into which those items fall. The number of bins can be chosen automatically or specified explicitly.
type HistogramChartSpec struct {
	BucketSize float64 `json:"bucketSize,omitempty\"` // By default the bucket size (the range of values stacked in a single column) is chosen automatically, but it may be overridden here. E.g., A bucket size of 1.5 results in buckets from 0 - 1.5, 1.5 - 3.0, etc. Cannot be negative. This field is optional.

	LegendPosition string `json:"legendPosition,omitempty\"` // The position of the chart legend.

	OutlierPercentile float64 `json:"outlierPercentile,omitempty\"` // The outlier percentile is used to ensure that outliers do not adversely affect the calculation of bucket sizes. For example, setting an outlier percentile of 0.05 indicates that the top and bottom 5% of values when calculating buckets. The values are still included in the chart, they will be added to the first or last buckets instead of their own buckets. Must be between 0.0 and 0.5.

	Series []HistogramSeries `json:"series,omitempty\"` // The series for a histogram may be either a single series of values to be bucketed or multiple series, each of the same length, containing the name of the series followed by the values to be bucketed for that series.

	ShowItemDividers bool `json:"showItemDividers,omitempty\"` // Whether horizontal divider lines should be displayed between items in each column.

}

// Allows you to organize the numeric values in a source data column into buckets of a constant size. All values from HistogramRule.start to HistogramRule.end are placed into groups of size HistogramRule.interval. In addition, all values below HistogramRule.start are placed in one group, and all values above HistogramRule.end are placed in another. Only HistogramRule.interval is required, though if HistogramRule.start and HistogramRule.end are both provided, HistogramRule.start must be less than HistogramRule.end. For example, a pivot table showing average purchase amount by age that has 50+ rows: +-----+-------------------+ | Age | AVERAGE of Amount | +-----+-------------------+ | 16 | $27.13 | | 17 | $5.24 | | 18 | $20.15 | ... +-----+-------------------+ could be turned into a pivot table that looks like the one below by applying a histogram group rule with a HistogramRule.start of 25, an HistogramRule.interval of 20, and an HistogramRule.end of 65. +-------------+-------------------+ | Grouped Age | AVERAGE of Amount | +-------------+-------------------+ | < 25 | $19.34 | | 25-45 | $31.43 | | 45-65 | $35.87 | | > 65 | $27.55 | +-------------+-------------------+ | Grand Total | $29.12 | +-------------+-------------------+
type HistogramRule struct {
	End float64 `json:"end,omitempty\"` // The maximum value at which items are placed into buckets of constant size. Values above end are lumped into a single bucket. This field is optional.

	Interval float64 `json:"interval,omitempty\"` // The size of the buckets that are created. Must be positive.

	Start float64 `json:"start,omitempty\"` // The minimum value at which items are placed into buckets of constant size. Values below start are lumped into a single bucket. This field is optional.

}

// A histogram series containing the series color and data.
type HistogramSeries struct {
	BarColor Color `json:"barColor,omitempty\"` // The color of the column representing this series in each bucket. This field is optional. Deprecated: Use bar_color_style.

	BarColorStyle ColorStyle `json:"barColorStyle,omitempty\"` // The color of the column representing this series in each bucket. This field is optional. If bar_color is also set, this field takes precedence.

	Data ChartData `json:"data,omitempty\"` // The data for this histogram series.

}

// Inserts rows or columns in a sheet at a particular index.
type InsertDimensionRequest struct {
	InheritFromBefore bool `json:"inheritFromBefore,omitempty\"` // Whether dimension properties should be extended from the dimensions before or after the newly inserted dimensions. True to inherit from the dimensions before (in which case the start index must be greater than 0), and false to inherit from the dimensions after. For example, if row index 0 has red background and row index 1 has a green background, then inserting 2 rows at index 1 can inherit either the green or red background. If `inheritFromBefore` is true, the two new rows will be red (because the row before the insertion point was red), whereas if `inheritFromBefore` is false, the two new rows will be green (because the row after the insertion point was green).

	RangeValue DimensionRange `json:"range,omitempty\"` // The dimensions to insert. Both the start and end indexes must be bounded.

}

// Inserts cells into a range, shifting the existing cells over or down.
type InsertRangeRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range to insert new cells into. The range is constrained to the current sheet boundaries.

	ShiftDimension string `json:"shiftDimension,omitempty\"` // The dimension which will be shifted when inserting cells. If ROWS, existing cells will be shifted down. If COLUMNS, existing cells will be shifted right.

}

// A single interpolation point on a gradient conditional format. These pin the gradient color scale according to the color, type and value chosen.
type InterpolationPoint struct {
	Color Color `json:"color,omitempty\"` // The color this interpolation point should use. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // The color this interpolation point should use. If color is also set, this field takes precedence.

	TypeValue string `json:"type,omitempty\"` // How the value should be interpreted.

	Value string `json:"value,omitempty\"` // The value this interpolation point uses. May be a formula. Unused if type is MIN or MAX.

}

// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start and end are unspecified, the interval matches any time.
type Interval struct {
	EndTime string `json:"endTime,omitempty\"` // Optional. Exclusive end of the interval. If specified, a Timestamp matching this interval will have to be before the end.

	StartTime string `json:"startTime,omitempty\"` // Optional. Inclusive start of the interval. If specified, a Timestamp matching this interval will have to be the same or after the start.

}

// Settings to control how circular dependencies are resolved with iterative calculation.
type IterativeCalculationSettings struct {
	ConvergenceThreshold float64 `json:"convergenceThreshold,omitempty\"` // When iterative calculation is enabled and successive results differ by less than this threshold value, the calculation rounds stop.

	MaxIterations int `json:"maxIterations,omitempty\"` // When iterative calculation is enabled, the maximum number of calculation rounds to perform.

}

// Formatting options for key value.
type KeyValueFormat struct {
	Position TextPosition `json:"position,omitempty\"` // Specifies the horizontal text positioning of key value. This field is optional. If not specified, default positioning is used.

	TextFormat TextFormat `json:"textFormat,omitempty\"` // Text formatting options for key value. The link field is not supported.

}

// Properties that describe the style of a line.
type LineStyle struct {
	TypeValue string `json:"type,omitempty\"` // The dash type of the line.

	Width int `json:"width,omitempty\"` // The thickness of the line, in px.

}

// An external or local reference.
type Link struct {
	Uri string `json:"uri,omitempty\"` // The link identifier.

}

// The specification of a Looker data source.
type LookerDataSourceSpec struct {
	Explore string `json:"explore,omitempty\"` // Name of a Looker model explore.

	InstanceUri string `json:"instanceUri,omitempty\"` // A Looker instance URL.

	Model string `json:"model,omitempty\"` // Name of a Looker model.

}

// Allows you to manually organize the values in a source data column into buckets with names of your choosing. For example, a pivot table that aggregates population by state: +-------+-------------------+ | State | SUM of Population | +-------+-------------------+ | AK | 0.7 | | AL | 4.8 | | AR | 2.9 | ... +-------+-------------------+ could be turned into a pivot table that aggregates population by time zone by providing a list of groups (for example, groupName = 'Central', items = ['AL', 'AR', 'IA', ...]) to a manual group rule. Note that a similar effect could be achieved by adding a time zone column to the source data and adjusting the pivot table. +-----------+-------------------+ | Time Zone | SUM of Population | +-----------+-------------------+ | Central | 106.3 | | Eastern | 151.9 | | Mountain | 17.4 | ... +-----------+-------------------+
type ManualRule struct {
	Groups []ManualRuleGroup `json:"groups,omitempty\"` // The list of group names and the corresponding items from the source data that map to each group name.

}

// A group name and a list of items from the source data that should be placed in the group with this name.
type ManualRuleGroup struct {
	GroupName ExtendedValue `json:"groupName,omitempty\"` // The group name, which must be a string. Each group in a given ManualRule must have a unique group name.

	Items []ExtendedValue `json:"items,omitempty\"` // The items in the source data that should be placed into this group. Each item may be a string, number, or boolean. Items may appear in at most one group within a given ManualRule. Items that do not appear in any group will appear on their own.

}

// A developer metadata entry and the data filters specified in the original request that matched it.
type MatchedDeveloperMetadata struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // All filters matching the returned developer metadata.

	DeveloperMetadata DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata matching the specified filters.

}

// A value range that was matched by one or more data filers.
type MatchedValueRange struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The DataFilters from the request that matched the range of values.

	ValueRange ValueRange `json:"valueRange,omitempty\"` // The values matched by the DataFilter.

}

// Merges all cells in the range.
type MergeCellsRequest struct {
	MergeType string `json:"mergeType,omitempty\"` // How the cells should be merged.

	RangeValue GridRange `json:"range,omitempty\"` // The range of cells to merge.

}

// Moves one or more rows or columns.
type MoveDimensionRequest struct {
	DestinationIndex int `json:"destinationIndex,omitempty\"` // The zero-based start index of where to move the source data to, based on the coordinates *before* the source data is removed from the grid. Existing data will be shifted down or right (depending on the dimension) to make room for the moved dimensions. The source dimensions are removed from the grid, so the the data may end up in a different index than specified. For example, given `A1..A5` of `0, 1, 2, 3, 4` and wanting to move `"1"` and `"2"` to between `"3"` and `"4"`, the source would be `ROWS [1..3)`,and the destination index would be `"4"` (the zero-based index of row 5). The end result would be `A1..A5` of `0, 3, 1, 2, 4`.

	Source DimensionRange `json:"source,omitempty\"` // The source dimensions to move.

}

// A named range.
type NamedRange struct {
	Name string `json:"name,omitempty\"` // The name of the named range.

	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the named range.

	RangeValue GridRange `json:"range,omitempty\"` // The range this represents.

}

// The number format of a cell.
type NumberFormat struct {
	Pattern string `json:"pattern,omitempty\"` // Pattern string used for formatting. If not set, a default pattern based on the spreadsheet's locale will be used if necessary for the given type. See the [Date and Number Formats guide](https://developers.google.com/workspace/sheets/api/guides/formats) for more information about the supported patterns.

	TypeValue string `json:"type,omitempty\"` // The type of the number format. When writing, this field must be set.

}

// An org chart. Org charts require a unique set of labels in labels and may optionally include parent_labels and tooltips. parent_labels contain, for each node, the label identifying the parent node. tooltips contain, for each node, an optional tooltip. For example, to describe an OrgChart with Alice as the CEO, Bob as the President (reporting to Alice) and Cathy as VP of Sales (also reporting to Alice), have labels contain "Alice", "Bob", "Cathy", parent_labels contain "", "Alice", "Alice" and tooltips contain "CEO", "President", "VP Sales".
type OrgChartSpec struct {
	Labels ChartData `json:"labels,omitempty\"` // The data containing the labels for all the nodes in the chart. Labels must be unique.

	NodeColor Color `json:"nodeColor,omitempty\"` // The color of the org chart nodes. Deprecated: Use node_color_style.

	NodeColorStyle ColorStyle `json:"nodeColorStyle,omitempty\"` // The color of the org chart nodes. If node_color is also set, this field takes precedence.

	NodeSize string `json:"nodeSize,omitempty\"` // The size of the org chart nodes.

	ParentLabels ChartData `json:"parentLabels,omitempty\"` // The data containing the label of the parent for the corresponding node. A blank value indicates that the node has no parent and is a top-level node. This field is optional.

	SelectedNodeColor Color `json:"selectedNodeColor,omitempty\"` // The color of the selected org chart nodes. Deprecated: Use selected_node_color_style.

	SelectedNodeColorStyle ColorStyle `json:"selectedNodeColorStyle,omitempty\"` // The color of the selected org chart nodes. If selected_node_color is also set, this field takes precedence.

	Tooltips ChartData `json:"tooltips,omitempty\"` // The data containing the tooltip for the corresponding node. A blank value results in no tooltip being displayed for the node. This field is optional.

}

// The location an object is overlaid on top of a grid.
type OverlayPosition struct {
	AnchorCell GridCoordinate `json:"anchorCell,omitempty\"` // The cell the object is anchored to.

	HeightPixels int `json:"heightPixels,omitempty\"` // The height of the object, in pixels. Defaults to 371.

	OffsetXPixels int `json:"offsetXPixels,omitempty\"` // The horizontal offset, in pixels, that the object is offset from the anchor cell.

	OffsetYPixels int `json:"offsetYPixels,omitempty\"` // The vertical offset, in pixels, that the object is offset from the anchor cell.

	WidthPixels int `json:"widthPixels,omitempty\"` // The width of the object, in pixels. Defaults to 600.

}

// The amount of padding around the cell, in pixels. When updating padding, every field must be specified.
type Padding struct {
	Bottom int `json:"bottom,omitempty\"` // The bottom padding of the cell.

	Left int `json:"left,omitempty\"` // The left padding of the cell.

	Right int `json:"right,omitempty\"` // The right padding of the cell.

	Top int `json:"top,omitempty\"` // The top padding of the cell.

}

// Inserts data into the spreadsheet starting at the specified coordinate.
type PasteDataRequest struct {
	Coordinate GridCoordinate `json:"coordinate,omitempty\"` // The coordinate at which the data should start being inserted.

	Data string `json:"data,omitempty\"` // The data to insert.

	Delimiter string `json:"delimiter,omitempty\"` // The delimiter in the data.

	Html bool `json:"html,omitempty\"` // True if the data is HTML.

	TypeValue string `json:"type,omitempty\"` // How the data should be pasted.

}

// Properties specific to a linked person.
type PersonProperties struct {
	DisplayFormat string `json:"displayFormat,omitempty\"` // Optional. The display format of the person chip. If not set, the default display format is used.

	Email string `json:"email,omitempty\"` // Required. The email address linked to this person. This field is always present.

}

// A pie chart.
type PieChartSpec struct {
	Domain ChartData `json:"domain,omitempty\"` // The data that covers the domain of the pie chart.

	LegendPosition string `json:"legendPosition,omitempty\"` // Where the legend of the pie chart should be drawn.

	PieHole float64 `json:"pieHole,omitempty\"` // The size of the hole in the pie chart.

	Series ChartData `json:"series,omitempty\"` // The data that covers the one and only series of the pie chart.

	ThreeDimensional bool `json:"threeDimensional,omitempty\"` // True if the pie is three dimensional.

}

// Criteria for showing/hiding rows in a pivot table.
type PivotFilterCriteria struct {
	Condition BooleanCondition `json:"condition,omitempty\"` // A condition that must be true for values to be shown. (`visibleValues` does not override this -- even if a value is listed there, it is still hidden if it does not meet the condition.) Condition values that refer to ranges in A1-notation are evaluated relative to the pivot table sheet. References are treated absolutely, so are not filled down the pivot table. For example, a condition value of `=A1` on "Pivot Table 1" is treated as `'Pivot Table 1'!$A$1`. The source data of the pivot table can be referenced by column header name. For example, if the source data has columns named "Revenue" and "Cost" and a condition is applied to the "Revenue" column with type `NUMBER_GREATER` and value `=Cost`, then only columns where "Revenue" > "Cost" are included.

	VisibleByDefault bool `json:"visibleByDefault,omitempty\"` // Whether values are visible by default. If true, the visible_values are ignored, all values that meet condition (if specified) are shown. If false, values that are both in visible_values and meet condition are shown.

	VisibleValues []string `json:"visibleValues,omitempty\"` // Values that should be included. Values not listed here are excluded.

}

// The pivot table filter criteria associated with a specific source column offset.
type PivotFilterSpec struct {
	ColumnOffsetIndex int `json:"columnOffsetIndex,omitempty\"` // The zero-based column offset of the source range.

	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // The reference to the data source column.

	FilterCriteria PivotFilterCriteria `json:"filterCriteria,omitempty\"` // The criteria for the column.

}

// A single grouping (either row or column) in a pivot table.
type PivotGroup struct {
	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // The reference to the data source column this grouping is based on.

	GroupLimit PivotGroupLimit `json:"groupLimit,omitempty\"` // The count limit on rows or columns to apply to this pivot group.

	GroupRule PivotGroupRule `json:"groupRule,omitempty\"` // The group rule to apply to this row/column group.

	Label string `json:"label,omitempty\"` // The labels to use for the row/column groups which can be customized. For example, in the following pivot table, the row label is `Region` (which could be renamed to `State`) and the column label is `Product` (which could be renamed `Item`). Pivot tables created before December 2017 do not have header labels. If you'd like to add header labels to an existing pivot table, please delete the existing pivot table and then create a new pivot table with same parameters. +--------------+---------+-------+ | SUM of Units | Product | | | Region | Pen | Paper | +--------------+---------+-------+ | New York | 345 | 98 | | Oregon | 234 | 123 | | Tennessee | 531 | 415 | +--------------+---------+-------+ | Grand Total | 1110 | 636 | +--------------+---------+-------+

	RepeatHeadings bool `json:"repeatHeadings,omitempty\"` // True if the headings in this pivot group should be repeated. This is only valid for row groupings and is ignored by columns. By default, we minimize repetition of headings by not showing higher level headings where they are the same. For example, even though the third row below corresponds to "Q1 Mar", "Q1" is not shown because it is redundant with previous rows. Setting repeat_headings to true would cause "Q1" to be repeated for "Feb" and "Mar". +--------------+ | Q1 | Jan | | | Feb | | | Mar | +--------+-----+ | Q1 Total | +--------------+

	ShowTotals bool `json:"showTotals,omitempty\"` // True if the pivot table should include the totals for this grouping.

	SortOrder string `json:"sortOrder,omitempty\"` // The order the values in this group should be sorted.

	SourceColumnOffset int `json:"sourceColumnOffset,omitempty\"` // The column offset of the source range that this grouping is based on. For example, if the source was `C10:E15`, a `sourceColumnOffset` of `0` means this group refers to column `C`, whereas the offset `1` would refer to column `D`.

	ValueBucket PivotGroupSortValueBucket `json:"valueBucket,omitempty\"` // The bucket of the opposite pivot group to sort by. If not specified, sorting is alphabetical by this group's values.

	ValueMetadata []PivotGroupValueMetadata `json:"valueMetadata,omitempty\"` // Metadata about values in the grouping.

}

// The count limit on rows or columns in the pivot group.
type PivotGroupLimit struct {
	ApplyOrder int `json:"applyOrder,omitempty\"` // The order in which the group limit is applied to the pivot table. Pivot group limits are applied from lower to higher order number. Order numbers are normalized to consecutive integers from 0. For write request, to fully customize the applying orders, all pivot group limits should have this field set with an unique number. Otherwise, the order is determined by the index in the PivotTable.rows list and then the PivotTable.columns list.

	CountLimit int `json:"countLimit,omitempty\"` // The count limit.

}

// An optional setting on a PivotGroup that defines buckets for the values in the source data column rather than breaking out each individual value. Only one PivotGroup with a group rule may be added for each column in the source data, though on any given column you may add both a PivotGroup that has a rule and a PivotGroup that does not.
type PivotGroupRule struct {
	DateTimeRule DateTimeRule `json:"dateTimeRule,omitempty\"` // A DateTimeRule.

	HistogramRule HistogramRule `json:"histogramRule,omitempty\"` // A HistogramRule.

	ManualRule ManualRule `json:"manualRule,omitempty\"` // A ManualRule.

}

// Information about which values in a pivot group should be used for sorting.
type PivotGroupSortValueBucket struct {
	Buckets []ExtendedValue `json:"buckets,omitempty\"` // Determines the bucket from which values are chosen to sort. For example, in a pivot table with one row group & two column groups, the row group can list up to two values. The first value corresponds to a value within the first column group, and the second value corresponds to a value in the second column group. If no values are listed, this would indicate that the row should be sorted according to the "Grand Total" over the column groups. If a single value is listed, this would correspond to using the "Total" of that bucket.

	ValuesIndex int `json:"valuesIndex,omitempty\"` // The offset in the PivotTable.values list which the values in this grouping should be sorted by.

}

// Metadata about a value in a pivot grouping.
type PivotGroupValueMetadata struct {
	Collapsed bool `json:"collapsed,omitempty\"` // True if the data corresponding to the value is collapsed.

	Value ExtendedValue `json:"value,omitempty\"` // The calculated value the metadata corresponds to. (Note that formulaValue is not valid, because the values will be calculated.)

}

// A pivot table.
type PivotTable struct {
	Columns []PivotGroup `json:"columns,omitempty\"` // Each column grouping in the pivot table.

	Criteria map[string]interface{} `json:"criteria,omitempty\"` // An optional mapping of filters per source column offset. The filters are applied before aggregating data into the pivot table. The map's key is the column offset of the source range that you want to filter, and the value is the criteria for that column. For example, if the source was `C10:E15`, a key of `0` will have the filter for column `C`, whereas the key `1` is for column `D`. This field is deprecated in favor of filter_specs.

	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // Output only. The data execution status for data source pivot tables.

	DataSourceId string `json:"dataSourceId,omitempty\"` // The ID of the data source the pivot table is reading data from.

	FilterSpecs []PivotFilterSpec `json:"filterSpecs,omitempty\"` // The filters applied to the source columns before aggregating data for the pivot table. Both criteria and filter_specs are populated in responses. If both fields are specified in an update request, this field takes precedence.

	Rows []PivotGroup `json:"rows,omitempty\"` // Each row grouping in the pivot table.

	Source GridRange `json:"source,omitempty\"` // The range the pivot table is reading data from.

	ValueLayout string `json:"valueLayout,omitempty\"` // Whether values should be listed horizontally (as columns) or vertically (as rows).

	Values []PivotValue `json:"values,omitempty\"` // A list of values to include in the pivot table.

}

// The definition of how a value in a pivot table should be calculated.
type PivotValue struct {
	CalculatedDisplayType string `json:"calculatedDisplayType,omitempty\"` // If specified, indicates that pivot values should be displayed as the result of a calculation with another pivot value. For example, if calculated_display_type is specified as PERCENT_OF_GRAND_TOTAL, all the pivot values are displayed as the percentage of the grand total. In the Sheets editor, this is referred to as "Show As" in the value section of a pivot table.

	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // The reference to the data source column that this value reads from.

	Formula string `json:"formula,omitempty\"` // A custom formula to calculate the value. The formula must start with an `=` character.

	Name string `json:"name,omitempty\"` // A name to use for the value.

	SourceColumnOffset int `json:"sourceColumnOffset,omitempty\"` // The column offset of the source range that this value reads from. For example, if the source was `C10:E15`, a `sourceColumnOffset` of `0` means this value refers to column `C`, whereas the offset `1` would refer to column `D`.

	SummarizeFunction string `json:"summarizeFunction,omitempty\"` // A function to summarize the value. If formula is set, the only supported values are SUM and CUSTOM. If sourceColumnOffset is set, then `CUSTOM` is not supported.

}

// The style of a point on the chart.
type PointStyle struct {
	Shape string `json:"shape,omitempty\"` // The point shape. If empty or unspecified, a default shape is used.

	Size float64 `json:"size,omitempty\"` // The point size. If empty, a default size is used.

}

// A protected range.
type ProtectedRange struct {
	Description string `json:"description,omitempty\"` // The description of this protected range.

	Editors Editors `json:"editors,omitempty\"` // The users and groups with edit access to the protected range. This field is only visible to users with edit access to the protected range and the document. Editors are not supported with warning_only protection.

	NamedRangeId string `json:"namedRangeId,omitempty\"` // The named range this protected range is backed by, if any. When writing, only one of range or named_range_id or table_id may be set.

	ProtectedRangeId int `json:"protectedRangeId,omitempty\"` // The ID of the protected range. This field is read-only.

	RangeValue GridRange `json:"range,omitempty\"` // The range that is being protected. The range may be fully unbounded, in which case this is considered a protected sheet. When writing, only one of range or named_range_id or table_id may be set.

	RequestingUserCanEdit bool `json:"requestingUserCanEdit,omitempty\"` // True if the user who requested this protected range can edit the protected area. This field is read-only.

	TableId string `json:"tableId,omitempty\"` // The table this protected range is backed by, if any. When writing, only one of range or named_range_id or table_id may be set.

	UnprotectedRanges []GridRange `json:"unprotectedRanges,omitempty\"` // The list of unprotected ranges within a protected sheet. Unprotected ranges are only supported on protected sheets.

	WarningOnly bool `json:"warningOnly,omitempty\"` // True if this protected range will show a warning when editing. Warning-based protection means that every user can edit data in the protected range, except editing will prompt a warning asking the user to confirm the edit. When writing: if this field is true, then editors are ignored. Additionally, if this field is changed from true to false and the `editors` field is not set (nor included in the field mask), then the editors will be set to all the editors in the document.

}

// Randomizes the order of the rows in a range.
type RandomizeRangeRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range to randomize.

}

// The status of a refresh cancellation. You can send a cancel request to explicitly cancel one or multiple data source object refreshes.
type RefreshCancellationStatus struct {
	ErrorCode string `json:"errorCode,omitempty\"` // The error code.

	State string `json:"state,omitempty\"` // The state of a call to cancel a refresh in Sheets.

}

// The execution status of refreshing one data source object.
type RefreshDataSourceObjectExecutionStatus struct {
	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // The data execution status.

	Reference DataSourceObjectReference `json:"reference,omitempty\"` // Reference to a data source object being refreshed.

}

// Refreshes one or multiple data source objects in the spreadsheet by the specified references. The request requires an additional `bigquery.readonly` OAuth scope if you are refreshing a BigQuery data source. If there are multiple refresh requests referencing the same data source objects in one batch, only the last refresh request is processed, and all those requests will have the same response accordingly.
type RefreshDataSourceRequest struct {
	DataSourceId string `json:"dataSourceId,omitempty\"` // Reference to a DataSource. If specified, refreshes all associated data source objects for the data source.

	Force bool `json:"force,omitempty\"` // Refreshes the data source objects regardless of the current state. If not set and a referenced data source object was in error state, the refresh will fail immediately.

	IsAll bool `json:"isAll,omitempty\"` // Refreshes all existing data source objects in the spreadsheet.

	References DataSourceObjectReferences `json:"references,omitempty\"` // References to data source objects to refresh.

}

// The response from refreshing one or multiple data source objects.
type RefreshDataSourceResponse struct {
	Statuses []RefreshDataSourceObjectExecutionStatus `json:"statuses,omitempty\"` // All the refresh status for the data source object references specified in the request. If is_all is specified, the field contains only those in failure status.

}

// Updates all cells in the range to the values in the given Cell object. Only the fields listed in the fields field are updated; others are unchanged. If writing a cell with a formula, the formula's ranges will automatically increment for each field in the range. For example, if writing a cell with formula `=A1` into range B2:C4, B2 would be `=A1`, B3 would be `=A2`, B4 would be `=A3`, C2 would be `=B1`, C3 would be `=B2`, C4 would be `=B3`. To keep the formula's ranges static, use the `$` indicator. For example, use the formula `=$A$1` to prevent both the row and the column from incrementing.
type RepeatCellRequest struct {
	Cell CellData `json:"cell,omitempty\"` // The data to write.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `cell` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	RangeValue GridRange `json:"range,omitempty\"` // The range to repeat the cell in.

}

// A single kind of update to apply to a spreadsheet.
type Request struct {
	AddBanding AddBandingRequest `json:"addBanding,omitempty\"` // Adds a new banded range

	AddChart AddChartRequest `json:"addChart,omitempty\"` // Adds a chart.

	AddConditionalFormatRule AddConditionalFormatRuleRequest `json:"addConditionalFormatRule,omitempty\"` // Adds a new conditional format rule.

	AddDataSource AddDataSourceRequest `json:"addDataSource,omitempty\"` // Adds a data source.

	AddDimensionGroup AddDimensionGroupRequest `json:"addDimensionGroup,omitempty\"` // Creates a group over the specified range.

	AddFilterView AddFilterViewRequest `json:"addFilterView,omitempty\"` // Adds a filter view.

	AddNamedRange AddNamedRangeRequest `json:"addNamedRange,omitempty\"` // Adds a named range.

	AddProtectedRange AddProtectedRangeRequest `json:"addProtectedRange,omitempty\"` // Adds a protected range.

	AddSheet AddSheetRequest `json:"addSheet,omitempty\"` // Adds a sheet.

	AddSlicer AddSlicerRequest `json:"addSlicer,omitempty\"` // Adds a slicer.

	AddTable AddTableRequest `json:"addTable,omitempty\"` // Adds a table.

	AppendCells AppendCellsRequest `json:"appendCells,omitempty\"` // Appends cells after the last row with data in a sheet.

	AppendDimension AppendDimensionRequest `json:"appendDimension,omitempty\"` // Appends dimensions to the end of a sheet.

	AutoFill AutoFillRequest `json:"autoFill,omitempty\"` // Automatically fills in more data based on existing data.

	AutoResizeDimensions AutoResizeDimensionsRequest `json:"autoResizeDimensions,omitempty\"` // Automatically resizes one or more dimensions based on the contents of the cells in that dimension.

	CancelDataSourceRefresh CancelDataSourceRefreshRequest `json:"cancelDataSourceRefresh,omitempty\"` // Cancels refreshes of one or multiple data sources and associated dbobjects.

	ClearBasicFilter ClearBasicFilterRequest `json:"clearBasicFilter,omitempty\"` // Clears the basic filter on a sheet.

	CopyPaste CopyPasteRequest `json:"copyPaste,omitempty\"` // Copies data from one area and pastes it to another.

	CreateDeveloperMetadata CreateDeveloperMetadataRequest `json:"createDeveloperMetadata,omitempty\"` // Creates new developer metadata

	CutPaste CutPasteRequest `json:"cutPaste,omitempty\"` // Cuts data from one area and pastes it to another.

	DeleteBanding DeleteBandingRequest `json:"deleteBanding,omitempty\"` // Removes a banded range

	DeleteConditionalFormatRule DeleteConditionalFormatRuleRequest `json:"deleteConditionalFormatRule,omitempty\"` // Deletes an existing conditional format rule.

	DeleteDataSource DeleteDataSourceRequest `json:"deleteDataSource,omitempty\"` // Deletes a data source.

	DeleteDeveloperMetadata DeleteDeveloperMetadataRequest `json:"deleteDeveloperMetadata,omitempty\"` // Deletes developer metadata

	DeleteDimension DeleteDimensionRequest `json:"deleteDimension,omitempty\"` // Deletes rows or columns in a sheet.

	DeleteDimensionGroup DeleteDimensionGroupRequest `json:"deleteDimensionGroup,omitempty\"` // Deletes a group over the specified range.

	DeleteDuplicates DeleteDuplicatesRequest `json:"deleteDuplicates,omitempty\"` // Removes rows containing duplicate values in specified columns of a cell range.

	DeleteEmbeddedObject DeleteEmbeddedObjectRequest `json:"deleteEmbeddedObject,omitempty\"` // Deletes an embedded object (e.g, chart, image) in a sheet.

	DeleteFilterView DeleteFilterViewRequest `json:"deleteFilterView,omitempty\"` // Deletes a filter view from a sheet.

	DeleteNamedRange DeleteNamedRangeRequest `json:"deleteNamedRange,omitempty\"` // Deletes a named range.

	DeleteProtectedRange DeleteProtectedRangeRequest `json:"deleteProtectedRange,omitempty\"` // Deletes a protected range.

	DeleteRange DeleteRangeRequest `json:"deleteRange,omitempty\"` // Deletes a range of cells from a sheet, shifting the remaining cells.

	DeleteSheet DeleteSheetRequest `json:"deleteSheet,omitempty\"` // Deletes a sheet.

	DeleteTable DeleteTableRequest `json:"deleteTable,omitempty\"` // A request for deleting a table.

	DuplicateFilterView DuplicateFilterViewRequest `json:"duplicateFilterView,omitempty\"` // Duplicates a filter view.

	DuplicateSheet DuplicateSheetRequest `json:"duplicateSheet,omitempty\"` // Duplicates a sheet.

	FindReplace FindReplaceRequest `json:"findReplace,omitempty\"` // Finds and replaces occurrences of some text with other text.

	InsertDimension InsertDimensionRequest `json:"insertDimension,omitempty\"` // Inserts new rows or columns in a sheet.

	InsertRange InsertRangeRequest `json:"insertRange,omitempty\"` // Inserts new cells in a sheet, shifting the existing cells.

	MergeCells MergeCellsRequest `json:"mergeCells,omitempty\"` // Merges cells together.

	MoveDimension MoveDimensionRequest `json:"moveDimension,omitempty\"` // Moves rows or columns to another location in a sheet.

	PasteData PasteDataRequest `json:"pasteData,omitempty\"` // Pastes data (HTML or delimited) into a sheet.

	RandomizeRange RandomizeRangeRequest `json:"randomizeRange,omitempty\"` // Randomizes the order of the rows in a range.

	RefreshDataSource RefreshDataSourceRequest `json:"refreshDataSource,omitempty\"` // Refreshes one or multiple data sources and associated dbobjects.

	RepeatCell RepeatCellRequest `json:"repeatCell,omitempty\"` // Repeats a single cell across a range.

	SetBasicFilter SetBasicFilterRequest `json:"setBasicFilter,omitempty\"` // Sets the basic filter on a sheet.

	SetDataValidation SetDataValidationRequest `json:"setDataValidation,omitempty\"` // Sets data validation for one or more cells.

	SortRange SortRangeRequest `json:"sortRange,omitempty\"` // Sorts data in a range.

	TextToColumns TextToColumnsRequest `json:"textToColumns,omitempty\"` // Converts a column of text into many columns of text.

	TrimWhitespace TrimWhitespaceRequest `json:"trimWhitespace,omitempty\"` // Trims cells of whitespace (such as spaces, tabs, or new lines).

	UnmergeCells UnmergeCellsRequest `json:"unmergeCells,omitempty\"` // Unmerges merged cells.

	UpdateBanding UpdateBandingRequest `json:"updateBanding,omitempty\"` // Updates a banded range

	UpdateBorders UpdateBordersRequest `json:"updateBorders,omitempty\"` // Updates the borders in a range of cells.

	UpdateCells UpdateCellsRequest `json:"updateCells,omitempty\"` // Updates many cells at once.

	UpdateChartSpec UpdateChartSpecRequest `json:"updateChartSpec,omitempty\"` // Updates a chart's specifications.

	UpdateConditionalFormatRule UpdateConditionalFormatRuleRequest `json:"updateConditionalFormatRule,omitempty\"` // Updates an existing conditional format rule.

	UpdateDataSource UpdateDataSourceRequest `json:"updateDataSource,omitempty\"` // Updates a data source.

	UpdateDeveloperMetadata UpdateDeveloperMetadataRequest `json:"updateDeveloperMetadata,omitempty\"` // Updates an existing developer metadata entry

	UpdateDimensionGroup UpdateDimensionGroupRequest `json:"updateDimensionGroup,omitempty\"` // Updates the state of the specified group.

	UpdateDimensionProperties UpdateDimensionPropertiesRequest `json:"updateDimensionProperties,omitempty\"` // Updates dimensions' properties.

	UpdateEmbeddedObjectBorder UpdateEmbeddedObjectBorderRequest `json:"updateEmbeddedObjectBorder,omitempty\"` // Updates an embedded object's border.

	UpdateEmbeddedObjectPosition UpdateEmbeddedObjectPositionRequest `json:"updateEmbeddedObjectPosition,omitempty\"` // Updates an embedded object's (e.g. chart, image) position.

	UpdateFilterView UpdateFilterViewRequest `json:"updateFilterView,omitempty\"` // Updates the properties of a filter view.

	UpdateNamedRange UpdateNamedRangeRequest `json:"updateNamedRange,omitempty\"` // Updates a named range.

	UpdateProtectedRange UpdateProtectedRangeRequest `json:"updateProtectedRange,omitempty\"` // Updates a protected range.

	UpdateSheetProperties UpdateSheetPropertiesRequest `json:"updateSheetProperties,omitempty\"` // Updates a sheet's properties.

	UpdateSlicerSpec UpdateSlicerSpecRequest `json:"updateSlicerSpec,omitempty\"` // Updates a slicer's specifications.

	UpdateSpreadsheetProperties UpdateSpreadsheetPropertiesRequest `json:"updateSpreadsheetProperties,omitempty\"` // Updates the spreadsheet's properties.

	UpdateTable UpdateTableRequest `json:"updateTable,omitempty\"` // Updates a table.

}

// A single response from an update.
type Response struct {
	AddBanding AddBandingResponse `json:"addBanding,omitempty\"` // A reply from adding a banded range.

	AddChart AddChartResponse `json:"addChart,omitempty\"` // A reply from adding a chart.

	AddDataSource AddDataSourceResponse `json:"addDataSource,omitempty\"` // A reply from adding a data source.

	AddDimensionGroup AddDimensionGroupResponse `json:"addDimensionGroup,omitempty\"` // A reply from adding a dimension group.

	AddFilterView AddFilterViewResponse `json:"addFilterView,omitempty\"` // A reply from adding a filter view.

	AddNamedRange AddNamedRangeResponse `json:"addNamedRange,omitempty\"` // A reply from adding a named range.

	AddProtectedRange AddProtectedRangeResponse `json:"addProtectedRange,omitempty\"` // A reply from adding a protected range.

	AddSheet AddSheetResponse `json:"addSheet,omitempty\"` // A reply from adding a sheet.

	AddSlicer AddSlicerResponse `json:"addSlicer,omitempty\"` // A reply from adding a slicer.

	AddTable AddTableResponse `json:"addTable,omitempty\"` // A reply from adding a table.

	CancelDataSourceRefresh CancelDataSourceRefreshResponse `json:"cancelDataSourceRefresh,omitempty\"` // A reply from cancelling data source object refreshes.

	CreateDeveloperMetadata CreateDeveloperMetadataResponse `json:"createDeveloperMetadata,omitempty\"` // A reply from creating a developer metadata entry.

	DeleteConditionalFormatRule DeleteConditionalFormatRuleResponse `json:"deleteConditionalFormatRule,omitempty\"` // A reply from deleting a conditional format rule.

	DeleteDeveloperMetadata DeleteDeveloperMetadataResponse `json:"deleteDeveloperMetadata,omitempty\"` // A reply from deleting a developer metadata entry.

	DeleteDimensionGroup DeleteDimensionGroupResponse `json:"deleteDimensionGroup,omitempty\"` // A reply from deleting a dimension group.

	DeleteDuplicates DeleteDuplicatesResponse `json:"deleteDuplicates,omitempty\"` // A reply from removing rows containing duplicate values.

	DuplicateFilterView DuplicateFilterViewResponse `json:"duplicateFilterView,omitempty\"` // A reply from duplicating a filter view.

	DuplicateSheet DuplicateSheetResponse `json:"duplicateSheet,omitempty\"` // A reply from duplicating a sheet.

	FindReplace FindReplaceResponse `json:"findReplace,omitempty\"` // A reply from doing a find/replace.

	RefreshDataSource RefreshDataSourceResponse `json:"refreshDataSource,omitempty\"` // A reply from refreshing data source objects.

	TrimWhitespace TrimWhitespaceResponse `json:"trimWhitespace,omitempty\"` // A reply from trimming whitespace.

	UpdateConditionalFormatRule UpdateConditionalFormatRuleResponse `json:"updateConditionalFormatRule,omitempty\"` // A reply from updating a conditional format rule.

	UpdateDataSource UpdateDataSourceResponse `json:"updateDataSource,omitempty\"` // A reply from updating a data source.

	UpdateDeveloperMetadata UpdateDeveloperMetadataResponse `json:"updateDeveloperMetadata,omitempty\"` // A reply from updating a developer metadata entry.

	UpdateEmbeddedObjectPosition UpdateEmbeddedObjectPositionResponse `json:"updateEmbeddedObjectPosition,omitempty\"` // A reply from updating an embedded object's position.

}

// Properties of a link to a Google resource (such as a file in Drive, a YouTube video, a Maps address, or a Calendar event). Only Drive files can be written as chips. All other rich link types are read only. URIs cannot exceed 2000 bytes when writing. NOTE: Writing Drive file chips requires at least one of the `drive.file`, `drive.readonly`, or `drive` OAuth scopes.
type RichLinkProperties struct {
	MimeType string `json:"mimeType,omitempty\"` // Output only. The [MIME type](https://developers.google.com/drive/api/v3/mime-types) of the link, if there's one (for example, when it's a file in Drive).

	Uri string `json:"uri,omitempty\"` // Required. The URI to the link. This is always present.

}

// Data about each cell in a row.
type RowData struct {
	Values []CellData `json:"values,omitempty\"` // The values in the row, one per column.

}

// A scorecard chart. Scorecard charts are used to highlight key performance indicators, known as KPIs, on the spreadsheet. A scorecard chart can represent things like total sales, average cost, or a top selling item. You can specify a single data value, or aggregate over a range of data. Percentage or absolute difference from a baseline value can be highlighted, like changes over time.
type ScorecardChartSpec struct {
	AggregateType string `json:"aggregateType,omitempty\"` // The aggregation type for key and baseline chart data in scorecard chart. This field is not supported for data source charts. Use the ChartData.aggregateType field of the key_value_data or baseline_value_data instead for data source charts. This field is optional.

	BaselineValueData ChartData `json:"baselineValueData,omitempty\"` // The data for scorecard baseline value. This field is optional.

	BaselineValueFormat BaselineValueFormat `json:"baselineValueFormat,omitempty\"` // Formatting options for baseline value. This field is needed only if baseline_value_data is specified.

	CustomFormatOptions ChartCustomNumberFormatOptions `json:"customFormatOptions,omitempty\"` // Custom formatting options for numeric key/baseline values in scorecard chart. This field is used only when number_format_source is set to CUSTOM. This field is optional.

	KeyValueData ChartData `json:"keyValueData,omitempty\"` // The data for scorecard key value.

	KeyValueFormat KeyValueFormat `json:"keyValueFormat,omitempty\"` // Formatting options for key value.

	NumberFormatSource string `json:"numberFormatSource,omitempty\"` // The number format source used in the scorecard chart. This field is optional.

	ScaleFactor float64 `json:"scaleFactor,omitempty\"` // Value to scale scorecard key and baseline value. For example, a factor of 10 can be used to divide all values in the chart by 10. This field is optional.

}

// A request to retrieve all developer metadata matching the set of specified criteria.
type SearchDeveloperMetadataRequest struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The data filters describing the criteria used to determine which DeveloperMetadata entries to return. DeveloperMetadata matching any of the specified filters are included in the response.

}

// A reply to a developer metadata search request.
type SearchDeveloperMetadataResponse struct {
	MatchedDeveloperMetadata []MatchedDeveloperMetadata `json:"matchedDeveloperMetadata,omitempty\"` // The metadata matching the criteria of the search request.

}

// Sets the basic filter associated with a sheet.
type SetBasicFilterRequest struct {
	Filter BasicFilter `json:"filter,omitempty\"` // The filter to set.

}

// Sets a data validation rule to every cell in the range. To clear validation in a range, call this with no rule specified.
type SetDataValidationRequest struct {
	FilteredRowsIncluded bool `json:"filteredRowsIncluded,omitempty\"` // Optional. If true, the data validation rule will be applied to the filtered rows as well.

	RangeValue GridRange `json:"range,omitempty\"` // The range the data validation rule should apply to.

	Rule DataValidationRule `json:"rule,omitempty\"` // The data validation rule to set on each cell in the range, or empty to clear the data validation in the range.

}

// A sheet in a spreadsheet.
type Sheet struct {
	BandedRanges []BandedRange `json:"bandedRanges,omitempty\"` // The banded (alternating colors) ranges on this sheet.

	BasicFilter BasicFilter `json:"basicFilter,omitempty\"` // The filter on this sheet, if any.

	Charts []EmbeddedChart `json:"charts,omitempty\"` // The specifications of every chart on this sheet.

	ColumnGroups []DimensionGroup `json:"columnGroups,omitempty\"` // All column groups on this sheet, ordered by increasing range start index, then by group depth.

	ConditionalFormats []ConditionalFormatRule `json:"conditionalFormats,omitempty\"` // The conditional format rules in this sheet.

	Data []GridData `json:"data,omitempty\"` // Data in the grid, if this is a grid sheet. The number of GridData objects returned is dependent on the number of ranges requested on this sheet. For example, if this is representing `Sheet1`, and the spreadsheet was requested with ranges `Sheet1!A1:C10` and `Sheet1!D15:E20`, then the first GridData will have a startRow/startColumn of `0`, while the second one will have `startRow 14` (zero-based row 15), and `startColumn 3` (zero-based column D). For a DATA_SOURCE sheet, you can not request a specific range, the GridData contains all the values.

	DeveloperMetadata []DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata associated with a sheet.

	FilterViews []FilterView `json:"filterViews,omitempty\"` // The filter views in this sheet.

	Merges []GridRange `json:"merges,omitempty\"` // The ranges that are merged together.

	Properties SheetProperties `json:"properties,omitempty\"` // The properties of the sheet.

	ProtectedRanges []ProtectedRange `json:"protectedRanges,omitempty\"` // The protected ranges in this sheet.

	RowGroups []DimensionGroup `json:"rowGroups,omitempty\"` // All row groups on this sheet, ordered by increasing range start index, then by group depth.

	Slicers []Slicer `json:"slicers,omitempty\"` // The slicers on this sheet.

	Tables []Table `json:"tables,omitempty\"` // The tables on this sheet.

}

// Properties of a sheet.
type SheetProperties struct {
	DataSourceSheetProperties DataSourceSheetProperties `json:"dataSourceSheetProperties,omitempty\"` // Output only. If present, the field contains DATA_SOURCE sheet specific properties.

	GridProperties GridProperties `json:"gridProperties,omitempty\"` // Additional properties of the sheet if this sheet is a grid. (If the sheet is an object sheet, containing a chart or image, then this field will be absent.) When writing it is an error to set any grid properties on non-grid sheets. If this sheet is a DATA_SOURCE sheet, this field is output only but contains the properties that reflect how a data source sheet is rendered in the UI, e.g. row_count.

	Hidden bool `json:"hidden,omitempty\"` // True if the sheet is hidden in the UI, false if it's visible.

	Index int `json:"index,omitempty\"` // The index of the sheet within the spreadsheet. When adding or updating sheet properties, if this field is excluded then the sheet is added or moved to the end of the sheet list. When updating sheet indices or inserting sheets, movement is considered in "before the move" indexes. For example, if there were three sheets (S1, S2, S3) in order to move S1 ahead of S2 the index would have to be set to 2. A sheet index update request is ignored if the requested index is identical to the sheets current index or if the requested new index is equal to the current sheet index + 1.

	RightToLeft bool `json:"rightToLeft,omitempty\"` // True if the sheet is an RTL sheet instead of an LTR sheet.

	SheetId int `json:"sheetId,omitempty\"` // The ID of the sheet. Must be non-negative. This field cannot be changed once set.

	SheetType string `json:"sheetType,omitempty\"` // The type of sheet. Defaults to GRID. This field cannot be changed once set.

	TabColor Color `json:"tabColor,omitempty\"` // The color of the tab in the UI. Deprecated: Use tab_color_style.

	TabColorStyle ColorStyle `json:"tabColorStyle,omitempty\"` // The color of the tab in the UI. If tab_color is also set, this field takes precedence.

	Title string `json:"title,omitempty\"` // The name of the sheet.

}

// A slicer in a sheet.
type Slicer struct {
	Position EmbeddedObjectPosition `json:"position,omitempty\"` // The position of the slicer. Note that slicer can be positioned only on existing sheet. Also, width and height of slicer can be automatically adjusted to keep it within permitted limits.

	SlicerId int `json:"slicerId,omitempty\"` // The ID of the slicer.

	Spec SlicerSpec `json:"spec,omitempty\"` // The specification of the slicer.

}

// The specifications of a slicer.
type SlicerSpec struct {
	ApplyToPivotTables bool `json:"applyToPivotTables,omitempty\"` // True if the filter should apply to pivot tables. If not set, default to `True`.

	BackgroundColor Color `json:"backgroundColor,omitempty\"` // The background color of the slicer. Deprecated: Use background_color_style.

	BackgroundColorStyle ColorStyle `json:"backgroundColorStyle,omitempty\"` // The background color of the slicer. If background_color is also set, this field takes precedence.

	ColumnIndex int `json:"columnIndex,omitempty\"` // The zero-based column index in the data table on which the filter is applied to.

	DataRange GridRange `json:"dataRange,omitempty\"` // The data range of the slicer.

	FilterCriteria FilterCriteria `json:"filterCriteria,omitempty\"` // The filtering criteria of the slicer.

	HorizontalAlignment string `json:"horizontalAlignment,omitempty\"` // The horizontal alignment of title in the slicer. If unspecified, defaults to `LEFT`

	TextFormat TextFormat `json:"textFormat,omitempty\"` // The text format of title in the slicer. The link field is not supported.

	Title string `json:"title,omitempty\"` // The title of the slicer.

}

// Sorts data in rows based on a sort order per column.
type SortRangeRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range to sort.

	SortSpecs []SortSpec `json:"sortSpecs,omitempty\"` // The sort order per column. Later specifications are used when values are equal in the earlier specifications.

}

// A sort order associated with a specific column or row.
type SortSpec struct {
	BackgroundColor Color `json:"backgroundColor,omitempty\"` // The background fill color to sort by; cells with this fill color are sorted to the top. Mutually exclusive with foreground_color. Deprecated: Use background_color_style.

	BackgroundColorStyle ColorStyle `json:"backgroundColorStyle,omitempty\"` // The background fill color to sort by; cells with this fill color are sorted to the top. Mutually exclusive with foreground_color, and must be an RGB-type color. If background_color is also set, this field takes precedence.

	DataSourceColumnReference DataSourceColumnReference `json:"dataSourceColumnReference,omitempty\"` // Reference to a data source column.

	DimensionIndex int `json:"dimensionIndex,omitempty\"` // The dimension the sort should be applied to.

	ForegroundColor Color `json:"foregroundColor,omitempty\"` // The foreground color to sort by; cells with this foreground color are sorted to the top. Mutually exclusive with background_color. Deprecated: Use foreground_color_style.

	ForegroundColorStyle ColorStyle `json:"foregroundColorStyle,omitempty\"` // The foreground color to sort by; cells with this foreground color are sorted to the top. Mutually exclusive with background_color, and must be an RGB-type color. If foreground_color is also set, this field takes precedence.

	SortOrder string `json:"sortOrder,omitempty\"` // The order data should be sorted.

}

// A combination of a source range and how to extend that source.
type SourceAndDestination struct {
	Dimension string `json:"dimension,omitempty\"` // The dimension that data should be filled into.

	FillLength int `json:"fillLength,omitempty\"` // The number of rows or columns that data should be filled into. Positive numbers expand beyond the last row or last column of the source. Negative numbers expand before the first row or first column of the source.

	Source GridRange `json:"source,omitempty\"` // The location of the data to use as the source of the autofill.

}

// Resource that represents a spreadsheet.
type Spreadsheet struct {
	DataSourceSchedules []DataSourceRefreshSchedule `json:"dataSourceSchedules,omitempty\"` // Output only. A list of data source refresh schedules.

	DataSources []DataSource `json:"dataSources,omitempty\"` // A list of external data sources connected with the spreadsheet.

	DeveloperMetadata []DeveloperMetadata `json:"developerMetadata,omitempty\"` // The developer metadata associated with a spreadsheet.

	NamedRanges []NamedRange `json:"namedRanges,omitempty\"` // The named ranges defined in a spreadsheet.

	Properties SpreadsheetProperties `json:"properties,omitempty\"` // Overall properties of a spreadsheet.

	Sheets []Sheet `json:"sheets,omitempty\"` // The sheets that are part of a spreadsheet.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the spreadsheet. This field is read-only.

	SpreadsheetUrl string `json:"spreadsheetUrl,omitempty\"` // The url of the spreadsheet. This field is read-only.

}

// Properties of a spreadsheet.
type SpreadsheetProperties struct {
	AutoRecalc string `json:"autoRecalc,omitempty\"` // The amount of time to wait before volatile functions are recalculated.

	DefaultFormat CellFormat `json:"defaultFormat,omitempty\"` // The default format of all cells in the spreadsheet. CellData.effectiveFormat will not be set if the cell's format is equal to this default format. This field is read-only.

	ImportFunctionsExternalUrlAccessAllowed bool `json:"importFunctionsExternalUrlAccessAllowed,omitempty\"` // Whether to allow external URL access for image and import functions. Read only when true. When false, you can set to true. This value will be bypassed and always return true if the admin has enabled the [allowlisting feature](https://support.google.com/a?p=url_allowlist).

	IterativeCalculationSettings IterativeCalculationSettings `json:"iterativeCalculationSettings,omitempty\"` // Determines whether and how circular references are resolved with iterative calculation. Absence of this field means that circular references result in calculation errors.

	Locale string `json:"locale,omitempty\"` // The locale of the spreadsheet in one of the following formats: * an ISO 639-1 language code such as `en` * an ISO 639-2 language code such as `fil`, if no 639-1 code exists * a combination of the ISO language code and country code, such as `en_US` Note: when updating this field, not all locales/languages are supported.

	SpreadsheetTheme SpreadsheetTheme `json:"spreadsheetTheme,omitempty\"` // Theme applied to the spreadsheet.

	TimeZone string `json:"timeZone,omitempty\"` // The time zone of the spreadsheet, in CLDR format such as `America/New_York`. If the time zone isn't recognized, this may be a custom time zone such as `GMT-07:00`.

	Title string `json:"title,omitempty\"` // The title of the spreadsheet.

}

// Represents spreadsheet theme
type SpreadsheetTheme struct {
	PrimaryFontFamily string `json:"primaryFontFamily,omitempty\"` // Name of the primary font family.

	ThemeColors []ThemeColorPair `json:"themeColors,omitempty\"` // The spreadsheet theme color pairs. To update you must provide all theme color pairs.

}

// A table.
type Table struct {
	ColumnProperties []TableColumnProperties `json:"columnProperties,omitempty\"` // The table column properties.

	Name string `json:"name,omitempty\"` // The table name. This is unique to all tables in the same spreadsheet.

	RangeValue GridRange `json:"range,omitempty\"` // The table range.

	RowsProperties TableRowsProperties `json:"rowsProperties,omitempty\"` // The table rows properties.

	TableId string `json:"tableId,omitempty\"` // The id of the table.

}

// A data validation rule for a column in a table.
type TableColumnDataValidationRule struct {
	Condition BooleanCondition `json:"condition,omitempty\"` // The condition that data in the cell must match. Valid only if the [BooleanCondition.type] is ONE_OF_LIST.

}

// The table column.
type TableColumnProperties struct {
	ColumnIndex int `json:"columnIndex,omitempty\"` // The 0-based column index. This index is relative to its position in the table and is not necessarily the same as the column index in the sheet.

	ColumnName string `json:"columnName,omitempty\"` // The column name.

	ColumnType string `json:"columnType,omitempty\"` // The column type.

	DataValidationRule TableColumnDataValidationRule `json:"dataValidationRule,omitempty\"` // The column data validation rule. Only set for dropdown column type.

}

// The table row properties.
type TableRowsProperties struct {
	FirstBandColorStyle ColorStyle `json:"firstBandColorStyle,omitempty\"` // The first color that is alternating. If this field is set, the first banded row is filled with the specified color. Otherwise, the first banded row is filled with a default color.

	FooterColorStyle ColorStyle `json:"footerColorStyle,omitempty\"` // The color of the last row. If this field is not set a footer is not added, the last row is filled with either first_band_color_style or second_band_color_style, depending on the color of the previous row. If updating an existing table without a footer to have a footer, the range will be expanded by 1 row. If updating an existing table with a footer and removing a footer, the range will be shrunk by 1 row.

	HeaderColorStyle ColorStyle `json:"headerColorStyle,omitempty\"` // The color of the header row. If this field is set, the header row is filled with the specified color. Otherwise, the header row is filled with a default color.

	SecondBandColorStyle ColorStyle `json:"secondBandColorStyle,omitempty\"` // The second color that is alternating. If this field is set, the second banded row is filled with the specified color. Otherwise, the second banded row is filled with a default color.

}

// The format of a run of text in a cell. Absent values indicate that the field isn't specified.
type TextFormat struct {
	Bold bool `json:"bold,omitempty\"` // True if the text is bold.

	FontFamily string `json:"fontFamily,omitempty\"` // The font family.

	FontSize int `json:"fontSize,omitempty\"` // The size of the font.

	ForegroundColor Color `json:"foregroundColor,omitempty\"` // The foreground color of the text. Deprecated: Use foreground_color_style.

	ForegroundColorStyle ColorStyle `json:"foregroundColorStyle,omitempty\"` // The foreground color of the text. If foreground_color is also set, this field takes precedence.

	Italic bool `json:"italic,omitempty\"` // True if the text is italicized.

	Link Link `json:"link,omitempty\"` // The link destination of the text, if any. Setting the link field in a TextFormatRun will clear the cell's existing links or a cell-level link set in the same request. When a link is set, the text foreground color will be set to the default link color and the text will be underlined. If these fields are modified in the same request, those values will be used instead of the link defaults.

	Strikethrough bool `json:"strikethrough,omitempty\"` // True if the text has a strikethrough.

	Underline bool `json:"underline,omitempty\"` // True if the text is underlined.

}

// A run of a text format. The format of this run continues until the start index of the next run. When updating, all fields must be set.
type TextFormatRun struct {
	Format TextFormat `json:"format,omitempty\"` // The format of this run. Absent values inherit the cell's format.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based character index where this run starts, in UTF-16 code units.

}

// Position settings for text.
type TextPosition struct {
	HorizontalAlignment string `json:"horizontalAlignment,omitempty\"` // Horizontal alignment setting for the piece of text.

}

// The rotation applied to text in a cell.
type TextRotation struct {
	Angle int `json:"angle,omitempty\"` // The angle between the standard orientation and the desired orientation. Measured in degrees. Valid values are between -90 and 90. Positive angles are angled upwards, negative are angled downwards. Note: For LTR text direction positive angles are in the counterclockwise direction, whereas for RTL they are in the clockwise direction

	Vertical bool `json:"vertical,omitempty\"` // If true, text reads top to bottom, but the orientation of individual characters is unchanged. For example: | V | | e | | r | | t | | i | | c | | a | | l |

}

// Splits a column of text into multiple columns, based on a delimiter in each cell.
type TextToColumnsRequest struct {
	Delimiter string `json:"delimiter,omitempty\"` // The delimiter to use. Used only if delimiterType is CUSTOM.

	DelimiterType string `json:"delimiterType,omitempty\"` // The delimiter type to use.

	Source GridRange `json:"source,omitempty\"` // The source data range. This must span exactly one column.

}

// A pair mapping a spreadsheet theme color type to the concrete color it represents.
type ThemeColorPair struct {
	Color ColorStyle `json:"color,omitempty\"` // The concrete color corresponding to the theme color type.

	ColorType string `json:"colorType,omitempty\"` // The type of the spreadsheet theme color.

}

// Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	Hours int `json:"hours,omitempty\"` // Hours of a day in 24 hour format. Must be greater than or equal to 0 and typically must be less than or equal to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.

	Minutes int `json:"minutes,omitempty\"` // Minutes of an hour. Must be greater than or equal to 0 and less than or equal to 59.

	Nanos int `json:"nanos,omitempty\"` // Fractions of seconds, in nanoseconds. Must be greater than or equal to 0 and less than or equal to 999,999,999.

	Seconds int `json:"seconds,omitempty\"` // Seconds of a minute. Must be greater than or equal to 0 and typically must be less than or equal to 59. An API may allow the value 60 if it allows leap-seconds.

}

// A color scale for a treemap chart.
type TreemapChartColorScale struct {
	MaxValueColor Color `json:"maxValueColor,omitempty\"` // The background color for cells with a color value greater than or equal to maxValue. Defaults to #109618 if not specified. Deprecated: Use max_value_color_style.

	MaxValueColorStyle ColorStyle `json:"maxValueColorStyle,omitempty\"` // The background color for cells with a color value greater than or equal to maxValue. Defaults to #109618 if not specified. If max_value_color is also set, this field takes precedence.

	MidValueColor Color `json:"midValueColor,omitempty\"` // The background color for cells with a color value at the midpoint between minValue and maxValue. Defaults to #efe6dc if not specified. Deprecated: Use mid_value_color_style.

	MidValueColorStyle ColorStyle `json:"midValueColorStyle,omitempty\"` // The background color for cells with a color value at the midpoint between minValue and maxValue. Defaults to #efe6dc if not specified. If mid_value_color is also set, this field takes precedence.

	MinValueColor Color `json:"minValueColor,omitempty\"` // The background color for cells with a color value less than or equal to minValue. Defaults to #dc3912 if not specified. Deprecated: Use min_value_color_style.

	MinValueColorStyle ColorStyle `json:"minValueColorStyle,omitempty\"` // The background color for cells with a color value less than or equal to minValue. Defaults to #dc3912 if not specified. If min_value_color is also set, this field takes precedence.

	NoDataColor Color `json:"noDataColor,omitempty\"` // The background color for cells that have no color data associated with them. Defaults to #000000 if not specified. Deprecated: Use no_data_color_style.

	NoDataColorStyle ColorStyle `json:"noDataColorStyle,omitempty\"` // The background color for cells that have no color data associated with them. Defaults to #000000 if not specified. If no_data_color is also set, this field takes precedence.

}

// A Treemap chart.
type TreemapChartSpec struct {
	ColorData ChartData `json:"colorData,omitempty\"` // The data that determines the background color of each treemap data cell. This field is optional. If not specified, size_data is used to determine background colors. If specified, the data is expected to be numeric. color_scale will determine how the values in this data map to data cell background colors.

	ColorScale TreemapChartColorScale `json:"colorScale,omitempty\"` // The color scale for data cells in the treemap chart. Data cells are assigned colors based on their color values. These color values come from color_data, or from size_data if color_data is not specified. Cells with color values less than or equal to min_value will have minValueColor as their background color. Cells with color values greater than or equal to max_value will have maxValueColor as their background color. Cells with color values between min_value and max_value will have background colors on a gradient between minValueColor and maxValueColor, the midpoint of the gradient being midValueColor. Cells with missing or non-numeric color values will have noDataColor as their background color.

	HeaderColor Color `json:"headerColor,omitempty\"` // The background color for header cells. Deprecated: Use header_color_style.

	HeaderColorStyle ColorStyle `json:"headerColorStyle,omitempty\"` // The background color for header cells. If header_color is also set, this field takes precedence.

	HideTooltips bool `json:"hideTooltips,omitempty\"` // True to hide tooltips.

	HintedLevels int `json:"hintedLevels,omitempty\"` // The number of additional data levels beyond the labeled levels to be shown on the treemap chart. These levels are not interactive and are shown without their labels. Defaults to 0 if not specified.

	Labels ChartData `json:"labels,omitempty\"` // The data that contains the treemap cell labels.

	Levels int `json:"levels,omitempty\"` // The number of data levels to show on the treemap chart. These levels are interactive and are shown with their labels. Defaults to 2 if not specified.

	MaxValue float64 `json:"maxValue,omitempty\"` // The maximum possible data value. Cells with values greater than this will have the same color as cells with this value. If not specified, defaults to the actual maximum value from color_data, or the maximum value from size_data if color_data is not specified.

	MinValue float64 `json:"minValue,omitempty\"` // The minimum possible data value. Cells with values less than this will have the same color as cells with this value. If not specified, defaults to the actual minimum value from color_data, or the minimum value from size_data if color_data is not specified.

	ParentLabels ChartData `json:"parentLabels,omitempty\"` // The data the contains the treemap cells' parent labels.

	SizeData ChartData `json:"sizeData,omitempty\"` // The data that determines the size of each treemap data cell. This data is expected to be numeric. The cells corresponding to non-numeric or missing data will not be rendered. If color_data is not specified, this data is used to determine data cell background colors as well.

	TextFormat TextFormat `json:"textFormat,omitempty\"` // The text format for all labels on the chart. The link field is not supported.

}

// Trims the whitespace (such as spaces, tabs, or new lines) in every cell in the specified range. This request removes all whitespace from the start and end of each cell's text, and reduces any subsequence of remaining whitespace characters to a single space. If the resulting trimmed text starts with a '+' or '=' character, the text remains as a string value and isn't interpreted as a formula.
type TrimWhitespaceRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range whose cells to trim.

}

// The result of trimming whitespace in cells.
type TrimWhitespaceResponse struct {
	CellsChangedCount int `json:"cellsChangedCount,omitempty\"` // The number of cells that were trimmed of whitespace.

}

// Unmerges cells in the given range.
type UnmergeCellsRequest struct {
	RangeValue GridRange `json:"range,omitempty\"` // The range within which all cells should be unmerged. If the range spans multiple merges, all will be unmerged. The range must not partially span any merge.

}

// Updates properties of the supplied banded range.
type UpdateBandingRequest struct {
	BandedRange BandedRange `json:"bandedRange,omitempty\"` // The banded range to update with the new properties.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `bandedRange` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

}

// Updates the borders of a range. If a field is not set in the request, that means the border remains as-is. For example, with two subsequent UpdateBordersRequest: 1. range: A1:A5 `{ top: RED, bottom: WHITE }` 2. range: A1:A5 `{ left: BLUE }` That would result in A1:A5 having a borders of `{ top: RED, bottom: WHITE, left: BLUE }`. If you want to clear a border, explicitly set the style to NONE.
type UpdateBordersRequest struct {
	Bottom Border `json:"bottom,omitempty\"` // The border to put at the bottom of the range.

	InnerHorizontal Border `json:"innerHorizontal,omitempty\"` // The horizontal border to put within the range.

	InnerVertical Border `json:"innerVertical,omitempty\"` // The vertical border to put within the range.

	Left Border `json:"left,omitempty\"` // The border to put at the left of the range.

	RangeValue GridRange `json:"range,omitempty\"` // The range whose borders should be updated.

	Right Border `json:"right,omitempty\"` // The border to put at the right of the range.

	Top Border `json:"top,omitempty\"` // The border to put at the top of the range.

}

// Updates all cells in a range with new data.
type UpdateCellsRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields of CellData that should be updated. At least one field must be specified. The root is the CellData; 'row.values.' should not be specified. A single `"*"` can be used as short-hand for listing every field.

	RangeValue GridRange `json:"range,omitempty\"` // The range to write data to. If the data in rows does not cover the entire requested range, the fields matching those set in fields will be cleared.

	Rows []RowData `json:"rows,omitempty\"` // The data to write.

	Start GridCoordinate `json:"start,omitempty\"` // The coordinate to start writing data at. Any number of rows and columns (including a different number of columns per row) may be written.

}

// Updates a chart's specifications. (This does not move or resize a chart. To move or resize a chart, use UpdateEmbeddedObjectPositionRequest.)
type UpdateChartSpecRequest struct {
	ChartId int `json:"chartId,omitempty\"` // The ID of the chart to update.

	Spec ChartSpec `json:"spec,omitempty\"` // The specification to apply to the chart.

}

// Updates a conditional format rule at the given index, or moves a conditional format rule to another index.
type UpdateConditionalFormatRuleRequest struct {
	Index int `json:"index,omitempty\"` // The zero-based index of the rule that should be replaced or moved.

	NewIndex int `json:"newIndex,omitempty\"` // The zero-based new index the rule should end up at.

	Rule ConditionalFormatRule `json:"rule,omitempty\"` // The rule that should replace the rule at the given index.

	SheetId int `json:"sheetId,omitempty\"` // The sheet of the rule to move. Required if new_index is set, unused otherwise.

}

// The result of updating a conditional format rule.
type UpdateConditionalFormatRuleResponse struct {
	NewIndex int `json:"newIndex,omitempty\"` // The index of the new rule.

	NewRule ConditionalFormatRule `json:"newRule,omitempty\"` // The new rule that replaced the old rule (if replacing), or the rule that was moved (if moved)

	OldIndex int `json:"oldIndex,omitempty\"` // The old index of the rule. Not set if a rule was replaced (because it is the same as new_index).

	OldRule ConditionalFormatRule `json:"oldRule,omitempty\"` // The old (deleted) rule. Not set if a rule was moved (because it is the same as new_rule).

}

// Updates a data source. After the data source is updated successfully, an execution is triggered to refresh the associated DATA_SOURCE sheet to read data from the updated data source. The request requires an additional `bigquery.readonly` OAuth scope if you are updating a BigQuery data source.
type UpdateDataSourceRequest struct {
	DataSource DataSource `json:"dataSource,omitempty\"` // The data source to update.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `dataSource` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

}

// The response from updating data source.
type UpdateDataSourceResponse struct {
	DataExecutionStatus DataExecutionStatus `json:"dataExecutionStatus,omitempty\"` // The data execution status.

	DataSource DataSource `json:"dataSource,omitempty\"` // The updated data source.

}

// A request to update properties of developer metadata. Updates the properties of the developer metadata selected by the filters to the values provided in the DeveloperMetadata resource. Callers must specify the properties they wish to update in the fields parameter, as well as specify at least one DataFilter matching the metadata they wish to update.
type UpdateDeveloperMetadataRequest struct {
	DataFilters []DataFilter `json:"dataFilters,omitempty\"` // The filters matching the developer metadata entries to update.

	DeveloperMetadata DeveloperMetadata `json:"developerMetadata,omitempty\"` // The value that all metadata matched by the data filters will be updated to.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `developerMetadata` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

}

// The response from updating developer metadata.
type UpdateDeveloperMetadataResponse struct {
	DeveloperMetadata []DeveloperMetadata `json:"developerMetadata,omitempty\"` // The updated developer metadata.

}

// Updates the state of the specified group.
type UpdateDimensionGroupRequest struct {
	DimensionGroup DimensionGroup `json:"dimensionGroup,omitempty\"` // The group whose state should be updated. The range and depth of the group should specify a valid group on the sheet, and all other fields updated.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `dimensionGroup` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

}

// Updates properties of dimensions within the specified range.
type UpdateDimensionPropertiesRequest struct {
	DataSourceSheetRange DataSourceSheetDimensionRange `json:"dataSourceSheetRange,omitempty\"` // The columns on a data source sheet to update.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `properties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Properties DimensionProperties `json:"properties,omitempty\"` // Properties to update.

	RangeValue DimensionRange `json:"range,omitempty\"` // The rows or columns to update.

}

// Updates an embedded object's border property.
type UpdateEmbeddedObjectBorderRequest struct {
	Border EmbeddedObjectBorder `json:"border,omitempty\"` // The border that applies to the embedded object.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `border` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	ObjectId int `json:"objectId,omitempty\"` // The ID of the embedded object to update.

}

// Update an embedded object's position (such as a moving or resizing a chart or image).
type UpdateEmbeddedObjectPositionRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields of OverlayPosition that should be updated when setting a new position. Used only if newPosition.overlayPosition is set, in which case at least one field must be specified. The root `newPosition.overlayPosition` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	NewPosition EmbeddedObjectPosition `json:"newPosition,omitempty\"` // An explicit position to move the embedded object to. If newPosition.sheetId is set, a new sheet with that ID will be created. If newPosition.newSheet is set to true, a new sheet will be created with an ID that will be chosen for you.

	ObjectId int `json:"objectId,omitempty\"` // The ID of the object to moved.

}

// The result of updating an embedded object's position.
type UpdateEmbeddedObjectPositionResponse struct {
	Position EmbeddedObjectPosition `json:"position,omitempty\"` // The new position of the embedded object.

}

// Updates properties of the filter view.
type UpdateFilterViewRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `filter` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Filter FilterView `json:"filter,omitempty\"` // The new properties of the filter view.

}

// Updates properties of the named range with the specified namedRangeId.
type UpdateNamedRangeRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `namedRange` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	NamedRange NamedRange `json:"namedRange,omitempty\"` // The named range to update with the new properties.

}

// Updates an existing protected range with the specified protectedRangeId.
type UpdateProtectedRangeRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `protectedRange` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	ProtectedRange ProtectedRange `json:"protectedRange,omitempty\"` // The protected range to update with the new properties.

}

// Updates properties of the sheet with the specified sheetId.
type UpdateSheetPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `properties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Properties SheetProperties `json:"properties,omitempty\"` // The properties to update.

}

// Updates a slicer's specifications. (This does not move or resize a slicer. To move or resize a slicer use UpdateEmbeddedObjectPositionRequest.
type UpdateSlicerSpecRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `SlicerSpec` is implied and should not be specified. A single "*"` can be used as short-hand for listing every field.

	SlicerId int `json:"slicerId,omitempty\"` // The id of the slicer to update.

	Spec SlicerSpec `json:"spec,omitempty\"` // The specification to apply to the slicer.

}

// Updates properties of a spreadsheet.
type UpdateSpreadsheetPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root 'properties' is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Properties SpreadsheetProperties `json:"properties,omitempty\"` // The properties to update.

}

// Updates a table in the spreadsheet.
type UpdateTableRequest struct {
	Fields string `json:"fields,omitempty\"` // Required. The fields that should be updated. At least one field must be specified. The root `table` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	Table Table `json:"table,omitempty\"` // Required. The table to update.

}

// The response when updating a range of values by a data filter in a spreadsheet.
type UpdateValuesByDataFilterResponse struct {
	DataFilter DataFilter `json:"dataFilter,omitempty\"` // The data filter that selected the range that was updated.

	UpdatedCells int `json:"updatedCells,omitempty\"` // The number of cells updated.

	UpdatedColumns int `json:"updatedColumns,omitempty\"` // The number of columns where at least one cell in the column was updated.

	UpdatedData ValueRange `json:"updatedData,omitempty\"` // The values of the cells in the range matched by the dataFilter after all updates were applied. This is only included if the request's `includeValuesInResponse` field was `true`.

	UpdatedRange string `json:"updatedRange,omitempty\"` // The range (in [A1 notation](https://developers.google.com/workspace/sheets/api/guides/concepts#cell)) that updates were applied to.

	UpdatedRows int `json:"updatedRows,omitempty\"` // The number of rows where at least one cell in the row was updated.

}

// The response when updating a range of values in a spreadsheet.
type UpdateValuesResponse struct {
	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The spreadsheet the updates were applied to.

	UpdatedCells int `json:"updatedCells,omitempty\"` // The number of cells updated.

	UpdatedColumns int `json:"updatedColumns,omitempty\"` // The number of columns where at least one cell in the column was updated.

	UpdatedData ValueRange `json:"updatedData,omitempty\"` // The values of the cells after updates were applied. This is only included if the request's `includeValuesInResponse` field was `true`.

	UpdatedRange string `json:"updatedRange,omitempty\"` // The range (in A1 notation) that updates were applied to.

	UpdatedRows int `json:"updatedRows,omitempty\"` // The number of rows where at least one cell in the row was updated.

}

// Data within a range of the spreadsheet.
type ValueRange struct {
	MajorDimension string `json:"majorDimension,omitempty\"` // The major dimension of the values. For output, if the spreadsheet data is: `A1=1,B1=2,A2=3,B2=4`, then requesting `range=A1:B2,majorDimension=ROWS` will return `[[1,2],[3,4]]`, whereas requesting `range=A1:B2,majorDimension=COLUMNS` will return `[[1,3],[2,4]]`. For input, with `range=A1:B2,majorDimension=ROWS` then `[[1,2],[3,4]]` will set `A1=1,B1=2,A2=3,B2=4`. With `range=A1:B2,majorDimension=COLUMNS` then `[[1,2],[3,4]]` will set `A1=1,B1=3,A2=2,B2=4`. When writing, if this field is not set, it defaults to ROWS.

	RangeValue string `json:"range,omitempty\"` // The range the values cover, in [A1 notation](https://developers.google.com/workspace/sheets/api/guides/concepts#cell). For output, this range indicates the entire requested range, even though the values will exclude trailing rows and columns. When appending values, this field represents the range to search for a table, after which values will be appended.

	Values [][]interface{} `json:"values,omitempty\"` // The data that was read or to be written. This is an array of arrays, the outer array representing all the data and each inner array representing a major dimension. Each item in the inner array corresponds with one cell. For output, empty trailing rows and columns will not be included. For input, supported value types are: bool, string, and double. Null values will be skipped. To set a cell to an empty value, set the string value to an empty string.

}

// Styles for a waterfall chart column.
type WaterfallChartColumnStyle struct {
	Color Color `json:"color,omitempty\"` // The color of the column. Deprecated: Use color_style.

	ColorStyle ColorStyle `json:"colorStyle,omitempty\"` // The color of the column. If color is also set, this field takes precedence.

	Label string `json:"label,omitempty\"` // The label of the column's legend.

}

// A custom subtotal column for a waterfall chart series.
type WaterfallChartCustomSubtotal struct {
	DataIsSubtotal bool `json:"dataIsSubtotal,omitempty\"` // True if the data point at subtotal_index is the subtotal. If false, the subtotal will be computed and appear after the data point.

	Label string `json:"label,omitempty\"` // A label for the subtotal column.

	SubtotalIndex int `json:"subtotalIndex,omitempty\"` // The zero-based index of a data point within the series. If data_is_subtotal is true, the data point at this index is the subtotal. Otherwise, the subtotal appears after the data point with this index. A series can have multiple subtotals at arbitrary indices, but subtotals do not affect the indices of the data points. For example, if a series has three data points, their indices will always be 0, 1, and 2, regardless of how many subtotals exist on the series or what data points they are associated with.

}

// The domain of a waterfall chart.
type WaterfallChartDomain struct {
	Data ChartData `json:"data,omitempty\"` // The data of the WaterfallChartDomain.

	Reversed bool `json:"reversed,omitempty\"` // True to reverse the order of the domain values (horizontal axis).

}

// A single series of data for a waterfall chart.
type WaterfallChartSeries struct {
	CustomSubtotals []WaterfallChartCustomSubtotal `json:"customSubtotals,omitempty\"` // Custom subtotal columns appearing in this series. The order in which subtotals are defined is not significant. Only one subtotal may be defined for each data point.

	Data ChartData `json:"data,omitempty\"` // The data being visualized in this series.

	DataLabel DataLabel `json:"dataLabel,omitempty\"` // Information about the data labels for this series.

	HideTrailingSubtotal bool `json:"hideTrailingSubtotal,omitempty\"` // True to hide the subtotal column from the end of the series. By default, a subtotal column will appear at the end of each series. Setting this field to true will hide that subtotal column for this series.

	NegativeColumnsStyle WaterfallChartColumnStyle `json:"negativeColumnsStyle,omitempty\"` // Styles for all columns in this series with negative values.

	PositiveColumnsStyle WaterfallChartColumnStyle `json:"positiveColumnsStyle,omitempty\"` // Styles for all columns in this series with positive values.

	SubtotalColumnsStyle WaterfallChartColumnStyle `json:"subtotalColumnsStyle,omitempty\"` // Styles for all subtotal columns in this series.

}

// A waterfall chart.
type WaterfallChartSpec struct {
	ConnectorLineStyle LineStyle `json:"connectorLineStyle,omitempty\"` // The line style for the connector lines.

	Domain WaterfallChartDomain `json:"domain,omitempty\"` // The domain data (horizontal axis) for the waterfall chart.

	FirstValueIsTotal bool `json:"firstValueIsTotal,omitempty\"` // True to interpret the first value as a total.

	HideConnectorLines bool `json:"hideConnectorLines,omitempty\"` // True to hide connector lines between columns.

	Series []WaterfallChartSeries `json:"series,omitempty\"` // The data this waterfall chart is visualizing.

	StackedType string `json:"stackedType,omitempty\"` // The stacked type.

	TotalDataLabel DataLabel `json:"totalDataLabel,omitempty\"` // Controls whether to display additional data labels on stacked charts which sum the total value of all stacked values at each value along the domain axis. stacked_type must be STACKED and neither CUSTOM nor placement can be set on the total_data_label.

}
