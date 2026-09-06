// Google Docs API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package docs

// Adds a document tab. When a tab is added at a given index, all subsequent tabs' indexes are incremented.
type AddDocumentTabRequest struct {
	TabProperties TabProperties `json:"tabProperties,omitempty\"` // The properties of the tab to add. All properties are optional.

}

// The result of adding a document tab.
type AddDocumentTabResponse struct {
	TabProperties TabProperties `json:"tabProperties,omitempty\"` // The properties of the newly added tab.

}

// A ParagraphElement representing a spot in the text that's dynamically replaced with content that can change over time, like a page number.
type AutoText struct {
	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. An AutoText may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this AutoText, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this AutoText.

	TypeValue string `json:"type,omitempty\"` // The type of this auto text.

}

// Represents the background of a document.
type Background struct {
	Color OptionalColor `json:"color,omitempty\"` // The background color.

}

// A mask that indicates which of the fields on the base Background have been changed in this suggestion. For any field set to true, the Backgound has a new suggested value.
type BackgroundSuggestionState struct {
	BackgroundColorSuggested bool `json:"backgroundColorSuggested,omitempty\"` // Indicates whether the current background color has been modified in this suggestion.

}

// Request message for BatchUpdateDocument.
type BatchUpdateDocumentRequest struct {
	Requests []Request `json:"requests,omitempty\"` // A list of updates to apply to the document.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed.

}

// Response message from a BatchUpdateDocument request.
type BatchUpdateDocumentResponse struct {
	DocumentId string `json:"documentId,omitempty\"` // The ID of the document to which the updates were applied to.

	Replies []Response `json:"replies,omitempty\"` // The reply of the updates. This maps 1:1 with the updates, although replies to some requests may be empty.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // The updated write control after applying the request.

}

// The document body. The body typically contains the full document contents except for headers, footers, and footnotes.
type Body struct {
	Content []StructuralElement `json:"content,omitempty\"` // The contents of the body. The indexes for the body's content begin at zero.

}

// A reference to a bookmark in this document.
type BookmarkLink struct {
	Id string `json:"id,omitempty\"` // The ID of a bookmark in this document.

	TabId string `json:"tabId,omitempty\"` // The ID of the tab containing this bookmark.

}

// Describes the bullet of a paragraph.
type Bullet struct {
	ListId string `json:"listId,omitempty\"` // The ID of the list this paragraph belongs to.

	NestingLevel int `json:"nestingLevel,omitempty\"` // The nesting level of this paragraph in the list.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The paragraph-specific text style applied to this bullet.

}

// A mask that indicates which of the fields on the base Bullet have been changed in this suggestion. For any field set to true, there's a new suggested value.
type BulletSuggestionState struct {
	ListIdSuggested bool `json:"listIdSuggested,omitempty\"` // Indicates if there was a suggested change to the list_id.

	NestingLevelSuggested bool `json:"nestingLevelSuggested,omitempty\"` // Indicates if there was a suggested change to the nesting_level.

	TextStyleSuggestionState TextStyleSuggestionState `json:"textStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields in text style have been changed in this suggestion.

}

// A solid color.
type Color struct {
	RgbColor RgbColor `json:"rgbColor,omitempty\"` // The RGB color value.

}

// A ParagraphElement representing a column break. A column break makes the subsequent text start at the top of the next column.
type ColumnBreak struct {
	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A ColumnBreak may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this ColumnBreak, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this ColumnBreak. Similar to text content, like text runs and footnote references, the text style of a column break can affect content layout as well as the styling of text inserted next to it.

}

// Creates a Footer. The new footer is applied to the SectionStyle at the location of the SectionBreak if specified, otherwise it is applied to the DocumentStyle. If a footer of the specified type already exists, a 400 bad request error is returned.
type CreateFooterRequest struct {
	SectionBreakLocation Location `json:"sectionBreakLocation,omitempty\"` // The location of the SectionBreak immediately preceding the section whose SectionStyle this footer should belong to. If this is unset or refers to the first section break in the document, the footer applies to the document style.

	TypeValue string `json:"type,omitempty\"` // The type of footer to create.

}

// The result of creating a footer.
type CreateFooterResponse struct {
	FooterId string `json:"footerId,omitempty\"` // The ID of the created footer.

}

// Creates a Footnote segment and inserts a new FootnoteReference to it at the given location. The new Footnote segment will contain a space followed by a newline character.
type CreateFootnoteRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the footnote reference at the end of the document body. Footnote references cannot be inserted inside a header, footer or footnote. Since footnote references can only be inserted in the body, the segment ID field must be empty.

	Location Location `json:"location,omitempty\"` // Inserts the footnote reference at a specific index in the document. The footnote reference must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). Footnote references cannot be inserted inside an equation, header, footer or footnote. Since footnote references can only be inserted in the body, the segment ID field must be empty.

}

// The result of creating a footnote.
type CreateFootnoteResponse struct {
	FootnoteId string `json:"footnoteId,omitempty\"` // The ID of the created footnote.

}

// Creates a Header. The new header is applied to the SectionStyle at the location of the SectionBreak if specified, otherwise it is applied to the DocumentStyle. If a header of the specified type already exists, a 400 bad request error is returned.
type CreateHeaderRequest struct {
	SectionBreakLocation Location `json:"sectionBreakLocation,omitempty\"` // The location of the SectionBreak which begins the section this header should belong to. If `section_break_location' is unset or if it refers to the first section break in the document body, the header applies to the DocumentStyle

	TypeValue string `json:"type,omitempty\"` // The type of header to create.

}

// The result of creating a header.
type CreateHeaderResponse struct {
	HeaderId string `json:"headerId,omitempty\"` // The ID of the created header.

}

// Creates a NamedRange referencing the given range.
type CreateNamedRangeRequest struct {
	Name string `json:"name,omitempty\"` // The name of the NamedRange. Names do not need to be unique. Names must be at least 1 character and no more than 256 characters, measured in UTF-16 code units.

	RangeValue RangeValue `json:"range,omitempty\"` // The range to apply the name to.

}

// The result of creating a named range.
type CreateNamedRangeResponse struct {
	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the created named range.

}

// Creates bullets for all of the paragraphs that overlap with the given range. The nesting level of each paragraph will be determined by counting leading tabs in front of each paragraph. To avoid excess space between the bullet and the corresponding paragraph, these leading tabs are removed by this request. This may change the indices of parts of the text. If the paragraph immediately before paragraphs being updated is in a list with a matching preset, the paragraphs being updated are added to that preceding list.
type CreateParagraphBulletsRequest struct {
	BulletPreset string `json:"bulletPreset,omitempty\"` // The kinds of bullet glyphs to be used.

	RangeValue RangeValue `json:"range,omitempty\"` // The range to apply the bullet preset to.

}

// The crop properties of an image. The crop rectangle is represented using fractional offsets from the original content's 4 edges. - If the offset is in the interval (0, 1), the corresponding edge of crop rectangle is positioned inside of the image's original bounding rectangle. - If the offset is negative or greater than 1, the corresponding edge of crop rectangle is positioned outside of the image's original bounding rectangle. - If all offsets and rotation angles are 0, the image is not cropped.
type CropProperties struct {
	Angle float64 `json:"angle,omitempty\"` // The clockwise rotation angle of the crop rectangle around its center, in radians. Rotation is applied after the offsets.

	OffsetBottom float64 `json:"offsetBottom,omitempty\"` // The offset specifies how far inwards the bottom edge of the crop rectangle is from the bottom edge of the original content as a fraction of the original content's height.

	OffsetLeft float64 `json:"offsetLeft,omitempty\"` // The offset specifies how far inwards the left edge of the crop rectangle is from the left edge of the original content as a fraction of the original content's width.

	OffsetRight float64 `json:"offsetRight,omitempty\"` // The offset specifies how far inwards the right edge of the crop rectangle is from the right edge of the original content as a fraction of the original content's width.

	OffsetTop float64 `json:"offsetTop,omitempty\"` // The offset specifies how far inwards the top edge of the crop rectangle is from the top edge of the original content as a fraction of the original content's height.

}

// A mask that indicates which of the fields on the base CropProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type CropPropertiesSuggestionState struct {
	AngleSuggested bool `json:"angleSuggested,omitempty\"` // Indicates if there was a suggested change to angle.

	OffsetBottomSuggested bool `json:"offsetBottomSuggested,omitempty\"` // Indicates if there was a suggested change to offset_bottom.

	OffsetLeftSuggested bool `json:"offsetLeftSuggested,omitempty\"` // Indicates if there was a suggested change to offset_left.

	OffsetRightSuggested bool `json:"offsetRightSuggested,omitempty\"` // Indicates if there was a suggested change to offset_right.

	OffsetTopSuggested bool `json:"offsetTopSuggested,omitempty\"` // Indicates if there was a suggested change to offset_top.

}

// A date instance mentioned in a document.
type DateElement struct {
	DateElementProperties DateElementProperties `json:"dateElementProperties,omitempty\"` // The properties of this DateElement.

	DateId string `json:"dateId,omitempty\"` // Output only. The unique ID of this date.

	SuggestedDateElementPropertiesChanges map[string]interface{} `json:"suggestedDateElementPropertiesChanges,omitempty\"` // The suggested changes to the date element properties, keyed by suggestion ID.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // IDs for suggestions that remove this date from the document. A DateElement might have multiple deletion IDs if, for example, multiple users suggest deleting it. If empty, then this date isn't suggested for deletion.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // IDs for suggestions that insert this date into the document. A DateElement might have multiple insertion IDs if it's a nested suggested change (a suggestion within a suggestion made by a different user, for example). If empty, then this date isn't a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this DateElement, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this DateElement.

}

// Properties of a DateElement.
type DateElementProperties struct {
	DateFormat string `json:"dateFormat,omitempty\"` // Determines how the date part of the DateElement will be displayed in the document. If unset, the default value is DATE_FORMAT_MONTH_DAY_YEAR_ABBREVIATED, indicating the DateElement will be formatted as `MMM d, y` in `en`, or locale specific equivalent.

	DisplayText string `json:"displayText,omitempty\"` // Output only. Indicates how the DateElement is displayed in the document.

	Locale string `json:"locale,omitempty\"` // The language code of the DateElement. For example, `en`. If unset, the default locale is `en`. Limited to the following locales: `af`, `am`, `ar`, `as`, `az`, `be`, `bg`, `bn`, `ca`, `cs`, `da`, `de`, `el`, `en`, `en-CA`, `en-GB`, `es`, `es-419`, `et`, `eu`, `fa`, `fi`, `fil`, `fr`, `fr-CA`, `gl`, `gu`, `hi`, `hr`, `hu`, `hy`, `id`, `is`, `it`, `iw`, `ja`, `ka`, `kk`, `km`, `kn`, `ko`, `lo`, `lt`, `lv`, `mk`, `ml`, `mn`, `mr`, `ms`, `ne`, `nl`, `no`, `or`, `pa`, `pl`, `pt-BR`, `pt-PT`, `ro`, `ru`, `si`, `sk`, `sl`, `sq`, `sr`, `sv`, `sw`, `ta`, `te`, `th`, `tr`, `uk`, `ur`, `uz`, `vi`, `zh-CN`, `zh-HK`, `zh-TW`, `zu`, `cy`, `my`.

	TimeFormat string `json:"timeFormat,omitempty\"` // Determines how the time part of the DateElement will be displayed in the document. If unset, the default value is TIME_FORMAT_DISABLED, indicating no time should be shown.

	TimeZoneId string `json:"timeZoneId,omitempty\"` // The time zone of the DateElement, as defined by the Unicode Common Locale Data Repository (CLDR) project. For example, `America/New_York`. If unset, the default time zone is `etc/UTC`.

	Timestamp string `json:"timestamp,omitempty\"` // The point in time to represent, in seconds and nanoseconds since Unix epoch: January 1, 1970 at midnight UTC. Timestamp is expected to be in UTC. If time_zone_id is set, the timestamp is adjusted according to the time zone. For example, a timestamp of `18000` with a date format of `DATE_FORMAT_ISO8601` and time format of `TIME_FORMAT_HOUR_MINUTE` would be displayed as `1970-01-01 5:00 AM`. A timestamp of `18000` with date format of `DATE_FORMAT_ISO8601`, time format of `TIME_FORMAT_HOUR_MINUTE`, and time zone set to `America/New_York` will instead be `1970-01-01 12:00 AM`.

}

// A mask that indicates which of the fields on the base DateElementProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type DateElementPropertiesSuggestionState struct {
	DateFormatSuggested bool `json:"dateFormatSuggested,omitempty\"` // Indicates if there was a suggested change to date_format.

	LocaleSuggested bool `json:"localeSuggested,omitempty\"` // Indicates if there was a suggested change to locale.

	TimeFormatSuggested bool `json:"timeFormatSuggested,omitempty\"` // Indicates if there was a suggested change to time_format.

	TimeZoneIdSuggested bool `json:"timeZoneIdSuggested,omitempty\"` // Indicates if there was a suggested change to time_zone_id.

	TimestampSuggested bool `json:"timestampSuggested,omitempty\"` // Indicates if there was a suggested change to timestamp.

}

// Deletes content from the document.
type DeleteContentRangeRequest struct {
	RangeValue RangeValue `json:"range,omitempty\"` // The range of content to delete. Deleting text that crosses a paragraph boundary may result in changes to paragraph styles, lists, positioned objects and bookmarks as the two paragraphs are merged. Attempting to delete certain ranges can result in an invalid document structure in which case a 400 bad request error is returned. Some examples of invalid delete requests include: * Deleting one code unit of a surrogate pair. * Deleting the last newline character of a Body, Header, Footer, Footnote, TableCell or TableOfContents. * Deleting the start or end of a Table, TableOfContents or Equation without deleting the entire element. * Deleting the newline character before a Table, TableOfContents or SectionBreak without deleting the element. * Deleting individual rows or cells of a table. Deleting the content within a table cell is allowed.

}

