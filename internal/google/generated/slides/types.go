// Google Slides API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package slides

// AffineTransform uses a 3x3 matrix with an implied last row of [ 0 0 1 ] to transform source coordinates (x,y) into destination coordinates (x', y') according to: x' x = shear_y scale_y translate_y 1 [ 1 ] After transformation, x' = scale_x * x + shear_x * y + translate_x; y' = scale_y * y + shear_y * x + translate_y; This message is therefore composed of these six matrix elements.
type AffineTransform struct {
	ScaleX float64 `json:"scaleX,omitempty\"` // The X coordinate scaling element.

	ScaleY float64 `json:"scaleY,omitempty\"` // The Y coordinate scaling element.

	ShearX float64 `json:"shearX,omitempty\"` // The X coordinate shearing element.

	ShearY float64 `json:"shearY,omitempty\"` // The Y coordinate shearing element.

	TranslateX float64 `json:"translateX,omitempty\"` // The X coordinate translation element.

	TranslateY float64 `json:"translateY,omitempty\"` // The Y coordinate translation element.

	Unit string `json:"unit,omitempty\"` // The units for translate elements.

}

// A TextElement kind that represents auto text.
type AutoText struct {
	Content string `json:"content,omitempty\"` // The rendered content of this auto text, if available.

	Style TextStyle `json:"style,omitempty\"` // The styling applied to this auto text.

	TypeValue string `json:"type,omitempty\"` // The type of this auto text.

}

// The autofit properties of a Shape. This property is only set for shapes that allow text.
type Autofit struct {
	AutofitType string `json:"autofitType,omitempty\"` // The autofit type of the shape. If the autofit type is AUTOFIT_TYPE_UNSPECIFIED, the autofit type is inherited from a parent placeholder if it exists. The field is automatically set to NONE if a request is made that might affect text fitting within its bounding text box. In this case, the font_scale is applied to the font_size and the line_spacing_reduction is applied to the line_spacing. Both properties are also reset to default values.

	FontScale float64 `json:"fontScale,omitempty\"` // The font scale applied to the shape. For shapes with autofit_type NONE or SHAPE_AUTOFIT, this value is the default value of 1. For TEXT_AUTOFIT, this value multiplied by the font_size gives the font size that's rendered in the editor. This property is read-only.

	LineSpacingReduction float64 `json:"lineSpacingReduction,omitempty\"` // The line spacing reduction applied to the shape. For shapes with autofit_type NONE or SHAPE_AUTOFIT, this value is the default value of 0. For TEXT_AUTOFIT, this value subtracted from the line_spacing gives the line spacing that's rendered in the editor. This property is read-only.

}

// Request message for PresentationsService.BatchUpdatePresentation.
type BatchUpdatePresentationRequest struct {
	Requests []Request `json:"requests,omitempty\"` // A list of updates to apply to the presentation.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed.

}

// Response message from a batch update.
type BatchUpdatePresentationResponse struct {
	PresentationId string `json:"presentationId,omitempty\"` // The presentation the updates were applied to.

	Replies []Response `json:"replies,omitempty\"` // The reply of the updates. This maps 1:1 with the updates, although replies to some requests may be empty.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // The updated write control after applying the request.

}

// Describes the bullet of a paragraph.
type Bullet struct {
	BulletStyle TextStyle `json:"bulletStyle,omitempty\"` // The paragraph specific text style applied to this bullet.

	Glyph string `json:"glyph,omitempty\"` // The rendered bullet glyph for this paragraph.

	ListId string `json:"listId,omitempty\"` // The ID of the list this paragraph belongs to.

	NestingLevel int `json:"nestingLevel,omitempty\"` // The nesting level of this paragraph in the list.

}

// The palette of predefined colors for a page.
type ColorScheme struct {
	Colors []ThemeColorPair `json:"colors,omitempty\"` // The ThemeColorType and corresponding concrete color pairs.

}

// A color and position in a gradient band.
type ColorStop struct {
	Alpha float64 `json:"alpha,omitempty\"` // The alpha value of this color in the gradient band. Defaults to 1.0, fully opaque.

	Color OpaqueColor `json:"color,omitempty\"` // The color of the gradient stop.

	Position float64 `json:"position,omitempty\"` // The relative position of the color stop in the gradient band measured in percentage. The value should be in the interval [0.0, 1.0].

}

// Creates an image.
type CreateImageRequest struct {
	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the image. When the aspect ratio of the provided size does not match the image aspect ratio, the image is scaled and centered with respect to the size in order to maintain the aspect ratio. The provided transform is applied after this operation. The PageElementProperties.size property is optional. If you don't specify the size, the default size of the image is used. The PageElementProperties.transform property is optional. If you don't specify a transform, the image will be placed at the top-left corner of the page.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

	Url string `json:"url,omitempty\"` // The image URL. The image is fetched once at insertion time and a copy is stored for display inside the presentation. Images must be less than 50 MB in size, can't exceed 25 megapixels, and must be in one of PNG, JPEG, or GIF formats. The provided URL must be publicly accessible and up to 2 KB in length. The URL is saved with the image, and exposed through the Image.source_url field.

}

// The result of creating an image.
type CreateImageResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created image.

}

// Creates a line.
type CreateLineRequest struct {
	Category string `json:"category,omitempty\"` // The category of the line to be created. The exact line type created is determined based on the category and how it's routed to connect to other page elements. If you specify both a `category` and a `line_category`, the `category` takes precedence. If you do not specify a value for `category`, but specify a value for `line_category`, then the specified `line_category` value is used. If you do not specify either, then STRAIGHT is used.

	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the line.

	LineCategory string `json:"lineCategory,omitempty\"` // The category of the line to be created. *Deprecated*: use `category` instead. The exact line type created is determined based on the category and how it's routed to connect to other page elements. If you specify both a `category` and a `line_category`, the `category` takes precedence.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

}

// The result of creating a line.
type CreateLineResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created line.

}

// Creates bullets for all of the paragraphs that overlap with the given text index range. The nesting level of each paragraph will be determined by counting leading tabs in front of each paragraph. To avoid excess space between the bullet and the corresponding paragraph, these leading tabs are removed by this request. This may change the indices of parts of the text. If the paragraph immediately before paragraphs being updated is in a list with a matching preset, the paragraphs being updated are added to that preceding list.
type CreateParagraphBulletsRequest struct {
	BulletPreset string `json:"bulletPreset,omitempty\"` // The kinds of bullet glyphs to be used. Defaults to the `BULLET_DISC_CIRCLE_SQUARE` preset.

	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The optional table cell location if the text to be modified is in a table cell. If present, the object_id must refer to a table.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table containing the text to add bullets to.

	TextRange RangeValue `json:"textRange,omitempty\"` // The range of text to apply the bullet presets to, based on TextElement indexes.

}

// Creates a new shape.
type CreateShapeRequest struct {
	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the shape.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If empty, a unique identifier will be generated.

	ShapeType string `json:"shapeType,omitempty\"` // The shape type.

}

// The result of creating a shape.
type CreateShapeResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created shape.

}

// Creates an embedded Google Sheets chart. NOTE: Chart creation requires at least one of the spreadsheets.readonly, spreadsheets, drive.readonly, drive.file, or drive OAuth scopes.
type CreateSheetsChartRequest struct {
	ChartId int `json:"chartId,omitempty\"` // The ID of the specific chart in the Google Sheets spreadsheet.

	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the chart. When the aspect ratio of the provided size does not match the chart aspect ratio, the chart is scaled and centered with respect to the size in order to maintain aspect ratio. The provided transform is applied after this operation.

	LinkingMode string `json:"linkingMode,omitempty\"` // The mode with which the chart is linked to the source spreadsheet. When not specified, the chart will be an image that is not linked.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If specified, the ID must be unique among all pages and page elements in the presentation. The ID should start with a word character [a-zA-Z0-9_] and then followed by any number of the following characters [a-zA-Z0-9_-:]. The length of the ID should not be less than 5 or greater than 50. If empty, a unique identifier will be generated.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the Google Sheets spreadsheet that contains the chart. You might need to add a resource key to the HTTP header for a subset of old files. For more information, see [Access link-shared files using resource keys](https://developers.google.com/drive/api/v3/resource-keys).

}

// The result of creating an embedded Google Sheets chart.
type CreateSheetsChartResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created chart.

}

// Creates a slide.
type CreateSlideRequest struct {
	InsertionIndex int `json:"insertionIndex,omitempty\"` // The optional zero-based index indicating where to insert the slides. If you don't specify an index, the slide is created at the end.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The ID length must be between 5 and 50 characters, inclusive. If you don't specify an ID, a unique one is generated.

	PlaceholderIdMappings []LayoutPlaceholderIdMapping `json:"placeholderIdMappings,omitempty\"` // An optional list of object ID mappings from the placeholder(s) on the layout to the placeholders that are created on the slide from the specified layout. Can only be used when `slide_layout_reference` is specified.

	SlideLayoutReference LayoutReference `json:"slideLayoutReference,omitempty\"` // Layout reference of the slide to be inserted, based on the *current master*, which is one of the following: - The master of the previous slide index. - The master of the first slide, if the insertion_index is zero. - The first master in the presentation, if there are no slides. If the LayoutReference is not found in the current master, a 400 bad request error is returned. If you don't specify a layout reference, the slide uses the predefined `BLANK` layout.

}

// The result of creating a slide.
type CreateSlideResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created slide.

}

// Creates a new table.
type CreateTableRequest struct {
	Columns int `json:"columns,omitempty\"` // Number of columns in the table.

	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the table. The table will be created at the provided size, subject to a minimum size. If no size is provided, the table will be automatically sized. Table transforms must have a scale of 1 and no shear components. If no transform is provided, the table will be centered on the page.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

	Rows int `json:"rows,omitempty\"` // Number of rows in the table.

}

// The result of creating a table.
type CreateTableResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created table.

}

// Creates a video. NOTE: Creating a video from Google Drive requires that the requesting app have at least one of the drive, drive.readonly, or drive.file OAuth scopes.
type CreateVideoRequest struct {
	ElementProperties PageElementProperties `json:"elementProperties,omitempty\"` // The element properties for the video. The PageElementProperties.size property is optional. If you don't specify a size, a default size is chosen by the server. The PageElementProperties.transform property is optional. The transform must not have shear components. If you don't specify a transform, the video will be placed at the top left corner of the page.

	Id string `json:"id,omitempty\"` // The video source's unique identifier for this video. e.g. For YouTube video https://www.youtube.com/watch?v=7U3axjORYZ0, the ID is 7U3axjORYZ0. For a Google Drive video https://drive.google.com/file/d/1xCgQLFTJi5_Xl8DgW_lcUYq5e-q6Hi5Q the ID is 1xCgQLFTJi5_Xl8DgW_lcUYq5e-q6Hi5Q. To access a Google Drive video file, you might need to add a resource key to the HTTP header for a subset of old files. For more information, see [Access link-shared files using resource keys](https://developers.google.com/drive/api/v3/resource-keys).

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

	Source string `json:"source,omitempty\"` // The video source.

}

