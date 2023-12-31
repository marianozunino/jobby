// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
)

func (u *ApplicantInterestUpdate) SetOrClearCreatedAt(value *time.Time) *ApplicantInterestUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicantInterestUpdate) SetOrClearUpdatedAt(value *time.Time) *ApplicantInterestUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicantInterestUpdate) SetOrClearDeletedAt(value *time.Time) *ApplicantInterestUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicantInterestUpdateOne) SetOrClearCreatedAt(value *time.Time) *ApplicantInterestUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicantInterestUpdateOne) SetOrClearUpdatedAt(value *time.Time) *ApplicantInterestUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicantInterestUpdateOne) SetOrClearDeletedAt(value *time.Time) *ApplicantInterestUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearCv(value *string) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearCv()
	}
	return u.SetCv(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearCreatedAt(value *time.Time) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearUpdatedAt(value *time.Time) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearDeletedAt(value *time.Time) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearUserID(value *uuid.UUID) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearUserID()
	}
	return u.SetUserID(*value)
}

func (u *ApplicantProfileUpdate) SetOrClearExtraSkills(value *struct{}) *ApplicantProfileUpdate {
	if value == nil {
		return u.ClearExtraSkills()
	}
	return u.SetExtraSkills(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearCv(value *string) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearCv()
	}
	return u.SetCv(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearCreatedAt(value *time.Time) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearUpdatedAt(value *time.Time) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearDeletedAt(value *time.Time) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearUserID(value *uuid.UUID) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearUserID()
	}
	return u.SetUserID(*value)
}

func (u *ApplicantProfileUpdateOne) SetOrClearExtraSkills(value *struct{}) *ApplicantProfileUpdateOne {
	if value == nil {
		return u.ClearExtraSkills()
	}
	return u.SetExtraSkills(*value)
}

func (u *ApplicantProfileSkillUpdate) SetOrClearSkillID(value *uuid.UUID) *ApplicantProfileSkillUpdate {
	if value == nil {
		return u.ClearSkillID()
	}
	return u.SetSkillID(*value)
}

func (u *ApplicantProfileSkillUpdate) SetOrClearDeletedAt(value *time.Time) *ApplicantProfileSkillUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicantProfileSkillUpdateOne) SetOrClearSkillID(value *uuid.UUID) *ApplicantProfileSkillUpdateOne {
	if value == nil {
		return u.ClearSkillID()
	}
	return u.SetSkillID(*value)
}

func (u *ApplicantProfileSkillUpdateOne) SetOrClearDeletedAt(value *time.Time) *ApplicantProfileSkillUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicationUpdate) SetOrClearCreatedAt(value *time.Time) *ApplicationUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicationUpdate) SetOrClearUpdatedAt(value *time.Time) *ApplicationUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicationUpdate) SetOrClearDeletedAt(value *time.Time) *ApplicationUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ApplicationUpdateOne) SetOrClearCreatedAt(value *time.Time) *ApplicationUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ApplicationUpdateOne) SetOrClearUpdatedAt(value *time.Time) *ApplicationUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ApplicationUpdateOne) SetOrClearDeletedAt(value *time.Time) *ApplicationUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *CategoryUpdate) SetOrClearParentID(value *uuid.UUID) *CategoryUpdate {
	if value == nil {
		return u.ClearParentID()
	}
	return u.SetParentID(*value)
}

func (u *CategoryUpdate) SetOrClearCreatedAt(value *time.Time) *CategoryUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *CategoryUpdate) SetOrClearUpdatedAt(value *time.Time) *CategoryUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *CategoryUpdate) SetOrClearDeletedAt(value *time.Time) *CategoryUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *CategoryUpdateOne) SetOrClearParentID(value *uuid.UUID) *CategoryUpdateOne {
	if value == nil {
		return u.ClearParentID()
	}
	return u.SetParentID(*value)
}

