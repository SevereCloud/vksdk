package object_test

import (
	"reflect"
	"testing"

	"github.com/Derad6709/vksdk/v2/object"
)

func TestPhotosPhoto_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhoto, want string) {
		t.Helper()

		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhoto.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhoto{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhoto{ID: 20, OwnerID: -10}, "photo-10_20")
}

func TestPhotosPhotoFull_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFull, want string) {
		t.Helper()

		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFull{ID: 10, OwnerID: 20}, "photo20_10")
	f(object.PhotosPhotoFull{ID: 20, OwnerID: -10}, "photo-10_20")
}

func TestPhotosPhotoFullXtrRealOffset_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFullXtrRealOffset, want string) {
		t.Helper()

		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoFullXtrRealOffset.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFullXtrRealOffset{PhotosPhotoFull: object.PhotosPhotoFull{ID: 10, OwnerID: 20}}, "photo20_10")
	f(object.PhotosPhotoFullXtrRealOffset{PhotosPhotoFull: object.PhotosPhotoFull{ID: 20, OwnerID: -10}}, "photo-10_20")
}

func TestPhotosPhotoXtrRealOffset_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrRealOffset, want string) {
		t.Helper()

		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoXtrRealOffset.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrRealOffset{PhotosPhoto: object.PhotosPhoto{ID: 10, OwnerID: 20}}, "photo20_10")
	f(object.PhotosPhotoXtrRealOffset{PhotosPhoto: object.PhotosPhoto{ID: 20, OwnerID: -10}}, "photo-10_20")
}

func Test_PhotosPhotoAlbum_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(album object.PhotosPhotoAlbum, want string) {
		t.Helper()

		if got := album.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoAlbum.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbum{ID: 10, OwnerID: 20}, "album20_10")
	f(object.PhotosPhotoAlbum{ID: 20, OwnerID: -10}, "album-10_20")
}

func TestPhotosPhotoAlbumFull_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(album object.PhotosPhotoAlbumFull, want string) {
		t.Helper()

		if got := album.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoAlbumFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbumFull{ID: 10, OwnerID: 20}, "album20_10")
	f(object.PhotosPhotoAlbumFull{ID: 20, OwnerID: -10}, "album-10_20")
}

func TestPhotosPhotoXtrTagInfo_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrTagInfo, want string) {
		t.Helper()

		if got := photo.ToAttachment(); got != want {
			t.Errorf("PhotosPhotoXtrTagInfo.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrTagInfo{PhotosPhoto: object.PhotosPhoto{ID: 10, OwnerID: 20}}, "photo20_10")
	f(object.PhotosPhotoXtrTagInfo{PhotosPhoto: object.PhotosPhoto{ID: 20, OwnerID: -10}}, "photo-10_20")
}

func TestPhotosPhoto_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhoto, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhoto.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhoto{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhoto{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhoto_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhoto, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhoto.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhoto{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhoto{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}

func TestPhotosPhotoFull_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFull, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoFull.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFull{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoFull{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhotoFull_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFull, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoFull.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFull{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoFull{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}

func TestPhotosPhotoAlbumFull_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoAlbumFull, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoAlbumFull.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbumFull{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoAlbumFull{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhotoAlbumFull_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoAlbumFull, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoAlbumFull.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoAlbumFull{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoAlbumFull{
		Sizes: []object.PhotosPhotoSizes{
			{
				object.BaseImage{Width: 10, Height: 20},
			},
			{
				object.BaseImage{Width: 100, Height: 200},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}

func TestPhotosPhotoFullXtrRealOffset_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFullXtrRealOffset, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoFullXtrRealOffset.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFullXtrRealOffset{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoFullXtrRealOffset{
		PhotosPhotoFull: object.PhotosPhotoFull{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhotoFullXtrRealOffset_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoFullXtrRealOffset, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoFullXtrRealOffset.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoFullXtrRealOffset{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoFullXtrRealOffset{
		PhotosPhotoFull: object.PhotosPhotoFull{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}

func TestPhotosPhotoXtrRealOffset_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrRealOffset, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoXtrRealOffset.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrRealOffset{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoXtrRealOffset{
		PhotosPhoto: object.PhotosPhoto{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhotoXtrRealOffset_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrRealOffset, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoXtrRealOffset.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrRealOffset{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoXtrRealOffset{
		PhotosPhoto: object.PhotosPhoto{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}

func TestPhotosPhotoXtrTagInfo_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrTagInfo, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoXtrTagInfo.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrTagInfo{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoXtrTagInfo{
		PhotosPhoto: object.PhotosPhoto{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 100, Height: 200},
	})
}

func TestPhotosPhotoXtrTagInfo_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.PhotosPhotoXtrTagInfo, want object.PhotosPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("PhotosPhotoXtrTagInfo.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.PhotosPhotoXtrTagInfo{}, object.PhotosPhotoSizes{})
	f(object.PhotosPhotoXtrTagInfo{
		PhotosPhoto: object.PhotosPhoto{
			Sizes: []object.PhotosPhotoSizes{
				{
					object.BaseImage{Width: 10, Height: 20},
				},
				{
					object.BaseImage{Width: 100, Height: 200},
				},
			},
		},
	}, object.PhotosPhotoSizes{
		BaseImage: object.BaseImage{Width: 10, Height: 20},
	})
}