// The result of creating a video.
type CreateVideoResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created video.

}

// The crop properties of an object enclosed in a container. For example, an Image. The crop properties is represented by the offsets of four edges which define a crop rectangle. The offsets are measured in percentage from the corresponding edges of the object's original bounding rectangle towards inside, relative to the object's original dimensions. - If the offset is in the interval (0, 1), the corresponding edge of crop rectangle is positioned inside of the object's original bounding rectangle. - If the offset is negative or greater than 1, the corresponding edge of crop rectangle is positioned outside of the object's original bounding rectangle. - If the left edge of the crop rectangle is on the right side of its right edge, the object will be flipped horizontally. - If the top edge of the crop rectangle is below its bottom edge, the object will be flipped vertically. - If all offsets and rotation angle is 0, the object is not cropped. After cropping, the content in the crop rectangle will be stretched to fit its container.
type CropProperties struct {
	Angle float64 `json:"angle,omitempty\"` // The rotation angle of the crop window around its center, in radians. Rotation angle is applied after the offset.

	BottomOffset float64 `json:"bottomOffset,omitempty\"` // The offset specifies the bottom edge of the crop rectangle that is located above the original bounding rectangle bottom edge, relative to the object's original height.

	LeftOffset float64 `json:"leftOffset,omitempty\"` // The offset specifies the left edge of the crop rectangle that is located to the right of the original bounding rectangle left edge, relative to the object's original width.

	RightOffset float64 `json:"rightOffset,omitempty\"` // The offset specifies the right edge of the crop rectangle that is located to the left of the original bounding rectangle right edge, relative to the object's original width.

	TopOffset float64 `json:"topOffset,omitempty\"` // The offset specifies the top edge of the crop rectangle that is located below the original bounding rectangle top edge, relative to the object's original height.

}

// Deletes an object, either pages or page elements, from the presentation.
type DeleteObjectRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the page or page element to delete. If after a delete operation a group contains only 1 or no page elements, the group is also deleted. If a placeholder is deleted on a layout, any empty inheriting placeholders are also deleted.

}

// Deletes bullets from all of the paragraphs that overlap with the given text index range. The nesting level of each paragraph will be visually preserved by adding indent to the start of the corresponding paragraph.
type DeleteParagraphBulletsRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The optional table cell location if the text to be modified is in a table cell. If present, the object_id must refer to a table.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table containing the text to delete bullets from.

	TextRange RangeValue `json:"textRange,omitempty\"` // The range of text to delete bullets from, based on TextElement indexes.

}

// Deletes a column from a table.
type DeleteTableColumnRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The reference table cell location from which a column will be deleted. The column this cell spans will be deleted. If this is a merged cell, multiple columns will be deleted. If no columns remain in the table after this deletion, the whole table is deleted.

	TableObjectId string `json:"tableObjectId,omitempty\"` // The table to delete columns from.

}

// Deletes a row from a table.
type DeleteTableRowRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The reference table cell location from which a row will be deleted. The row this cell spans will be deleted. If this is a merged cell, multiple rows will be deleted. If no rows remain in the table after this deletion, the whole table is deleted.

	TableObjectId string `json:"tableObjectId,omitempty\"` // The table to delete rows from.

}

// Deletes text from a shape or a table cell.
type DeleteTextRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The optional table cell location if the text is to be deleted from a table cell. If present, the object_id must refer to a table.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table from which the text will be deleted.

	TextRange RangeValue `json:"textRange,omitempty\"` // The range of text to delete, based on TextElement indexes. There is always an implicit newline character at the end of a shape's or table cell's text that cannot be deleted. `Range.Type.ALL` will use the correct bounds, but care must be taken when specifying explicit bounds for range types `FROM_START_INDEX` and `FIXED_RANGE`. For example, if the text is "ABC", followed by an implicit newline, then the maximum value is 2 for `text_range.start_index` and 3 for `text_range.end_index`. Deleting text that crosses a paragraph boundary may result in changes to paragraph styles and lists as the two paragraphs are merged. Ranges that include only one code unit of a surrogate pair are expanded to include both code units.

}

// A magnitude in a single direction in the specified units.
type Dimension struct {
	Magnitude float64 `json:"magnitude,omitempty\"` // The magnitude.

	Unit string `json:"unit,omitempty\"` // The units for magnitude.

}

// Duplicates a slide or page element. When duplicating a slide, the duplicate slide will be created immediately following the specified slide. When duplicating a page element, the duplicate will be placed on the same page at the same position as the original.
type DuplicateObjectRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The ID of the object to duplicate.

	ObjectIds map[string]interface{} `json:"objectIds,omitempty\"` // The object being duplicated may contain other objects, for example when duplicating a slide or a group page element. This map defines how the IDs of duplicated objects are generated: the keys are the IDs of the original objects and its values are the IDs that will be assigned to the corresponding duplicate object. The ID of the source object's duplicate may be specified in this map as well, using the same value of the `object_id` field as a key and the newly desired ID as the value. All keys must correspond to existing IDs in the presentation. All values must be unique in the presentation and must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the new ID must not be less than 5 or greater than 50. If any IDs of source objects are omitted from the map, a new random ID will be assigned. If the map is empty or unset, all duplicate objects will receive a new random ID.

}

// The response of duplicating an object.
type DuplicateObjectResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The ID of the new duplicate object.

}

// A PageElement kind representing a joined collection of PageElements.
type Group struct {
	Children []PageElement `json:"children,omitempty\"` // The collection of elements in the group. The minimum size of a group is 2.

}

// Groups objects to create an object group. For example, groups PageElements to create a Group on the same page as all the children.
type GroupObjectsRequest struct {
	ChildrenObjectIds []string `json:"childrenObjectIds,omitempty\"` // The object IDs of the objects to group. Only page elements can be grouped. There should be at least two page elements on the same page that are not already in another group. Some page elements, such as videos, tables and placeholders cannot be grouped.

	GroupObjectId string `json:"groupObjectId,omitempty\"` // A user-supplied object ID for the group to be created. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

}

// The result of grouping objects.
type GroupObjectsResponse struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the created group.

}

// A PageElement kind representing an image.
type Image struct {
	ContentUrl string `json:"contentUrl,omitempty\"` // An URL to an image with a default lifetime of 30 minutes. This URL is tagged with the account of the requester. Anyone with the URL effectively accesses the image as the original requester. Access to the image may be lost if the presentation's sharing settings change.

	ImageProperties ImageProperties `json:"imageProperties,omitempty\"` // The properties of the image.

	Placeholder Placeholder `json:"placeholder,omitempty\"` // Placeholders are page elements that inherit from corresponding placeholders on layouts and masters. If set, the image is a placeholder image and any inherited properties can be resolved by looking at the parent placeholder identified by the Placeholder.parent_object_id field.

	SourceUrl string `json:"sourceUrl,omitempty\"` // The source URL is the URL used to insert the image. The source URL can be empty.

}

// The properties of the Image.
type ImageProperties struct {
	Brightness float64 `json:"brightness,omitempty\"` // The brightness effect of the image. The value should be in the interval [-1.0, 1.0], where 0 means no effect. This property is read-only.

	Contrast float64 `json:"contrast,omitempty\"` // The contrast effect of the image. The value should be in the interval [-1.0, 1.0], where 0 means no effect. This property is read-only.

	CropProperties CropProperties `json:"cropProperties,omitempty\"` // The crop properties of the image. If not set, the image is not cropped. This property is read-only.

	Link Link `json:"link,omitempty\"` // The hyperlink destination of the image. If unset, there is no link.

	Outline Outline `json:"outline,omitempty\"` // The outline of the image. If not set, the image has no outline.

	Recolor Recolor `json:"recolor,omitempty\"` // The recolor effect of the image. If not set, the image is not recolored. This property is read-only.

	Shadow Shadow `json:"shadow,omitempty\"` // The shadow of the image. If not set, the image has no shadow. This property is read-only.

	Transparency float64 `json:"transparency,omitempty\"` // The transparency effect of the image. The value should be in the interval [0.0, 1.0], where 0 means no effect and 1 means completely transparent. This property is read-only.

}

// Inserts columns into a table. Other columns in the table will be resized to fit the new column.
type InsertTableColumnsRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The reference table cell location from which columns will be inserted. A new column will be inserted to the left (or right) of the column where the reference cell is. If the reference cell is a merged cell, a new column will be inserted to the left (or right) of the merged cell.

	InsertRight bool `json:"insertRight,omitempty\"` // Whether to insert new columns to the right of the reference cell location. - `True`: insert to the right. - `False`: insert to the left.

	Number int `json:"number,omitempty\"` // The number of columns to be inserted. Maximum 20 per request.

	TableObjectId string `json:"tableObjectId,omitempty\"` // The table to insert columns into.

}

// Inserts rows into a table.
type InsertTableRowsRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The reference table cell location from which rows will be inserted. A new row will be inserted above (or below) the row where the reference cell is. If the reference cell is a merged cell, a new row will be inserted above (or below) the merged cell.

	InsertBelow bool `json:"insertBelow,omitempty\"` // Whether to insert new rows below the reference cell location. - `True`: insert below the cell. - `False`: insert above the cell.

	Number int `json:"number,omitempty\"` // The number of rows to be inserted. Maximum 20 per request.

	TableObjectId string `json:"tableObjectId,omitempty\"` // The table to insert rows into.

}

// Inserts text into a shape or a table cell.
type InsertTextRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The optional table cell location if the text is to be inserted into a table cell. If present, the object_id must refer to a table.

	InsertionIndex int `json:"insertionIndex,omitempty\"` // The index where the text will be inserted, in Unicode code units, based on TextElement indexes. The index is zero-based and is computed from the start of the string. The index may be adjusted to prevent insertions inside Unicode grapheme clusters. In these cases, the text will be inserted immediately after the grapheme cluster.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table where the text will be inserted.

	Text string `json:"text,omitempty\"` // The text to be inserted. Inserting a newline character will implicitly create a new ParagraphMarker at that index. The paragraph style of the new paragraph will be copied from the paragraph at the current insertion index, including lists and bullets. Text styles for inserted text will be determined automatically, generally preserving the styling of neighboring text. In most cases, the text will be added to the TextRun that exists at the insertion index. Some control characters (U+0000-U+0008, U+000C-U+001F) and characters from the Unicode Basic Multilingual Plane Private Use Area (U+E000-U+F8FF) will be stripped out of the inserted text.

}