// Deletes a Footer from the document.
type DeleteFooterRequest struct {
	FooterId string `json:"footerId,omitempty\"` // The id of the footer to delete. If this footer is defined on DocumentStyle, the reference to this footer is removed, resulting in no footer of that type for the first section of the document. If this footer is defined on a SectionStyle, the reference to this footer is removed and the footer of that type is now continued from the previous section.

	TabId string `json:"tabId,omitempty\"` // The tab that contains the footer to delete. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// Deletes a Header from the document.
type DeleteHeaderRequest struct {
	HeaderId string `json:"headerId,omitempty\"` // The id of the header to delete. If this header is defined on DocumentStyle, the reference to this header is removed, resulting in no header of that type for the first section of the document. If this header is defined on a SectionStyle, the reference to this header is removed and the header of that type is now continued from the previous section.

	TabId string `json:"tabId,omitempty\"` // The tab containing the header to delete. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// Deletes a NamedRange.
type DeleteNamedRangeRequest struct {
	Name string `json:"name,omitempty\"` // The name of the range(s) to delete. All named ranges with the given name will be deleted.

	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the named range to delete.

	TabsCriteria TabsCriteria `json:"tabsCriteria,omitempty\"` // Optional. The criteria used to specify which tab(s) the range deletion should occur in. When omitted, the range deletion is applied to all tabs. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the range deletion applies to the singular tab. In a document containing multiple tabs: - If provided, the range deletion applies to the specified tabs. - If not provided, the range deletion applies to all tabs.

}

// Deletes bullets from all of the paragraphs that overlap with the given range. The nesting level of each paragraph will be visually preserved by adding indent to the start of the corresponding paragraph.
type DeleteParagraphBulletsRequest struct {
	RangeValue RangeValue `json:"range,omitempty\"` // The range to delete bullets from.

}

// Deletes a PositionedObject from the document.
type DeletePositionedObjectRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The ID of the positioned object to delete.

	TabId string `json:"tabId,omitempty\"` // The tab that the positioned object to delete is in. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// Deletes a tab. If the tab has child tabs, they are deleted as well.
type DeleteTabRequest struct {
	TabId string `json:"tabId,omitempty\"` // The ID of the tab to delete.

}

// Deletes a column from a table.
type DeleteTableColumnRequest struct {
	TableCellLocation TableCellLocation `json:"tableCellLocation,omitempty\"` // The reference table cell location from which the column will be deleted. The column this cell spans will be deleted. If this is a merged cell that spans multiple columns, all columns that the cell spans will be deleted. If no columns remain in the table after this deletion, the whole table is deleted.

}

// Deletes a row from a table.
type DeleteTableRowRequest struct {
	TableCellLocation TableCellLocation `json:"tableCellLocation,omitempty\"` // The reference table cell location from which the row will be deleted. The row this cell spans will be deleted. If this is a merged cell that spans multiple rows, all rows that the cell spans will be deleted. If no rows remain in the table after this deletion, the whole table is deleted.

}

// A magnitude in a single direction in the specified units.
type Dimension struct {
	Magnitude float64 `json:"magnitude,omitempty\"` // The magnitude.

	Unit string `json:"unit,omitempty\"` // The units for magnitude.

}

// A Google Docs document.
type Document struct {
	Body Body `json:"body,omitempty\"` // Output only. The main body of the document. Legacy field: Instead, use Document.tabs.documentTab.body, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	DocumentId string `json:"documentId,omitempty\"` // Output only. The ID of the document.

	DocumentStyle DocumentStyle `json:"documentStyle,omitempty\"` // Output only. The style of the document. Legacy field: Instead, use Document.tabs.documentTab.documentStyle, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	Footers map[string]interface{} `json:"footers,omitempty\"` // Output only. The footers in the document, keyed by footer ID. Legacy field: Instead, use Document.tabs.documentTab.footers, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	Footnotes map[string]interface{} `json:"footnotes,omitempty\"` // Output only. The footnotes in the document, keyed by footnote ID. Legacy field: Instead, use Document.tabs.documentTab.footnotes, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	Headers map[string]interface{} `json:"headers,omitempty\"` // Output only. The headers in the document, keyed by header ID. Legacy field: Instead, use Document.tabs.documentTab.headers, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	InlineObjects map[string]interface{} `json:"inlineObjects,omitempty\"` // Output only. The inline objects in the document, keyed by object ID. Legacy field: Instead, use Document.tabs.documentTab.inlineObjects, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	Lists map[string]interface{} `json:"lists,omitempty\"` // Output only. The lists in the document, keyed by list ID. Legacy field: Instead, use Document.tabs.documentTab.lists, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	NamedRanges map[string]interface{} `json:"namedRanges,omitempty\"` // Output only. The named ranges in the document, keyed by name. Legacy field: Instead, use Document.tabs.documentTab.namedRanges, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	NamedStyles NamedStyles `json:"namedStyles,omitempty\"` // Output only. The named styles of the document. Legacy field: Instead, use Document.tabs.documentTab.namedStyles, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	PositionedObjects map[string]interface{} `json:"positionedObjects,omitempty\"` // Output only. The positioned objects in the document, keyed by object ID. Legacy field: Instead, use Document.tabs.documentTab.positionedObjects, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	RevisionId string `json:"revisionId,omitempty\"` // Output only. The revision ID of the document. Can be used in update requests to specify which revision of a document to apply updates to and how the request should behave if the document has been edited since that revision. Only populated if the user has edit access to the document. The revision ID is not a sequential number but an opaque string. The format of the revision ID might change over time. A returned revision ID is only guaranteed to be valid for 24 hours after it has been returned and cannot be shared across users. If the revision ID is unchanged between calls, then the document has not changed. Conversely, a changed ID (for the same document and user) usually means the document has been updated. However, a changed ID can also be due to internal factors such as ID format changes.

	SuggestedDocumentStyleChanges map[string]interface{} `json:"suggestedDocumentStyleChanges,omitempty\"` // Output only. The suggested changes to the style of the document, keyed by suggestion ID. Legacy field: Instead, use Document.tabs.documentTab.suggestedDocumentStyleChanges, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	SuggestedNamedStylesChanges map[string]interface{} `json:"suggestedNamedStylesChanges,omitempty\"` // Output only. The suggested changes to the named styles of the document, keyed by suggestion ID. Legacy field: Instead, use Document.tabs.documentTab.suggestedNamedStylesChanges, which exposes the actual document content from all tabs when the includeTabsContent parameter is set to `true`. If `false` or unset, this field contains information about the first tab in the document.

	SuggestionsViewMode string `json:"suggestionsViewMode,omitempty\"` // Output only. The suggestions view mode applied to the document. Note: When editing a document, changes must be based on a document with SUGGESTIONS_INLINE.

	Tabs []Tab `json:"tabs,omitempty\"` // Tabs that are part of a document. Tabs can contain child tabs, a tab nested within another tab. Child tabs are represented by the Tab.childTabs field.

	Title string `json:"title,omitempty\"` // The title of the document.

}

// Represents document-level format settings.
type DocumentFormat struct {
	DocumentMode string `json:"documentMode,omitempty\"` // Whether the document has pages or is pageless.

}

// The style of the document.
type DocumentStyle struct {
	Background Background `json:"background,omitempty\"` // The background of the document. Documents cannot have a transparent background color.

	DefaultFooterId string `json:"defaultFooterId,omitempty\"` // The ID of the default footer. If not set, there's no default footer. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	DefaultHeaderId string `json:"defaultHeaderId,omitempty\"` // The ID of the default header. If not set, there's no default header. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	DocumentFormat DocumentFormat `json:"documentFormat,omitempty\"` // Specifies document-level format settings, such as the document mode (pages vs pageless).

	EvenPageFooterId string `json:"evenPageFooterId,omitempty\"` // The ID of the footer used only for even pages. The value of use_even_page_header_footer determines whether to use the default_footer_id or this value for the footer on even pages. If not set, there's no even page footer. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	EvenPageHeaderId string `json:"evenPageHeaderId,omitempty\"` // The ID of the header used only for even pages. The value of use_even_page_header_footer determines whether to use the default_header_id or this value for the header on even pages. If not set, there's no even page header. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FirstPageFooterId string `json:"firstPageFooterId,omitempty\"` // The ID of the footer used only for the first page. If not set then a unique footer for the first page does not exist. The value of use_first_page_header_footer determines whether to use the default_footer_id or this value for the footer on the first page. If not set, there's no first page footer. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FirstPageHeaderId string `json:"firstPageHeaderId,omitempty\"` // The ID of the header used only for the first page. If not set then a unique header for the first page does not exist. The value of use_first_page_header_footer determines whether to use the default_header_id or this value for the header on the first page. If not set, there's no first page header. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FlipPageOrientation bool `json:"flipPageOrientation,omitempty\"` // Optional. Indicates whether to flip the dimensions of the page_size, which allows changing the page orientation between portrait and landscape. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginBottom Dimension `json:"marginBottom,omitempty\"` // The bottom page margin. Updating the bottom page margin on the document style clears the bottom page margin on all section styles. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginFooter Dimension `json:"marginFooter,omitempty\"` // The amount of space between the bottom of the page and the contents of the footer. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginHeader Dimension `json:"marginHeader,omitempty\"` // The amount of space between the top of the page and the contents of the header. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginLeft Dimension `json:"marginLeft,omitempty\"` // The left page margin. Updating the left page margin on the document style clears the left page margin on all section styles. It may also cause columns to resize in all sections. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginRight Dimension `json:"marginRight,omitempty\"` // The right page margin. Updating the right page margin on the document style clears the right page margin on all section styles. It may also cause columns to resize in all sections. If DocumentMode is PAGELESS, this property will not be rendered.

	MarginTop Dimension `json:"marginTop,omitempty\"` // The top page margin. Updating the top page margin on the document style clears the top page margin on all section styles. If DocumentMode is PAGELESS, this property will not be rendered.

	PageNumberStart int `json:"pageNumberStart,omitempty\"` // The page number from which to start counting the number of pages. If DocumentMode is PAGELESS, this property will not be rendered.

	PageSize Size `json:"pageSize,omitempty\"` // The size of a page in the document. If DocumentMode is PAGELESS, this property will not be rendered.

	UseCustomHeaderFooterMargins bool `json:"useCustomHeaderFooterMargins,omitempty\"` // Indicates whether DocumentStyle margin_header, SectionStyle margin_header and DocumentStyle margin_footer, SectionStyle margin_footer are respected. When false, the default values in the Docs editor for header and footer margin is used. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	UseEvenPageHeaderFooter bool `json:"useEvenPageHeaderFooter,omitempty\"` // Indicates whether to use the even page header / footer IDs for the even pages. If DocumentMode is PAGELESS, this property will not be rendered.

	UseFirstPageHeaderFooter bool `json:"useFirstPageHeaderFooter,omitempty\"` // Indicates whether to use the first page header / footer IDs for the first page. If DocumentMode is PAGELESS, this property will not be rendered.

}

// A mask that indicates which of the fields on the base DocumentStyle have been changed in this suggestion. For any field set to true, there's a new suggested value.
type DocumentStyleSuggestionState struct {
	BackgroundSuggestionState BackgroundSuggestionState `json:"backgroundSuggestionState,omitempty\"` // A mask that indicates which of the fields in background have been changed in this suggestion.

	DefaultFooterIdSuggested bool `json:"defaultFooterIdSuggested,omitempty\"` // Indicates if there was a suggested change to default_footer_id.

	DefaultHeaderIdSuggested bool `json:"defaultHeaderIdSuggested,omitempty\"` // Indicates if there was a suggested change to default_header_id.

	EvenPageFooterIdSuggested bool `json:"evenPageFooterIdSuggested,omitempty\"` // Indicates if there was a suggested change to even_page_footer_id.

	EvenPageHeaderIdSuggested bool `json:"evenPageHeaderIdSuggested,omitempty\"` // Indicates if there was a suggested change to even_page_header_id.

	FirstPageFooterIdSuggested bool `json:"firstPageFooterIdSuggested,omitempty\"` // Indicates if there was a suggested change to first_page_footer_id.

	FirstPageHeaderIdSuggested bool `json:"firstPageHeaderIdSuggested,omitempty\"` // Indicates if there was a suggested change to first_page_header_id.

	FlipPageOrientationSuggested bool `json:"flipPageOrientationSuggested,omitempty\"` // Optional. Indicates if there was a suggested change to flip_page_orientation.

	MarginBottomSuggested bool `json:"marginBottomSuggested,omitempty\"` // Indicates if there was a suggested change to margin_bottom.

	MarginFooterSuggested bool `json:"marginFooterSuggested,omitempty\"` // Indicates if there was a suggested change to margin_footer.

	MarginHeaderSuggested bool `json:"marginHeaderSuggested,omitempty\"` // Indicates if there was a suggested change to margin_header.

	MarginLeftSuggested bool `json:"marginLeftSuggested,omitempty\"` // Indicates if there was a suggested change to margin_left.

	MarginRightSuggested bool `json:"marginRightSuggested,omitempty\"` // Indicates if there was a suggested change to margin_right.

	MarginTopSuggested bool `json:"marginTopSuggested,omitempty\"` // Indicates if there was a suggested change to margin_top.

	PageNumberStartSuggested bool `json:"pageNumberStartSuggested,omitempty\"` // Indicates if there was a suggested change to page_number_start.

	PageSizeSuggestionState SizeSuggestionState `json:"pageSizeSuggestionState,omitempty\"` // A mask that indicates which of the fields in size have been changed in this suggestion.

	UseCustomHeaderFooterMarginsSuggested bool `json:"useCustomHeaderFooterMarginsSuggested,omitempty\"` // Indicates if there was a suggested change to use_custom_header_footer_margins.

	UseEvenPageHeaderFooterSuggested bool `json:"useEvenPageHeaderFooterSuggested,omitempty\"` // Indicates if there was a suggested change to use_even_page_header_footer.

	UseFirstPageHeaderFooterSuggested bool `json:"useFirstPageHeaderFooterSuggested,omitempty\"` // Indicates if there was a suggested change to use_first_page_header_footer.

}

// A tab with document contents.
type DocumentTab struct {
	Body Body `json:"body,omitempty\"` // The main body of the document tab.

	DocumentStyle DocumentStyle `json:"documentStyle,omitempty\"` // The style of the document tab.

	Footers map[string]interface{} `json:"footers,omitempty\"` // The footers in the document tab, keyed by footer ID.

	Footnotes map[string]interface{} `json:"footnotes,omitempty\"` // The footnotes in the document tab, keyed by footnote ID.

	Headers map[string]interface{} `json:"headers,omitempty\"` // The headers in the document tab, keyed by header ID.

	InlineObjects map[string]interface{} `json:"inlineObjects,omitempty\"` // The inline objects in the document tab, keyed by object ID.

	Lists map[string]interface{} `json:"lists,omitempty\"` // The lists in the document tab, keyed by list ID.

	NamedRanges map[string]interface{} `json:"namedRanges,omitempty\"` // The named ranges in the document tab, keyed by name.

	NamedStyles NamedStyles `json:"namedStyles,omitempty\"` // The named styles of the document tab.

	PositionedObjects map[string]interface{} `json:"positionedObjects,omitempty\"` // The positioned objects in the document tab, keyed by object ID.

	SuggestedDocumentStyleChanges map[string]interface{} `json:"suggestedDocumentStyleChanges,omitempty\"` // The suggested changes to the style of the document tab, keyed by suggestion ID.

	SuggestedNamedStylesChanges map[string]interface{} `json:"suggestedNamedStylesChanges,omitempty\"` // The suggested changes to the named styles of the document tab, keyed by suggestion ID.

}

// The properties of an embedded drawing and used to differentiate the object type. An embedded drawing is one that's created and edited within a document. Note that extensive details are not supported.
type EmbeddedDrawingProperties struct {
}

// A mask that indicates which of the fields on the base EmbeddedDrawingProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type EmbeddedDrawingPropertiesSuggestionState struct {
}

// An embedded object in the document.
type EmbeddedObject struct {
	Description string `json:"description,omitempty\"` // The description of the embedded object. The `title` and `description` are both combined to display alt text.

	EmbeddedDrawingProperties EmbeddedDrawingProperties `json:"embeddedDrawingProperties,omitempty\"` // The properties of an embedded drawing.

	EmbeddedObjectBorder EmbeddedObjectBorder `json:"embeddedObjectBorder,omitempty\"` // The border of the embedded object.

	ImageProperties ImageProperties `json:"imageProperties,omitempty\"` // The properties of an image.

	LinkedContentReference LinkedContentReference `json:"linkedContentReference,omitempty\"` // A reference to the external linked source content. For example, it contains a reference to the source Google Sheets chart when the embedded object is a linked chart. If unset, then the embedded object is not linked.

	MarginBottom Dimension `json:"marginBottom,omitempty\"` // The bottom margin of the embedded object.

	MarginLeft Dimension `json:"marginLeft,omitempty\"` // The left margin of the embedded object.

	MarginRight Dimension `json:"marginRight,omitempty\"` // The right margin of the embedded object.

	MarginTop Dimension `json:"marginTop,omitempty\"` // The top margin of the embedded object.

	Size Size `json:"size,omitempty\"` // The visible size of the image after cropping.

	Title string `json:"title,omitempty\"` // The title of the embedded object. The `title` and `description` are both combined to display alt text.

}

// A border around an EmbeddedObject.
type EmbeddedObjectBorder struct {
	Color OptionalColor `json:"color,omitempty\"` // The color of the border.

	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the border.

	PropertyState string `json:"propertyState,omitempty\"` // The property state of the border property.

	Width Dimension `json:"width,omitempty\"` // The width of the border.

}

// A mask that indicates which of the fields on the base EmbeddedObjectBorder have been changed in this suggestion. For any field set to true, there's a new suggested value.
type EmbeddedObjectBorderSuggestionState struct {
	ColorSuggested bool `json:"colorSuggested,omitempty\"` // Indicates if there was a suggested change to color.

	DashStyleSuggested bool `json:"dashStyleSuggested,omitempty\"` // Indicates if there was a suggested change to dash_style.

	PropertyStateSuggested bool `json:"propertyStateSuggested,omitempty\"` // Indicates if there was a suggested change to property_state.

	WidthSuggested bool `json:"widthSuggested,omitempty\"` // Indicates if there was a suggested change to width.

}

// A mask that indicates which of the fields on the base EmbeddedObject have been changed in this suggestion. For any field set to true, there's a new suggested value.
type EmbeddedObjectSuggestionState struct {
	DescriptionSuggested bool `json:"descriptionSuggested,omitempty\"` // Indicates if there was a suggested change to description.

	EmbeddedDrawingPropertiesSuggestionState EmbeddedDrawingPropertiesSuggestionState `json:"embeddedDrawingPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields in embedded_drawing_properties have been changed in this suggestion.

	EmbeddedObjectBorderSuggestionState EmbeddedObjectBorderSuggestionState `json:"embeddedObjectBorderSuggestionState,omitempty\"` // A mask that indicates which of the fields in embedded_object_border have been changed in this suggestion.

	ImagePropertiesSuggestionState ImagePropertiesSuggestionState `json:"imagePropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields in image_properties have been changed in this suggestion.

	LinkedContentReferenceSuggestionState LinkedContentReferenceSuggestionState `json:"linkedContentReferenceSuggestionState,omitempty\"` // A mask that indicates which of the fields in linked_content_reference have been changed in this suggestion.

	MarginBottomSuggested bool `json:"marginBottomSuggested,omitempty\"` // Indicates if there was a suggested change to margin_bottom.

	MarginLeftSuggested bool `json:"marginLeftSuggested,omitempty\"` // Indicates if there was a suggested change to margin_left.

	MarginRightSuggested bool `json:"marginRightSuggested,omitempty\"` // Indicates if there was a suggested change to margin_right.

	MarginTopSuggested bool `json:"marginTopSuggested,omitempty\"` // Indicates if there was a suggested change to margin_top.

	SizeSuggestionState SizeSuggestionState `json:"sizeSuggestionState,omitempty\"` // A mask that indicates which of the fields in size have been changed in this suggestion.

	TitleSuggested bool `json:"titleSuggested,omitempty\"` // Indicates if there was a suggested change to title.

}

// Location at the end of a body, header, footer or footnote. The location is immediately before the last newline in the document segment.
type EndOfSegmentLocation struct {
	SegmentId string `json:"segmentId,omitempty\"` // The ID of the header, footer or footnote the location is in. An empty segment ID signifies the document's body.

	TabId string `json:"tabId,omitempty\"` // The tab that the location is in. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// A ParagraphElement representing an equation.
type Equation struct {
	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. An Equation may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

}

// A document footer.
type Footer struct {
	Content []StructuralElement `json:"content,omitempty\"` // The contents of the footer. The indexes for a footer's content begin at zero.

	FooterId string `json:"footerId,omitempty\"` // The ID of the footer.

}

// A document footnote.
type Footnote struct {
	Content []StructuralElement `json:"content,omitempty\"` // The contents of the footnote. The indexes for a footnote's content begin at zero.

	FootnoteId string `json:"footnoteId,omitempty\"` // The ID of the footnote.

}

// A ParagraphElement representing a footnote reference. A footnote reference is the inline content rendered with a number and is used to identify the footnote.
type FootnoteReference struct {
	FootnoteId string `json:"footnoteId,omitempty\"` // The ID of the footnote that contains the content of this footnote reference.

	FootnoteNumber string `json:"footnoteNumber,omitempty\"` // The rendered number of this footnote.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A FootnoteReference may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this FootnoteReference, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this FootnoteReference.

}

// A document header.
type Header struct {
	Content []StructuralElement `json:"content,omitempty\"` // The contents of the header. The indexes for a header's content begin at zero.

	HeaderId string `json:"headerId,omitempty\"` // The ID of the header.

}

// A reference to a heading in this document.
type HeadingLink struct {
	Id string `json:"id,omitempty\"` // The ID of a heading in this document.

	TabId string `json:"tabId,omitempty\"` // The ID of the tab containing this heading.

}

// A ParagraphElement representing a horizontal line.
type HorizontalRule struct {
	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A HorizontalRule may have multiple insertion IDs if it is a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this HorizontalRule, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this HorizontalRule. Similar to text content, like text runs and footnote references, the text style of a horizontal rule can affect content layout as well as the styling of text inserted next to it.

}

// The properties of an image.
type ImageProperties struct {
	Angle float64 `json:"angle,omitempty\"` // The clockwise rotation angle of the image, in radians.

	Brightness float64 `json:"brightness,omitempty\"` // The brightness effect of the image. The value should be in the interval [-1.0, 1.0], where 0 means no effect.

	ContentUri string `json:"contentUri,omitempty\"` // A URI to the image with a default lifetime of 30 minutes. This URI is tagged with the account of the requester. Anyone with the URI effectively accesses the image as the original requester. Access to the image may be lost if the document's sharing settings change.

	Contrast float64 `json:"contrast,omitempty\"` // The contrast effect of the image. The value should be in the interval [-1.0, 1.0], where 0 means no effect.

	CropProperties CropProperties `json:"cropProperties,omitempty\"` // The crop properties of the image.

	SourceUri string `json:"sourceUri,omitempty\"` // The source URI is the URI used to insert the image. The source URI can be empty.

	Transparency float64 `json:"transparency,omitempty\"` // The transparency effect of the image. The value should be in the interval [0.0, 1.0], where 0 means no effect and 1 means transparent.

}

// A mask that indicates which of the fields on the base ImageProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type ImagePropertiesSuggestionState struct {
	AngleSuggested bool `json:"angleSuggested,omitempty\"` // Indicates if there was a suggested change to angle.

	BrightnessSuggested bool `json:"brightnessSuggested,omitempty\"` // Indicates if there was a suggested change to brightness.

	ContentUriSuggested bool `json:"contentUriSuggested,omitempty\"` // Indicates if there was a suggested change to content_uri.

	ContrastSuggested bool `json:"contrastSuggested,omitempty\"` // Indicates if there was a suggested change to contrast.

	CropPropertiesSuggestionState CropPropertiesSuggestionState `json:"cropPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields in crop_properties have been changed in this suggestion.

	SourceUriSuggested bool `json:"sourceUriSuggested,omitempty\"` // Indicates if there was a suggested change to source_uri.

	TransparencySuggested bool `json:"transparencySuggested,omitempty\"` // Indicates if there was a suggested change to transparency.

}

// An object that appears inline with text. An InlineObject contains an EmbeddedObject such as an image.
type InlineObject struct {
	InlineObjectProperties InlineObjectProperties `json:"inlineObjectProperties,omitempty\"` // The properties of this inline object.

	ObjectId string `json:"objectId,omitempty\"` // The ID of this inline object. Can be used to update an object’s properties.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInlineObjectPropertiesChanges map[string]interface{} `json:"suggestedInlineObjectPropertiesChanges,omitempty\"` // The suggested changes to the inline object properties, keyed by suggestion ID.

	SuggestedInsertionId string `json:"suggestedInsertionId,omitempty\"` // The suggested insertion ID. If empty, then this is not a suggested insertion.

}

// A ParagraphElement that contains an InlineObject.
type InlineObjectElement struct {
	InlineObjectId string `json:"inlineObjectId,omitempty\"` // The ID of the InlineObject this element contains.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. An InlineObjectElement may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this InlineObject, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this InlineObjectElement. Similar to text content, like text runs and footnote references, the text style of an inline object element can affect content layout as well as the styling of text inserted next to it.

}

// Properties of an InlineObject.
type InlineObjectProperties struct {
	EmbeddedObject EmbeddedObject `json:"embeddedObject,omitempty\"` // The embedded object of this inline object.

}

// A mask that indicates which of the fields on the base InlineObjectProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type InlineObjectPropertiesSuggestionState struct {
	EmbeddedObjectSuggestionState EmbeddedObjectSuggestionState `json:"embeddedObjectSuggestionState,omitempty\"` // A mask that indicates which of the fields in embedded_object have been changed in this suggestion.

}

