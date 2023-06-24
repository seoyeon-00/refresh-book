package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User는 사용자 엔티티에 대한 스키마 정의를 담고 있습니다.
type User struct {
	ent.Schema
}

// Fields는 User의 필드를 정의합니다.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 사용자의 고유 식별자
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("사용자의 고유 식별자"),
		// 사용자의 이름
		field.String("name").
			NotEmpty().
			Comment("사용자의 이름"),
		// 사용자의 이메일 주소 (고유해야 함)
		field.String("email").
			NotEmpty().
			Unique().
			Comment("사용자의 이메일 주소"),
		// 사용자의 비밀번호 (API 응답에 노출되지 않음)
		field.String("password").
			Sensitive().
			Comment("사용자의 비밀번호"),
		// 사용자의 주소 (선택 사항)
		field.String("address").
			Optional().
			Comment("사용자의 주소"),
		// 사용자의 전화번호 (선택 사항)
		field.Int("phone").
			Optional().
			Comment("사용자의 전화번호"),
		// 사용자가 관리자인지 여부
		field.Bool("is_admin").
			Default(false).
			Comment("사용자가 관리자인지 여부"),
		// 사용자 계정 생성 시각
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("사용자 계정 생성 시각"),
	}
}

// Indexes는 User의 색인을 정의합니다.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 이메일 필드에 대해 고유 색인 생성
		index.Fields("email").Unique(),
	}
}