// The user-specified ID mapping for a placeholder that will be created on a slide from a specified layout.
type LayoutPlaceholderIdMapping struct {
	LayoutPlaceholder Placeholder `json:"layoutPlaceholder,omitempty\"` // The placeholder on a layout that will be applied to a slide. Only type and index are needed. For example, a predefined `TITLE_AND_BODY` layout may usually have a TITLE placeholder with index 0 and a BODY placeholder with index 0.

	LayoutPlaceholderObjectId string `json:"layoutPlaceholderObjectId,omitempty\"` // The object ID of the placeholder on a layout that will be applied to a slide.

	ObjectId string `json:"objectId,omitempty\"` // A user-supplied object ID for the placeholder identified above that to be created onto a slide. If you specify an ID, it must be unique among all pages and page elements in the presentation. The ID must start with an alphanumeric character or an underscore (matches regex `[a-zA-Z0-9_]`); remaining characters may include those as well as a hyphen or colon (matches regex `[a-zA-Z0-9_-:]`). The length of the ID must not be less than 5 or greater than 50. If you don't specify an ID, a unique one is generated.

}

// The properties of Page are only relevant for pages with page_type LAYOUT.
type LayoutProperties struct {
	DisplayName string `json:"displayName,omitempty\"` // The human-readable name of the layout.

	MasterObjectId string `json:"masterObjectId,omitempty\"` // The object ID of the master that this layout is based on.

	Name string `json:"name,omitempty\"` // The name of the layout.

}

// Slide layout reference. This may reference either: - A predefined layout - One of the layouts in the presentation.
type LayoutReference struct {
	LayoutId string `json:"layoutId,omitempty\"` // Layout ID: the object ID of one of the layouts in the presentation.

	PredefinedLayout string `json:"predefinedLayout,omitempty\"` // Predefined layout.

}

// A PageElement kind representing a non-connector line, straight connector, curved connector, or bent connector.
type Line struct {
	LineCategory string `json:"lineCategory,omitempty\"` // The category of the line. It matches the `category` specified in CreateLineRequest, and can be updated with UpdateLineCategoryRequest.

	LineProperties LineProperties `json:"lineProperties,omitempty\"` // The properties of the line.

	LineType string `json:"lineType,omitempty\"` // The type of the line.

}

// The properties for one end of a Line connection.
type LineConnection struct {
	ConnectedObjectId string `json:"connectedObjectId,omitempty\"` // The object ID of the connected page element. Some page elements, such as groups, tables, and lines do not have connection sites and therefore cannot be connected to a connector line.

	ConnectionSiteIndex int `json:"connectionSiteIndex,omitempty\"` // The index of the connection site on the connected page element. In most cases, it corresponds to the predefined connection site index from the ECMA-376 standard. More information on those connection sites can be found in both the description of the "cxn" attribute in section 20.1.9.9 and "Annex H. Example Predefined DrawingML Shape and Text Geometries" of "Office Open XML File Formats - Fundamentals and Markup Language Reference", part 1 of [ECMA-376 5th edition](https://ecma-international.org/publications-and-standards/standards/ecma-376/). The position of each connection site can also be viewed from Slides editor.

}

// The fill of the line.
type LineFill struct {
	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid color fill.

}

// The properties of the Line. When unset, these fields default to values that match the appearance of new lines created in the Slides editor.
type LineProperties struct {
	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the line.

	EndArrow string `json:"endArrow,omitempty\"` // The style of the arrow at the end of the line.

	EndConnection LineConnection `json:"endConnection,omitempty\"` // The connection at the end of the line. If unset, there is no connection. Only lines with a Type indicating it is a "connector" can have an `end_connection`.

	LineFill LineFill `json:"lineFill,omitempty\"` // The fill of the line. The default line fill matches the defaults for new lines created in the Slides editor.

	Link Link `json:"link,omitempty\"` // The hyperlink destination of the line. If unset, there is no link.

	StartArrow string `json:"startArrow,omitempty\"` // The style of the arrow at the beginning of the line.

	StartConnection LineConnection `json:"startConnection,omitempty\"` // The connection at the beginning of the line. If unset, there is no connection. Only lines with a Type indicating it is a "connector" can have a `start_connection`.

	Weight Dimension `json:"weight,omitempty\"` // The thickness of the line.

}

// A hypertext link.
type Link struct {
	PageObjectId string `json:"pageObjectId,omitempty\"` // If set, indicates this is a link to the specific page in this presentation with this ID. A page with this ID may not exist.

	RelativeLink string `json:"relativeLink,omitempty\"` // If set, indicates this is a link to a slide in this presentation, addressed by its position.

	SlideIndex int `json:"slideIndex,omitempty\"` // If set, indicates this is a link to the slide at this zero-based index in the presentation. There may not be a slide at this index.

	Url string `json:"url,omitempty\"` // If set, indicates this is a link to the external web page at this URL.

}

// A List describes the look and feel of bullets belonging to paragraphs associated with a list. A paragraph that is part of a list has an implicit reference to that list's ID.
type List struct {
	ListId string `json:"listId,omitempty\"` // The ID of the list.

	NestingLevel map[string]interface{} `json:"nestingLevel,omitempty\"` // A map of nesting levels to the properties of bullets at the associated level. A list has at most nine levels of nesting, so the possible values for the keys of this map are 0 through 8, inclusive.

}

// The properties of Page that are only relevant for pages with page_type MASTER.
type MasterProperties struct {
	DisplayName string `json:"displayName,omitempty\"` // The human-readable name of the master.

}

// Merges cells in a Table.
type MergeTableCellsRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	TableRange TableRange `json:"tableRange,omitempty\"` // The table range specifying which cells of the table to merge. Any text in the cells being merged will be concatenated and stored in the upper-left ("head") cell of the range. If the range is non-rectangular (which can occur in some cases where the range covers cells that are already merged), a 400 bad request error is returned.

}

// Contains properties describing the look and feel of a list bullet at a given level of nesting.
type NestingLevel struct {
	BulletStyle TextStyle `json:"bulletStyle,omitempty\"` // The style of a bullet at this level of nesting.

}

// The properties of Page that are only relevant for pages with page_type NOTES.
type NotesProperties struct {
	SpeakerNotesObjectId string `json:"speakerNotesObjectId,omitempty\"` // The object ID of the shape on this notes page that contains the speaker notes for the corresponding slide. The actual shape may not always exist on the notes page. Inserting text using this object ID will automatically create the shape. In this case, the actual shape may have different object ID. The `GetPresentation` or `GetPage` action will always return the latest object ID.

}

// A themeable solid color value.
type OpaqueColor struct {
	RgbColor RgbColor `json:"rgbColor,omitempty\"` // An opaque RGB color.

	ThemeColor string `json:"themeColor,omitempty\"` // An opaque theme color.

}

// A color that can either be fully opaque or fully transparent.
type OptionalColor struct {
	OpaqueColor OpaqueColor `json:"opaqueColor,omitempty\"` // If set, this will be used as an opaque color. If unset, this represents a transparent color.

}

// The outline of a PageElement. If these fields are unset, they may be inherited from a parent placeholder if it exists. If there is no parent, the fields will default to the value used for new page elements created in the Slides editor, which may depend on the page element kind.
type Outline struct {
	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the outline.

	OutlineFill OutlineFill `json:"outlineFill,omitempty\"` // The fill of the outline.

	PropertyState string `json:"propertyState,omitempty\"` // The outline property state. Updating the outline on a page element will implicitly update this field to `RENDERED`, unless another value is specified in the same request. To have no outline on a page element, set this field to `NOT_RENDERED`. In this case, any other outline fields set in the same request will be ignored.

	Weight Dimension `json:"weight,omitempty\"` // The thickness of the outline.

}

// The fill of the outline.
type OutlineFill struct {
	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid color fill.

}

// A page in a presentation.
type Page struct {
	LayoutProperties LayoutProperties `json:"layoutProperties,omitempty\"` // Layout specific properties. Only set if page_type = LAYOUT.

	MasterProperties MasterProperties `json:"masterProperties,omitempty\"` // Master specific properties. Only set if page_type = MASTER.

	NotesProperties NotesProperties `json:"notesProperties,omitempty\"` // Notes specific properties. Only set if page_type = NOTES.

	ObjectId string `json:"objectId,omitempty\"` // The object ID for this page. Object IDs used by Page and PageElement share the same namespace.

	PageElements []PageElement `json:"pageElements,omitempty\"` // The page elements rendered on the page.

	PageProperties PageProperties `json:"pageProperties,omitempty\"` // The properties of the page.

	PageType string `json:"pageType,omitempty\"` // The type of the page.

	RevisionId string `json:"revisionId,omitempty\"` // Output only. The revision ID of the presentation. Can be used in update requests to assert the presentation revision hasn't changed since the last read operation. Only populated if the user has edit access to the presentation. The revision ID is not a sequential number but an opaque string. The format of the revision ID might change over time. A returned revision ID is only guaranteed to be valid for 24 hours after it has been returned and cannot be shared across users. If the revision ID is unchanged between calls, then the presentation has not changed. Conversely, a changed ID (for the same presentation and user) usually means the presentation has been updated. However, a changed ID can also be due to internal factors such as ID format changes.

	SlideProperties SlideProperties `json:"slideProperties,omitempty\"` // Slide specific properties. Only set if page_type = SLIDE.

}

// The page background fill.
type PageBackgroundFill struct {
	PropertyState string `json:"propertyState,omitempty\"` // The background fill property state. Updating the fill on a page will implicitly update this field to `RENDERED`, unless another value is specified in the same request. To have no fill on a page, set this field to `NOT_RENDERED`. In this case, any other fill fields set in the same request will be ignored.

	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid color fill.

	StretchedPictureFill StretchedPictureFill `json:"stretchedPictureFill,omitempty\"` // Stretched picture fill.

}