// Inserts a date at the specified location.
type InsertDateRequest struct {
	DateElementProperties DateElementProperties `json:"dateElementProperties,omitempty\"` // The properties of the date to insert.

	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the date at the end of the given header, footer or document body.

	Location Location `json:"location,omitempty\"` // Inserts the date at a specific index in the document. The date must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between an existing table and its preceding paragraph).

}

// Inserts an InlineObject containing an image at the given location.
type InsertInlineImageRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the text at the end of a header, footer or the document body. Inline images cannot be inserted inside a footnote.

	Location Location `json:"location,omitempty\"` // Inserts the image at a specific index in the document. The image must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). Inline images cannot be inserted inside a footnote or equation.

	ObjectSize Size `json:"objectSize,omitempty\"` // The size that the image should appear as in the document. This property is optional and the final size of the image in the document is determined by the following rules: * If neither width nor height is specified, then a default size of the image is calculated based on its resolution. * If one dimension is specified then the other dimension is calculated to preserve the aspect ratio of the image. * If both width and height are specified, the image is scaled to fit within the provided dimensions while maintaining its aspect ratio.

	Uri string `json:"uri,omitempty\"` // The image URI. The image is fetched once at insertion time and a copy is stored for display inside the document. Images must be less than 50MB in size, cannot exceed 25 megapixels, and must be in one of PNG, JPEG, or GIF format. The provided URI must be publicly accessible and at most 2 kB in length. The URI itself is saved with the image, and exposed via the ImageProperties.content_uri field.

}

// The result of inserting an inline image.
type InsertInlineImageResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The ID of the created InlineObject.

}

// The result of inserting an embedded Google Sheets chart.
type InsertInlineSheetsChartResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the inserted chart.

}

// Inserts a page break followed by a newline at the specified location.
type InsertPageBreakRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the page break at the end of the document body. Page breaks cannot be inserted inside a footnote, header or footer. Since page breaks can only be inserted inside the body, the segment ID field must be empty.

	Location Location `json:"location,omitempty\"` // Inserts the page break at a specific index in the document. The page break must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). Page breaks cannot be inserted inside a table, equation, footnote, header or footer. Since page breaks can only be inserted inside the body, the segment ID field must be empty.

}

// Inserts a person mention.
type InsertPersonRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the person mention at the end of a header, footer, footnote or the document body.

	Location Location `json:"location,omitempty\"` // Inserts the person mention at a specific index in the document. The person mention must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). Person mentions cannot be inserted inside an equation.

	PersonProperties PersonProperties `json:"personProperties,omitempty\"` // The properties of the person mention to insert.

}

// Inserts a RichLink at the specified location.
type InsertRichLinkRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the rich link at the end of a header, footer, footnote or the document body.

	Location Location `json:"location,omitempty\"` // Inserts the rich link at a specific index in the document. The rich link must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). The rich link cannot be inserted inside an equation.

	RichLinkProperties RichLinkProperties `json:"richLinkProperties,omitempty\"` // The properties of the rich link to insert.

}

// Inserts a section break at the given location. A newline character will be inserted before the section break.
type InsertSectionBreakRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts a newline and a section break at the end of the document body. Section breaks cannot be inserted inside a footnote, header or footer. Because section breaks can only be inserted inside the body, the segment ID field must be empty.

	Location Location `json:"location,omitempty\"` // Inserts a newline and a section break at a specific index in the document. The section break must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). Section breaks cannot be inserted inside a table, equation, footnote, header, or footer. Since section breaks can only be inserted inside the body, the segment ID field must be empty.

	SectionType string `json:"sectionType,omitempty\"` // The type of section to insert.

}

// Inserts an empty column into a table.
type InsertTableColumnRequest struct {
	InsertRight bool `json:"insertRight,omitempty\"` // Whether to insert new column to the right of the reference cell location. - `True`: insert to the right. - `False`: insert to the left.

	TableCellLocation TableCellLocation `json:"tableCellLocation,omitempty\"` // The reference table cell location from which columns will be inserted. A new column will be inserted to the left (or right) of the column where the reference cell is. If the reference cell is a merged cell, a new column will be inserted to the left (or right) of the merged cell.

}

// Inserts a table at the specified location. A newline character will be inserted before the inserted table.
type InsertTableRequest struct {
	Columns int `json:"columns,omitempty\"` // The number of columns in the table.

	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the table at the end of the given header, footer or document body. A newline character will be inserted before the inserted table. Tables cannot be inserted inside a footnote.

	Location Location `json:"location,omitempty\"` // Inserts the table at a specific model index. A newline character will be inserted before the inserted table, therefore the table start index will be at the specified location index + 1. The table must be inserted inside the bounds of an existing Paragraph. For instance, it cannot be inserted at a table's start index (i.e. between an existing table and its preceding paragraph). Tables cannot be inserted inside a footnote or equation.

	Rows int `json:"rows,omitempty\"` // The number of rows in the table.

}

// Inserts an empty row into a table.
type InsertTableRowRequest struct {
	InsertBelow bool `json:"insertBelow,omitempty\"` // Whether to insert new row below the reference cell location. - `True`: insert below the cell. - `False`: insert above the cell.

	TableCellLocation TableCellLocation `json:"tableCellLocation,omitempty\"` // The reference table cell location from which rows will be inserted. A new row will be inserted above (or below) the row where the reference cell is. If the reference cell is a merged cell, a new row will be inserted above (or below) the merged cell.

}

