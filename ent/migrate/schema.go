// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApplicantInterestsColumns holds the columns for the "applicant_interests" table.
	ApplicantInterestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID},
		{Name: "category_id", Type: field.TypeUUID},
	}
	// ApplicantInterestsTable holds the schema information for the "applicant_interests" table.
	ApplicantInterestsTable = &schema.Table{
		Name:       "applicant_interests",
		Columns:    ApplicantInterestsColumns,
		PrimaryKey: []*schema.Column{ApplicantInterestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "applicant_interests_applicant_profiles_applicant_interests",
				Columns:    []*schema.Column{ApplicantInterestsColumns[4]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "applicant_interests_categories_applicant_interests",
				Columns:    []*schema.Column{ApplicantInterestsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ApplicantProfilesColumns holds the columns for the "applicant_profiles" table.
	ApplicantProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "gender", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "address1", Type: field.TypeString},
		{Name: "address2", Type: field.TypeString},
		{Name: "cv", Type: field.TypeString, Nullable: true},
		{Name: "internal_comments", Type: field.TypeString},
		{Name: "receive_emails", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "extra_skills", Type: field.TypeJSON, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// ApplicantProfilesTable holds the schema information for the "applicant_profiles" table.
	ApplicantProfilesTable = &schema.Table{
		Name:       "applicant_profiles",
		Columns:    ApplicantProfilesColumns,
		PrimaryKey: []*schema.Column{ApplicantProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "applicant_profiles_users_applicant_profiles",
				Columns:    []*schema.Column{ApplicantProfilesColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ApplicantProfileSkillsColumns holds the columns for the "applicant_profile_skills" table.
	ApplicantProfileSkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"Basico", "Intermedio", "Avanzado"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID},
		{Name: "skill_id", Type: field.TypeUUID, Nullable: true},
	}
	// ApplicantProfileSkillsTable holds the schema information for the "applicant_profile_skills" table.
	ApplicantProfileSkillsTable = &schema.Table{
		Name:       "applicant_profile_skills",
		Columns:    ApplicantProfileSkillsColumns,
		PrimaryKey: []*schema.Column{ApplicantProfileSkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "applicant_profile_skills_applicant_profiles_applicant_profile_skills",
				Columns:    []*schema.Column{ApplicantProfileSkillsColumns[5]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "applicant_profile_skills_skills_applicant_profile_skills",
				Columns:    []*schema.Column{ApplicantProfileSkillsColumns[6]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ApplicationsColumns holds the columns for the "applications" table.
	ApplicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID},
		{Name: "job_offer_id", Type: field.TypeUUID},
	}
	// ApplicationsTable holds the schema information for the "applications" table.
	ApplicationsTable = &schema.Table{
		Name:       "applications",
		Columns:    ApplicationsColumns,
		PrimaryKey: []*schema.Column{ApplicationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "applications_applicant_profiles_applications",
				Columns:    []*schema.Column{ApplicationsColumns[4]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "applications_job_offers_applications",
				Columns:    []*schema.Column{ApplicationsColumns[5]},
				RefColumns: []*schema.Column{JobOffersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_root", Type: field.TypeBool},
		{Name: "parent_id", Type: field.TypeUUID, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_categories_child_categories",
				Columns:    []*schema.Column{CategoriesColumns[7]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ContactUsMessagesColumns holds the columns for the "contact_us_messages" table.
	ContactUsMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// ContactUsMessagesTable holds the schema information for the "contact_us_messages" table.
	ContactUsMessagesTable = &schema.Table{
		Name:       "contact_us_messages",
		Columns:    ContactUsMessagesColumns,
		PrimaryKey: []*schema.Column{ContactUsMessagesColumns[0]},
	}
	// DegreeLevelsColumns holds the columns for the "degree_levels" table.
	DegreeLevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// DegreeLevelsTable holds the schema information for the "degree_levels" table.
	DegreeLevelsTable = &schema.Table{
		Name:       "degree_levels",
		Columns:    DegreeLevelsColumns,
		PrimaryKey: []*schema.Column{DegreeLevelsColumns[0]},
	}
	// EducationColumns holds the columns for the "education" table.
	EducationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "institution", Type: field.TypeString, Nullable: true},
		{Name: "date_obtained", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID, Nullable: true},
		{Name: "applicant_profile_educations", Type: field.TypeUUID, Nullable: true},
		{Name: "degree_level_id", Type: field.TypeUUID, Nullable: true},
	}
	// EducationTable holds the schema information for the "education" table.
	EducationTable = &schema.Table{
		Name:       "education",
		Columns:    EducationColumns,
		PrimaryKey: []*schema.Column{EducationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "education_applicant_profiles_educations",
				Columns:    []*schema.Column{EducationColumns[9]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "education_degree_levels_educations",
				Columns:    []*schema.Column{EducationColumns[10]},
				RefColumns: []*schema.Column{DegreeLevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InterviewsColumns holds the columns for the "interviews" table.
	InterviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "comment", Type: field.TypeString},
		{Name: "interview_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "application_id", Type: field.TypeUUID, Nullable: true},
		{Name: "interviewer_id", Type: field.TypeUUID, Nullable: true},
	}
	// InterviewsTable holds the schema information for the "interviews" table.
	InterviewsTable = &schema.Table{
		Name:       "interviews",
		Columns:    InterviewsColumns,
		PrimaryKey: []*schema.Column{InterviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "interviews_applications_interviews",
				Columns:    []*schema.Column{InterviewsColumns[7]},
				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "interviews_users_interviews",
				Columns:    []*schema.Column{InterviewsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobOffersColumns holds the columns for the "job_offers" table.
	JobOffersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "reference", Type: field.TypeInt32},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "address1", Type: field.TypeString, Nullable: true},
		{Name: "address2", Type: field.TypeString, Nullable: true},
		{Name: "department", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "working_hours", Type: field.TypeString},
		{Name: "salary", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "is_featured", Type: field.TypeBool, Nullable: true},
		{Name: "has_been_emailed", Type: field.TypeBool, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "status_id", Type: field.TypeUUID, Nullable: true},
	}
	// JobOffersTable holds the schema information for the "job_offers" table.
	JobOffersTable = &schema.Table{
		Name:       "job_offers",
		Columns:    JobOffersColumns,
		PrimaryKey: []*schema.Column{JobOffersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_offers_status_job_offers",
				Columns:    []*schema.Column{JobOffersColumns[17]},
				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobOfferCategoriesColumns holds the columns for the "job_offer_categories" table.
	JobOfferCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "category_id", Type: field.TypeUUID},
		{Name: "job_offer_id", Type: field.TypeUUID},
	}
	// JobOfferCategoriesTable holds the schema information for the "job_offer_categories" table.
	JobOfferCategoriesTable = &schema.Table{
		Name:       "job_offer_categories",
		Columns:    JobOfferCategoriesColumns,
		PrimaryKey: []*schema.Column{JobOfferCategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_offer_categories_categories_job_offer_categories",
				Columns:    []*schema.Column{JobOfferCategoriesColumns[4]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "job_offer_categories_job_offers_job_offer_categories",
				Columns:    []*schema.Column{JobOfferCategoriesColumns[5]},
				RefColumns: []*schema.Column{JobOffersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LanguagesColumns holds the columns for the "languages" table.
	LanguagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true, SchemaType: map[string]string{"postgres": "serial"}},
		{Name: "name", Type: field.TypeEnum, Enums: []string{"Español", "Inglés", "Portugués", "Francés", "Alemán", "Italiano", "Chino", "Japonés", "Ruso", "Árabe", "Otro"}},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"Básico", "Intermedio", "Avanzado"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID},
	}
	// LanguagesTable holds the schema information for the "languages" table.
	LanguagesTable = &schema.Table{
		Name:       "languages",
		Columns:    LanguagesColumns,
		PrimaryKey: []*schema.Column{LanguagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "languages_applicant_profiles_languages",
				Columns:    []*schema.Column{LanguagesColumns[7]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "is_highlighted", Type: field.TypeBool},
		{Name: "is_published", Type: field.TypeBool},
		{Name: "published_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "preview_image", Type: field.TypeString, Nullable: true},
		{Name: "author_id", Type: field.TypeUUID, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostCategoriesColumns holds the columns for the "post_categories" table.
	PostCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "post_post_categories", Type: field.TypeUUID, Nullable: true},
	}
	// PostCategoriesTable holds the schema information for the "post_categories" table.
	PostCategoriesTable = &schema.Table{
		Name:       "post_categories",
		Columns:    PostCategoriesColumns,
		PrimaryKey: []*schema.Column{PostCategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_categories_posts_post_categories",
				Columns:    []*schema.Column{PostCategoriesColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:       "status",
		Columns:    StatusColumns,
		PrimaryKey: []*schema.Column{StatusColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "APPLICANT", "RECRUITER"}},
		{Name: "reset_password_token", Type: field.TypeString, Nullable: true},
		{Name: "reset_password_expires", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "full_name", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "external_id", Type: field.TypeInt32, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WorkExperienceColumns holds the columns for the "work_experience" table.
	WorkExperienceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "company", Type: field.TypeString, Nullable: true},
		{Name: "position", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "internal_comments", Type: field.TypeString, Nullable: true},
		{Name: "applicant_profile_id", Type: field.TypeUUID, Nullable: true},
	}
	// WorkExperienceTable holds the schema information for the "work_experience" table.
	WorkExperienceTable = &schema.Table{
		Name:       "work_experience",
		Columns:    WorkExperienceColumns,
		PrimaryKey: []*schema.Column{WorkExperienceColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "work_experience_applicant_profiles_work_experiences",
				Columns:    []*schema.Column{WorkExperienceColumns[10]},
				RefColumns: []*schema.Column{ApplicantProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostCategoryPostCategoriesColumns holds the columns for the "post_category_post_categories" table.
	PostCategoryPostCategoriesColumns = []*schema.Column{
		{Name: "post_category_id", Type: field.TypeUUID},
		{Name: "post_category_id", Type: field.TypeUUID},
	}
	// PostCategoryPostCategoriesTable holds the schema information for the "post_category_post_categories" table.
	PostCategoryPostCategoriesTable = &schema.Table{
		Name:       "post_category_post_categories",
		Columns:    PostCategoryPostCategoriesColumns,
		PrimaryKey: []*schema.Column{PostCategoryPostCategoriesColumns[0], PostCategoryPostCategoriesColumns[1], PostCategoryPostCategoriesColumns[0], PostCategoryPostCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_category_post_categories_post_category_id",
				Columns:    []*schema.Column{PostCategoryPostCategoriesColumns[0], PostCategoryPostCategoriesColumns[1]},
				RefColumns: []*schema.Column{PostCategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_category_post_categories_post_category_id",
				Columns:    []*schema.Column{PostCategoryPostCategoriesColumns[0], PostCategoryPostCategoriesColumns[1]},
				RefColumns: []*schema.Column{PostCategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicantInterestsTable,
		ApplicantProfilesTable,
		ApplicantProfileSkillsTable,
		ApplicationsTable,
		CategoriesTable,
		ContactUsMessagesTable,
		DegreeLevelsTable,
		EducationTable,
		InterviewsTable,
		JobOffersTable,
		JobOfferCategoriesTable,
		LanguagesTable,
		PostsTable,
		PostCategoriesTable,
		SkillsTable,
		StatusTable,
		UsersTable,
		WorkExperienceTable,
		PostCategoryPostCategoriesTable,
	}
)

func init() {
	ApplicantInterestsTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	ApplicantInterestsTable.ForeignKeys[1].RefTable = CategoriesTable
	ApplicantProfilesTable.ForeignKeys[0].RefTable = UsersTable
	ApplicantProfileSkillsTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	ApplicantProfileSkillsTable.ForeignKeys[1].RefTable = SkillsTable
	ApplicationsTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	ApplicationsTable.ForeignKeys[1].RefTable = JobOffersTable
	CategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	EducationTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	EducationTable.ForeignKeys[1].RefTable = DegreeLevelsTable
	EducationTable.Annotation = &entsql.Annotation{
		Table: "education",
	}
	InterviewsTable.ForeignKeys[0].RefTable = ApplicationsTable
	InterviewsTable.ForeignKeys[1].RefTable = UsersTable
	JobOffersTable.ForeignKeys[0].RefTable = StatusTable
	JobOfferCategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	JobOfferCategoriesTable.ForeignKeys[1].RefTable = JobOffersTable
	LanguagesTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	PostCategoriesTable.ForeignKeys[0].RefTable = PostsTable
	StatusTable.Annotation = &entsql.Annotation{
		Table: "status",
	}
	WorkExperienceTable.ForeignKeys[0].RefTable = ApplicantProfilesTable
	WorkExperienceTable.Annotation = &entsql.Annotation{
		Table: "work_experience",
	}
	PostCategoryPostCategoriesTable.ForeignKeys[0].RefTable = PostCategoriesTable
	PostCategoryPostCategoriesTable.ForeignKeys[1].RefTable = PostCategoriesTable
}
