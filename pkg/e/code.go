package e

const (
	SUCCESS_BOOL = 1
	ERROR_BOOL   = 0

	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	ERR            = 50000

	ERROR_INVALID_NUMBER_OF_COLUMN = 10001
	ERROR_CAN_NOT_SAVE_FILE        = 10002
	ERROR_INVALID_UNIX_TIMESTAMP   = 10003
	ERROR_GET_TAGS_FAIL            = 10004
	ERROR_COUNT_TAG_FAIL           = 10005
	ERROR_ADD_TAG_FAIL             = 10006
	ERROR_EDIT_TAG_FAIL            = 10007
	ERROR_DELETE_TAG_FAIL          = 10008
	ERROR_EXPORT_TAG_FAIL          = 10009
	ERROR_IMPORT_TAG_FAIL          = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH_CAPTCHA             = 20004
	ERROR_AUTH                     = 20005
	ERROR_UNAUTHORIZES             = 20006
	ERROR_AUTH_PASSWORD            = 20007
	ERROR_AUTH_SECONDARY_PASSWORD  = 20008
	ERROR_AUTH_FEATURE_PASSWORD    = 20009
	ERROR_AUTH_IDENTITY_CODE       = 20010
	ERROR_IN_UPDATE_BANK_PROCESS   = 20011

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	ERROR_IN_REGISTER_MEMBER_PROCESS                = 40001
	ERROR_IN_CHECK_PASSWORD_PROCESS                 = 40002
	ERROR_IN_CHECK_SEC_PASSWORD_PROCESS             = 40003
	ERROR_IN_CHANGE_PASSWORD_PROCESS                = 40004
	ERROR_IN_CHANGE_SEC_PASSWORD_PROCESS            = 40005
	ERROR_IN_TRANSFER_PROCESS                       = 40006
	ERROR_IN_TOPUP_MW_PROCESS                       = 40007
	ERROR_DUPLICATE_USERNAME                        = 40008
	ERROR_GET_COUNTRIES                             = 40009
	ERROR_IN_GET_MEMBER                             = 40010
	ERROR_IN_GET_SPONSOR                            = 40011
	ERROR_IN_GET_WALLET                             = 40012
	ERROR_IN_GET_PAYMENT_TYPE                       = 40013
	ERROR_IN_GET_PACKAGE                            = 40014
	ERROR_WALLET_LOCK                               = 40015
	ERROR_IN_CHECK_PAY_AMOUNT                       = 40016
	ERROR_IN_CHECK_PAY_AMOUNT_GREATER_THAN_MAX      = 40017
	ERROR_IN_CHECK_PAY_AMOUNT_LOWER_THAN_MIN        = 40018
	ERROR_IN_CHECK_WALLET_BALLANCE                  = 40019
	ERROR_IN_CHECK_TOPUP_AMOUNT_MIN                 = 40020
	ERROR_IN_UPDATE_PROFILE_PROCESS                 = 40021
	ERROR_IN_PURCHASE_PIN_PROCESS                   = 40022
	ERROR_IN_PURCHASE_PIN_AGENT                     = 40023
	ERROR_IN_GET_PIN                                = 40024
	ERROR_PIN_CODE_USED                             = 40025
	ERROR_PIN_CODE_CANCELLED                        = 40026
	ERROR_PIN_CHECK                                 = 40027
	ERROR_IN_PUBLIC_REGISTER_ERR_SPONSOR_CODE       = 40028
	ERROR_IN_REGISTER_ERR_SPONSOR_CODE              = 40029
	ERROR_IN_REGISTER_ERR_UPLINE_CODE               = 40030
	ERROR_IN_REGISTER_ERR_SPONSOR_NOT_DONWLINE      = 40031
	ERROR_IN_ACTIVATE_REQUIRED_REFERRAL_CODE        = 40032
	ERROR_IN_GENERATE_SALES_ORDER                   = 40033
	ERROR_IN_PAYMENT_PROCESS                        = 40034
	ERROR_IN_ACTIVATE_REFERRAL_CODE_NOT_MEM         = 40035
	ERROR_IN_ACTIVATE_PACKAGE_NOT_FOUND             = 40036
	ERROR_IN_GET_BONUS                              = 40037
	ERROR_IN_GET_PENDING_PLACEMENT                  = 40038
	ERROR_IN_GET_MEMBERGUIDES                       = 40039
	ERROR_IN_GET_MEMBERANNOUNCEMENT                 = 40040
	ERROR_IN_GET_DETAIL_MEMBERANNOUNCEMENT          = 40041
	ERROR_IN_TRANSFER_TO_YOURSELF                   = 40042
	ERROR_IN_TRANSFER_TO_INACTIVE_MEMBER            = 40043
	ERROR_IN_GET_PLACEMENT_TREE                     = 40044
	ERROR_IN_GET_INFO_PLACEMENT_TREE                = 40045
	ERROR_IN_GET_DETAIL_MEMBERGUIDES                = 40046
	ERROR_IN_GET_ADMIN                              = 40047
	ERROR_IN_GET_USERGROUP                          = 40048
	ERROR_IN_CREATE_ADMIN_GROUP                     = 40049
	ERROR_IN_GET_SYSTEM_MENU                        = 40050
	ERROR_IN_CREATE_SYSTEM_MENU                     = 40051
	ERROR_IN_UPDATE_SYSTEM_MENU                     = 40052
	ERROR_IN_GET_COUNTRIES                          = 40053
	ERROR_IN_CREATE_COUNTRY_PROCESS                 = 40054
	ERROR_DUPLICATE_COUNTRY_CODE                    = 40055
	ERROR_IN_UPDATE_COUNTRY_PROCESS                 = 40056
	ERROR_IN_UPDATE_ADMIN_GROUP                     = 40057
	ERROR_IN_GET_PAYMENT_SETTING                    = 40058
	ERROR_IN_CREATE_PAYMENT_SETTING                 = 40059
	ERROR_IN_UPDATE_PAYMENT_SETTING                 = 40060
	ERROR_IN_GET_WALLET_TYPE                        = 40061
	ERROR_IN_CREATE_WALLET_TYPE                     = 40062
	ERROR_IN_UPDATE_WALLET_TYPE                     = 40063
	ERROR_IN_GET_CONTACT_TYPE                       = 40064
	ERROR_IN_CREATE_CONTACT_TYPE                    = 40065
	ERROR_IN_UPDATE_CONTACT_TYPE                    = 40066
	ERROR_IN_GET_MEMBER_STATUS                      = 40067
	ERROR_IN_CREATE_MEMBER_STATUS                   = 40068
	ERROR_IN_UPDATE_MEMBER_STATUS                   = 40069
	ERROR_IN_GET_MEMBER_TYPE                        = 40070
	ERROR_IN_CREATE_MEMBER_TYPE                     = 40071
	ERROR_IN_UPDATE_MEMBER_TYPE                     = 40072
	ERROR_IN_CREATE_PAYMENT_TYPE                    = 40073
	ERROR_IN_UPDATE_PAYMENT_TYPE                    = 40074
	ERROR_IN_GET_ANNOUNCEMENT_LIST                  = 40075
	ERROR_IN_CREATE_ANNOUNCEMENT                    = 40076
	ERROR_IN_UPDATE_ANNOUNCEMENT_LIST               = 40077
	ERROR_IN_GET_GENDER                             = 40078
	ERROR_IN_CREATE_GENDER                          = 40079
	ERROR_IN_UPDATE_GENDER                          = 40080
	ERROR_IN_GET_BANK                               = 40081
	ERROR_IN_CREATE_BANK                            = 40082
	ERROR_IN_UPDATE_BANK                            = 40083
	ERROR_IN_GET_RELATIONSHIP                       = 40084
	ERROR_IN_CREATE_RELATIONSHIP                    = 40085
	ERROR_IN_UPDATE_RELATIONSHIP                    = 40086
	ERROR_IN_GET_CENTERTYPE                         = 40087
	ERROR_IN_CREATE_CENTERTYPE                      = 40088
	ERROR_IN_UPDATE_CENTERTYPE                      = 40089
	ERROR_IN_GET_MARITAL_STATUS                     = 40090
	ERROR_IN_CREATE_MARITAL_STATUS                  = 40091
	ERROR_IN_UPDATE_MARITAL_STATUS                  = 40092
	ERROR_IN_GET_RACE                               = 40093
	ERROR_IN_CREATE_RACE                            = 40094
	ERROR_IN_UPDATE_RACE                            = 40095
	ERROR_IN_GET_TERRITORY_TYPE                     = 40096
	ERROR_IN_CREATE_TERRITORY_TYPE                  = 40097
	ERROR_IN_UPDATE_TERRITORY_TYPE                  = 40098
	ERROR_IN_GET_TERRITORY                          = 40099
	ERROR_IN_CREATE_TERRITORY                       = 40100
	ERROR_IN_UPDATE_TERRITORY                       = 40101
	ERROR_IN_GET_WITHDRAW_PAYMENT_TYPE              = 40102
	ERROR_IN_CREATE_WITHDRAW_PAYMENT_TYPE           = 40103
	ERROR_IN_UPDATE_WITHDRAW_PAYMENT_TYPE           = 40104
	ERROR_IN_GET_CURRENCY                           = 40105
	ERROR_IN_CREATE_CURRENCY                        = 40106
	ERROR_IN_UPDATE_CURRENCY                        = 40107
	ERROR_IN_GET_SYS_CURRENCY_RATE                  = 40108
	ERROR_IN_CREATE_SYS_CURRENCY_RATE               = 40109
	ERROR_IN_UPDATE_SYS_CURRENCY_RATE               = 40110
	ERROR_IN_GET_MODULE                             = 40111
	ERROR_IN_CREATE_MODULE                          = 40112
	ERROR_IN_UPDATE_MODULE                          = 40113
	ERROR_IN_UPDATE_ADMIN_GROUP_PERMISSION          = 40114
	ERROR_IN_DELETE_ADMIN_GROUP                     = 40115
	ERROR_IN_GET_SALES_LIST                         = 40116
	ERROR_IN_GET_PIN_LIST                           = 40117
	ERROR_IN_GET_GROUP_TYPE                         = 40118
	ERROR_IN_CANCEL_SALES                           = 40119
	ERROR_IN_CANCEL_PIN                             = 40120
	ERROR_IN_GET_GENEALOGY_UPLINE_TREE              = 40121
	ERROR_IN_GET_DOWNLOAD_LIST                      = 40122
	ERROR_IN_CREATE_DOWNLOAD_LIST                   = 40123
	ERROR_IN_UPDATE_DOWNLOAD_LIST                   = 40124
	ERROR_IN_GET_DOCUMENT_TYPE                      = 40125
	ERROR_IN_CREATE_DOCUMENT_TYPE                   = 40126
	ERROR_IN_UPDATE_DOCUMENT_TYPE                   = 40127
	ERROR_IN_GET_DOC_TYPE                           = 40128
	ERROR_IN_GET_TABLE_NAME                         = 40129
	ERROR_IN_GET_COMPANY                            = 40130
	ERROR_IN_GET_CHANGE_LOT_LIST                    = 40131
	ERROR_IN_CHANGE_SPONSOR                         = 40132
	ERROR_IN_CHANGE_UPLINE                          = 40133
	ERROR_IN_GET_UPLINE                             = 40134
	ERROR_IN_CREATE_ADMIN                           = 40135
	ERROR_IN_CREATE_AGENT                           = 40136
	ERROR_IN_SUSPEND_MEMBER                         = 40137
	ERROR_IN_REACTIVATE_MEMBER                      = 40138
	ERROR_IN_TERMINATE_MEMBER                       = 40139
	ERROR_CAUSE_BY_MEMBER_STATUS_NOT_ACTIVE         = 40140
	ERROR_CAUSE_BY_MEMBER_STATUS_NOT_SUSPEND        = 40141
	ERROR_CAUSE_BY_MEMBER_STATUS_ALREADY_TERMINATED = 40142
	ERROR_CAUSE_BY_MEMBER_STATUS_ALREADY_CANCELLED  = 40143
	ERROR_INVALID_MEMBER_SETTING_GROUP              = 40145
	ERROR_IN_GENERATE_DOCUMENT_NUMBER               = 40146
	ERROR_IN_GET_SUSPENDED_LIST                     = 40147
	ERROR_IN_GET_CANCELLED_LIST                     = 40148
	ERROR_IN_GET_TERMINATED_LIST                    = 40149
	ERROR_IN_GET_MEMBER_STATUS_CHANGE_RECORD        = 40150
	ERROR_IN_ADJUST_WALLET_PROCESS                  = 40151
	ERROR_REQUIRED_PAYMENT_TYPE                     = 40152
	ERROR_PAYMENT_TYPE_WRONG                        = 40153
	ERROR_IN_GET_TRANSACTION                        = 40154
	ERROR_IN_GET_BALANCE                            = 40155
	ERROR_IN_GET_WALLET_SETTING                     = 40156
	ERROR_IN_CREATE_WALLET_SETTING                  = 40157
	ERROR_IN_UPDATE_WALLET_SETTING                  = 40158
	ERROR_IN_GET_MEMBER_WALLET_LOCK                 = 40159
	ERROR_IN_UPDATE_MEMBER_WALLET_LOCK              = 40160
	ERROR_IN_GET_APPROVAL_LIST                      = 40161
	ERROR_IN_GET_MEMBER_DETAIL                      = 40162
	ERROR_IN_CHECK_WALLET_TYPE                      = 40163
	ERROR_IN_DELETE_ADMIN_GROUP_REMAIN_CHILD        = 40164
	WALLET_MEMBER_ALREADY_HIDE                      = 40165
	WALLET_MEMBER_ALREADY_UN_HIDE                   = 40166
	ERROR_IN_SETTING_WALLET_MEMBER                  = 40167
	ERROR_IN_GET_MEMBER_WALLET_HIDE                 = 40168
	ERROR_IN_LOCK_WALLET_MEMBER                     = 40169
	ERROR_IN_GET_CATEGORY                           = 40170
	ERROR_IN_GET_PRODUCT                            = 40171
	ERROR_IN_CREATE_PRODUCT                         = 40172
	ERROR_IN_GET_TAX_TYPE                           = 40173
	ERROR_IN_GET_PRICE_TYPE                         = 40174
	ERROR_IN_CREATE_PRICE_TYPE                      = 40175
	ERROR_IN_UPDATE_PRICE_TYPE                      = 40176
	ERROR_IN_GET_BRAND                              = 40177
	ERROR_IN_CREATE_PRICE_COUNTRY                   = 40178
	ERROR_IN_UPDATE_PRICE_COUNTRY                   = 40179
	ERROR_IN_UPDATE_PRODUCT                         = 40180
	INVALID_ADJUSTMENT_AMOUNT                       = 40181
	INVALID_TRANSFER_AMOUNT                         = 40182
	ERROR_IN_CREATE_ALBUM_PHOTO                     = 40183
	ERROR_IN_UPDATE_ALBUM_PHOTO                     = 40184
	ERROR_IN_DELETE_ALBUM_PHOTO                     = 40185
	ERROR_IN_GET_ALBUM                              = 40186
	ERROR_IN_GET_ALBUM_VIDEO                        = 40187
	ERROR_IN_CREATE_ALBUM_VIDEO                     = 40189
	ERROR_IN_UPDATE_ALBUM_VIDEO                     = 40190
	ERROR_IN_DELETE_ALBUM_VIDEO                     = 40191
	ERROR_IN_GET_ACTIVATION_CODE                    = 40192
	ERROR_IN_CREATE_SKU_CODE                        = 40193
	ERROR_IN_EDIT_SKU_CODE                          = 40194
	ERROR_IN_DELETE_SKU_CODE                        = 40195
	ERROR_IN_GET_SKU_CODE                           = 40196
	ERROR_DUPLICATE_SKU_CODE                        = 40197
	ERROR_IN_GET_STORE                              = 40198
	ERROR_IN_GET_REGISTRATION_PACKAGE               = 40199
	ERROR_IN_CREATE_REGISTRATION_PACKAGE            = 40200
	ERROR_IN_GET_CUSTOMER_GROUP                     = 40201
	ERROR_IN_CREATE_CATEGORY                        = 40202
	ERROR_IN_GET_ATTRIBUTE_SET                      = 40203
	ERROR_INVALID_PARENT_CATEGORY                   = 40204
	ERROR_IN_GET_ATTRIBUTE                          = 40205
	ERROR_IN_GET_ATTRIBUTE_OPTION                   = 40206
	ERROR_IN_CREATE_PACKAGE                         = 40207
	ERROR_IN_GET_SIMPLE_PRODUCT                     = 40208
	ERROR_IN_BINDING                                = 40209
	ERROR_IN_PIN_CODE_NOT_FOUND                     = 40210
	ERROR_PACKAGE_NOT_FOUND                         = 40211
	ERROR_IN_MAINTENANCE_PROCESS                    = 40212
	ERROR_IN_INVOICE_DETAIL                         = 40213
	ERROR_IN_GET_TOP_UP_REQUEST                     = 40214
	ERROR_IN_TOPUP_PROCESS                          = 40215
	ERROR_IN_TOP_UP_REQUEST_ALREAY_PROCESS          = 40216
	ERROR_IN_ACTIVATION_CODE                        = 40217
	ERROR_IN_SET_FEATURE_PASSWORD                   = 40218
	ERROR_IN_UPDATE_FEATURE_PASSWORD                = 40219
	ERROR_IN_CREATE_POTENTIAL_MEMBER                = 40220
	ERROR_IN_UPDATE_SYSTEM_SETTING                  = 40221
	ERROR_IN_GET_SYSTEM_STATIC_SETTING              = 40222
	INVALID_HASH                                    = 40223
	ERROR_IN_MATCH_ACTIVATE_CODE                    = 40224
	ERROR_IN_UPDATED_USED_ACTIVATE_CODE             = 40225
	ERROR_IN_GET_AVAILABLE_SLOT                     = 40226
	ERROR_IN_QUERY_MEMBER                           = 40227
	ERROR_IN_CREATE_ATTRIBUTE_OPTION_VALUE          = 40228
	ERROR_IN_UPDATE_ATTRIBUTE_OPTION_VALUE          = 40229
	ERROR_IN_UPDATE_BENEFICIARY                     = 40230
	ERROR_IN_WITHDRAW_CHECK_GREATER_THAN_MIN        = 40231
	ERROR_IN_CHECK_ACTIVATION_EXPIRY_DATE           = 40232
	ERROR_CHECK_PLACEMENT_TREE_NETWORK              = 40233
	ERROR_IN_ASSIGN_MEMBER_RANK                     = 40234
	ERROR_IN_ALREADY_ASSIGN_MEMBER_RANK             = 40235
	ERROR_IN_UNDEFINED_MEMBER_RANK                  = 40236
	ERROR_IN_GET_TIMEZONE                           = 40237
	ERROR_IN_GET_TRANSFER_MEMBER                    = 40238
	ERROR_IN_MEMBER_NOT_EXISTS                      = 40239
	ERROR_IN_UPDATE_LOCKED_ECODE                    = 40240
	ERROR_IN_GET_INVOICE                            = 40241
	ERROR_IN_CHECK_INVENTORY                        = 40242
	ERROR_IN_CHECK_PRODUCT_BUY_QUANTITY             = 40243
	ERROR_IN_UPDATE_PACKAGE                         = 40244
	ERROR_IN_GET_DEFAULT_STOCK                      = 40245
	ERROR_IN_PRODUCT_INVENTORY_NEGATIVE             = 40246
	ERROR_IN_CANCEL_REGISTER                        = 40247
	ERROR_IN_NO_CANCELLATION_ABILITY                = 40248
	ERROR_IN_INVOICE_ALREADY_CANCELED               = 40249
	ERROR_IN_BONUS_RECALCULATION_ABILITY            = 40250
	ERROR_IN_CREAT_EVENT                            = 40251
	ERROR_IN_CHECK_TRAINING_KEY_NOT_FOUND           = 40252
	ERROR_IN_CHECK_TRAINING_KEY_USED                = 40253
	ERROR_IN_CHECK_EVENT_NOT_FOUND                  = 40254
	ERROR_IN_CHECK_EVENT_TICKET                     = 40255
	ERROR_IN_CHECK_EVENT_TICKET_ALREADY_JOIN        = 40256
	ERROR_IN_CHECK_TICKET_NOT_FOUND                 = 40257
	ERROR_IN_CHECK_REFRESHER_NOT_FOUND              = 40258
	ERROR_IN_WITHDRAW_CHECK_KEEP_WALLET_AMOUNT      = 40259
	ERROR_IN_DELETE_ANNOUNCEMENT                    = 40260
	ERROR_IN_UPDATE_LOCKED_ECODE_OF_INACTIVE_MEMBER = 40261
	ERROR_IN_WITHDRAW_PROCESS                       = 40262
	ERROR_IN_CREAT_DISCOUNT_CODE                    = 40263
	ERROR_IN_UPDATE_DISCOUNT_CODE                   = 40264
	ERROR_DUPLICATE_DISCOUNT_CODE                   = 40265
	ERROR_DUPLICATE_IDENTITY                        = 40266
	ERROR_IN_CREATE_DOCUMENT                        = 40267
	ERROR_IN_DELETE_EVENT                           = 40268
	ERROR_PERIOD_DATE                               = 40269
	ERROR_IN_GET_STOCK                              = 40270
	ERROR_IN_GET_TEAM_MEMBER                        = 40271
	ERROR_IN_GET_SLIDER                             = 40272
	ERROR_IN_GET_POTENTIAL_MEMBER                   = 40273
	ERROR_IN_GET_INTRO_VIDEO                        = 40274
	ERROR_IN_CHECK_INVALID_UPLINE                   = 40275
	ERROR_IN_CHECK_SLOT                             = 40276
	ERROR_IN_DELETE_SELF                            = 40277
	ERROR_IN_DELETE_ROOT                            = 40278
	ERROR_IN_INVALID_SPONSOR_CODE                   = 40279
	ERROR_IN_CREATE_SUPPLIER                        = 40280
	ERROR_IN_GET_SUPPLIER                           = 40281
	ERROR_IN_DELETE_SUPPLIER                        = 40282
	ERROR_IN_UPDATE_SUPPLIER                        = 40283
	ERROR_IN_CHANGE_UPLINE_VALIDATE_SELF_SLOT       = 40284
	ERROR_IN_CREATE_GOOD_RECEIVE                    = 40285
	ERROR_IN_GET_GOOD_RECEIVE                       = 40286
	ERROR_IN_GET_ADJUSTMENT                         = 40287
	ERROR_IN_CREATE_ADJUSTMENT                      = 40288
	ERROR_IN_CHANGE_UPLINE_VALIDATE_DOWNLINE_SLOT   = 40289
	ERROR_IN_INVALID_TYPE                           = 40290
	ERROR_IN_INVALID_SITE                           = 40291
	ERROR_IN_INVALID_ACTION                         = 40292
	ERROR_IN_UPDATE_WALLET_TRANSACTION              = 40293
	ERROR_IN_GET_WALLET_TRANSACTION                 = 40294
	ERROR_IN_CREATE_DELIVERY_SETTING                = 40295
	ERROR_IN_EDIT_DELIVERY_SETTING                  = 40296
	ERROR_IN_GET_DELIVERY_SETTING                   = 40297
	ERROR_IN_ACTIVE_RANGE_DESCOUNT_CODE             = 40298
	ERROR_IN_ACTIVE_RANGE_DELIVERY_SETTING          = 40299
	ERROR_IN_CHECK_DELIVERY_FEE_RANGE               = 40300
	ERROR_IN_GET_COMMISSION_SETTING                 = 40301
	ERROR_IN_GET_BONUS_SETTING_LIST                 = 40302
	ERROR_IN_CREAT_BONUS_SETTING                    = 40303
	ERROR_IN_GET_BONUS_SETTING_DETAIL               = 40304
	ERROR_IN_UPDATE_BONUS_SETTING                   = 40305
	ERROR_IN_CHECK_SLOT_IS_BUSY                     = 40306
	ERROR_GET_LANGUAGES                             = 40307
	ERROR_IN_CHECK_TRANSFER_WALLET_BALLANCE         = 40308
	ERROR_IN_DELETE_ADMIN_UEER                      = 40309
	ERROR_IN_UPDATE_SCHOOL_CATEGORY                 = 40310
	ERROR_IN_UPDATE_SUBJECT                         = 40311

	MEM_STATUS_TRACK_MESSAGE_NEW                              = 60000
	MEM_STATUS_TRACK_MESSAGE_E_CODE_OVER                      = 60001
	MEM_STATUS_TRACK_MESSAGE_LAST_MAINTENANCE_OVER_2YEARS     = 60002
	MEM_STATUS_TRACK_MESSAGE_MAINTENANCE_EXPIRED              = 60003
	MEM_STATUS_TRACK_MESSAGE_MEMBERSHIP_OVER                  = 60004
	MEM_STATUS_TRACK_MESSAGE_ADMIN_MANUALLY                   = 60005
	MEM_STATUS_TRACK_MESSAGE_ADMIN_UPD_MEMBERSHIP             = 60006
	MEM_STATUS_TRACK_MESSAGE_ADMIN_UPD_MAINTENANCE            = 60007
	MEM_STATUS_TRACK_MESSAGE_ADMIN_UPD_UPGRADE                = 60008
	MEM_STATUS_TRACK_MESSAGE_MEMBER_PACKAGE_PURCHASE_COMEBACK = 60009
	MEM_STATUS_TRACK_MESSAGE_MEMBER_REGISTER_COMEBACK         = 60010
	MEM_STATUS_TRACK_MESSAGE_MEMBER_MAINTENANCE_COMEBACK      = 60011
	MEM_STATUS_TRACK_MESSAGE_MEMBER_UPGRADE_COMEBACK          = 60012

	MEM_PACKAGE_TRACK_MESSAGE_NEW              = 70000
	MEM_PACKAGE_TRACK_MESSAGE_MANUALLY_UPGRADE = 70001
	MEM_PACKAGE_TRACK_MESSAGE_ADMIN_REVERT     = 70002

	EXPIRED_DATE_TRACK_MESSAGE_BY_PURCHASE    = 61000
	EXPIRED_DATE_TRACK_MESSAGE_BY_REGISTER    = 61001
	EXPIRED_DATE_TRACK_MESSAGE_BY_MAINTENANCE = 61002
	EXPIRED_DATE_TRACK_MESSAGE_BY_UPGRADE     = 61003
	EXPIRED_DATE_TRACK_MESSAGE_BY_ADMIN       = 61004
	EXPIRED_DATE_TRACK_MESSAGE_BY_REVERT      = 61005

	SUCCESS_LOGIN                  = 20200
	SUCCESS_LOGOUT                 = 20201
	SUCCESS_REGISTER               = 20202
	SUCCESS_UPDATE_PASSWORD        = 20203
	SUCCESS_UPDATE_PROFILE         = 20204
	SUCCESS_UPDATE_BENEFICIARY     = 20205
	SUCCESS_UPDATE_AVATAR          = 20206
	SUCCESS_UPDATE_SECOND_PASSWORD = 20207
	SUCCESS_IN_CHECKING_PASSWORD   = 20208
	SUCCESS_UPDATE_BANK            = 20209
)

var ConstLookup = map[uint64]string{
	ERROR_AUTH_FEATURE_PASSWORD: "ERROR_AUTH_FEATURE_PASSWORD",
	ERROR_AUTH_IDENTITY_CODE:    "ERROR_AUTH_IDENTITY_CODE",
	ERROR_AUTH_PASSWORD:         "ERROR_AUTH_PASSWORD",
}

var ConstTracedBack = map[string]int{
	"ERROR_AUTH_FEATURE_PASSWORD":   ERROR_AUTH_FEATURE_PASSWORD,
	"ERROR_AUTH_IDENTITY_CODE":      ERROR_AUTH_IDENTITY_CODE,
	"ERROR_AUTH_SECONDARY_PASSWORD": ERROR_AUTH_SECONDARY_PASSWORD,
}