// Inserts text at the specified location.
type InsertTextRequest struct {
	EndOfSegmentLocation EndOfSegmentLocation `json:"endOfSegmentLocation,omitempty\"` // Inserts the text at the end of a header, footer, footnote or the document body.

	Location Location `json:"location,omitempty\"` // Inserts the text at a specific index in the document. Text must be inserted inside the bounds of an existing Paragraph. For instance, text cannot be inserted at a table's start index (i.e. between the table and its preceding paragraph). The text must be inserted in the preceding paragraph.

	Text string `json:"text,omitempty\"` // The text to be inserted. Inserting a newline character will implicitly create a new Paragraph at that index. The paragraph style of the new paragraph will be copied from the paragraph at the current insertion index, including lists and bullets. Text styles for inserted text will be determined automatically, generally preserving the styling of neighboring text. In most cases, the text style for the inserted text will match the text immediately before the insertion index. Some control characters (U+0000-U+0008, U+000C-U+001F) and characters from the Unicode Basic Multilingual Plane Private Use Area (U+E000-U+F8FF) will be stripped out of the inserted text.

}

// A reference to another portion of a document or an external URL resource.
type Link struct {
	Bookmark BookmarkLink `json:"bookmark,omitempty\"` // A bookmark in this document. In documents containing a single tab, links to bookmarks within the singular tab continue to return Link.bookmarkId when the includeTabsContent parameter is set to `false` or unset. Otherwise, this field is returned.

	BookmarkId string `json:"bookmarkId,omitempty\"` // The ID of a bookmark in this document. Legacy field: Instead, set includeTabsContent to `true` and use Link.bookmark for read and write operations. This field is only returned when includeTabsContent is set to `false` in documents containing a single tab and links to a bookmark within the singular tab. Otherwise, Link.bookmark is returned. If this field is used in a write request, the bookmark is considered to be from the tab ID specified in the request. If a tab ID is not specified in the request, it is considered to be from the first tab in the document.

	Heading HeadingLink `json:"heading,omitempty\"` // A heading in this document. In documents containing a single tab, links to headings within the singular tab continue to return Link.headingId when the includeTabsContent parameter is set to `false` or unset. Otherwise, this field is returned.

	HeadingId string `json:"headingId,omitempty\"` // The ID of a heading in this document. Legacy field: Instead, set includeTabsContent to `true` and use Link.heading for read and write operations. This field is only returned when includeTabsContent is set to `false` in documents containing a single tab and links to a heading within the singular tab. Otherwise, Link.heading is returned. If this field is used in a write request, the heading is considered to be from the tab ID specified in the request. If a tab ID is not specified in the request, it is considered to be from the first tab in the document.

	TabId string `json:"tabId,omitempty\"` // The ID of a tab in this document.

	Url string `json:"url,omitempty\"` // An external URL.

}

// A reference to the external linked source content.
type LinkedContentReference struct {
	SheetsChartReference SheetsChartReference `json:"sheetsChartReference,omitempty\"` // A reference to the linked chart.

}

// A mask that indicates which of the fields on the base LinkedContentReference have been changed in this suggestion. For any field set to true, there's a new suggested value.
type LinkedContentReferenceSuggestionState struct {
	SheetsChartReferenceSuggestionState SheetsChartReferenceSuggestionState `json:"sheetsChartReferenceSuggestionState,omitempty\"` // A mask that indicates which of the fields in sheets_chart_reference have been changed in this suggestion.

}

// A List represents the list attributes for a group of paragraphs that all belong to the same list. A paragraph that's part of a list has a reference to the list's ID in its bullet.
type List struct {
	ListProperties ListProperties `json:"listProperties,omitempty\"` // The properties of the list.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this list.

	SuggestedInsertionId string `json:"suggestedInsertionId,omitempty\"` // The suggested insertion ID. If empty, then this is not a suggested insertion.

	SuggestedListPropertiesChanges map[string]interface{} `json:"suggestedListPropertiesChanges,omitempty\"` // The suggested changes to the list properties, keyed by suggestion ID.

}

// The properties of a list that describe the look and feel of bullets belonging to paragraphs associated with a list.
type ListProperties struct {
	NestingLevels []NestingLevel `json:"nestingLevels,omitempty\"` // Describes the properties of the bullets at the associated level. A list has at most 9 levels of nesting with nesting level 0 corresponding to the top-most level and nesting level 8 corresponding to the most nested level. The nesting levels are returned in ascending order with the least nested returned first.

}

// A mask that indicates which of the fields on the base ListProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type ListPropertiesSuggestionState struct {
	NestingLevelsSuggestionStates []NestingLevelSuggestionState `json:"nestingLevelsSuggestionStates,omitempty\"` // A mask that indicates which of the fields on the corresponding NestingLevel in nesting_levels have been changed in this suggestion. The nesting level suggestion states are returned in ascending order of the nesting level with the least nested returned first.

}