func (u *CategoryUpdateOne) SetOrClearCreatedAt(value *time.Time) *CategoryUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *CategoryUpdateOne) SetOrClearUpdatedAt(value *time.Time) *CategoryUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *CategoryUpdateOne) SetOrClearDeletedAt(value *time.Time) *CategoryUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ContactUsMessageUpdate) SetOrClearPhone(value *string) *ContactUsMessageUpdate {
	if value == nil {
		return u.ClearPhone()
	}
	return u.SetPhone(*value)
}

func (u *ContactUsMessageUpdate) SetOrClearCreatedAt(value *time.Time) *ContactUsMessageUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ContactUsMessageUpdate) SetOrClearUpdatedAt(value *time.Time) *ContactUsMessageUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ContactUsMessageUpdate) SetOrClearDeletedAt(value *time.Time) *ContactUsMessageUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *ContactUsMessageUpdateOne) SetOrClearPhone(value *string) *ContactUsMessageUpdateOne {
	if value == nil {
		return u.ClearPhone()
	}
	return u.SetPhone(*value)
}

func (u *ContactUsMessageUpdateOne) SetOrClearCreatedAt(value *time.Time) *ContactUsMessageUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *ContactUsMessageUpdateOne) SetOrClearUpdatedAt(value *time.Time) *ContactUsMessageUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *ContactUsMessageUpdateOne) SetOrClearDeletedAt(value *time.Time) *ContactUsMessageUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *DegreeLevelUpdate) SetOrClearCreatedAt(value *time.Time) *DegreeLevelUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *DegreeLevelUpdate) SetOrClearUpdatedAt(value *time.Time) *DegreeLevelUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *DegreeLevelUpdate) SetOrClearDeletedAt(value *time.Time) *DegreeLevelUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *DegreeLevelUpdateOne) SetOrClearCreatedAt(value *time.Time) *DegreeLevelUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *DegreeLevelUpdateOne) SetOrClearUpdatedAt(value *time.Time) *DegreeLevelUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *DegreeLevelUpdateOne) SetOrClearDeletedAt(value *time.Time) *DegreeLevelUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *EducationUpdate) SetOrClearInstitution(value *string) *EducationUpdate {
	if value == nil {
		return u.ClearInstitution()
	}
	return u.SetInstitution(*value)
}

func (u *EducationUpdate) SetOrClearDateObtained(value *time.Time) *EducationUpdate {
	if value == nil {
		return u.ClearDateObtained()
	}
	return u.SetDateObtained(*value)
}

func (u *EducationUpdate) SetOrClearDegreeLevelID(value *uuid.UUID) *EducationUpdate {
	if value == nil {
		return u.ClearDegreeLevelID()
	}
	return u.SetDegreeLevelID(*value)
}

func (u *EducationUpdate) SetOrClearCreatedAt(value *time.Time) *EducationUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *EducationUpdate) SetOrClearUpdatedAt(value *time.Time) *EducationUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *EducationUpdate) SetOrClearDeletedAt(value *time.Time) *EducationUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *EducationUpdate) SetOrClearComments(value *string) *EducationUpdate {
	if value == nil {
		return u.ClearComments()
	}
	return u.SetComments(*value)
}

func (u *EducationUpdate) SetOrClearApplicantProfileID(value *uuid.UUID) *EducationUpdate {
	if value == nil {
		return u.ClearApplicantProfileID()
	}
	return u.SetApplicantProfileID(*value)
}

func (u *EducationUpdateOne) SetOrClearInstitution(value *string) *EducationUpdateOne {
	if value == nil {
		return u.ClearInstitution()
	}
	return u.SetInstitution(*value)
}

func (u *EducationUpdateOne) SetOrClearDateObtained(value *time.Time) *EducationUpdateOne {
	if value == nil {
		return u.ClearDateObtained()
	}
	return u.SetDateObtained(*value)
}

func (u *EducationUpdateOne) SetOrClearDegreeLevelID(value *uuid.UUID) *EducationUpdateOne {
	if value == nil {
		return u.ClearDegreeLevelID()
	}
	return u.SetDegreeLevelID(*value)
}