// A visual element rendered on a page.
type PageElement struct {
	Description string `json:"description,omitempty\"` // The description of the page element. Combined with title to display alt text. The field is not supported for Group elements.

	ElementGroup Group `json:"elementGroup,omitempty\"` // A collection of page elements joined as a single unit.

	Image Image `json:"image,omitempty\"` // An image page element.

	Line Line `json:"line,omitempty\"` // A line page element.

	ObjectId string `json:"objectId,omitempty\"` // The object ID for this page element. Object IDs used by google.apps.slides.v1.Page and google.apps.slides.v1.PageElement share the same namespace.

	Shape Shape `json:"shape,omitempty\"` // A generic shape.

	SheetsChart SheetsChart `json:"sheetsChart,omitempty\"` // A linked chart embedded from Google Sheets. Unlinked charts are represented as images.

	Size Size `json:"size,omitempty\"` // The size of the page element.

	SpeakerSpotlight SpeakerSpotlight `json:"speakerSpotlight,omitempty\"` // A Speaker Spotlight.

	Table Table `json:"table,omitempty\"` // A table page element.

	Title string `json:"title,omitempty\"` // The title of the page element. Combined with description to display alt text. The field is not supported for Group elements.

	Transform AffineTransform `json:"transform,omitempty\"` // The transform of the page element. The visual appearance of the page element is determined by its absolute transform. To compute the absolute transform, preconcatenate a page element's transform with the transforms of all of its parent groups. If the page element is not in a group, its absolute transform is the same as the value in this field. The initial transform for the newly created Group is always the identity transform.

	Video Video `json:"video,omitempty\"` // A video page element.

	WordArt WordArt `json:"wordArt,omitempty\"` // A word art page element.

}

// Common properties for a page element. Note: When you initially create a PageElement, the API may modify the values of both `size` and `transform`, but the visual size will be unchanged.
type PageElementProperties struct {
	PageObjectId string `json:"pageObjectId,omitempty\"` // The object ID of the page where the element is located.

	Size Size `json:"size,omitempty\"` // The size of the element.

	Transform AffineTransform `json:"transform,omitempty\"` // The transform for the element.

}

// The properties of the Page. The page will inherit properties from the parent page. Depending on the page type the hierarchy is defined in either SlideProperties or LayoutProperties.
type PageProperties struct {
	ColorScheme ColorScheme `json:"colorScheme,omitempty\"` // The color scheme of the page. If unset, the color scheme is inherited from a parent page. If the page has no parent, the color scheme uses a default Slides color scheme, matching the defaults in the Slides editor. Only the concrete colors of the first 12 ThemeColorTypes are editable. In addition, only the color scheme on `Master` pages can be updated. To update the field, a color scheme containing mappings from all the first 12 ThemeColorTypes to their concrete colors must be provided. Colors for the remaining ThemeColorTypes will be ignored.

	PageBackgroundFill PageBackgroundFill `json:"pageBackgroundFill,omitempty\"` // The background fill of the page. If unset, the background fill is inherited from a parent page if it exists. If the page has no parent, then the background fill defaults to the corresponding fill in the Slides editor.

}

// A TextElement kind that represents the beginning of a new paragraph.
type ParagraphMarker struct {
	Bullet Bullet `json:"bullet,omitempty\"` // The bullet for this paragraph. If not present, the paragraph does not belong to a list.

	Style ParagraphStyle `json:"style,omitempty\"` // The paragraph's style

}

// Styles that apply to a whole paragraph. If this text is contained in a shape with a parent placeholder, then these paragraph styles may be inherited from the parent. Which paragraph styles are inherited depend on the nesting level of lists: * A paragraph not in a list will inherit its paragraph style from the paragraph at the 0 nesting level of the list inside the parent placeholder. * A paragraph in a list will inherit its paragraph style from the paragraph at its corresponding nesting level of the list inside the parent placeholder. Inherited paragraph styles are represented as unset fields in this message.
type ParagraphStyle struct {
	Alignment string `json:"alignment,omitempty\"` // The text alignment for this paragraph.

	Direction string `json:"direction,omitempty\"` // The text direction of this paragraph. If unset, the value defaults to LEFT_TO_RIGHT since text direction is not inherited.

	IndentEnd Dimension `json:"indentEnd,omitempty\"` // The amount indentation for the paragraph on the side that corresponds to the end of the text, based on the current text direction. If unset, the value is inherited from the parent.

	IndentFirstLine Dimension `json:"indentFirstLine,omitempty\"` // The amount of indentation for the start of the first line of the paragraph. If unset, the value is inherited from the parent.

	IndentStart Dimension `json:"indentStart,omitempty\"` // The amount indentation for the paragraph on the side that corresponds to the start of the text, based on the current text direction. If unset, the value is inherited from the parent.

	LineSpacing float64 `json:"lineSpacing,omitempty\"` // The amount of space between lines, as a percentage of normal, where normal is represented as 100.0. If unset, the value is inherited from the parent.

	SpaceAbove Dimension `json:"spaceAbove,omitempty\"` // The amount of extra space above the paragraph. If unset, the value is inherited from the parent.

	SpaceBelow Dimension `json:"spaceBelow,omitempty\"` // The amount of extra space below the paragraph. If unset, the value is inherited from the parent.

	SpacingMode string `json:"spacingMode,omitempty\"` // The spacing mode for the paragraph.

}

// The placeholder information that uniquely identifies a placeholder shape.
type Placeholder struct {
	Index int `json:"index,omitempty\"` // The index of the placeholder. If the same placeholder types are present in the same page, they would have different index values.

	ParentObjectId string `json:"parentObjectId,omitempty\"` // The object ID of this shape's parent placeholder. If unset, the parent placeholder shape does not exist, so the shape does not inherit properties from any other shape.

	TypeValue string `json:"type,omitempty\"` // The type of the placeholder.

}

// A Google Slides presentation.
type Presentation struct {
	Layouts []Page `json:"layouts,omitempty\"` // The layouts in the presentation. A layout is a template that determines how content is arranged and styled on the slides that inherit from that layout.

	Locale string `json:"locale,omitempty\"` // The locale of the presentation, as an IETF BCP 47 language tag.

	Masters []Page `json:"masters,omitempty\"` // The slide masters in the presentation. A slide master contains all common page elements and the common properties for a set of layouts. They serve three purposes: - Placeholder shapes on a master contain the default text styles and shape properties of all placeholder shapes on pages that use that master. - The master page properties define the common page properties inherited by its layouts. - Any other shapes on the master slide appear on all slides using that master, regardless of their layout.

	NotesMaster Page `json:"notesMaster,omitempty\"` // The notes master in the presentation. It serves three purposes: - Placeholder shapes on a notes master contain the default text styles and shape properties of all placeholder shapes on notes pages. Specifically, a `SLIDE_IMAGE` placeholder shape contains the slide thumbnail, and a `BODY` placeholder shape contains the speaker notes. - The notes master page properties define the common page properties inherited by all notes pages. - Any other shapes on the notes master appear on all notes pages. The notes master is read-only.

	PageSize Size `json:"pageSize,omitempty\"` // The size of pages in the presentation.

	PresentationId string `json:"presentationId,omitempty\"` // The ID of the presentation.

	RevisionId string `json:"revisionId,omitempty\"` // Output only. The revision ID of the presentation. Can be used in update requests to assert the presentation revision hasn't changed since the last read operation. Only populated if the user has edit access to the presentation. The revision ID is not a sequential number but a nebulous string. The format of the revision ID may change over time, so it should be treated opaquely. A returned revision ID is only guaranteed to be valid for 24 hours after it has been returned and cannot be shared across users. If the revision ID is unchanged between calls, then the presentation has not changed. Conversely, a changed ID (for the same presentation and user) usually means the presentation has been updated. However, a changed ID can also be due to internal factors such as ID format changes.

	Slides []Page `json:"slides,omitempty\"` // The slides in the presentation. A slide inherits properties from a slide layout.

	Title string `json:"title,omitempty\"` // The title of the presentation.

}

// Specifies a contiguous range of an indexed collection, such as characters in text.
type RangeValue struct {
	EndIndex int `json:"endIndex,omitempty\"` // The optional zero-based index of the end of the collection. Required for `FIXED_RANGE` ranges.

	StartIndex int `json:"startIndex,omitempty\"` // The optional zero-based index of the beginning of the collection. Required for `FIXED_RANGE` and `FROM_START_INDEX` ranges.

	TypeValue string `json:"type,omitempty\"` // The type of range.

}

// A recolor effect applied on an image.
type Recolor struct {
	Name string `json:"name,omitempty\"` // The name of the recolor effect. The name is determined from the `recolor_stops` by matching the gradient against the colors in the page's current color scheme. This property is read-only.

	RecolorStops []ColorStop `json:"recolorStops,omitempty\"` // The recolor effect is represented by a gradient, which is a list of color stops. The colors in the gradient will replace the corresponding colors at the same position in the color palette and apply to the image. This property is read-only.

}

// Refreshes an embedded Google Sheets chart by replacing it with the latest version of the chart from Google Sheets. NOTE: Refreshing charts requires at least one of the spreadsheets.readonly, spreadsheets, drive.readonly, or drive OAuth scopes.
type RefreshSheetsChartRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the chart to refresh.

}

// Replaces all shapes that match the given criteria with the provided image. The images replacing the shapes are rectangular after being inserted into the presentation and do not take on the forms of the shapes.
type ReplaceAllShapesWithImageRequest struct {
	ContainsText SubstringMatchCriteria `json:"containsText,omitempty\"` // If set, this request will replace all of the shapes that contain the given text.

	ImageReplaceMethod string `json:"imageReplaceMethod,omitempty\"` // The image replace method. If you specify both a `replace_method` and an `image_replace_method`, the `image_replace_method` takes precedence. If you do not specify a value for `image_replace_method`, but specify a value for `replace_method`, then the specified `replace_method` value is used. If you do not specify either, then CENTER_INSIDE is used.

	ImageUrl string `json:"imageUrl,omitempty\"` // The image URL. The image is fetched once at insertion time and a copy is stored for display inside the presentation. Images must be less than 50MB in size, cannot exceed 25 megapixels, and must be in one of PNG, JPEG, or GIF format. The provided URL can be at most 2 kB in length. The URL itself is saved with the image, and exposed via the Image.source_url field.

	PageObjectIds []string `json:"pageObjectIds,omitempty\"` // If non-empty, limits the matches to page elements only on the given pages. Returns a 400 bad request error if given the page object ID of a notes page or a notes master, or if a page with that object ID doesn't exist in the presentation.

	ReplaceMethod string `json:"replaceMethod,omitempty\"` // The replace method. *Deprecated*: use `image_replace_method` instead. If you specify both a `replace_method` and an `image_replace_method`, the `image_replace_method` takes precedence.

}

// The result of replacing shapes with an image.
type ReplaceAllShapesWithImageResponse struct {
	OccurrencesChanged int `json:"occurrencesChanged,omitempty\"` // The number of shapes replaced with images.

}