// A particular location in the document.
type Location struct {
	Index int `json:"index,omitempty\"` // The zero-based index, in UTF-16 code units. The index is relative to the beginning of the segment specified by segment_id.

	SegmentId string `json:"segmentId,omitempty\"` // The ID of the header, footer or footnote the location is in. An empty segment ID signifies the document's body.

	TabId string `json:"tabId,omitempty\"` // The tab that the location is in. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// Merges cells in a Table.
type MergeTableCellsRequest struct {
	TableRange TableRange `json:"tableRange,omitempty\"` // The table range specifying which cells of the table to merge. Any text in the cells being merged will be concatenated and stored in the "head" cell of the range. This is the upper-left cell of the range when the content direction is left to right, and the upper-right cell of the range otherwise. If the range is non-rectangular (which can occur in some cases where the range covers cells that are already merged or where the table is non-rectangular), a 400 bad request error is returned.

}

// A collection of Ranges with the same named range ID. Named ranges allow developers to associate parts of a document with an arbitrary user-defined label so their contents can be programmatically read or edited later. A document can contain multiple named ranges with the same name, but every named range has a unique ID. A named range is created with a single Range, and content inserted inside a named range generally expands that range. However, certain document changes can cause the range to be split into multiple ranges. Named ranges are not private. All applications and collaborators that have access to the document can see its named ranges.
type NamedRange struct {
	Name string `json:"name,omitempty\"` // The name of the named range.

	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the named range.

	Ranges []RangeValue `json:"ranges,omitempty\"` // The ranges that belong to this named range.

}

// A collection of all the NamedRanges in the document that share a given name.
type NamedRanges struct {
	Name string `json:"name,omitempty\"` // The name that all the named ranges share.

	NamedRanges []NamedRange `json:"namedRanges,omitempty\"` // The NamedRanges that share the same name.

}

// A named style. Paragraphs in the document can inherit their TextStyle and ParagraphStyle from this named style when they have the same named style type.
type NamedStyle struct {
	NamedStyleType string `json:"namedStyleType,omitempty\"` // The type of this named style.

	ParagraphStyle ParagraphStyle `json:"paragraphStyle,omitempty\"` // The paragraph style of this named style.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this named style.

}

// A suggestion state of a NamedStyle message.
type NamedStyleSuggestionState struct {
	NamedStyleType string `json:"namedStyleType,omitempty\"` // The named style type that this suggestion state corresponds to. This field is provided as a convenience for matching the NamedStyleSuggestionState with its corresponding NamedStyle.

	ParagraphStyleSuggestionState ParagraphStyleSuggestionState `json:"paragraphStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields in paragraph style have been changed in this suggestion.

	TextStyleSuggestionState TextStyleSuggestionState `json:"textStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields in text style have been changed in this suggestion.

}

// The named styles. Paragraphs in the document can inherit their TextStyle and ParagraphStyle from these named styles.
type NamedStyles struct {
	Styles []NamedStyle `json:"styles,omitempty\"` // The named styles. There's an entry for each of the possible named style types.

}

// The suggestion state of a NamedStyles message.
type NamedStylesSuggestionState struct {
	StylesSuggestionStates []NamedStyleSuggestionState `json:"stylesSuggestionStates,omitempty\"` // A mask that indicates which of the fields on the corresponding NamedStyle in styles have been changed in this suggestion. The order of these named style suggestion states matches the order of the corresponding named style within the named styles suggestion.

}

// Contains properties describing the look and feel of a list bullet at a given level of nesting.
type NestingLevel struct {
	BulletAlignment string `json:"bulletAlignment,omitempty\"` // The alignment of the bullet within the space allotted for rendering the bullet.

	GlyphFormat string `json:"glyphFormat,omitempty\"` // The format string used by bullets at this level of nesting. The glyph format contains one or more placeholders, and these placeholders are replaced with the appropriate values depending on the glyph_type or glyph_symbol. The placeholders follow the pattern `%[nesting_level]`. Furthermore, placeholders can have prefixes and suffixes. Thus, the glyph format follows the pattern `%[nesting_level]`. Note that the prefix and suffix are optional and can be arbitrary strings. For example, the glyph format `%0.` indicates that the rendered glyph will replace the placeholder with the corresponding glyph for nesting level 0 followed by a period as the suffix. So a list with a glyph type of UPPER_ALPHA and glyph format `%0.` at nesting level 0 will result in a list with rendered glyphs `A.` `B.` `C.` The glyph format can contain placeholders for the current nesting level as well as placeholders for parent nesting levels. For example, a list can have a glyph format of `%0.` at nesting level 0 and a glyph format of `%0.%1.` at nesting level 1. Assuming both nesting levels have DECIMAL glyph types, this would result in a list with rendered glyphs `1.` `2.` ` 2.1.` ` 2.2.` `3.` For nesting levels that are ordered, the string that replaces a placeholder in the glyph format for a particular paragraph depends on the paragraph's order within the list.

	GlyphSymbol string `json:"glyphSymbol,omitempty\"` // A custom glyph symbol used by bullets when paragraphs at this level of nesting is unordered. The glyph symbol replaces placeholders within the glyph_format. For example, if the glyph_symbol is the solid circle corresponding to Unicode U+25cf code point and the glyph_format is `%0`, the rendered glyph would be the solid circle.

	GlyphType string `json:"glyphType,omitempty\"` // The type of glyph used by bullets when paragraphs at this level of nesting is ordered. The glyph type determines the type of glyph used to replace placeholders within the glyph_format when paragraphs at this level of nesting are ordered. For example, if the nesting level is 0, the glyph_format is `%0.` and the glyph type is DECIMAL, then the rendered glyph would replace the placeholder `%0` in the glyph format with a number corresponding to the list item's order within the list.

	IndentFirstLine Dimension `json:"indentFirstLine,omitempty\"` // The amount of indentation for the first line of paragraphs at this level of nesting.

	IndentStart Dimension `json:"indentStart,omitempty\"` // The amount of indentation for paragraphs at this level of nesting. Applied to the side that corresponds to the start of the text, based on the paragraph's content direction.

	StartNumber int `json:"startNumber,omitempty\"` // The number of the first list item at this nesting level. A value of 0 is treated as a value of 1 for lettered lists and Roman numeral lists. For values of both 0 and 1, lettered and Roman numeral lists will begin at `a` and `i` respectively. This value is ignored for nesting levels with unordered glyphs.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of bullets at this level of nesting.

}

// A mask that indicates which of the fields on the base NestingLevel have been changed in this suggestion. For any field set to true, there's a new suggested value.
type NestingLevelSuggestionState struct {
	BulletAlignmentSuggested bool `json:"bulletAlignmentSuggested,omitempty\"` // Indicates if there was a suggested change to bullet_alignment.

	GlyphFormatSuggested bool `json:"glyphFormatSuggested,omitempty\"` // Indicates if there was a suggested change to glyph_format.

	GlyphSymbolSuggested bool `json:"glyphSymbolSuggested,omitempty\"` // Indicates if there was a suggested change to glyph_symbol.

	GlyphTypeSuggested bool `json:"glyphTypeSuggested,omitempty\"` // Indicates if there was a suggested change to glyph_type.

	IndentFirstLineSuggested bool `json:"indentFirstLineSuggested,omitempty\"` // Indicates if there was a suggested change to indent_first_line.

	IndentStartSuggested bool `json:"indentStartSuggested,omitempty\"` // Indicates if there was a suggested change to indent_start.

	StartNumberSuggested bool `json:"startNumberSuggested,omitempty\"` // Indicates if there was a suggested change to start_number.

	TextStyleSuggestionState TextStyleSuggestionState `json:"textStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields in text style have been changed in this suggestion.

}

// A collection of object IDs.
type ObjectReferences struct {
	ObjectIds []string `json:"objectIds,omitempty\"` // The object IDs.

}

// A color that can either be fully opaque or fully transparent.
type OptionalColor struct {
	Color Color `json:"color,omitempty\"` // If set, this will be used as an opaque color. If unset, this represents a transparent color.

}

// A ParagraphElement representing a page break. A page break makes the subsequent text start at the top of the next page.
type PageBreak struct {
	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A PageBreak may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this PageBreak, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this PageBreak. Similar to text content, like text runs and footnote references, the text style of a page break can affect content layout as well as the styling of text inserted next to it.

}

// A StructuralElement representing a paragraph. A paragraph is a range of content that's terminated with a newline character.
type Paragraph struct {
	Bullet Bullet `json:"bullet,omitempty\"` // The bullet for this paragraph. If not present, the paragraph does not belong to a list.

	Elements []ParagraphElement `json:"elements,omitempty\"` // The content of the paragraph, broken down into its component parts.

	ParagraphStyle ParagraphStyle `json:"paragraphStyle,omitempty\"` // The style of this paragraph.

	PositionedObjectIds []string `json:"positionedObjectIds,omitempty\"` // The IDs of the positioned objects tethered to this paragraph.

	SuggestedBulletChanges map[string]interface{} `json:"suggestedBulletChanges,omitempty\"` // The suggested changes to this paragraph's bullet.

	SuggestedParagraphStyleChanges map[string]interface{} `json:"suggestedParagraphStyleChanges,omitempty\"` // The suggested paragraph style changes to this paragraph, keyed by suggestion ID.

	SuggestedPositionedObjectIds map[string]interface{} `json:"suggestedPositionedObjectIds,omitempty\"` // The IDs of the positioned objects suggested to be attached to this paragraph, keyed by suggestion ID.

}

// A border around a paragraph.
type ParagraphBorder struct {
	Color OptionalColor `json:"color,omitempty\"` // The color of the border.

	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the border.

	Padding Dimension `json:"padding,omitempty\"` // The padding of the border.

	Width Dimension `json:"width,omitempty\"` // The width of the border.

}

// A ParagraphElement describes content within a Paragraph.
type ParagraphElement struct {
	AutoText AutoText `json:"autoText,omitempty\"` // An auto text paragraph element.

	ColumnBreak ColumnBreak `json:"columnBreak,omitempty\"` // A column break paragraph element.

	DateElement DateElement `json:"dateElement,omitempty\"` // A paragraph element that represents a date.

	EndIndex int `json:"endIndex,omitempty\"` // The zero-base end index of this paragraph element, exclusive, in UTF-16 code units.

	Equation Equation `json:"equation,omitempty\"` // An equation paragraph element.

	FootnoteReference FootnoteReference `json:"footnoteReference,omitempty\"` // A footnote reference paragraph element.

	HorizontalRule HorizontalRule `json:"horizontalRule,omitempty\"` // A horizontal rule paragraph element.

	InlineObjectElement InlineObjectElement `json:"inlineObjectElement,omitempty\"` // An inline object paragraph element.

	PageBreak PageBreak `json:"pageBreak,omitempty\"` // A page break paragraph element.

	Person Person `json:"person,omitempty\"` // A paragraph element that links to a person or email address.

	RichLink RichLink `json:"richLink,omitempty\"` // A paragraph element that links to a Google resource (such as a file in Google Drive, a YouTube video, or a Calendar event.)

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this paragraph element, in UTF-16 code units.

	TextRun TextRun `json:"textRun,omitempty\"` // A text run paragraph element.

}

// Styles that apply to a whole paragraph. Inherited paragraph styles are represented as unset fields in this message. A paragraph style's parent depends on where the paragraph style is defined: * The ParagraphStyle on a Paragraph inherits from the paragraph's corresponding named style type. * The ParagraphStyle on a named style inherits from the normal text named style. * The ParagraphStyle of the normal text named style inherits from the default paragraph style in the Docs editor. * The ParagraphStyle on a Paragraph element that's contained in a table may inherit its paragraph style from the table style. If the paragraph style does not inherit from a parent, unsetting fields will revert the style to a value matching the defaults in the Docs editor.
type ParagraphStyle struct {
	Alignment string `json:"alignment,omitempty\"` // The text alignment for this paragraph.

	AvoidWidowAndOrphan bool `json:"avoidWidowAndOrphan,omitempty\"` // Whether to avoid widows and orphans for the paragraph. If unset, the value is inherited from the parent.

	BorderBetween ParagraphBorder `json:"borderBetween,omitempty\"` // The border between this paragraph and the next and previous paragraphs. If unset, the value is inherited from the parent. The between border is rendered when the adjacent paragraph has the same border and indent properties. Paragraph borders cannot be partially updated. When changing a paragraph border, the new border must be specified in its entirety.

	BorderBottom ParagraphBorder `json:"borderBottom,omitempty\"` // The border at the bottom of this paragraph. If unset, the value is inherited from the parent. The bottom border is rendered when the paragraph below has different border and indent properties. Paragraph borders cannot be partially updated. When changing a paragraph border, the new border must be specified in its entirety.

	BorderLeft ParagraphBorder `json:"borderLeft,omitempty\"` // The border to the left of this paragraph. If unset, the value is inherited from the parent. Paragraph borders cannot be partially updated. When changing a paragraph border, the new border must be specified in its entirety.

	BorderRight ParagraphBorder `json:"borderRight,omitempty\"` // The border to the right of this paragraph. If unset, the value is inherited from the parent. Paragraph borders cannot be partially updated. When changing a paragraph border, the new border must be specified in its entirety.

	BorderTop ParagraphBorder `json:"borderTop,omitempty\"` // The border at the top of this paragraph. If unset, the value is inherited from the parent. The top border is rendered when the paragraph above has different border and indent properties. Paragraph borders cannot be partially updated. When changing a paragraph border, the new border must be specified in its entirety.

	Direction string `json:"direction,omitempty\"` // The text direction of this paragraph. If unset, the value defaults to LEFT_TO_RIGHT since paragraph direction is not inherited.

	HeadingId string `json:"headingId,omitempty\"` // The heading ID of the paragraph. If empty, then this paragraph is not a heading. This property is read-only.

	IndentEnd Dimension `json:"indentEnd,omitempty\"` // The amount of indentation for the paragraph on the side that corresponds to the end of the text, based on the current paragraph direction. If unset, the value is inherited from the parent.

	IndentFirstLine Dimension `json:"indentFirstLine,omitempty\"` // The amount of indentation for the first line of the paragraph. If unset, the value is inherited from the parent.

	IndentStart Dimension `json:"indentStart,omitempty\"` // The amount of indentation for the paragraph on the side that corresponds to the start of the text, based on the current paragraph direction. If unset, the value is inherited from the parent.

	KeepLinesTogether bool `json:"keepLinesTogether,omitempty\"` // Whether all lines of the paragraph should be laid out on the same page or column if possible. If unset, the value is inherited from the parent.

	KeepWithNext bool `json:"keepWithNext,omitempty\"` // Whether at least a part of this paragraph should be laid out on the same page or column as the next paragraph if possible. If unset, the value is inherited from the parent.

	LineSpacing float64 `json:"lineSpacing,omitempty\"` // The amount of space between lines, as a percentage of normal, where normal is represented as 100.0. If unset, the value is inherited from the parent.

	NamedStyleType string `json:"namedStyleType,omitempty\"` // The named style type of the paragraph. Since updating the named style type affects other properties within ParagraphStyle, the named style type is applied before the other properties are updated.

	PageBreakBefore bool `json:"pageBreakBefore,omitempty\"` // Whether the current paragraph should always start at the beginning of a page. If unset, the value is inherited from the parent. Attempting to update page_break_before for paragraphs in unsupported regions, including Table, Header, Footer and Footnote, can result in an invalid document state that returns a 400 bad request error.

	Shading Shading `json:"shading,omitempty\"` // The shading of the paragraph. If unset, the value is inherited from the parent.

	SpaceAbove Dimension `json:"spaceAbove,omitempty\"` // The amount of extra space above the paragraph. If unset, the value is inherited from the parent.

	SpaceBelow Dimension `json:"spaceBelow,omitempty\"` // The amount of extra space below the paragraph. If unset, the value is inherited from the parent.

	SpacingMode string `json:"spacingMode,omitempty\"` // The spacing mode for the paragraph.

	TabStops []TabStop `json:"tabStops,omitempty\"` // A list of the tab stops for this paragraph. The list of tab stops is not inherited. This property is read-only.

}

// A mask that indicates which of the fields on the base ParagraphStyle have been changed in this suggestion. For any field set to true, there's a new suggested value.
type ParagraphStyleSuggestionState struct {
	AlignmentSuggested bool `json:"alignmentSuggested,omitempty\"` // Indicates if there was a suggested change to alignment.

	AvoidWidowAndOrphanSuggested bool `json:"avoidWidowAndOrphanSuggested,omitempty\"` // Indicates if there was a suggested change to avoid_widow_and_orphan.

	BorderBetweenSuggested bool `json:"borderBetweenSuggested,omitempty\"` // Indicates if there was a suggested change to border_between.

	BorderBottomSuggested bool `json:"borderBottomSuggested,omitempty\"` // Indicates if there was a suggested change to border_bottom.

	BorderLeftSuggested bool `json:"borderLeftSuggested,omitempty\"` // Indicates if there was a suggested change to border_left.

	BorderRightSuggested bool `json:"borderRightSuggested,omitempty\"` // Indicates if there was a suggested change to border_right.

	BorderTopSuggested bool `json:"borderTopSuggested,omitempty\"` // Indicates if there was a suggested change to border_top.

	DirectionSuggested bool `json:"directionSuggested,omitempty\"` // Indicates if there was a suggested change to direction.

	HeadingIdSuggested bool `json:"headingIdSuggested,omitempty\"` // Indicates if there was a suggested change to heading_id.

	IndentEndSuggested bool `json:"indentEndSuggested,omitempty\"` // Indicates if there was a suggested change to indent_end.

	IndentFirstLineSuggested bool `json:"indentFirstLineSuggested,omitempty\"` // Indicates if there was a suggested change to indent_first_line.

	IndentStartSuggested bool `json:"indentStartSuggested,omitempty\"` // Indicates if there was a suggested change to indent_start.

	KeepLinesTogetherSuggested bool `json:"keepLinesTogetherSuggested,omitempty\"` // Indicates if there was a suggested change to keep_lines_together.

	KeepWithNextSuggested bool `json:"keepWithNextSuggested,omitempty\"` // Indicates if there was a suggested change to keep_with_next.

	LineSpacingSuggested bool `json:"lineSpacingSuggested,omitempty\"` // Indicates if there was a suggested change to line_spacing.

	NamedStyleTypeSuggested bool `json:"namedStyleTypeSuggested,omitempty\"` // Indicates if there was a suggested change to named_style_type.

	PageBreakBeforeSuggested bool `json:"pageBreakBeforeSuggested,omitempty\"` // Indicates if there was a suggested change to page_break_before.

	ShadingSuggestionState ShadingSuggestionState `json:"shadingSuggestionState,omitempty\"` // A mask that indicates which of the fields in shading have been changed in this suggestion.

	SpaceAboveSuggested bool `json:"spaceAboveSuggested,omitempty\"` // Indicates if there was a suggested change to space_above.

	SpaceBelowSuggested bool `json:"spaceBelowSuggested,omitempty\"` // Indicates if there was a suggested change to space_below.

	SpacingModeSuggested bool `json:"spacingModeSuggested,omitempty\"` // Indicates if there was a suggested change to spacing_mode.

}

// A person or email address mentioned in a document. These mentions behave as a single, immutable element containing the person's name or email address.
type Person struct {
	PersonId string `json:"personId,omitempty\"` // Output only. The unique ID of this link.

	PersonProperties PersonProperties `json:"personProperties,omitempty\"` // Output only. The properties of this Person. This field is always present.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // IDs for suggestions that remove this person link from the document. A Person might have multiple deletion IDs if, for example, multiple users suggest deleting it. If empty, then this person link isn't suggested for deletion.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // IDs for suggestions that insert this person link into the document. A Person might have multiple insertion IDs if it's a nested suggested change (a suggestion within a suggestion made by a different user, for example). If empty, then this person link isn't a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this Person, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this Person.

}

// Properties specific to a linked Person.
type PersonProperties struct {
	Email string `json:"email,omitempty\"` // The email address linked to this Person. This field is always present.

	Name string `json:"name,omitempty\"` // The name of the person if it's displayed in the link text instead of the person's email address.

}

// Updates the number of pinned table header rows in a table.
type PinTableHeaderRowsRequest struct {
	PinnedHeaderRowsCount int `json:"pinnedHeaderRowsCount,omitempty\"` // The number of table rows to pin, where 0 implies that all rows are unpinned.

	TableStartLocation Location `json:"tableStartLocation,omitempty\"` // The location where the table starts in the document.

}

// An object that's tethered to a Paragraph and positioned relative to the beginning of the paragraph. A PositionedObject contains an EmbeddedObject such as an image.
type PositionedObject struct {
	ObjectId string `json:"objectId,omitempty\"` // The ID of this positioned object.

	PositionedObjectProperties PositionedObjectProperties `json:"positionedObjectProperties,omitempty\"` // The properties of this positioned object.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionId string `json:"suggestedInsertionId,omitempty\"` // The suggested insertion ID. If empty, then this is not a suggested insertion.

	SuggestedPositionedObjectPropertiesChanges map[string]interface{} `json:"suggestedPositionedObjectPropertiesChanges,omitempty\"` // The suggested changes to the positioned object properties, keyed by suggestion ID.

}

// The positioning of a PositionedObject. The positioned object is positioned relative to the beginning of the Paragraph it's tethered to.
type PositionedObjectPositioning struct {
	Layout string `json:"layout,omitempty\"` // The layout of this positioned object.

	LeftOffset Dimension `json:"leftOffset,omitempty\"` // The offset of the left edge of the positioned object relative to the beginning of the Paragraph it's tethered to. The exact positioning of the object can depend on other content in the document and the document's styling.

	TopOffset Dimension `json:"topOffset,omitempty\"` // The offset of the top edge of the positioned object relative to the beginning of the Paragraph it's tethered to. The exact positioning of the object can depend on other content in the document and the document's styling.

}

// A mask that indicates which of the fields on the base PositionedObjectPositioning have been changed in this suggestion. For any field set to true, there's a new suggested value.
type PositionedObjectPositioningSuggestionState struct {
	LayoutSuggested bool `json:"layoutSuggested,omitempty\"` // Indicates if there was a suggested change to layout.

	LeftOffsetSuggested bool `json:"leftOffsetSuggested,omitempty\"` // Indicates if there was a suggested change to left_offset.

	TopOffsetSuggested bool `json:"topOffsetSuggested,omitempty\"` // Indicates if there was a suggested change to top_offset.

}

// Properties of a PositionedObject.
type PositionedObjectProperties struct {
	EmbeddedObject EmbeddedObject `json:"embeddedObject,omitempty\"` // The embedded object of this positioned object.

	Positioning PositionedObjectPositioning `json:"positioning,omitempty\"` // The positioning of this positioned object relative to the newline of the Paragraph that references this positioned object.

}

// A mask that indicates which of the fields on the base PositionedObjectProperties have been changed in this suggestion. For any field set to true, there's a new suggested value.
type PositionedObjectPropertiesSuggestionState struct {
	EmbeddedObjectSuggestionState EmbeddedObjectSuggestionState `json:"embeddedObjectSuggestionState,omitempty\"` // A mask that indicates which of the fields in embedded_object have been changed in this suggestion.

	PositioningSuggestionState PositionedObjectPositioningSuggestionState `json:"positioningSuggestionState,omitempty\"` // A mask that indicates which of the fields in positioning have been changed in this suggestion.

}

// Specifies a contiguous range of text.
type RangeValue struct {
	EndIndex int `json:"endIndex,omitempty\"` // The zero-based end index of this range, exclusive, in UTF-16 code units. In all current uses, an end index must be provided. This field is an Int32Value in order to accommodate future use cases with open-ended ranges.

	SegmentId string `json:"segmentId,omitempty\"` // The ID of the header, footer, or footnote that this range is contained in. An empty segment ID signifies the document's body.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this range, in UTF-16 code units. In all current uses, a start index must be provided. This field is an Int32Value in order to accommodate future use cases with open-ended ranges.

	TabId string `json:"tabId,omitempty\"` // The tab that contains this range. When omitted, the request applies to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

}

// Replaces all instances of text matching a criteria with replace text.
type ReplaceAllTextRequest struct {
	ContainsText SubstringMatchCriteria `json:"containsText,omitempty\"` // Finds text in the document matching this substring.

	ReplaceText string `json:"replaceText,omitempty\"` // The text that will replace the matched text.

	TabsCriteria TabsCriteria `json:"tabsCriteria,omitempty\"` // Optional. The criteria used to specify in which tabs the replacement occurs. When omitted, the replacement applies to all tabs. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the replacement applies to the singular tab. In a document containing multiple tabs: - If provided, the replacement applies to the specified tabs. - If omitted, the replacement applies to all tabs.

}

// The result of replacing text.
type ReplaceAllTextResponse struct {
	OccurrencesChanged int `json:"occurrencesChanged,omitempty\"` // The number of occurrences changed by replacing all text.

}

// Replaces an existing image with a new image. Replacing an image removes some image effects from the existing image in order to mirror the behavior of the Docs editor.
type ReplaceImageRequest struct {
	ImageObjectId string `json:"imageObjectId,omitempty\"` // The ID of the existing image that will be replaced. The ID can be retrieved from the response of a get request.

	ImageReplaceMethod string `json:"imageReplaceMethod,omitempty\"` // The replacement method.

	TabId string `json:"tabId,omitempty\"` // The tab that the image to be replaced is in. When omitted, the request is applied to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If omitted, the request applies to the first tab in the document.

	Uri string `json:"uri,omitempty\"` // The URI of the new image. The image is fetched once at insertion time and a copy is stored for display inside the document. Images must be less than 50MB, cannot exceed 25 megapixels, and must be in PNG, JPEG, or GIF format. The provided URI can't surpass 2 KB in length. The URI is saved with the image, and exposed through the ImageProperties.source_uri field.

}

// Replaces the contents of the specified NamedRange or NamedRanges with the given replacement content. Note that an individual NamedRange may consist of multiple discontinuous ranges. In this case, only the content in the first range will be replaced. The other ranges and their content will be deleted. In cases where replacing or deleting any ranges would result in an invalid document structure, a 400 bad request error is returned.
type ReplaceNamedRangeContentRequest struct {
	NamedRangeId string `json:"namedRangeId,omitempty\"` // The ID of the named range whose content will be replaced. If there is no named range with the given ID a 400 bad request error is returned.

	NamedRangeName string `json:"namedRangeName,omitempty\"` // The name of the NamedRanges whose content will be replaced. If there are multiple named ranges with the given name, then the content of each one will be replaced. If there are no named ranges with the given name, then the request will be a no-op.

	TabsCriteria TabsCriteria `json:"tabsCriteria,omitempty\"` // Optional. The criteria used to specify in which tabs the replacement occurs. When omitted, the replacement applies to all tabs. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the replacement applies to the singular tab. In a document containing multiple tabs: - If provided, the replacement applies to the specified tabs. - If omitted, the replacement applies to all tabs.

	Text string `json:"text,omitempty\"` // Replaces the content of the specified named range(s) with the given text.

}

// A single update to apply to a document.
type Request struct {
	AddDocumentTab AddDocumentTabRequest `json:"addDocumentTab,omitempty\"` // Adds a document tab.

	CreateFooter CreateFooterRequest `json:"createFooter,omitempty\"` // Creates a footer.

	CreateFootnote CreateFootnoteRequest `json:"createFootnote,omitempty\"` // Creates a footnote.

	CreateHeader CreateHeaderRequest `json:"createHeader,omitempty\"` // Creates a header.

	CreateNamedRange CreateNamedRangeRequest `json:"createNamedRange,omitempty\"` // Creates a named range.

	CreateParagraphBullets CreateParagraphBulletsRequest `json:"createParagraphBullets,omitempty\"` // Creates bullets for paragraphs.

	DeleteContentRange DeleteContentRangeRequest `json:"deleteContentRange,omitempty\"` // Deletes content from the document.

	DeleteFooter DeleteFooterRequest `json:"deleteFooter,omitempty\"` // Deletes a footer from the document.

	DeleteHeader DeleteHeaderRequest `json:"deleteHeader,omitempty\"` // Deletes a header from the document.

	DeleteNamedRange DeleteNamedRangeRequest `json:"deleteNamedRange,omitempty\"` // Deletes a named range.

	DeleteParagraphBullets DeleteParagraphBulletsRequest `json:"deleteParagraphBullets,omitempty\"` // Deletes bullets from paragraphs.

	DeletePositionedObject DeletePositionedObjectRequest `json:"deletePositionedObject,omitempty\"` // Deletes a positioned object from the document.

	DeleteTab DeleteTabRequest `json:"deleteTab,omitempty\"` // Deletes a document tab.

	DeleteTableColumn DeleteTableColumnRequest `json:"deleteTableColumn,omitempty\"` // Deletes a column from a table.

	DeleteTableRow DeleteTableRowRequest `json:"deleteTableRow,omitempty\"` // Deletes a row from a table.

	InsertDate InsertDateRequest `json:"insertDate,omitempty\"` // Inserts a date.

	InsertInlineImage InsertInlineImageRequest `json:"insertInlineImage,omitempty\"` // Inserts an inline image at the specified location.

	InsertPageBreak InsertPageBreakRequest `json:"insertPageBreak,omitempty\"` // Inserts a page break at the specified location.

	InsertPerson InsertPersonRequest `json:"insertPerson,omitempty\"` // Inserts a person mention.

	InsertRichLink InsertRichLinkRequest `json:"insertRichLink,omitempty\"` // Insert a rich link.

	InsertSectionBreak InsertSectionBreakRequest `json:"insertSectionBreak,omitempty\"` // Inserts a section break at the specified location.

	InsertTable InsertTableRequest `json:"insertTable,omitempty\"` // Inserts a table at the specified location.

	InsertTableColumn InsertTableColumnRequest `json:"insertTableColumn,omitempty\"` // Inserts an empty column into a table.

	InsertTableRow InsertTableRowRequest `json:"insertTableRow,omitempty\"` // Inserts an empty row into a table.

	InsertText InsertTextRequest `json:"insertText,omitempty\"` // Inserts text at the specified location.

	MergeTableCells MergeTableCellsRequest `json:"mergeTableCells,omitempty\"` // Merges cells in a table.

	PinTableHeaderRows PinTableHeaderRowsRequest `json:"pinTableHeaderRows,omitempty\"` // Updates the number of pinned header rows in a table.

	ReplaceAllText ReplaceAllTextRequest `json:"replaceAllText,omitempty\"` // Replaces all instances of the specified text.

	ReplaceImage ReplaceImageRequest `json:"replaceImage,omitempty\"` // Replaces an image in the document.

	ReplaceNamedRangeContent ReplaceNamedRangeContentRequest `json:"replaceNamedRangeContent,omitempty\"` // Replaces the content in a named range.

	UnmergeTableCells UnmergeTableCellsRequest `json:"unmergeTableCells,omitempty\"` // Unmerges cells in a table.

	UpdateDocumentStyle UpdateDocumentStyleRequest `json:"updateDocumentStyle,omitempty\"` // Updates the style of the document.

	UpdateDocumentTabProperties UpdateDocumentTabPropertiesRequest `json:"updateDocumentTabProperties,omitempty\"` // Updates the properties of a document tab.

	UpdateNamedStyle UpdateNamedStyleRequest `json:"updateNamedStyle,omitempty\"` // Updates a named style.

	UpdateParagraphStyle UpdateParagraphStyleRequest `json:"updateParagraphStyle,omitempty\"` // Updates the paragraph style at the specified range.

	UpdateSectionStyle UpdateSectionStyleRequest `json:"updateSectionStyle,omitempty\"` // Updates the section style of the specified range.

	UpdateTableCellStyle UpdateTableCellStyleRequest `json:"updateTableCellStyle,omitempty\"` // Updates the style of table cells.

	UpdateTableColumnProperties UpdateTableColumnPropertiesRequest `json:"updateTableColumnProperties,omitempty\"` // Updates the properties of columns in a table.

	UpdateTableRowStyle UpdateTableRowStyleRequest `json:"updateTableRowStyle,omitempty\"` // Updates the row style in a table.

	UpdateTextStyle UpdateTextStyleRequest `json:"updateTextStyle,omitempty\"` // Updates the text style at the specified range.

}

// A single response from an update.
type Response struct {
	AddDocumentTab AddDocumentTabResponse `json:"addDocumentTab,omitempty\"` // The result of adding a document tab.

	CreateFooter CreateFooterResponse `json:"createFooter,omitempty\"` // The result of creating a footer.

	CreateFootnote CreateFootnoteResponse `json:"createFootnote,omitempty\"` // The result of creating a footnote.

	CreateHeader CreateHeaderResponse `json:"createHeader,omitempty\"` // The result of creating a header.

	CreateNamedRange CreateNamedRangeResponse `json:"createNamedRange,omitempty\"` // The result of creating a named range.

	InsertInlineImage InsertInlineImageResponse `json:"insertInlineImage,omitempty\"` // The result of inserting an inline image.

	InsertInlineSheetsChart InsertInlineSheetsChartResponse `json:"insertInlineSheetsChart,omitempty\"` // The result of inserting an inline Google Sheets chart.

	ReplaceAllText ReplaceAllTextResponse `json:"replaceAllText,omitempty\"` // The result of replacing text.

}

// An RGB color.
type RgbColor struct {
	Blue float64 `json:"blue,omitempty\"` // The blue component of the color, from 0.0 to 1.0.

	Green float64 `json:"green,omitempty\"` // The green component of the color, from 0.0 to 1.0.

	Red float64 `json:"red,omitempty\"` // The red component of the color, from 0.0 to 1.0.

}

// A link to a Google resource (such as a file in Drive, a YouTube video, or a Calendar event).
type RichLink struct {
	RichLinkId string `json:"richLinkId,omitempty\"` // Output only. The ID of this link.

	RichLinkProperties RichLinkProperties `json:"richLinkProperties,omitempty\"` // Output only. The properties of this RichLink. This field is always present.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // IDs for suggestions that remove this link from the document. A RichLink might have multiple deletion IDs if, for example, multiple users suggest deleting it. If empty, then this person link isn't suggested for deletion.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // IDs for suggestions that insert this link into the document. A RichLink might have multiple insertion IDs if it's a nested suggested change (a suggestion within a suggestion made by a different user, for example). If empty, then this person link isn't a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this RichLink, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this RichLink.

}

// Properties specific to a RichLink.
type RichLinkProperties struct {
	MimeType string `json:"mimeType,omitempty\"` // The [MIME type](https://developers.google.com/drive/api/v3/mime-types) of the RichLink, if there's one (for example, when it's a file in Drive).

	Title string `json:"title,omitempty\"` // The title of the RichLink as displayed in the link. This title matches the title of the linked resource at the time of the insertion or last update of the link. This field is always present.

	Uri string `json:"uri,omitempty\"` // The URI to the RichLink. This is always present.

}

// A StructuralElement representing a section break. A section is a range of content that has the same SectionStyle. A section break represents the start of a new section, and the section style applies to the section after the section break. The document body always begins with a section break.
type SectionBreak struct {
	SectionStyle SectionStyle `json:"sectionStyle,omitempty\"` // The style of the section after this section break.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A SectionBreak may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

}

// Properties that apply to a section's column.
type SectionColumnProperties struct {
	PaddingEnd Dimension `json:"paddingEnd,omitempty\"` // The padding at the end of the column.

	Width Dimension `json:"width,omitempty\"` // Output only. The width of the column.

}

// The styling that applies to a section.
type SectionStyle struct {
	ColumnProperties []SectionColumnProperties `json:"columnProperties,omitempty\"` // The section's columns properties. If empty, the section contains one column with the default properties in the Docs editor. A section can be updated to have no more than 3 columns. When updating this property, setting a concrete value is required. Unsetting this property will result in a 400 bad request error.

	ColumnSeparatorStyle string `json:"columnSeparatorStyle,omitempty\"` // The style of column separators. This style can be set even when there's one column in the section. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	ContentDirection string `json:"contentDirection,omitempty\"` // The content direction of this section. If unset, the value defaults to LEFT_TO_RIGHT. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	DefaultFooterId string `json:"defaultFooterId,omitempty\"` // The ID of the default footer. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's default_footer_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	DefaultHeaderId string `json:"defaultHeaderId,omitempty\"` // The ID of the default header. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's default_header_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	EvenPageFooterId string `json:"evenPageFooterId,omitempty\"` // The ID of the footer used only for even pages. If the value of DocumentStyle's use_even_page_header_footer is true, this value is used for the footers on even pages in the section. If it is false, the footers on even pages use the default_footer_id. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's even_page_footer_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	EvenPageHeaderId string `json:"evenPageHeaderId,omitempty\"` // The ID of the header used only for even pages. If the value of DocumentStyle's use_even_page_header_footer is true, this value is used for the headers on even pages in the section. If it is false, the headers on even pages use the default_header_id. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's even_page_header_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FirstPageFooterId string `json:"firstPageFooterId,omitempty\"` // The ID of the footer used only for the first page of the section. If use_first_page_header_footer is true, this value is used for the footer on the first page of the section. If it's false, the footer on the first page of the section uses the default_footer_id. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's first_page_footer_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FirstPageHeaderId string `json:"firstPageHeaderId,omitempty\"` // The ID of the header used only for the first page of the section. If use_first_page_header_footer is true, this value is used for the header on the first page of the section. If it's false, the header on the first page of the section uses the default_header_id. If unset, the value inherits from the previous SectionBreak's SectionStyle. If the value is unset in the first SectionBreak, it inherits from DocumentStyle's first_page_header_id. If DocumentMode is PAGELESS, this property will not be rendered. This property is read-only.

	FlipPageOrientation bool `json:"flipPageOrientation,omitempty\"` // Optional. Indicates whether to flip the dimensions of DocumentStyle's page_size for this section, which allows changing the page orientation between portrait and landscape. If unset, the value inherits from DocumentStyle's flip_page_orientation. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginBottom Dimension `json:"marginBottom,omitempty\"` // The bottom page margin of the section. If unset, the value defaults to margin_bottom from DocumentStyle. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginFooter Dimension `json:"marginFooter,omitempty\"` // The footer margin of the section. If unset, the value defaults to margin_footer from DocumentStyle. If updated, use_custom_header_footer_margins is set to true on DocumentStyle. The value of use_custom_header_footer_margins on DocumentStyle indicates if a footer margin is being respected for this section If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginHeader Dimension `json:"marginHeader,omitempty\"` // The header margin of the section. If unset, the value defaults to margin_header from DocumentStyle. If updated, use_custom_header_footer_margins is set to true on DocumentStyle. The value of use_custom_header_footer_margins on DocumentStyle indicates if a header margin is being respected for this section. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginLeft Dimension `json:"marginLeft,omitempty\"` // The left page margin of the section. If unset, the value defaults to margin_left from DocumentStyle. Updating the left margin causes columns in this section to resize. Since the margin affects column width, it's applied before column properties. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginRight Dimension `json:"marginRight,omitempty\"` // The right page margin of the section. If unset, the value defaults to margin_right from DocumentStyle. Updating the right margin causes columns in this section to resize. Since the margin affects column width, it's applied before column properties. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	MarginTop Dimension `json:"marginTop,omitempty\"` // The top page margin of the section. If unset, the value defaults to margin_top from DocumentStyle. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	PageNumberStart int `json:"pageNumberStart,omitempty\"` // The page number from which to start counting the number of pages for this section. If unset, page numbering continues from the previous section. If the value is unset in the first SectionBreak, refer to DocumentStyle's page_number_start. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

	SectionType string `json:"sectionType,omitempty\"` // Output only. The type of section.

	UseFirstPageHeaderFooter bool `json:"useFirstPageHeaderFooter,omitempty\"` // Indicates whether to use the first page header / footer IDs for the first page of the section. If unset, it inherits from DocumentStyle's use_first_page_header_footer for the first section. If the value is unset for subsequent sectors, it should be interpreted as false. If DocumentMode is PAGELESS, this property will not be rendered. When updating this property, setting a concrete value is required. Unsetting this property results in a 400 bad request error.

}

// The shading of a paragraph.
type Shading struct {
	BackgroundColor OptionalColor `json:"backgroundColor,omitempty\"` // The background color of this paragraph shading.

}

// A mask that indicates which of the fields on the base Shading have been changed in this suggested change. For any field set to true, there's a new suggested value.
type ShadingSuggestionState struct {
	BackgroundColorSuggested bool `json:"backgroundColorSuggested,omitempty\"` // Indicates if there was a suggested change to the Shading.

}

// A reference to a linked chart embedded from Google Sheets.
type SheetsChartReference struct {
	ChartId int `json:"chartId,omitempty\"` // The ID of the specific chart in the Google Sheets spreadsheet that's embedded.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the Google Sheets spreadsheet that contains the source chart.

}

// A mask that indicates which of the fields on the base SheetsChartReference have been changed in this suggestion. For any field set to true, there's a new suggested value.
type SheetsChartReferenceSuggestionState struct {
	ChartIdSuggested bool `json:"chartIdSuggested,omitempty\"` // Indicates if there was a suggested change to chart_id.

	SpreadsheetIdSuggested bool `json:"spreadsheetIdSuggested,omitempty\"` // Indicates if there was a suggested change to spreadsheet_id.

}

// A width and height.
type Size struct {
	Height Dimension `json:"height,omitempty\"` // The height of the object.

	Width Dimension `json:"width,omitempty\"` // The width of the object.

}

// A mask that indicates which of the fields on the base Size have been changed in this suggestion. For any field set to true, the Size has a new suggested value.
type SizeSuggestionState struct {
	HeightSuggested bool `json:"heightSuggested,omitempty\"` // Indicates if there was a suggested change to height.

	WidthSuggested bool `json:"widthSuggested,omitempty\"` // Indicates if there was a suggested change to width.

}

// A StructuralElement describes content that provides structure to the document.
type StructuralElement struct {
	EndIndex int `json:"endIndex,omitempty\"` // The zero-based end index of this structural element, exclusive, in UTF-16 code units.

	Paragraph Paragraph `json:"paragraph,omitempty\"` // A paragraph type of structural element.

	SectionBreak SectionBreak `json:"sectionBreak,omitempty\"` // A section break type of structural element.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this structural element, in UTF-16 code units.

	Table Table `json:"table,omitempty\"` // A table type of structural element.

	TableOfContents TableOfContents `json:"tableOfContents,omitempty\"` // A table of contents type of structural element.

}

// A criteria that matches a specific string of text in the document.
type SubstringMatchCriteria struct {
	MatchCase bool `json:"matchCase,omitempty\"` // Indicates whether the search should respect case: - `True`: the search is case sensitive. - `False`: the search is case insensitive.

	SearchByRegex bool `json:"searchByRegex,omitempty\"` // Optional. True if the find value should be treated as a regular expression. Any backslashes in the pattern should be escaped. - `True`: the search text is treated as a regular expressions. - `False`: the search text is treated as a substring for matching.

	Text string `json:"text,omitempty\"` // The text to search for in the document.

}

// A suggested change to a Bullet.
type SuggestedBullet struct {
	Bullet Bullet `json:"bullet,omitempty\"` // A Bullet that only includes the changes made in this suggestion. This can be used along with the bullet_suggestion_state to see which fields have changed and their new values.

	BulletSuggestionState BulletSuggestionState `json:"bulletSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base Bullet have been changed in this suggestion.

}

// A suggested change to a DateElementProperties.
type SuggestedDateElementProperties struct {
	DateElementProperties DateElementProperties `json:"dateElementProperties,omitempty\"` // DateElementProperties that only includes the changes made in this suggestion. This can be used along with the date_element_properties_suggestion_state to see which fields have changed and their new values.

	DateElementPropertiesSuggestionState DateElementPropertiesSuggestionState `json:"dateElementPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base DateElementProperties have been changed in this suggestion.

}

// A suggested change to the DocumentStyle.
type SuggestedDocumentStyle struct {
	DocumentStyle DocumentStyle `json:"documentStyle,omitempty\"` // A DocumentStyle that only includes the changes made in this suggestion. This can be used along with the document_style_suggestion_state to see which fields have changed and their new values.

	DocumentStyleSuggestionState DocumentStyleSuggestionState `json:"documentStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base DocumentStyle have been changed in this suggestion.

}

// A suggested change to InlineObjectProperties.
type SuggestedInlineObjectProperties struct {
	InlineObjectProperties InlineObjectProperties `json:"inlineObjectProperties,omitempty\"` // An InlineObjectProperties that only includes the changes made in this suggestion. This can be used along with the inline_object_properties_suggestion_state to see which fields have changed and their new values.

	InlineObjectPropertiesSuggestionState InlineObjectPropertiesSuggestionState `json:"inlineObjectPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base InlineObjectProperties have been changed in this suggestion.

}

// A suggested change to ListProperties.
type SuggestedListProperties struct {
	ListProperties ListProperties `json:"listProperties,omitempty\"` // A ListProperties that only includes the changes made in this suggestion. This can be used along with the list_properties_suggestion_state to see which fields have changed and their new values.

	ListPropertiesSuggestionState ListPropertiesSuggestionState `json:"listPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base ListProperties have been changed in this suggestion.

}

// A suggested change to the NamedStyles.
type SuggestedNamedStyles struct {
	NamedStyles NamedStyles `json:"namedStyles,omitempty\"` // A NamedStyles that only includes the changes made in this suggestion. This can be used along with the named_styles_suggestion_state to see which fields have changed and their new values.

	NamedStylesSuggestionState NamedStylesSuggestionState `json:"namedStylesSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base NamedStyles have been changed in this suggestion.

}

// A suggested change to a ParagraphStyle.
type SuggestedParagraphStyle struct {
	ParagraphStyle ParagraphStyle `json:"paragraphStyle,omitempty\"` // A ParagraphStyle that only includes the changes made in this suggestion. This can be used along with the paragraph_style_suggestion_state to see which fields have changed and their new values.

	ParagraphStyleSuggestionState ParagraphStyleSuggestionState `json:"paragraphStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base ParagraphStyle have been changed in this suggestion.

}

// A suggested change to PositionedObjectProperties.
type SuggestedPositionedObjectProperties struct {
	PositionedObjectProperties PositionedObjectProperties `json:"positionedObjectProperties,omitempty\"` // A PositionedObjectProperties that only includes the changes made in this suggestion. This can be used along with the positioned_object_properties_suggestion_state to see which fields have changed and their new values.

	PositionedObjectPropertiesSuggestionState PositionedObjectPropertiesSuggestionState `json:"positionedObjectPropertiesSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base PositionedObjectProperties have been changed in this suggestion.

}

// A suggested change to a TableCellStyle.
type SuggestedTableCellStyle struct {
	TableCellStyle TableCellStyle `json:"tableCellStyle,omitempty\"` // A TableCellStyle that only includes the changes made in this suggestion. This can be used along with the table_cell_style_suggestion_state to see which fields have changed and their new values.

	TableCellStyleSuggestionState TableCellStyleSuggestionState `json:"tableCellStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base TableCellStyle have been changed in this suggestion.

}

// A suggested change to a TableRowStyle.
type SuggestedTableRowStyle struct {
	TableRowStyle TableRowStyle `json:"tableRowStyle,omitempty\"` // A TableRowStyle that only includes the changes made in this suggestion. This can be used along with the table_row_style_suggestion_state to see which fields have changed and their new values.

	TableRowStyleSuggestionState TableRowStyleSuggestionState `json:"tableRowStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base TableRowStyle have been changed in this suggestion.

}

// A suggested change to a TextStyle.
type SuggestedTextStyle struct {
	TextStyle TextStyle `json:"textStyle,omitempty\"` // A TextStyle that only includes the changes made in this suggestion. This can be used along with the text_style_suggestion_state to see which fields have changed and their new values.

	TextStyleSuggestionState TextStyleSuggestionState `json:"textStyleSuggestionState,omitempty\"` // A mask that indicates which of the fields on the base TextStyle have been changed in this suggestion.

}

// A tab in a document.
type Tab struct {
	ChildTabs []Tab `json:"childTabs,omitempty\"` // The child tabs nested within this tab.

	DocumentTab DocumentTab `json:"documentTab,omitempty\"` // A tab with document contents, like text and images.

	TabProperties TabProperties `json:"tabProperties,omitempty\"` // The properties of the tab, like ID and title.

}

// Properties of a tab.
type TabProperties struct {
	IconEmoji string `json:"iconEmoji,omitempty\"` // Optional. The emoji icon displayed with the tab. A valid emoji icon is represented by a non-empty Unicode string. Any set of characters that don't represent a single emoji is invalid. If an emoji is invalid, a 400 bad request error is returned. If this value is unset or empty, the tab will display the default tab icon.

	Index int `json:"index,omitempty\"` // The zero-based index of the tab within the parent.

	NestingLevel int `json:"nestingLevel,omitempty\"` // Output only. The depth of the tab within the document. Root-level tabs start at 0.

	ParentTabId string `json:"parentTabId,omitempty\"` // Optional. The ID of the parent tab. Empty when the current tab is a root-level tab, which means it doesn't have any parents.

	TabId string `json:"tabId,omitempty\"` // The immutable ID of the tab.

	Title string `json:"title,omitempty\"` // The user-visible name of the tab.

}

// A tab stop within a paragraph.
type TabStop struct {
	Alignment string `json:"alignment,omitempty\"` // The alignment of this tab stop. If unset, the value defaults to START.

	Offset Dimension `json:"offset,omitempty\"` // The offset between this tab stop and the start margin.

}

// A StructuralElement representing a table.
type Table struct {
	Columns int `json:"columns,omitempty\"` // Number of columns in the table. It's possible for a table to be non-rectangular, so some rows may have a different number of cells.

	Rows int `json:"rows,omitempty\"` // Number of rows in the table.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A Table may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	TableRows []TableRow `json:"tableRows,omitempty\"` // The contents and style of each row.

	TableStyle TableStyle `json:"tableStyle,omitempty\"` // The style of the table.

}

// The contents and style of a cell in a Table.
type TableCell struct {
	Content []StructuralElement `json:"content,omitempty\"` // The content of the cell.

	EndIndex int `json:"endIndex,omitempty\"` // The zero-based end index of this cell, exclusive, in UTF-16 code units.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this cell, in UTF-16 code units.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A TableCell may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTableCellStyleChanges map[string]interface{} `json:"suggestedTableCellStyleChanges,omitempty\"` // The suggested changes to the table cell style, keyed by suggestion ID.

	TableCellStyle TableCellStyle `json:"tableCellStyle,omitempty\"` // The style of the cell.

}

// A border around a table cell. Table cell borders cannot be transparent. To hide a table cell border, make its width 0.
type TableCellBorder struct {
	Color OptionalColor `json:"color,omitempty\"` // The color of the border. This color cannot be transparent.

	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the border.

	Width Dimension `json:"width,omitempty\"` // The width of the border.

}

// Location of a single cell within a table.
type TableCellLocation struct {
	ColumnIndex int `json:"columnIndex,omitempty\"` // The zero-based column index. For example, the second column in the table has a column index of 1.

	RowIndex int `json:"rowIndex,omitempty\"` // The zero-based row index. For example, the second row in the table has a row index of 1.

	TableStartLocation Location `json:"tableStartLocation,omitempty\"` // The location where the table starts in the document.

}

// The style of a TableCell. Inherited table cell styles are represented as unset fields in this message. A table cell style can inherit from the table's style.
type TableCellStyle struct {
	BackgroundColor OptionalColor `json:"backgroundColor,omitempty\"` // The background color of the cell.

	BorderBottom TableCellBorder `json:"borderBottom,omitempty\"` // The bottom border of the cell.

	BorderLeft TableCellBorder `json:"borderLeft,omitempty\"` // The left border of the cell.

	BorderRight TableCellBorder `json:"borderRight,omitempty\"` // The right border of the cell.

	BorderTop TableCellBorder `json:"borderTop,omitempty\"` // The top border of the cell.

	ColumnSpan int `json:"columnSpan,omitempty\"` // The column span of the cell. This property is read-only.

	ContentAlignment string `json:"contentAlignment,omitempty\"` // The alignment of the content in the table cell. The default alignment matches the alignment for newly created table cells in the Docs editor.

	PaddingBottom Dimension `json:"paddingBottom,omitempty\"` // The bottom padding of the cell.

	PaddingLeft Dimension `json:"paddingLeft,omitempty\"` // The left padding of the cell.

	PaddingRight Dimension `json:"paddingRight,omitempty\"` // The right padding of the cell.

	PaddingTop Dimension `json:"paddingTop,omitempty\"` // The top padding of the cell.

	RowSpan int `json:"rowSpan,omitempty\"` // The row span of the cell. This property is read-only.

}

// A mask that indicates which of the fields on the base TableCellStyle have been changed in this suggestion. For any field set to true, there's a new suggested value.
type TableCellStyleSuggestionState struct {
	BackgroundColorSuggested bool `json:"backgroundColorSuggested,omitempty\"` // Indicates if there was a suggested change to background_color.

	BorderBottomSuggested bool `json:"borderBottomSuggested,omitempty\"` // Indicates if there was a suggested change to border_bottom.

	BorderLeftSuggested bool `json:"borderLeftSuggested,omitempty\"` // Indicates if there was a suggested change to border_left.

	BorderRightSuggested bool `json:"borderRightSuggested,omitempty\"` // Indicates if there was a suggested change to border_right.

	BorderTopSuggested bool `json:"borderTopSuggested,omitempty\"` // Indicates if there was a suggested change to border_top.

	ColumnSpanSuggested bool `json:"columnSpanSuggested,omitempty\"` // Indicates if there was a suggested change to column_span.

	ContentAlignmentSuggested bool `json:"contentAlignmentSuggested,omitempty\"` // Indicates if there was a suggested change to content_alignment.

	PaddingBottomSuggested bool `json:"paddingBottomSuggested,omitempty\"` // Indicates if there was a suggested change to padding_bottom.

	PaddingLeftSuggested bool `json:"paddingLeftSuggested,omitempty\"` // Indicates if there was a suggested change to padding_left.

	PaddingRightSuggested bool `json:"paddingRightSuggested,omitempty\"` // Indicates if there was a suggested change to padding_right.

	PaddingTopSuggested bool `json:"paddingTopSuggested,omitempty\"` // Indicates if there was a suggested change to padding_top.

	RowSpanSuggested bool `json:"rowSpanSuggested,omitempty\"` // Indicates if there was a suggested change to row_span.

}

// The properties of a column in a table.
type TableColumnProperties struct {
	Width Dimension `json:"width,omitempty\"` // The width of the column. Set when the column's `width_type` is FIXED_WIDTH.

	WidthType string `json:"widthType,omitempty\"` // The width type of the column.

}

// A StructuralElement representing a table of contents.
type TableOfContents struct {
	Content []StructuralElement `json:"content,omitempty\"` // The content of the table of contents.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A TableOfContents may have multiple insertion IDs if it is a nested suggested change. If empty, then this is not a suggested insertion.

}

// A table range represents a reference to a subset of a table. It's important to note that the cells specified by a table range do not necessarily form a rectangle. For example, let's say we have a 3 x 3 table where all the cells of the last row are merged together. The table looks like this: [ ] A table range with table cell location = (table_start_location, row = 0, column = 0), row span = 3 and column span = 2 specifies the following cells: x x [ x x x ]
type TableRange struct {
	ColumnSpan int `json:"columnSpan,omitempty\"` // The column span of the table range.

	RowSpan int `json:"rowSpan,omitempty\"` // The row span of the table range.

	TableCellLocation TableCellLocation `json:"tableCellLocation,omitempty\"` // The cell location where the table range starts.

}

// The contents and style of a row in a Table.
type TableRow struct {
	EndIndex int `json:"endIndex,omitempty\"` // The zero-based end index of this row, exclusive, in UTF-16 code units.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this row, in UTF-16 code units.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A TableRow may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTableRowStyleChanges map[string]interface{} `json:"suggestedTableRowStyleChanges,omitempty\"` // The suggested style changes to this row, keyed by suggestion ID.

	TableCells []TableCell `json:"tableCells,omitempty\"` // The contents and style of each cell in this row. It's possible for a table to be non-rectangular, so some rows may have a different number of cells than other rows in the same table.

	TableRowStyle TableRowStyle `json:"tableRowStyle,omitempty\"` // The style of the table row.

}

// Styles that apply to a table row.
type TableRowStyle struct {
	MinRowHeight Dimension `json:"minRowHeight,omitempty\"` // The minimum height of the row. The row will be rendered in the Docs editor at a height equal to or greater than this value in order to show all the content in the row's cells.

	PreventOverflow bool `json:"preventOverflow,omitempty\"` // Whether the row cannot overflow across page or column boundaries.

	TableHeader bool `json:"tableHeader,omitempty\"` // Whether the row is a table header.

}

// A mask that indicates which of the fields on the base TableRowStyle have been changed in this suggestion. For any field set to true, there's a new suggested value.
type TableRowStyleSuggestionState struct {
	MinRowHeightSuggested bool `json:"minRowHeightSuggested,omitempty\"` // Indicates if there was a suggested change to min_row_height.

}

// Styles that apply to a table.
type TableStyle struct {
	TableColumnProperties []TableColumnProperties `json:"tableColumnProperties,omitempty\"` // The properties of each column. Note that in Docs, tables contain rows and rows contain cells, similar to HTML. So the properties for a row can be found on the row's table_row_style.

}

// A criteria that specifies in which tabs a request executes.
type TabsCriteria struct {
	TabIds []string `json:"tabIds,omitempty\"` // The list of tab IDs in which the request executes.

}

// A ParagraphElement that represents a run of text that all has the same styling.
type TextRun struct {
	Content string `json:"content,omitempty\"` // The text of this run. Any non-text elements in the run are replaced with the Unicode character U+E907.

	SuggestedDeletionIds []string `json:"suggestedDeletionIds,omitempty\"` // The suggested deletion IDs. If empty, then there are no suggested deletions of this content.

	SuggestedInsertionIds []string `json:"suggestedInsertionIds,omitempty\"` // The suggested insertion IDs. A TextRun may have multiple insertion IDs if it's a nested suggested change. If empty, then this is not a suggested insertion.

	SuggestedTextStyleChanges map[string]interface{} `json:"suggestedTextStyleChanges,omitempty\"` // The suggested text style changes to this run, keyed by suggestion ID.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The text style of this run.

}

// Represents the styling that can be applied to text. Inherited text styles are represented as unset fields in this message. A text style's parent depends on where the text style is defined: * The TextStyle of text in a Paragraph inherits from the paragraph's corresponding named style type. * The TextStyle on a named style inherits from the normal text named style. * The TextStyle of the normal text named style inherits from the default text style in the Docs editor. * The TextStyle on a Paragraph element that's contained in a table may inherit its text style from the table style. If the text style does not inherit from a parent, unsetting fields will revert the style to a value matching the defaults in the Docs editor.
type TextStyle struct {
	BackgroundColor OptionalColor `json:"backgroundColor,omitempty\"` // The background color of the text. If set, the color is either an RGB color or transparent, depending on the `color` field.

	BaselineOffset string `json:"baselineOffset,omitempty\"` // The text's vertical offset from its normal position. Text with `SUPERSCRIPT` or `SUBSCRIPT` baseline offsets is automatically rendered in a smaller font size, computed based on the `font_size` field. Changes in this field don't affect the `font_size`.

	Bold bool `json:"bold,omitempty\"` // Whether or not the text is rendered as bold.

	FontSize Dimension `json:"fontSize,omitempty\"` // The size of the text's font.

	ForegroundColor OptionalColor `json:"foregroundColor,omitempty\"` // The foreground color of the text. If set, the color is either an RGB color or transparent, depending on the `color` field.

	Italic bool `json:"italic,omitempty\"` // Whether or not the text is italicized.

	Link Link `json:"link,omitempty\"` // The hyperlink destination of the text. If unset, there's no link. Links are not inherited from parent text. Changing the link in an update request causes some other changes to the text style of the range: * When setting a link, the text foreground color will be updated to the default link color and the text will be underlined. If these fields are modified in the same request, those values will be used instead of the link defaults. * Setting a link on a text range that overlaps with an existing link will also update the existing link to point to the new URL. * Links are not settable on newline characters. As a result, setting a link on a text range that crosses a paragraph boundary, such as `"ABC\n123"`, will separate the newline character(s) into their own text runs. The link will be applied separately to the runs before and after the newline. * Removing a link will update the text style of the range to match the style of the preceding text (or the default text styles if the preceding text is another link) unless different styles are being set in the same request.

	SmallCaps bool `json:"smallCaps,omitempty\"` // Whether or not the text is in small capital letters.

	Strikethrough bool `json:"strikethrough,omitempty\"` // Whether or not the text is struck through.

	Underline bool `json:"underline,omitempty\"` // Whether or not the text is underlined.

	WeightedFontFamily WeightedFontFamily `json:"weightedFontFamily,omitempty\"` // The font family and rendered weight of the text. If an update request specifies values for both `weighted_font_family` and `bold`, the `weighted_font_family` is applied first, then `bold`. If `weighted_font_family#weight` is not set, it defaults to `400`. If `weighted_font_family` is set, then `weighted_font_family#font_family` must also be set with a non-empty value. Otherwise, a 400 bad request error is returned.

}

// A mask that indicates which of the fields on the base TextStyle have been changed in this suggestion. For any field set to true, there's a new suggested value.
type TextStyleSuggestionState struct {
	BackgroundColorSuggested bool `json:"backgroundColorSuggested,omitempty\"` // Indicates if there was a suggested change to background_color.

	BaselineOffsetSuggested bool `json:"baselineOffsetSuggested,omitempty\"` // Indicates if there was a suggested change to baseline_offset.

	BoldSuggested bool `json:"boldSuggested,omitempty\"` // Indicates if there was a suggested change to bold.

	FontSizeSuggested bool `json:"fontSizeSuggested,omitempty\"` // Indicates if there was a suggested change to font_size.

	ForegroundColorSuggested bool `json:"foregroundColorSuggested,omitempty\"` // Indicates if there was a suggested change to foreground_color.

	ItalicSuggested bool `json:"italicSuggested,omitempty\"` // Indicates if there was a suggested change to italic.

	LinkSuggested bool `json:"linkSuggested,omitempty\"` // Indicates if there was a suggested change to link.

	SmallCapsSuggested bool `json:"smallCapsSuggested,omitempty\"` // Indicates if there was a suggested change to small_caps.

	StrikethroughSuggested bool `json:"strikethroughSuggested,omitempty\"` // Indicates if there was a suggested change to strikethrough.

	UnderlineSuggested bool `json:"underlineSuggested,omitempty\"` // Indicates if there was a suggested change to underline.

	WeightedFontFamilySuggested bool `json:"weightedFontFamilySuggested,omitempty\"` // Indicates if there was a suggested change to weighted_font_family.

}

// Unmerges cells in a Table.
type UnmergeTableCellsRequest struct {
	TableRange TableRange `json:"tableRange,omitempty\"` // The table range specifying which cells of the table to unmerge. All merged cells in this range will be unmerged, and cells that are already unmerged will not be affected. If the range has no merged cells, the request will do nothing. If there is text in any of the merged cells, the text will remain in the "head" cell of the resulting block of unmerged cells. The "head" cell is the upper-left cell when the content direction is from left to right, and the upper-right otherwise.

}

// Updates the DocumentStyle.
type UpdateDocumentStyleRequest struct {
	DocumentStyle DocumentStyle `json:"documentStyle,omitempty\"` // The styles to set on the document. Certain document style changes may cause other changes in order to mirror the behavior of the Docs editor. See the documentation of DocumentStyle for more information.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `document_style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the background, set `fields` to `"background"`.

	TabId string `json:"tabId,omitempty\"` // The tab that contains the style to update. When omitted, the request applies to the first tab. In a document containing a single tab: - If provided, must match the singular tab's ID. - If omitted, the request applies to the singular tab. In a document containing multiple tabs: - If provided, the request applies to the specified tab. - If not provided, the request applies to the first tab in the document.

}

// Update the properties of a document tab.
type UpdateDocumentTabPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tab_properties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field.

	TabProperties TabProperties `json:"tabProperties,omitempty\"` // The tab properties to update.

}

