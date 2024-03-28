package permission

type Permission struct {
	Users    int `json:"users"`
	Classes  int `json:"classes"`
	Profiles int `json:"profiles"`
	Lessons  int `json:"lessons"`
}

func NewPermission(users, classes, profiles, lessons int) *Permission {
	return &Permission{
		Users:    users,
		Classes:  classes,
		Profiles: profiles,
		Lessons:  lessons,
	}
}

func UpdatePermission(permission *Permission, users, classes, profiles, lesson *int) {
	if users != nil {
		permission.Users = *users
	}

	if classes != nil {
		permission.Classes = *classes
	}

	if profiles != nil {
		permission.Profiles = *profiles
	}

	if lesson != nil {
		permission.Lessons = *lesson
	}
}