// Replaces all shapes that match the given criteria with the provided Google Sheets chart. The chart will be scaled and centered to fit within the bounds of the original shape. NOTE: Replacing shapes with a chart requires at least one of the spreadsheets.readonly, spreadsheets, drive.readonly, or drive OAuth scopes.
type ReplaceAllShapesWithSheetsChartRequest struct {
	ChartId int `json:"chartId,omitempty\"` // The ID of the specific chart in the Google Sheets spreadsheet.

	ContainsText SubstringMatchCriteria `json:"containsText,omitempty\"` // The criteria that the shapes must match in order to be replaced. The request will replace all of the shapes that contain the given text.

	LinkingMode string `json:"linkingMode,omitempty\"` // The mode with which the chart is linked to the source spreadsheet. When not specified, the chart will be an image that is not linked.

	PageObjectIds []string `json:"pageObjectIds,omitempty\"` // If non-empty, limits the matches to page elements only on the given pages. Returns a 400 bad request error if given the page object ID of a notes page or a notes master, or if a page with that object ID doesn't exist in the presentation.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the Google Sheets spreadsheet that contains the chart.

}

// The result of replacing shapes with a Google Sheets chart.
type ReplaceAllShapesWithSheetsChartResponse struct {
	OccurrencesChanged int `json:"occurrencesChanged,omitempty\"` // The number of shapes replaced with charts.

}

// Replaces all instances of text matching a criteria with replace text.
type ReplaceAllTextRequest struct {
	ContainsText SubstringMatchCriteria `json:"containsText,omitempty\"` // Finds text in a shape matching this substring.

	PageObjectIds []string `json:"pageObjectIds,omitempty\"` // If non-empty, limits the matches to page elements only on the given pages. Returns a 400 bad request error if given the page object ID of a notes master, or if a page with that object ID doesn't exist in the presentation.

	ReplaceText string `json:"replaceText,omitempty\"` // The text that will replace the matched text.

}

// The result of replacing text.
type ReplaceAllTextResponse struct {
	OccurrencesChanged int `json:"occurrencesChanged,omitempty\"` // The number of occurrences changed by replacing all text.

}

// Replaces an existing image with a new image. Replacing an image removes some image effects from the existing image.
type ReplaceImageRequest struct {
	ImageObjectId string `json:"imageObjectId,omitempty\"` // The ID of the existing image that will be replaced. The ID can be retrieved from the response of a get request.

	ImageReplaceMethod string `json:"imageReplaceMethod,omitempty\"` // The replacement method.

	Url string `json:"url,omitempty\"` // The image URL. The image is fetched once at insertion time and a copy is stored for display inside the presentation. Images must be less than 50MB, cannot exceed 25 megapixels, and must be in PNG, JPEG, or GIF format. The provided URL can't surpass 2 KB in length. The URL is saved with the image, and exposed through the Image.source_url field.

}

// A single kind of update to apply to a presentation.
type Request struct {
	CreateImage CreateImageRequest `json:"createImage,omitempty\"` // Creates an image.

	CreateLine CreateLineRequest `json:"createLine,omitempty\"` // Creates a line.

	CreateParagraphBullets CreateParagraphBulletsRequest `json:"createParagraphBullets,omitempty\"` // Creates bullets for paragraphs.

	CreateShape CreateShapeRequest `json:"createShape,omitempty\"` // Creates a new shape.

	CreateSheetsChart CreateSheetsChartRequest `json:"createSheetsChart,omitempty\"` // Creates an embedded Google Sheets chart.

	CreateSlide CreateSlideRequest `json:"createSlide,omitempty\"` // Creates a new slide.

	CreateTable CreateTableRequest `json:"createTable,omitempty\"` // Creates a new table.

	CreateVideo CreateVideoRequest `json:"createVideo,omitempty\"` // Creates a video.

	DeleteObject DeleteObjectRequest `json:"deleteObject,omitempty\"` // Deletes a page or page element from the presentation.

	DeleteParagraphBullets DeleteParagraphBulletsRequest `json:"deleteParagraphBullets,omitempty\"` // Deletes bullets from paragraphs.

	DeleteTableColumn DeleteTableColumnRequest `json:"deleteTableColumn,omitempty\"` // Deletes a column from a table.

	DeleteTableRow DeleteTableRowRequest `json:"deleteTableRow,omitempty\"` // Deletes a row from a table.

	DeleteText DeleteTextRequest `json:"deleteText,omitempty\"` // Deletes text from a shape or a table cell.

	DuplicateObject DuplicateObjectRequest `json:"duplicateObject,omitempty\"` // Duplicates a slide or page element.

	GroupObjects GroupObjectsRequest `json:"groupObjects,omitempty\"` // Groups objects, such as page elements.

	InsertTableColumns InsertTableColumnsRequest `json:"insertTableColumns,omitempty\"` // Inserts columns into a table.

	InsertTableRows InsertTableRowsRequest `json:"insertTableRows,omitempty\"` // Inserts rows into a table.

	InsertText InsertTextRequest `json:"insertText,omitempty\"` // Inserts text into a shape or table cell.

	MergeTableCells MergeTableCellsRequest `json:"mergeTableCells,omitempty\"` // Merges cells in a Table.

	RefreshSheetsChart RefreshSheetsChartRequest `json:"refreshSheetsChart,omitempty\"` // Refreshes a Google Sheets chart.

	ReplaceAllShapesWithImage ReplaceAllShapesWithImageRequest `json:"replaceAllShapesWithImage,omitempty\"` // Replaces all shapes matching some criteria with an image.

	ReplaceAllShapesWithSheetsChart ReplaceAllShapesWithSheetsChartRequest `json:"replaceAllShapesWithSheetsChart,omitempty\"` // Replaces all shapes matching some criteria with a Google Sheets chart.

	ReplaceAllText ReplaceAllTextRequest `json:"replaceAllText,omitempty\"` // Replaces all instances of specified text.

	ReplaceImage ReplaceImageRequest `json:"replaceImage,omitempty\"` // Replaces an existing image with a new image.

	RerouteLine RerouteLineRequest `json:"rerouteLine,omitempty\"` // Reroutes a line such that it's connected at the two closest connection sites on the connected page elements.

	UngroupObjects UngroupObjectsRequest `json:"ungroupObjects,omitempty\"` // Ungroups objects, such as groups.

	UnmergeTableCells UnmergeTableCellsRequest `json:"unmergeTableCells,omitempty\"` // Unmerges cells in a Table.

	UpdateImageProperties UpdateImagePropertiesRequest `json:"updateImageProperties,omitempty\"` // Updates the properties of an Image.

	UpdateLineCategory UpdateLineCategoryRequest `json:"updateLineCategory,omitempty\"` // Updates the category of a line.

	UpdateLineProperties UpdateLinePropertiesRequest `json:"updateLineProperties,omitempty\"` // Updates the properties of a Line.

	UpdatePageElementAltText UpdatePageElementAltTextRequest `json:"updatePageElementAltText,omitempty\"` // Updates the alt text title and/or description of a page element.

	UpdatePageElementTransform UpdatePageElementTransformRequest `json:"updatePageElementTransform,omitempty\"` // Updates the transform of a page element.

	UpdatePageElementsZOrder UpdatePageElementsZOrderRequest `json:"updatePageElementsZOrder,omitempty\"` // Updates the Z-order of page elements.

	UpdatePageProperties UpdatePagePropertiesRequest `json:"updatePageProperties,omitempty\"` // Updates the properties of a Page.

	UpdateParagraphStyle UpdateParagraphStyleRequest `json:"updateParagraphStyle,omitempty\"` // Updates the styling of paragraphs within a Shape or Table.

	UpdateShapeProperties UpdateShapePropertiesRequest `json:"updateShapeProperties,omitempty\"` // Updates the properties of a Shape.

	UpdateSlideProperties UpdateSlidePropertiesRequest `json:"updateSlideProperties,omitempty\"` // Updates the properties of a Slide

	UpdateSlidesPosition UpdateSlidesPositionRequest `json:"updateSlidesPosition,omitempty\"` // Updates the position of a set of slides in the presentation.

	UpdateTableBorderProperties UpdateTableBorderPropertiesRequest `json:"updateTableBorderProperties,omitempty\"` // Updates the properties of the table borders in a Table.

	UpdateTableCellProperties UpdateTableCellPropertiesRequest `json:"updateTableCellProperties,omitempty\"` // Updates the properties of a TableCell.

	UpdateTableColumnProperties UpdateTableColumnPropertiesRequest `json:"updateTableColumnProperties,omitempty\"` // Updates the properties of a Table column.

	UpdateTableRowProperties UpdateTableRowPropertiesRequest `json:"updateTableRowProperties,omitempty\"` // Updates the properties of a Table row.

	UpdateTextStyle UpdateTextStyleRequest `json:"updateTextStyle,omitempty\"` // Updates the styling of text within a Shape or Table.

	UpdateVideoProperties UpdateVideoPropertiesRequest `json:"updateVideoProperties,omitempty\"` // Updates the properties of a Video.

}

// Reroutes a line such that it's connected at the two closest connection sites on the connected page elements.
type RerouteLineRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the line to reroute. Only a line with a category indicating it is a "connector" can be rerouted. The start and end connections of the line must be on different page elements.

}

// A single response from an update.
type Response struct {
	CreateImage CreateImageResponse `json:"createImage,omitempty\"` // The result of creating an image.

	CreateLine CreateLineResponse `json:"createLine,omitempty\"` // The result of creating a line.

	CreateShape CreateShapeResponse `json:"createShape,omitempty\"` // The result of creating a shape.

	CreateSheetsChart CreateSheetsChartResponse `json:"createSheetsChart,omitempty\"` // The result of creating a Google Sheets chart.

	CreateSlide CreateSlideResponse `json:"createSlide,omitempty\"` // The result of creating a slide.

	CreateTable CreateTableResponse `json:"createTable,omitempty\"` // The result of creating a table.

	CreateVideo CreateVideoResponse `json:"createVideo,omitempty\"` // The result of creating a video.

	DuplicateObject DuplicateObjectResponse `json:"duplicateObject,omitempty\"` // The result of duplicating an object.

	GroupObjects GroupObjectsResponse `json:"groupObjects,omitempty\"` // The result of grouping objects.

	ReplaceAllShapesWithImage ReplaceAllShapesWithImageResponse `json:"replaceAllShapesWithImage,omitempty\"` // The result of replacing all shapes matching some criteria with an image.

	ReplaceAllShapesWithSheetsChart ReplaceAllShapesWithSheetsChartResponse `json:"replaceAllShapesWithSheetsChart,omitempty\"` // The result of replacing all shapes matching some criteria with a Google Sheets chart.

	ReplaceAllText ReplaceAllTextResponse `json:"replaceAllText,omitempty\"` // The result of replacing text.

}

// An RGB color.
type RgbColor struct {
	Blue float64 `json:"blue,omitempty\"` // The blue component of the color, from 0.0 to 1.0.

	Green float64 `json:"green,omitempty\"` // The green component of the color, from 0.0 to 1.0.

	Red float64 `json:"red,omitempty\"` // The red component of the color, from 0.0 to 1.0.

}