// Updates a named style.
type UpdateNamedStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The NamedStyle fields that should be updated. At least `named_style_type` must be specified. The root `named_style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example, to update the text style to bold, set `fields` to include `"text_style"` and `"text_style.bold"`. To update the paragraph style's alignment property, set `fields` to include `"paragraph_style"` and `"paragraph_style.alignment"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset. Specifying `"text_style"` or `"paragraph_style"` with an empty TextStyle or ParagraphStyle will reset all of its nested fields.

	NamedStyle NamedStyle `json:"namedStyle,omitempty\"` // The document style to update.

	TabId string `json:"tabId,omitempty\"` // The document tab to update. By default, the update is applied to the first tab.

}

// Update the styling of all paragraphs that overlap with the given range.
type UpdateParagraphStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `paragraph_style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example, to update the paragraph style's alignment property, set `fields` to `"alignment"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ParagraphStyle ParagraphStyle `json:"paragraphStyle,omitempty\"` // The styles to set on the paragraphs. Certain paragraph style changes may cause other changes in order to mirror the behavior of the Docs editor. See the documentation of ParagraphStyle for more information.

	RangeValue RangeValue `json:"range,omitempty\"` // The range overlapping the paragraphs to style.

}

// Updates the SectionStyle.
type UpdateSectionStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `section_style` is implied and must not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the left margin, set `fields` to `"margin_left"`.

	RangeValue RangeValue `json:"range,omitempty\"` // The range overlapping the sections to style. Because section breaks can only be inserted inside the body, the segment ID field must be empty.

	SectionStyle SectionStyle `json:"sectionStyle,omitempty\"` // The styles to be set on the section. Certain section style changes may cause other changes in order to mirror the behavior of the Docs editor. See the documentation of SectionStyle for more information.

}

