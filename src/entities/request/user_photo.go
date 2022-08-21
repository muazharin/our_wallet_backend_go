package request

import "mime/multipart"

type UserPhotoReq struct {
	UserId    int64                 `form:"user_id" bson:"user_id" binding:"required"`
	UserPhoto *multipart.FileHeader `form:"user_photo" bson:"user_photo" binding:"required"`
}