// The shadow properties of a page element. If these fields are unset, they may be inherited from a parent placeholder if it exists. If there is no parent, the fields will default to the value used for new page elements created in the Slides editor, which may depend on the page element kind.
type Shadow struct {
	Alignment string `json:"alignment,omitempty\"` // The alignment point of the shadow, that sets the origin for translate, scale and skew of the shadow. This property is read-only.

	Alpha float64 `json:"alpha,omitempty\"` // The alpha of the shadow's color, from 0.0 to 1.0.

	BlurRadius Dimension `json:"blurRadius,omitempty\"` // The radius of the shadow blur. The larger the radius, the more diffuse the shadow becomes.

	Color OpaqueColor `json:"color,omitempty\"` // The shadow color value.

	PropertyState string `json:"propertyState,omitempty\"` // The shadow property state. Updating the shadow on a page element will implicitly update this field to `RENDERED`, unless another value is specified in the same request. To have no shadow on a page element, set this field to `NOT_RENDERED`. In this case, any other shadow fields set in the same request will be ignored.

	RotateWithShape bool `json:"rotateWithShape,omitempty\"` // Whether the shadow should rotate with the shape. This property is read-only.

	Transform AffineTransform `json:"transform,omitempty\"` // Transform that encodes the translate, scale, and skew of the shadow, relative to the alignment position.

	TypeValue string `json:"type,omitempty\"` // The type of the shadow. This property is read-only.

}

// A PageElement kind representing a generic shape that doesn't have a more specific classification. For more information, see [Size and position page elements](https://developers.google.com/workspace/slides/api/guides/transform).
type Shape struct {
	Placeholder Placeholder `json:"placeholder,omitempty\"` // Placeholders are page elements that inherit from corresponding placeholders on layouts and masters. If set, the shape is a placeholder shape and any inherited properties can be resolved by looking at the parent placeholder identified by the Placeholder.parent_object_id field.

	ShapeProperties ShapeProperties `json:"shapeProperties,omitempty\"` // The properties of the shape.

	ShapeType string `json:"shapeType,omitempty\"` // The type of the shape.

	Text TextContent `json:"text,omitempty\"` // The text content of the shape.

}

// The shape background fill.
type ShapeBackgroundFill struct {
	PropertyState string `json:"propertyState,omitempty\"` // The background fill property state. Updating the fill on a shape will implicitly update this field to `RENDERED`, unless another value is specified in the same request. To have no fill on a shape, set this field to `NOT_RENDERED`. In this case, any other fill fields set in the same request will be ignored.

	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid color fill.

}

// The properties of a Shape. If the shape is a placeholder shape as determined by the placeholder field, then these properties may be inherited from a parent placeholder shape. Determining the rendered value of the property depends on the corresponding property_state field value. Any text autofit settings on the shape are automatically deactivated by requests that can impact how text fits in the shape.
type ShapeProperties struct {
	Autofit Autofit `json:"autofit,omitempty\"` // The autofit properties of the shape. This property is only set for shapes that allow text.

	ContentAlignment string `json:"contentAlignment,omitempty\"` // The alignment of the content in the shape. If unspecified, the alignment is inherited from a parent placeholder if it exists. If the shape has no parent, the default alignment matches the alignment for new shapes created in the Slides editor.

	Link Link `json:"link,omitempty\"` // The hyperlink destination of the shape. If unset, there is no link. Links are not inherited from parent placeholders.

	Outline Outline `json:"outline,omitempty\"` // The outline of the shape. If unset, the outline is inherited from a parent placeholder if it exists. If the shape has no parent, then the default outline depends on the shape type, matching the defaults for new shapes created in the Slides editor.

	Shadow Shadow `json:"shadow,omitempty\"` // The shadow properties of the shape. If unset, the shadow is inherited from a parent placeholder if it exists. If the shape has no parent, then the default shadow matches the defaults for new shapes created in the Slides editor. This property is read-only.

	ShapeBackgroundFill ShapeBackgroundFill `json:"shapeBackgroundFill,omitempty\"` // The background fill of the shape. If unset, the background fill is inherited from a parent placeholder if it exists. If the shape has no parent, then the default background fill depends on the shape type, matching the defaults for new shapes created in the Slides editor.

}

// A PageElement kind representing a linked chart embedded from Google Sheets.
type SheetsChart struct {
	ChartId int `json:"chartId,omitempty\"` // The ID of the specific chart in the Google Sheets spreadsheet that is embedded.

	ContentUrl string `json:"contentUrl,omitempty\"` // The URL of an image of the embedded chart, with a default lifetime of 30 minutes. This URL is tagged with the account of the requester. Anyone with the URL effectively accesses the image as the original requester. Access to the image may be lost if the presentation's sharing settings change.

	SheetsChartProperties SheetsChartProperties `json:"sheetsChartProperties,omitempty\"` // The properties of the Sheets chart.

	SpreadsheetId string `json:"spreadsheetId,omitempty\"` // The ID of the Google Sheets spreadsheet that contains the source chart.

}

// The properties of the SheetsChart.
type SheetsChartProperties struct {
	ChartImageProperties ImageProperties `json:"chartImageProperties,omitempty\"` // The properties of the embedded chart image.

}

// A width and height.
type Size struct {
	Height Dimension `json:"height,omitempty\"` // The height of the object.

	Width Dimension `json:"width,omitempty\"` // The width of the object.

}

// The properties of Page that are only relevant for pages with page_type SLIDE.
type SlideProperties struct {
	IsSkipped bool `json:"isSkipped,omitempty\"` // Whether the slide is skipped in the presentation mode. Defaults to false.

	LayoutObjectId string `json:"layoutObjectId,omitempty\"` // The object ID of the layout that this slide is based on. This property is read-only.

	MasterObjectId string `json:"masterObjectId,omitempty\"` // The object ID of the master that this slide is based on. This property is read-only.

	NotesPage Page `json:"notesPage,omitempty\"` // The notes page that this slide is associated with. It defines the visual appearance of a notes page when printing or exporting slides with speaker notes. A notes page inherits properties from the notes master. The placeholder shape with type BODY on the notes page contains the speaker notes for this slide. The ID of this shape is identified by the speakerNotesObjectId field. The notes page is read-only except for the text content and styles of the speaker notes shape. This property is read-only.

}

// A solid color fill. The page or page element is filled entirely with the specified color value. If any field is unset, its value may be inherited from a parent placeholder if it exists.
type SolidFill struct {
	Alpha float64 `json:"alpha,omitempty\"` // The fraction of this `color` that should be applied to the pixel. That is, the final pixel color is defined by the equation: pixel color = alpha * (color) + (1.0 - alpha) * (background color) This means that a value of 1.0 corresponds to a solid color, whereas a value of 0.0 corresponds to a completely transparent color.

	Color OpaqueColor `json:"color,omitempty\"` // The color value of the solid fill.

}

// A PageElement kind representing a Speaker Spotlight.
type SpeakerSpotlight struct {
	SpeakerSpotlightProperties SpeakerSpotlightProperties `json:"speakerSpotlightProperties,omitempty\"` // The properties of the Speaker Spotlight.

}

// The properties of the SpeakerSpotlight.
type SpeakerSpotlightProperties struct {
	Outline Outline `json:"outline,omitempty\"` // The outline of the Speaker Spotlight. If not set, it has no outline.

	Shadow Shadow `json:"shadow,omitempty\"` // The shadow of the Speaker Spotlight. If not set, it has no shadow.

}

// The stretched picture fill. The page or page element is filled entirely with the specified picture. The picture is stretched to fit its container.
type StretchedPictureFill struct {
	ContentUrl string `json:"contentUrl,omitempty\"` // Reading the content_url: An URL to a picture with a default lifetime of 30 minutes. This URL is tagged with the account of the requester. Anyone with the URL effectively accesses the picture as the original requester. Access to the picture may be lost if the presentation's sharing settings change. Writing the content_url: The picture is fetched once at insertion time and a copy is stored for display inside the presentation. Pictures must be less than 50MB in size, cannot exceed 25 megapixels, and must be in one of PNG, JPEG, or GIF format. The provided URL can be at most 2 kB in length.

	Size Size `json:"size,omitempty\"` // The original size of the picture fill. This field is read-only.

}

// A criteria that matches a specific string of text in a shape or table.
type SubstringMatchCriteria struct {
	MatchCase bool `json:"matchCase,omitempty\"` // Indicates whether the search should respect case: - `True`: the search is case sensitive. - `False`: the search is case insensitive.

	SearchByRegex bool `json:"searchByRegex,omitempty\"` // Optional. True if the find value should be treated as a regular expression. Any backslashes in the pattern should be escaped. - `True`: the search text is treated as a regular expressions. - `False`: the search text is treated as a substring for matching.

	Text string `json:"text,omitempty\"` // The text to search for in the shape or table.

}

// A PageElement kind representing a table.
type Table struct {
	Columns int `json:"columns,omitempty\"` // Number of columns in the table.

	HorizontalBorderRows []TableBorderRow `json:"horizontalBorderRows,omitempty\"` // Properties of horizontal cell borders. A table's horizontal cell borders are represented as a grid. The grid has one more row than the number of rows in the table and the same number of columns as the table. For example, if the table is 3 x 3, its horizontal borders will be represented as a grid with 4 rows and 3 columns.

	Rows int `json:"rows,omitempty\"` // Number of rows in the table.

	TableColumns []TableColumnProperties `json:"tableColumns,omitempty\"` // Properties of each column.

	TableRows []TableRow `json:"tableRows,omitempty\"` // Properties and contents of each row. Cells that span multiple rows are contained in only one of these rows and have a row_span greater than 1.

	VerticalBorderRows []TableBorderRow `json:"verticalBorderRows,omitempty\"` // Properties of vertical cell borders. A table's vertical cell borders are represented as a grid. The grid has the same number of rows as the table and one more column than the number of columns in the table. For example, if the table is 3 x 3, its vertical borders will be represented as a grid with 3 rows and 4 columns.

}

// The properties of each border cell.
type TableBorderCell struct {
	Location TableCellLocation `json:"location,omitempty\"` // The location of the border within the border table.

	TableBorderProperties TableBorderProperties `json:"tableBorderProperties,omitempty\"` // The border properties.

}

// The fill of the border.
type TableBorderFill struct {
	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid fill.

}

// The border styling properties of the TableBorderCell.
type TableBorderProperties struct {
	DashStyle string `json:"dashStyle,omitempty\"` // The dash style of the border.

	TableBorderFill TableBorderFill `json:"tableBorderFill,omitempty\"` // The fill of the table border.

	Weight Dimension `json:"weight,omitempty\"` // The thickness of the border.

}

// Contents of each border row in a table.
type TableBorderRow struct {
	TableBorderCells []TableBorderCell `json:"tableBorderCells,omitempty\"` // Properties of each border cell. When a border's adjacent table cells are merged, it is not included in the response.

}