func (u *EducationUpdateOne) SetOrClearCreatedAt(value *time.Time) *EducationUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *EducationUpdateOne) SetOrClearUpdatedAt(value *time.Time) *EducationUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *EducationUpdateOne) SetOrClearDeletedAt(value *time.Time) *EducationUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *EducationUpdateOne) SetOrClearComments(value *string) *EducationUpdateOne {
	if value == nil {
		return u.ClearComments()
	}
	return u.SetComments(*value)
}

func (u *EducationUpdateOne) SetOrClearApplicantProfileID(value *uuid.UUID) *EducationUpdateOne {
	if value == nil {
		return u.ClearApplicantProfileID()
	}
	return u.SetApplicantProfileID(*value)
}

func (u *InterviewUpdate) SetOrClearApplicationID(value *uuid.UUID) *InterviewUpdate {
	if value == nil {
		return u.ClearApplicationID()
	}
	return u.SetApplicationID(*value)
}

func (u *InterviewUpdate) SetOrClearInterviewerID(value *uuid.UUID) *InterviewUpdate {
	if value == nil {
		return u.ClearInterviewerID()
	}
	return u.SetInterviewerID(*value)
}

func (u *InterviewUpdate) SetOrClearDeletedAt(value *time.Time) *InterviewUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *InterviewUpdateOne) SetOrClearApplicationID(value *uuid.UUID) *InterviewUpdateOne {
	if value == nil {
		return u.ClearApplicationID()
	}
	return u.SetApplicationID(*value)
}

func (u *InterviewUpdateOne) SetOrClearInterviewerID(value *uuid.UUID) *InterviewUpdateOne {
	if value == nil {
		return u.ClearInterviewerID()
	}
	return u.SetInterviewerID(*value)
}

func (u *InterviewUpdateOne) SetOrClearDeletedAt(value *time.Time) *InterviewUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *JobOfferUpdate) SetOrClearAddress1(value *string) *JobOfferUpdate {
	if value == nil {
		return u.ClearAddress1()
	}
	return u.SetAddress1(*value)
}

func (u *JobOfferUpdate) SetOrClearAddress2(value *string) *JobOfferUpdate {
	if value == nil {
		return u.ClearAddress2()
	}
	return u.SetAddress2(*value)
}

func (u *JobOfferUpdate) SetOrClearIsFeatured(value *bool) *JobOfferUpdate {
	if value == nil {
		return u.ClearIsFeatured()
	}
	return u.SetIsFeatured(*value)
}

func (u *JobOfferUpdate) SetOrClearHasBeenEmailed(value *bool) *JobOfferUpdate {
	if value == nil {
		return u.ClearHasBeenEmailed()
	}
	return u.SetHasBeenEmailed(*value)
}

func (u *JobOfferUpdate) SetOrClearStatusID(value *uuid.UUID) *JobOfferUpdate {
	if value == nil {
		return u.ClearStatusID()
	}
	return u.SetStatusID(*value)
}

func (u *JobOfferUpdate) SetOrClearCreatedAt(value *time.Time) *JobOfferUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *JobOfferUpdate) SetOrClearUpdatedAt(value *time.Time) *JobOfferUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *JobOfferUpdate) SetOrClearDeletedAt(value *time.Time) *JobOfferUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *JobOfferUpdateOne) SetOrClearAddress1(value *string) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearAddress1()
	}
	return u.SetAddress1(*value)
}

func (u *JobOfferUpdateOne) SetOrClearAddress2(value *string) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearAddress2()
	}
	return u.SetAddress2(*value)
}

func (u *JobOfferUpdateOne) SetOrClearIsFeatured(value *bool) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearIsFeatured()
	}
	return u.SetIsFeatured(*value)
}

func (u *JobOfferUpdateOne) SetOrClearHasBeenEmailed(value *bool) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearHasBeenEmailed()
	}
	return u.SetHasBeenEmailed(*value)
}

func (u *JobOfferUpdateOne) SetOrClearStatusID(value *uuid.UUID) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearStatusID()
	}
	return u.SetStatusID(*value)
}

