// This file was auto-generated by Fern from our API Definition.

package labor

type EmployeeWagesListRequest struct {
	// Filter the returned wages to only those that are associated with the specified employee.
	EmployeeID *string `json:"-" url:"employee_id,omitempty"`
	// The maximum number of `EmployeeWage` results to return per page. The number can range between
	// 1 and 200. The default is 200.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pointer to the next page of `EmployeeWage` results to fetch.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}