// Properties and contents of each table cell.
type TableCell struct {
	ColumnSpan int `json:"columnSpan,omitempty\"` // Column span of the cell.

	Location TableCellLocation `json:"location,omitempty\"` // The location of the cell within the table.

	RowSpan int `json:"rowSpan,omitempty\"` // Row span of the cell.

	TableCellProperties TableCellProperties `json:"tableCellProperties,omitempty\"` // The properties of the table cell.

	Text TextContent `json:"text,omitempty\"` // The text content of the cell.

}

// The table cell background fill.
type TableCellBackgroundFill struct {
	PropertyState string `json:"propertyState,omitempty\"` // The background fill property state. Updating the fill on a table cell will implicitly update this field to `RENDERED`, unless another value is specified in the same request. To have no fill on a table cell, set this field to `NOT_RENDERED`. In this case, any other fill fields set in the same request will be ignored.

	SolidFill SolidFill `json:"solidFill,omitempty\"` // Solid color fill.

}

// A location of a single table cell within a table.
type TableCellLocation struct {
	ColumnIndex int `json:"columnIndex,omitempty\"` // The 0-based column index.

	RowIndex int `json:"rowIndex,omitempty\"` // The 0-based row index.

}

// The properties of the TableCell.
type TableCellProperties struct {
	ContentAlignment string `json:"contentAlignment,omitempty\"` // The alignment of the content in the table cell. The default alignment matches the alignment for newly created table cells in the Slides editor.

	TableCellBackgroundFill TableCellBackgroundFill `json:"tableCellBackgroundFill,omitempty\"` // The background fill of the table cell. The default fill matches the fill for newly created table cells in the Slides editor.

}

// Properties of each column in a table.
type TableColumnProperties struct {
	ColumnWidth Dimension `json:"columnWidth,omitempty\"` // Width of a column.

}

// A table range represents a reference to a subset of a table. It's important to note that the cells specified by a table range do not necessarily form a rectangle. For example, let's say we have a 3 x 3 table where all the cells of the last row are merged together. The table looks like this: [ ] A table range with location = (0, 0), row span = 3 and column span = 2 specifies the following cells: x x [ x x x ]
type TableRange struct {
	ColumnSpan int `json:"columnSpan,omitempty\"` // The column span of the table range.

	Location TableCellLocation `json:"location,omitempty\"` // The starting location of the table range.

	RowSpan int `json:"rowSpan,omitempty\"` // The row span of the table range.

}

// Properties and contents of each row in a table.
type TableRow struct {
	RowHeight Dimension `json:"rowHeight,omitempty\"` // Height of a row.

	TableCells []TableCell `json:"tableCells,omitempty\"` // Properties and contents of each cell. Cells that span multiple columns are represented only once with a column_span greater than 1. As a result, the length of this collection does not always match the number of columns of the entire table.

	TableRowProperties TableRowProperties `json:"tableRowProperties,omitempty\"` // Properties of the row.

}

// Properties of each row in a table.
type TableRowProperties struct {
	MinRowHeight Dimension `json:"minRowHeight,omitempty\"` // Minimum height of the row. The row will be rendered in the Slides editor at a height equal to or greater than this value in order to show all the text in the row's cell(s).

}

// The general text content. The text must reside in a compatible shape (e.g. text box or rectangle) or a table cell in a page.
type TextContent struct {
	Lists map[string]interface{} `json:"lists,omitempty\"` // The bulleted lists contained in this text, keyed by list ID.

	TextElements []TextElement `json:"textElements,omitempty\"` // The text contents broken down into its component parts, including styling information. This property is read-only.

}

// A TextElement describes the content of a range of indices in the text content of a Shape or TableCell.
type TextElement struct {
	AutoText AutoText `json:"autoText,omitempty\"` // A TextElement representing a spot in the text that is dynamically replaced with content that can change over time.

	EndIndex int `json:"endIndex,omitempty\"` // The zero-based end index of this text element, exclusive, in Unicode code units.

	ParagraphMarker ParagraphMarker `json:"paragraphMarker,omitempty\"` // A marker representing the beginning of a new paragraph. The `start_index` and `end_index` of this TextElement represent the range of the paragraph. Other TextElements with an index range contained inside this paragraph's range are considered to be part of this paragraph. The range of indices of two separate paragraphs will never overlap.

	StartIndex int `json:"startIndex,omitempty\"` // The zero-based start index of this text element, in Unicode code units.

	TextRun TextRun `json:"textRun,omitempty\"` // A TextElement representing a run of text where all of the characters in the run have the same TextStyle. The `start_index` and `end_index` of TextRuns will always be fully contained in the index range of a single `paragraph_marker` TextElement. In other words, a TextRun will never span multiple paragraphs.

}

// A TextElement kind that represents a run of text that all has the same styling.
type TextRun struct {
	Content string `json:"content,omitempty\"` // The text of this run.

	Style TextStyle `json:"style,omitempty\"` // The styling applied to this run.

}

// Represents the styling that can be applied to a TextRun. If this text is contained in a shape with a parent placeholder, then these text styles may be inherited from the parent. Which text styles are inherited depend on the nesting level of lists: * A text run in a paragraph that is not in a list will inherit its text style from the the newline character in the paragraph at the 0 nesting level of the list inside the parent placeholder. * A text run in a paragraph that is in a list will inherit its text style from the newline character in the paragraph at its corresponding nesting level of the list inside the parent placeholder. Inherited text styles are represented as unset fields in this message. If text is contained in a shape without a parent placeholder, unsetting these fields will revert the style to a value matching the defaults in the Slides editor.
type TextStyle struct {
	BackgroundColor OptionalColor `json:"backgroundColor,omitempty\"` // The background color of the text. If set, the color is either opaque or transparent, depending on if the `opaque_color` field in it is set.

	BaselineOffset string `json:"baselineOffset,omitempty\"` // The text's vertical offset from its normal position. Text with `SUPERSCRIPT` or `SUBSCRIPT` baseline offsets is automatically rendered in a smaller font size, computed based on the `font_size` field. The `font_size` itself is not affected by changes in this field.

	Bold bool `json:"bold,omitempty\"` // Whether or not the text is rendered as bold.

	FontFamily string `json:"fontFamily,omitempty\"` // The font family of the text. The font family can be any font from the Font menu in Slides or from [Google Fonts] (https://fonts.google.com/). If the font name is unrecognized, the text is rendered in `Arial`. Some fonts can affect the weight of the text. If an update request specifies values for both `font_family` and `bold`, the explicitly-set `bold` value is used.

	FontSize Dimension `json:"fontSize,omitempty\"` // The size of the text's font. When read, the `font_size` will specified in points.

	ForegroundColor OptionalColor `json:"foregroundColor,omitempty\"` // The color of the text itself. If set, the color is either opaque or transparent, depending on if the `opaque_color` field in it is set.

	Italic bool `json:"italic,omitempty\"` // Whether or not the text is italicized.

	Link Link `json:"link,omitempty\"` // The hyperlink destination of the text. If unset, there is no link. Links are not inherited from parent text. Changing the link in an update request causes some other changes to the text style of the range: * When setting a link, the text foreground color will be set to ThemeColorType.HYPERLINK and the text will be underlined. If these fields are modified in the same request, those values will be used instead of the link defaults. * Setting a link on a text range that overlaps with an existing link will also update the existing link to point to the new URL. * Links are not settable on newline characters. As a result, setting a link on a text range that crosses a paragraph boundary, such as `"ABC\n123"`, will separate the newline character(s) into their own text runs. The link will be applied separately to the runs before and after the newline. * Removing a link will update the text style of the range to match the style of the preceding text (or the default text styles if the preceding text is another link) unless different styles are being set in the same request.

	SmallCaps bool `json:"smallCaps,omitempty\"` // Whether or not the text is in small capital letters.

	Strikethrough bool `json:"strikethrough,omitempty\"` // Whether or not the text is struck through.

	Underline bool `json:"underline,omitempty\"` // Whether or not the text is underlined.

	WeightedFontFamily WeightedFontFamily `json:"weightedFontFamily,omitempty\"` // The font family and rendered weight of the text. This field is an extension of `font_family` meant to support explicit font weights without breaking backwards compatibility. As such, when reading the style of a range of text, the value of `weighted_font_family#font_family` will always be equal to that of `font_family`. However, when writing, if both fields are included in the field mask (either explicitly or through the wildcard `"*"`), their values are reconciled as follows: * If `font_family` is set and `weighted_font_family` is not, the value of `font_family` is applied with weight `400` ("normal"). * If both fields are set, the value of `font_family` must match that of `weighted_font_family#font_family`. If so, the font family and weight of `weighted_font_family` is applied. Otherwise, a 400 bad request error is returned. * If `weighted_font_family` is set and `font_family` is not, the font family and weight of `weighted_font_family` is applied. * If neither field is set, the font family and weight of the text inherit from the parent. Note that these properties cannot inherit separately from each other. If an update request specifies values for both `weighted_font_family` and `bold`, the `weighted_font_family` is applied first, then `bold`. If `weighted_font_family#weight` is not set, it defaults to `400`. If `weighted_font_family` is set, then `weighted_font_family#font_family` must also be set with a non-empty value. Otherwise, a 400 bad request error is returned.

}

// A pair mapping a theme color type to the concrete color it represents.
type ThemeColorPair struct {
	Color RgbColor `json:"color,omitempty\"` // The concrete color corresponding to the theme color type above.

	TypeValue string `json:"type,omitempty\"` // The type of the theme color.

}

// The thumbnail of a page.
type Thumbnail struct {
	ContentUrl string `json:"contentUrl,omitempty\"` // The content URL of the thumbnail image. The URL to the image has a default lifetime of 30 minutes. This URL is tagged with the account of the requester. Anyone with the URL effectively accesses the image as the original requester. Access to the image may be lost if the presentation's sharing settings change. The mime type of the thumbnail image is the same as specified in the `GetPageThumbnailRequest`.

	Height int `json:"height,omitempty\"` // The positive height in pixels of the thumbnail image.

	Width int `json:"width,omitempty\"` // The positive width in pixels of the thumbnail image.

}

// Ungroups objects, such as groups.
type UngroupObjectsRequest struct {
	ObjectIds []string `json:"objectIds,omitempty\"` // The object IDs of the objects to ungroup. Only groups that are not inside other groups can be ungrouped. All the groups should be on the same page. The group itself is deleted. The visual sizes and positions of all the children are preserved.

}