// Updates the style of a range of table cells.
type UpdateTableCellStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableCellStyle` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the table cell background color, set `fields` to `"backgroundColor"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	TableCellStyle TableCellStyle `json:"tableCellStyle,omitempty\"` // The style to set on the table cells. When updating borders, if a cell shares a border with an adjacent cell, the corresponding border property of the adjacent cell is updated as well. Borders that are merged and invisible are not updated. Since updating a border shared by adjacent cells in the same request can cause conflicting border updates, border updates are applied in the following order: - `border_right` - `border_left` - `border_bottom` - `border_top`

	TableRange TableRange `json:"tableRange,omitempty\"` // The table range representing the subset of the table to which the updates are applied.

	TableStartLocation Location `json:"tableStartLocation,omitempty\"` // The location where the table starts in the document. When specified, the updates are applied to all the cells in the table.

}

// Updates the TableColumnProperties of columns in a table.
type UpdateTableColumnPropertiesRequest struct {
	ColumnIndices []int `json:"columnIndices,omitempty\"` // The list of zero-based column indices whose property should be updated. If no indices are specified, all columns will be updated.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableColumnProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the column width, set `fields` to `"width"`.

	TableColumnProperties TableColumnProperties `json:"tableColumnProperties,omitempty\"` // The table column properties to update. If the value of `table_column_properties#width` is less than 5 points (5/72 inch), a 400 bad request error is returned.

	TableStartLocation Location `json:"tableStartLocation,omitempty\"` // The location where the table starts in the document.

}

