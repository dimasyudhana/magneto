package main

import (
	"fmt"
	"time"
)

type Dashboard struct {
	Header    string
	Message   map[string]any
	WadocList map[string]map[bool][]UsersLeaveWADoc
}

type UsersLeaveWADoc struct {
	WADoc               WADoc
	Leave               Leave
	UserAvailability    []map[string]bool
	PostDashboard       map[int]bool
	PostHistoryApproval HistoryApproval
}

type WADoc struct {
	ID                   string
	ModuleID             int
	ModuleRefID          string
	ApprovalTypeID       int
	Section              string
	CurrentState         string
	LeavePolicy          int
	Username             string
	ReplacementUsers     []string
	IsBackToLastApproval bool
	ReplacementUserIdx   int
}

type Leave struct {
	LeaveStartDate      time.Time
	LeaveEndDate        time.Time
	OriginalUsername    string
	ReplacementUsername string
}

type HistoryApproval struct {
	User           string
	ApprovalAction string
	Note           string
}

func DashboardCompute() *Dashboard {
	return &Dashboard{
		Header:    "",
		Message:   make(map[string]any),
		WadocList: make(map[string]map[bool][]UsersLeaveWADoc),
	}
}

func (notify *Dashboard) DashboardUserLeaveWADoc(userLeave UsersLeaveWADoc, key string, isApproved bool) {
	// Check if the key exists
	if _, exists := notify.WadocList[key]; !exists {
		notify.WadocList[key] = make(map[bool][]UsersLeaveWADoc)
	}

	// Append UsersLeaveWADoc to the corresponding boolean key
	notify.WadocList[key][isApproved] = append(notify.WadocList[key][isApproved], userLeave)

}

func main() {
	// Create a new dashboard
	dashboard := DashboardCompute()

	// Sample data for WADoc and Leave
	wadoc1 := WADoc{
		ID:                   "WADOC-001",
		ModuleID:             101,
		ModuleRefID:          "REF-001",
		ApprovalTypeID:       1,
		Section:              "Finance",
		CurrentState:         "Pending",
		LeavePolicy:          2,
		Username:             "john.doe",
		ReplacementUsers:     []string{"jane.smith"},
		IsBackToLastApproval: false,
		ReplacementUserIdx:   0,
	}

	leave1 := Leave{
		LeaveStartDate:      time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		LeaveEndDate:        time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC),
		OriginalUsername:    "john.doe",
		ReplacementUsername: "jane.smith",
	}

	history1 := HistoryApproval{
		User:           "john.doe",
		ApprovalAction: "Forwarded",
		Note:           "User on leave, approval forwarded.",
	}

	userLeave1 := UsersLeaveWADoc{
		WADoc:               wadoc1,
		Leave:               leave1,
		UserAvailability:    []map[string]bool{{"john.doe": false}, {"jane.smith": true}},
		PostDashboard:       map[int]bool{1: true},
		PostHistoryApproval: history1,
	}

	// Add isApproved = true
	dashboard.DashboardUserLeaveWADoc(userLeave1, "AI870", true)

	// Add isApproved = false
	wadoc2 := WADoc{
		ID:                   "WADOC-002",
		ModuleID:             102,
		ModuleRefID:          "REF-002",
		ApprovalTypeID:       2,
		Section:              "HR",
		CurrentState:         "Rejected",
		LeavePolicy:          3,
		Username:             "jane.smith",
		ReplacementUsers:     []string{"john.doe"},
		IsBackToLastApproval: true,
		ReplacementUserIdx:   1,
	}

	leave2 := Leave{
		LeaveStartDate:      time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
		LeaveEndDate:        time.Date(2023, 11, 20, 0, 0, 0, 0, time.UTC),
		OriginalUsername:    "jane.smith",
		ReplacementUsername: "john.doe",
	}

	history2 := HistoryApproval{
		User:           "jane.smith",
		ApprovalAction: "Skipped",
		Note:           "User unavailable, approval skipped.",
	}

	userLeave2 := UsersLeaveWADoc{
		WADoc:               wadoc2,
		Leave:               leave2,
		UserAvailability:    []map[string]bool{{"jane.smith": false}, {"john.doe": true}},
		PostDashboard:       map[int]bool{2: false},
		PostHistoryApproval: history2,
	}

	dashboard.DashboardUserLeaveWADoc(userLeave2, "AI870", false)

	// Sample data for WADoc and Leave
	wadoc3 := WADoc{
		ID:                   "WADOC-001",
		ModuleID:             101,
		ModuleRefID:          "REF-001",
		ApprovalTypeID:       1,
		Section:              "Finance",
		CurrentState:         "Pending",
		LeavePolicy:          2,
		Username:             "john.doe",
		ReplacementUsers:     []string{"jane.smith"},
		IsBackToLastApproval: false,
		ReplacementUserIdx:   0,
	}

	leave3 := Leave{
		LeaveStartDate:      time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		LeaveEndDate:        time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC),
		OriginalUsername:    "john.doe",
		ReplacementUsername: "jane.smith",
	}

	history3 := HistoryApproval{
		User:           "john.doe",
		ApprovalAction: "Forwarded",
		Note:           "User on leave, approval forwarded.",
	}

	userLeave3 := UsersLeaveWADoc{
		WADoc:               wadoc3,
		Leave:               leave3,
		UserAvailability:    []map[string]bool{{"john.doe": false}, {"jane.smith": true}},
		PostDashboard:       map[int]bool{3: true},
		PostHistoryApproval: history3,
	}

	// Add isApproved = true
	dashboard.DashboardUserLeaveWADoc(userLeave3, "AI870", true)

	// Print the dashboard content
	for key, subMap := range dashboard.WadocList {
		fmt.Printf("Key: %s\n", key)
		for isApproved, docs := range subMap {
			fmt.Printf("  IsApproved: %v\n", isApproved)
			fmt.Printf("  [\n")
			for _, doc := range docs {
				fmt.Printf("    {\n")
				fmt.Printf("      WADoc: %+v,\n", doc.WADoc)
				fmt.Printf("      Leave: %+v,\n", doc.Leave)
				fmt.Printf("      History: %+v\n", doc.PostHistoryApproval)
				fmt.Printf("    },\n")
			}
			fmt.Printf("  ]\n")
		}
	}

}