// Unmerges cells in a Table.
type UnmergeTableCellsRequest struct {
	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	TableRange TableRange `json:"tableRange,omitempty\"` // The table range specifying which cells of the table to unmerge. All merged cells in this range will be unmerged, and cells that are already unmerged will not be affected. If the range has no merged cells, the request will do nothing. If there is text in any of the merged cells, the text will remain in the upper-left ("head") cell of the resulting block of unmerged cells.

}

// Update the properties of an Image.
type UpdateImagePropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `imageProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the image outline color, set `fields` to `"outline.outlineFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ImageProperties ImageProperties `json:"imageProperties,omitempty\"` // The image properties to update.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the image the updates are applied to.

}

// Updates the category of a line.
type UpdateLineCategoryRequest struct {
	LineCategory string `json:"lineCategory,omitempty\"` // The line category to update to. The exact line type is determined based on the category to update to and how it's routed to connect to other page elements.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the line the update is applied to. Only a line with a category indicating it is a "connector" can be updated. The line may be rerouted after updating its category.

}

// Updates the properties of a Line.
type UpdateLinePropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `lineProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the line solid fill color, set `fields` to `"lineFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	LineProperties LineProperties `json:"lineProperties,omitempty\"` // The line properties to update.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the line the update is applied to.

}

// Updates the alt text title and/or description of a page element.
type UpdatePageElementAltTextRequest struct {
	Description string `json:"description,omitempty\"` // The updated alt text description of the page element. If unset the existing value will be maintained. The description is exposed to screen readers and other accessibility interfaces. Only use human readable values related to the content of the page element.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the page element the updates are applied to.

	Title string `json:"title,omitempty\"` // The updated alt text title of the page element. If unset the existing value will be maintained. The title is exposed to screen readers and other accessibility interfaces. Only use human readable values related to the content of the page element.

}

// Updates the transform of a page element. Updating the transform of a group will change the absolute transform of the page elements in that group, which can change their visual appearance. See the documentation for PageElement.transform for more details.
type UpdatePageElementTransformRequest struct {
	ApplyMode string `json:"applyMode,omitempty\"` // The apply mode of the transform update.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the page element to update.

	Transform AffineTransform `json:"transform,omitempty\"` // The input transform matrix used to update the page element.

}

// Updates the Z-order of page elements. Z-order is an ordering of the elements on the page from back to front. The page element in the front may cover the elements that are behind it.
type UpdatePageElementsZOrderRequest struct {
	Operation string `json:"operation,omitempty\"` // The Z-order operation to apply on the page elements. When applying the operation on multiple page elements, the relative Z-orders within these page elements before the operation is maintained.

	PageElementObjectIds []string `json:"pageElementObjectIds,omitempty\"` // The object IDs of the page elements to update. All the page elements must be on the same page and must not be grouped.

}

// Updates the properties of a Page.
type UpdatePagePropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `pageProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the page background solid fill color, set `fields` to `"pageBackgroundFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the page the update is applied to.

	PageProperties PageProperties `json:"pageProperties,omitempty\"` // The page properties to update.

}

// Updates the styling for all of the paragraphs within a Shape or Table that overlap with the given text index range.
type UpdateParagraphStyleRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The location of the cell in the table containing the paragraph(s) to style. If `object_id` refers to a table, `cell_location` must have a value. Otherwise, it must not.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example, to update the paragraph alignment, set `fields` to `"alignment"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table with the text to be styled.

	Style ParagraphStyle `json:"style,omitempty\"` // The paragraph's style.

	TextRange RangeValue `json:"textRange,omitempty\"` // The range of text containing the paragraph(s) to style.

}

// Update the properties of a Shape.
type UpdateShapePropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `shapeProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the shape background solid fill color, set `fields` to `"shapeBackgroundFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape the updates are applied to.

	ShapeProperties ShapeProperties `json:"shapeProperties,omitempty\"` // The shape properties to update.

}

// Updates the properties of a Slide.
type UpdateSlidePropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root 'slideProperties' is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update whether a slide is skipped, set `fields` to `"isSkipped"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the slide the update is applied to.

	SlideProperties SlideProperties `json:"slideProperties,omitempty\"` // The slide properties to update.

}

// Updates the position of slides in the presentation.
type UpdateSlidesPositionRequest struct {
	InsertionIndex int `json:"insertionIndex,omitempty\"` // The index where the slides should be inserted, based on the slide arrangement before the move takes place. Must be between zero and the number of slides in the presentation, inclusive.

	SlideObjectIds []string `json:"slideObjectIds,omitempty\"` // The IDs of the slides in the presentation that should be moved. The slides in this list must be in existing presentation order, without duplicates.

}

// Updates the properties of the table borders in a Table.
type UpdateTableBorderPropertiesRequest struct {
	BorderPosition string `json:"borderPosition,omitempty\"` // The border position in the table range the updates should apply to. If a border position is not specified, the updates will apply to all borders in the table range.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableBorderProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the table border solid fill color, set `fields` to `"tableBorderFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	TableBorderProperties TableBorderProperties `json:"tableBorderProperties,omitempty\"` // The table border properties to update.

	TableRange TableRange `json:"tableRange,omitempty\"` // The table range representing the subset of the table to which the updates are applied. If a table range is not specified, the updates will apply to the entire table.

}

// Update the properties of a TableCell.
type UpdateTableCellPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableCellProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the table cell background solid fill color, set `fields` to `"tableCellBackgroundFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	TableCellProperties TableCellProperties `json:"tableCellProperties,omitempty\"` // The table cell properties to update.

	TableRange TableRange `json:"tableRange,omitempty\"` // The table range representing the subset of the table to which the updates are applied. If a table range is not specified, the updates will apply to the entire table.

}

// Updates the properties of a Table column.
type UpdateTableColumnPropertiesRequest struct {
	ColumnIndices []int `json:"columnIndices,omitempty\"` // The list of zero-based indices specifying which columns to update. If no indices are provided, all columns in the table will be updated.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableColumnProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the column width, set `fields` to `"column_width"`. If '"column_width"' is included in the field mask but the property is left unset, the column width will default to 406,400 EMU (32 points).

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	TableColumnProperties TableColumnProperties `json:"tableColumnProperties,omitempty\"` // The table column properties to update. If the value of `table_column_properties#column_width` in the request is less than 406,400 EMU (32 points), a 400 bad request error is returned.

}

// Updates the properties of a Table row.
type UpdateTableRowPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `tableRowProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the minimum row height, set `fields` to `"min_row_height"`. If '"min_row_height"' is included in the field mask but the property is left unset, the minimum row height will default to 0.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the table.

	RowIndices []int `json:"rowIndices,omitempty\"` // The list of zero-based indices specifying which rows to update. If no indices are provided, all rows in the table will be updated.

	TableRowProperties TableRowProperties `json:"tableRowProperties,omitempty\"` // The table row properties to update.

}

// Update the styling of text in a Shape or Table.
type UpdateTextStyleRequest struct {
	CellLocation TableCellLocation `json:"cellLocation,omitempty\"` // The location of the cell in the table containing the text to style. If `object_id` refers to a table, `cell_location` must have a value. Otherwise, it must not.

	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `style` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example, to update the text style to bold, set `fields` to `"bold"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the shape or table with the text to be styled.

	Style TextStyle `json:"style,omitempty\"` // The style(s) to set on the text. If the value for a particular style matches that of the parent, that style will be set to inherit. Certain text style changes may cause other changes meant to mirror the behavior of the Slides editor. See the documentation of TextStyle for more information.

	TextRange RangeValue `json:"textRange,omitempty\"` // The range of text to style. The range may be extended to include adjacent newlines. If the range fully contains a paragraph belonging to a list, the paragraph's bullet is also updated with the matching text style.

}

// Update the properties of a Video.
type UpdateVideoPropertiesRequest struct {
	Fields string `json:"fields,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `videoProperties` is implied and should not be specified. A single `"*"` can be used as short-hand for listing every field. For example to update the video outline color, set `fields` to `"outline.outlineFill.solidFill.color"`. To reset a property to its default value, include its field name in the field mask but leave the field itself unset.

	ObjectId string `json:"objectId,omitempty\"` // The object ID of the video the updates are applied to.

	VideoProperties VideoProperties `json:"videoProperties,omitempty\"` // The video properties to update.

}

// A PageElement kind representing a video.
type Video struct {
	Id string `json:"id,omitempty\"` // The video source's unique identifier for this video.

	Source string `json:"source,omitempty\"` // The video source.

	Url string `json:"url,omitempty\"` // An URL to a video. The URL is valid as long as the source video exists and sharing settings do not change.

	VideoProperties VideoProperties `json:"videoProperties,omitempty\"` // The properties of the video.

}

// The properties of the Video.
type VideoProperties struct {
	AutoPlay bool `json:"autoPlay,omitempty\"` // Whether to enable video autoplay when the page is displayed in present mode. Defaults to false.

	End int `json:"end,omitempty\"` // The time at which to end playback, measured in seconds from the beginning of the video. If set, the end time should be after the start time. If not set or if you set this to a value that exceeds the video's length, the video will be played until its end.

	Mute bool `json:"mute,omitempty\"` // Whether to mute the audio during video playback. Defaults to false.

	Outline Outline `json:"outline,omitempty\"` // The outline of the video. The default outline matches the defaults for new videos created in the Slides editor.

	Start int `json:"start,omitempty\"` // The time at which to start playback, measured in seconds from the beginning of the video. If set, the start time should be before the end time. If you set this to a value that exceeds the video's length in seconds, the video will be played from the last second. If not set, the video will be played from the beginning.

}

// Represents a font family and weight used to style a TextRun.
type WeightedFontFamily struct {
	FontFamily string `json:"fontFamily,omitempty\"` // The font family of the text. The font family can be any font from the Font menu in Slides or from [Google Fonts] (https://fonts.google.com/). If the font name is unrecognized, the text is rendered in `Arial`.

	Weight int `json:"weight,omitempty\"` // The rendered weight of the text. This field can have any value that is a multiple of `100` between `100` and `900`, inclusive. This range corresponds to the numerical values described in the CSS 2.1 Specification, [section 15.6](https://www.w3.org/TR/CSS21/fonts.html#font-boldness), with non-numerical values disallowed. Weights greater than or equal to `700` are considered bold, and weights less than `700`are not bold. The default value is `400` ("normal").

}

// A PageElement kind representing word art.
type WordArt struct {
	RenderedText string `json:"renderedText,omitempty\"` // The text rendered as word art.

}

// Provides control over how write requests are executed.
type WriteControl struct {
	RequiredRevisionId string `json:"requiredRevisionId,omitempty\"` // The revision ID of the presentation required for the write request. If specified and the required revision ID doesn't match the presentation's current revision ID, the request is not processed and returns a 400 bad request error. When a required revision ID is returned in a response, it indicates the revision ID of the document after the request was applied.

}