// Updates the TableRowStyle of rows in a table.
type UpdateTableRowStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableRowStyle` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the minimum row height, set `fields` to `"min_row_height"`.

	RowIndices []int `json:"rowIndices,omitempty\"` // The list of zero-based row indices whose style should be updated. If no indices are specified, all rows will be updated.

	TableRowStyle TableRowStyle `json:"tableRowStyle,omitempty\"` // The styles to be set on the rows.

	TableStartLocation Location `json:"tableStartLocation,omitempty\"` // The location where the table starts in the document.

}

// Update the styling of text.
type UpdateTextStyleRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `text_style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example, to update the text style to bold, set `fields` to `"bold"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	RangeValue RangeValue `json:"range,omitempty\"` // The range of text to style. The range may be extended to include adjacent newlines. If the range fully contains a paragraph belonging to a list, the paragraph's bullet is also updated with the matching text style. Ranges cannot be inserted inside a relative UpdateTextStyleRequest.

	TextStyle TextStyle `json:"textStyle,omitempty\"` // The styles to set on the text. If the value for a particular style matches that of the parent, that style will be set to inherit. Certain text style changes may cause other changes in order to to mirror the behavior of the Docs editor. See the documentation of TextStyle for more information.

}

// Represents a font family and weight of text.
type WeightedFontFamily struct {
	FontFamily string `json:"fontFamily,omitempty\"` // The font family of the text. The font family can be any font from the Font menu in Docs or from [Google Fonts] (https://fonts.google.com/). If the font name is unrecognized, the text is rendered in `Arial`.

	Weight int `json:"weight,omitempty\"` // The weight of the font. This field can have any value that's a multiple of `100` between `100` and `900`, inclusive. This range corresponds to the numerical values described in the CSS 2.1 Specification, [section 15.6](https://www.w3.org/TR/CSS21/fonts.html#font-boldness), with non-numerical values disallowed. The default value is `400` ("normal"). The font weight makes up just one component of the rendered font weight. A combination of the `weight` and the text style's resolved `bold` value determine the rendered weight, after accounting for inheritance: * If the text is bold and the weight is less than `400`, the rendered weight is 400. * If the text is bold and the weight is greater than or equal to `400` but is less than `700`, the rendered weight is `700`. * If the weight is greater than or equal to `700`, the rendered weight is equal to the weight. * If the text is not bold, the rendered weight is equal to the weight.

}

// Provides control over how write requests are executed.
type WriteControl struct {
	RequiredRevisionId string `json:"requiredRevisionId,omitempty\"` // The optional revision ID of the document the write request is applied to. If this is not the latest revision of the document, the request is not processed and returns a 400 bad request error. When a required revision ID is returned in a response, it indicates the revision ID of the document after the request was applied.

	TargetRevisionId string `json:"targetRevisionId,omitempty\"` // The optional target revision ID of the document the write request is applied to. If collaborator changes have occurred after the document was read using the API, the changes produced by this write request are applied against the collaborator changes. This results in a new revision of the document that incorporates both the collaborator changes and the changes in the request, with the Docs server resolving conflicting changes. When using target revision ID, the API client can be thought of as another collaborator of the document. The target revision ID can only be used to write to recent versions of a document. If the target revision is too far behind the latest revision, the request is not processed and returns a 400 bad request error. The request should be tried again after retrieving the latest version of the document. Usually a revision ID remains valid for use as a target revision for several minutes after it's read, but for frequently edited documents this window might be shorter.

}
