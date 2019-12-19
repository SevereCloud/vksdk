package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestPhotosPhoto_ToAttachment(t *testing.T) {
	f := func(photo object.PhotosPhoto, want string) {
		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhoto.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhoto{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhoto{ID: 20, OwnerID: -10}, "photo-10_20")
}

func TestPhotosPhotoFull_ToAttachment(t *testing.T) {
	f := func(photo object.PhotosPhotoFull, want string) {
		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFull{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhotoFull{ID: 20, OwnerID: -10}, "photo-10_20")
}

func TestPhotosPhotoFullXtrRealOffset_ToAttachment(t *testing.T) {
	f := func(photo object.PhotosPhotoFullXtrRealOffset, want string) {
		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoFullXtrRealOffset.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFullXtrRealOffset{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhotoFullXtrRealOffset{ID: 20, OwnerID: -10}, "photo-10_20")
}

func TestPhotosPhotoXtrRealOffset_ToAttachment(t *testing.T) {
	f := func(photo object.PhotosPhotoXtrRealOffset, want string) {
		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoXtrRealOffset.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrRealOffset{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhotoXtrRealOffset{ID: 20, OwnerID: -10}, "photo-10_20")
}

func Test_PhotosPhotoAlbum_ToAttachment(t *testing.T) {
	f := func(album object.PhotosPhotoAlbum, want string) {
		if got := album.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoAlbum.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbum{ID: "10", OwnerID: 20}, "album20_10")
	f(object.PhotosPhotoAlbum{ID: "20", OwnerID: -10}, "album-10_20")
}

func TestPhotosPhotoAlbumFull_ToAttachment(t *testing.T) {
	f := func(album object.PhotosPhotoAlbumFull, want string) {
		if got := album.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoAlbumFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbumFull{ID: 10, OwnerID: 20}, "album20_10")
	f(object.PhotosPhotoAlbumFull{ID: 20, OwnerID: -10}, "album-10_20")
}

func TestPhotosPhotoXtrTagInfo_ToAttachment(t *testing.T) {
	f := func(photo object.PhotosPhotoXtrTagInfo, want string) {
		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoXtrTagInfo.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrTagInfo{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhotoXtrTagInfo{ID: 20, OwnerID: -10}, "photo-10_20")
}