func (u *JobOfferUpdateOne) SetOrClearCreatedAt(value *time.Time) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *JobOfferUpdateOne) SetOrClearUpdatedAt(value *time.Time) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *JobOfferUpdateOne) SetOrClearDeletedAt(value *time.Time) *JobOfferUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *JobOfferCategoryUpdate) SetOrClearCreatedAt(value *time.Time) *JobOfferCategoryUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *JobOfferCategoryUpdate) SetOrClearUpdatedAt(value *time.Time) *JobOfferCategoryUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *JobOfferCategoryUpdate) SetOrClearDeletedAt(value *time.Time) *JobOfferCategoryUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *JobOfferCategoryUpdateOne) SetOrClearCreatedAt(value *time.Time) *JobOfferCategoryUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *JobOfferCategoryUpdateOne) SetOrClearUpdatedAt(value *time.Time) *JobOfferCategoryUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *JobOfferCategoryUpdateOne) SetOrClearDeletedAt(value *time.Time) *JobOfferCategoryUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *LanguageUpdate) SetOrClearComments(value *string) *LanguageUpdate {
	if value == nil {
		return u.ClearComments()
	}
	return u.SetComments(*value)
}

func (u *LanguageUpdate) SetOrClearDeletedAt(value *time.Time) *LanguageUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *LanguageUpdateOne) SetOrClearComments(value *string) *LanguageUpdateOne {
	if value == nil {
		return u.ClearComments()
	}
	return u.SetComments(*value)
}

func (u *LanguageUpdateOne) SetOrClearDeletedAt(value *time.Time) *LanguageUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *PostUpdate) SetOrClearPublishedAt(value *time.Time) *PostUpdate {
	if value == nil {
		return u.ClearPublishedAt()
	}
	return u.SetPublishedAt(*value)
}

func (u *PostUpdate) SetOrClearPreviewImage(value *string) *PostUpdate {
	if value == nil {
		return u.ClearPreviewImage()
	}
	return u.SetPreviewImage(*value)
}

func (u *PostUpdate) SetOrClearAuthorID(value *uuid.UUID) *PostUpdate {
	if value == nil {
		return u.ClearAuthorID()
	}
	return u.SetAuthorID(*value)
}

func (u *PostUpdate) SetOrClearDeletedAt(value *time.Time) *PostUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *PostUpdateOne) SetOrClearPublishedAt(value *time.Time) *PostUpdateOne {
	if value == nil {
		return u.ClearPublishedAt()
	}
	return u.SetPublishedAt(*value)
}

func (u *PostUpdateOne) SetOrClearPreviewImage(value *string) *PostUpdateOne {
	if value == nil {
		return u.ClearPreviewImage()
	}
	return u.SetPreviewImage(*value)
}

func (u *PostUpdateOne) SetOrClearAuthorID(value *uuid.UUID) *PostUpdateOne {
	if value == nil {
		return u.ClearAuthorID()
	}
	return u.SetAuthorID(*value)
}

func (u *PostUpdateOne) SetOrClearDeletedAt(value *time.Time) *PostUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *PostCategoryUpdate) SetOrClearDeletedAt(value *time.Time) *PostCategoryUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *PostCategoryUpdateOne) SetOrClearDeletedAt(value *time.Time) *PostCategoryUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *SkillUpdate) SetOrClearDeletedAt(value *time.Time) *SkillUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *SkillUpdateOne) SetOrClearDeletedAt(value *time.Time) *SkillUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *StatusUpdate) SetOrClearCreatedAt(value *time.Time) *StatusUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *StatusUpdate) SetOrClearUpdatedAt(value *time.Time) *StatusUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *StatusUpdate) SetOrClearDeletedAt(value *time.Time) *StatusUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *StatusUpdateOne) SetOrClearCreatedAt(value *time.Time) *StatusUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *StatusUpdateOne) SetOrClearUpdatedAt(value *time.Time) *StatusUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *StatusUpdateOne) SetOrClearDeletedAt(value *time.Time) *StatusUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *UserUpdate) SetOrClearResetPasswordToken(value *string) *UserUpdate {
	if value == nil {
		return u.ClearResetPasswordToken()
	}
	return u.SetResetPasswordToken(*value)
}

