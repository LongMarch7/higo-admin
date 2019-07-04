package utils

type UserLogin struct {
    UserName        string `form:"username" validate:"required"`
    Password        string `form:"password" validate:"required,ascii"`
}

type MenuPower struct {
    Pattern  string
    Method   string
    Status   bool
}

type LoginInfo struct {
    RoleName     string   `json:"role_name" validate:"required,alphanum"`
    RoleId       int      `json:"role_id" validate:"required,numeric"`
    UserLogin    string   `json:"user_login" validate:"required,alphanum"`
    UserPass     string   `json:"user_pass"`
    UserNickname string   `json:"user_nickname"`
    UserStatus   int      `json:"user_status"`
    UserId       int64    `json:"user_id" validate:"required,numeric"`
}

type UserPostData struct {
    Id                int64  `form:"id" validate:"omitempty,numeric"`
    RoleId            []int  `form:"role_id" validate:"required,gt=0,dive,required,numeric"`
    UserLogin         string `form:"user_login" validate:"required,alphanum"`
    UserPass          string `form:"user_pass" `
    PayPass           string `form:"pay_pass" `
    UserEmail         string `form:"user_email" validate:"omitempty,email"`
}

type MemberPostData struct {
    Id                int64  `form:"id" validate:"omitempty,numeric"`
    BindId            string `form:"bind_id" validate:"omitempty,alphanum"`
    UserType          int    `form:"user_type" validate:"omitempty,numeric"`
    Sex               int    `form:"sex" validate:"required,numeric"`
    Birthday          string `form:"birthday" `
    Score             int    `form:"score" validate:"omitempty,numeric"`
    Coin              int    `form:"coin" validate:"omitempty,numeric"`
    Balance           string `form:"balance" validate:"omitempty,numeric"`
    UserLogin         string `form:"user_login" validate:"required,alphanum"`
    UserPass          string `form:"user_pass" `
    PayPass           string `form:"pay_pass" `
    UserNickname      string `form:"user_nickname" `
    UserEmail         string `form:"user_email" validate:"omitempty,email"`
    Avatar            string `form:"avatar" `
    Signature         string `form:"signature" `
}

type MailPostData struct {
    SmtpServer        string `form:"smtp_server" validate:"required"`
    Cache             string `form:"cache" validate:"required"`
    SendMail          string `form:"send_email" validate:"required"`
    SendNickname      string `form:"send_nickname" `
    SendPwd           string `form:"send_pwd" validate:"required"`
}

type MailSendData struct {
    Receive           string `form:"receive" validate:"required,email"`
    SendSubject       string `form:"send_subject" validate:"required,alphanum"`
    SendData         string `form:"send_data" validate:"required"`
}

type LinkPostData struct {
    LinkId           int `form:"link_id" validate:"omitempty,numeric"`
    LinkName         string `form:"link_name" validate:"omitempty"`
    LinkUrl          string `form:"link_url" validate:"required,url"`
    LinkImage        string `form:"link_image" validate:"omitempty,url"`
    LinkDescription  string `form:"link_description"`
}

type SitePostData struct {
    SiteName           string `form:"site_name" validate:"omitempty,alphanum"`
    SiteDomain         string `form:"site_domain" validate:"omitempty,url"`
    SiteTitle          string `form:"site_title" validate:"required"`
    SiteKeywords       string `form:"site_keywords"`
    SiteDescription    string `form:"site_description"`
    SiteCopyright    string `form:"site_copyright"`
}

type UploadPostData struct {
    PictureSize        int    `form:"picture_size" validate:"omitempty,numeric"`
    PictureType        string `form:"picture_type"`
    VideoSize          int    `form:"video_size" validate:"omitempty,numeric"`
    VideoType          string `form:"video_type"`
    AudioSize          int    `form:"audio_size" validate:"omitempty,numeric"`
    AudioType          string `form:"audio_type"`
    AnnexSize          int    `form:"annex_size" validate:"omitempty,numeric"`
    AnnexType          string `form:"annex_type"`
}

type TagPostData struct {
    TagId           int64 `form:"tag_id" validate:"omitempty,numeric"`
    TagName         string `form:"tag_name" validate:"required"`
}

type CatePostData struct {
    Id              int64  `form:"id" validate:"omitempty,numeric"`
    ParentId        int64  `form:"parent_id" validate:"omitempty,numeric"`
    Level           int    `form:"level" validate:"required,numeric"`
    Name            string `form:"name" validate:"required"`
    Description     string `form:"description"`
    SeoTitle        string `form:"seo_title"`
    SeoKeywords     string `form:"seo_keywords"`
    SeoDescription  string `form:"seo_description"`
}
type ArticlePostData struct {
    Id                int64   `form:"id" validate:"omitempty,numeric"`
    ArticleHits       int64   `form:"article_hits" `
    ArticleFavorites  int64   `form:"article_favorites" `
    ArticleLike       int64   `form:"article_like" `
    CommentCount      int64   `form:"comment_count"`
    ArticleTitle      string  `form:"article_title" validate:"required"`
    ArticleKeywords   string  `form:"article_keywords"`
    ArticleExcerpt    string  `form:"article_excerpt"`
    Thumbnail         string  `form:"thumbnail"`
    ArticleContent    string  `form:"article_content"`
    Cate              int64   `form:"cate" validate:"required,gt=0"`
    Tag               []int64 `form:"tag"`
}