func (u *UserUpdate) SetOrClearResetPasswordExpires(value *time.Time) *UserUpdate {
	if value == nil {
		return u.ClearResetPasswordExpires()
	}
	return u.SetResetPasswordExpires(*value)
}

func (u *UserUpdate) SetOrClearCreatedAt(value *time.Time) *UserUpdate {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *UserUpdate) SetOrClearUpdatedAt(value *time.Time) *UserUpdate {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *UserUpdate) SetOrClearDeletedAt(value *time.Time) *UserUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *UserUpdate) SetOrClearExternalID(value *int32) *UserUpdate {
	if value == nil {
		return u.ClearExternalID()
	}
	return u.SetExternalID(*value)
}

func (u *UserUpdateOne) SetOrClearResetPasswordToken(value *string) *UserUpdateOne {
	if value == nil {
		return u.ClearResetPasswordToken()
	}
	return u.SetResetPasswordToken(*value)
}

func (u *UserUpdateOne) SetOrClearResetPasswordExpires(value *time.Time) *UserUpdateOne {
	if value == nil {
		return u.ClearResetPasswordExpires()
	}
	return u.SetResetPasswordExpires(*value)
}

func (u *UserUpdateOne) SetOrClearCreatedAt(value *time.Time) *UserUpdateOne {
	if value == nil {
		return u.ClearCreatedAt()
	}
	return u.SetCreatedAt(*value)
}

func (u *UserUpdateOne) SetOrClearUpdatedAt(value *time.Time) *UserUpdateOne {
	if value == nil {
		return u.ClearUpdatedAt()
	}
	return u.SetUpdatedAt(*value)
}

func (u *UserUpdateOne) SetOrClearDeletedAt(value *time.Time) *UserUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *UserUpdateOne) SetOrClearExternalID(value *int32) *UserUpdateOne {
	if value == nil {
		return u.ClearExternalID()
	}
	return u.SetExternalID(*value)
}

func (u *WorkExperienceUpdate) SetOrClearCompany(value *string) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearCompany()
	}
	return u.SetCompany(*value)
}

func (u *WorkExperienceUpdate) SetOrClearPosition(value *string) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearPosition()
	}
	return u.SetPosition(*value)
}

func (u *WorkExperienceUpdate) SetOrClearDescription(value *string) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearDescription()
	}
	return u.SetDescription(*value)
}

func (u *WorkExperienceUpdate) SetOrClearStartDate(value *time.Time) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearStartDate()
	}
	return u.SetStartDate(*value)
}

func (u *WorkExperienceUpdate) SetOrClearEndDate(value *time.Time) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearEndDate()
	}
	return u.SetEndDate(*value)
}

func (u *WorkExperienceUpdate) SetOrClearApplicantProfileID(value *uuid.UUID) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearApplicantProfileID()
	}
	return u.SetApplicantProfileID(*value)
}

func (u *WorkExperienceUpdate) SetOrClearDeletedAt(value *time.Time) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *WorkExperienceUpdate) SetOrClearInternalComments(value *string) *WorkExperienceUpdate {
	if value == nil {
		return u.ClearInternalComments()
	}
	return u.SetInternalComments(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearCompany(value *string) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearCompany()
	}
	return u.SetCompany(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearPosition(value *string) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearPosition()
	}
	return u.SetPosition(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearDescription(value *string) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearDescription()
	}
	return u.SetDescription(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearStartDate(value *time.Time) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearStartDate()
	}
	return u.SetStartDate(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearEndDate(value *time.Time) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearEndDate()
	}
	return u.SetEndDate(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearApplicantProfileID(value *uuid.UUID) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearApplicantProfileID()
	}
	return u.SetApplicantProfileID(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearDeletedAt(value *time.Time) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearDeletedAt()
	}
	return u.SetDeletedAt(*value)
}

func (u *WorkExperienceUpdateOne) SetOrClearInternalComments(value *string) *WorkExperienceUpdateOne {
	if value == nil {
		return u.ClearInternalComments()
	}
	return u.SetInternalComments(*value)
}
