package mtproto

import "fmt"

const (
	crc_boolFalse                                                         = 0xbc799737
	crc_boolTrue                                                          = 0x997275b5
	crc_true                                                              = 0x3fedd339
	crc_error                                                             = 0xc4b9f9bb
	crc_null                                                              = 0x56730bcc
	crc_inputPeerEmpty                                                    = 0x7f3b18ea
	crc_inputPeerSelf                                                     = 0x7da07ec9
	crc_inputPeerChat                                                     = 0x179be863
	crc_inputUserEmpty                                                    = 0xb98886cf
	crc_inputUserSelf                                                     = 0xf7c1b13f
	crc_inputPhoneContact                                                 = 0xf392b7f4
	crc_inputFile                                                         = 0xf52ff27f
	crc_inputMediaEmpty                                                   = 0x9664f57f
	crc_inputMediaUploadedPhoto                                           = 0x1e287d04
	crc_inputMediaPhoto                                                   = 0xb3ba0635
	crc_inputMediaGeoPoint                                                = 0xf9c44144
	crc_inputMediaContact                                                 = 0xf8ab7dfb
	crc_inputChatPhotoEmpty                                               = 0x1ca48f57
	crc_inputChatUploadedPhoto                                            = 0x927c55b4
	crc_inputChatPhoto                                                    = 0x8953ad37
	crc_inputGeoPointEmpty                                                = 0xe4c123d6
	crc_inputGeoPoint                                                     = 0xf3b7acc9
	crc_inputPhotoEmpty                                                   = 0x1cd7bf0d
	crc_inputPhoto                                                        = 0x3bb3b94a
	crc_inputFileLocation                                                 = 0xdfdaabe1
	crc_peerUser                                                          = 0x9db1bc6d
	crc_peerChat                                                          = 0xbad0e5bb
	crc_storage_fileUnknown                                               = 0xaa963b05
	crc_storage_filePartial                                               = 0x40bc6f52
	crc_storage_fileJpeg                                                  = 0x007efe0e
	crc_storage_fileGif                                                   = 0xcae1aadf
	crc_storage_filePng                                                   = 0x0a4f63c0
	crc_storage_filePdf                                                   = 0xae1e508d
	crc_storage_fileMp3                                                   = 0x528a0677
	crc_storage_fileMov                                                   = 0x4b09ebbc
	crc_storage_fileMp4                                                   = 0xb3cea0e4
	crc_storage_fileWebp                                                  = 0x1081464c
	crc_userEmpty                                                         = 0x200250ba
	crc_userProfilePhotoEmpty                                             = 0x4f11bae1
	crc_userProfilePhoto                                                  = 0xecd75d8c
	crc_userStatusEmpty                                                   = 0x09d05049
	crc_userStatusOnline                                                  = 0xedb93949
	crc_userStatusOffline                                                 = 0x008c703f
	crc_chatEmpty                                                         = 0x9ba2d800
	crc_chat                                                              = 0x3bda1bde
	crc_chatForbidden                                                     = 0x07328bdb
	crc_chatFull                                                          = 0x1b7c9db3
	crc_chatParticipant                                                   = 0xc8d7493e
	crc_chatParticipantsForbidden                                         = 0xfc900c2b
	crc_chatParticipants                                                  = 0x3f460fed
	crc_chatPhotoEmpty                                                    = 0x37c1011c
	crc_chatPhoto                                                         = 0x475cdbd5
	crc_messageEmpty                                                      = 0x83e5de54
	crc_message                                                           = 0x452c0e65
	crc_messageService                                                    = 0x9e19a1f6
	crc_messageMediaEmpty                                                 = 0x3ded6320
	crc_messageMediaPhoto                                                 = 0x695150d7
	crc_messageMediaGeo                                                   = 0x56e0d474
	crc_messageMediaContact                                               = 0xcbf24940
	crc_messageMediaUnsupported                                           = 0x9f84f49e
	crc_messageActionEmpty                                                = 0xb6aef7b0
	crc_messageActionChatCreate                                           = 0xa6638b9a
	crc_messageActionChatEditTitle                                        = 0xb5a1ce5a
	crc_messageActionChatEditPhoto                                        = 0x7fcb13a8
	crc_messageActionChatDeletePhoto                                      = 0x95e3fbef
	crc_messageActionChatAddUser                                          = 0x488a7337
	crc_messageActionChatDeleteUser                                       = 0xb2ae9b0c
	crc_dialog                                                            = 0x2c171f72
	crc_photoEmpty                                                        = 0x2331b22d
	crc_photo                                                             = 0xd07504a5
	crc_photoSizeEmpty                                                    = 0x0e17e23c
	crc_photoSize                                                         = 0x77bfb61b
	crc_photoCachedSize                                                   = 0xe9a734fa
	crc_geoPointEmpty                                                     = 0x1117dd5f
	crc_geoPoint                                                          = 0x0296f104
	crc_auth_sentCode                                                     = 0x5e002502
	crc_auth_authorization                                                = 0xcd050916
	crc_auth_exportedAuthorization                                        = 0xdf969c2d
	crc_inputNotifyPeer                                                   = 0xb8bc5b0c
	crc_inputNotifyUsers                                                  = 0x193b4417
	crc_inputNotifyChats                                                  = 0x4a95e84e
	crc_inputPeerNotifySettings                                           = 0x9c3d198e
	crc_peerNotifySettings                                                = 0xaf509d20
	crc_peerSettings                                                      = 0x818426cd
	crc_wallPaper                                                         = 0xa437c3ed
	crc_inputReportReasonSpam                                             = 0x58dbcab8
	crc_inputReportReasonViolence                                         = 0x1e22c78d
	crc_inputReportReasonPornography                                      = 0x2e59d922
	crc_inputReportReasonChildAbuse                                       = 0xadf44ee3
	crc_inputReportReasonOther                                            = 0xe1746d0a
	crc_userFull                                                          = 0xedf17c12
	crc_contact                                                           = 0xf911c994
	crc_importedContact                                                   = 0xd0028438
	crc_contactBlocked                                                    = 0x561bc879
	crc_contactStatus                                                     = 0xd3680c61
	crc_contacts_contactsNotModified                                      = 0xb74ba9d2
	crc_contacts_contacts                                                 = 0xeae87e42
	crc_contacts_importedContacts                                         = 0x77d01c3b
	crc_contacts_blocked                                                  = 0x1c138d15
	crc_contacts_blockedSlice                                             = 0x900802a1
	crc_messages_dialogs                                                  = 0x15ba6c40
	crc_messages_dialogsSlice                                             = 0x71e094f3
	crc_messages_messages                                                 = 0x8c718e87
	crc_messages_messagesSlice                                            = 0xc8edce1e
	crc_messages_chats                                                    = 0x64ff9fd5
	crc_messages_chatFull                                                 = 0xe5d7d19c
	crc_messages_affectedHistory                                          = 0xb45c69d1
	crc_inputMessagesFilterEmpty                                          = 0x57e2f66c
	crc_inputMessagesFilterPhotos                                         = 0x9609a51c
	crc_inputMessagesFilterVideo                                          = 0x9fc00e65
	crc_inputMessagesFilterPhotoVideo                                     = 0x56e9f0e4
	crc_inputMessagesFilterDocument                                       = 0x9eddf188
	crc_inputMessagesFilterUrl                                            = 0x7ef0dd87
	crc_inputMessagesFilterGif                                            = 0xffc86587
	crc_updateNewMessage                                                  = 0x1f2b0afd
	crc_updateMessageID                                                   = 0x4e90bfd6
	crc_updateDeleteMessages                                              = 0xa20db0e5
	crc_updateUserTyping                                                  = 0x5c486927
	crc_updateChatUserTyping                                              = 0x9a65ea1f
	crc_updateChatParticipants                                            = 0x07761198
	crc_updateUserStatus                                                  = 0x1bfbd823
	crc_updateUserName                                                    = 0xa7332b73
	crc_updateUserPhoto                                                   = 0x95313b0c
	crc_updates_state                                                     = 0xa56c2a3e
	crc_updates_differenceEmpty                                           = 0x5d75a138
	crc_updates_difference                                                = 0x00f49ca0
	crc_updates_differenceSlice                                           = 0xa8fb1981
	crc_updatesTooLong                                                    = 0xe317af7e
	crc_updateShortMessage                                                = 0x914fbf11
	crc_updateShortChatMessage                                            = 0x16812688
	crc_updateShort                                                       = 0x78d4dec1
	crc_updatesCombined                                                   = 0x725b04c3
	crc_updates                                                           = 0x74ae4240
	crc_photos_photos                                                     = 0x8dca6aa5
	crc_photos_photosSlice                                                = 0x15051f54
	crc_photos_photo                                                      = 0x20212ca8
	crc_upload_file                                                       = 0x096a18d5
	crc_dcOption                                                          = 0x18b7a10d
	crc_config                                                            = 0x330b4067
	crc_nearestDc                                                         = 0x8e1a1775
	crc_help_appUpdate                                                    = 0x1da7158f
	crc_help_noAppUpdate                                                  = 0xc45a6536
	crc_help_inviteText                                                   = 0x18cb9f78
	crc_updateNewEncryptedMessage                                         = 0x12bcbd9a
	crc_updateEncryptedChatTyping                                         = 0x1710f156
	crc_updateEncryption                                                  = 0xb4a2e88d
	crc_updateEncryptedMessagesRead                                       = 0x38fe25b7
	crc_encryptedChatEmpty                                                = 0xab7ec0a0
	crc_encryptedChatWaiting                                              = 0x3bf703dc
	crc_encryptedChatRequested                                            = 0xc878527e
	crc_encryptedChat                                                     = 0xfa56ce36
	crc_encryptedChatDiscarded                                            = 0x13d6dd27
	crc_inputEncryptedChat                                                = 0xf141b5e1
	crc_encryptedFileEmpty                                                = 0xc21f497e
	crc_encryptedFile                                                     = 0x4a70994c
	crc_inputEncryptedFileEmpty                                           = 0x1837c364
	crc_inputEncryptedFileUploaded                                        = 0x64bd0306
	crc_inputEncryptedFile                                                = 0x5a17b5e5
	crc_inputEncryptedFileLocation                                        = 0xf5235d55
	crc_encryptedMessage                                                  = 0xed18c118
	crc_encryptedMessageService                                           = 0x23734b06
	crc_messages_dhConfigNotModified                                      = 0xc0e24635
	crc_messages_dhConfig                                                 = 0x2c221edd
	crc_messages_sentEncryptedMessage                                     = 0x560f8935
	crc_messages_sentEncryptedFile                                        = 0x9493ff32
	crc_inputFileBig                                                      = 0xfa4f0bb5
	crc_inputEncryptedFileBigUploaded                                     = 0x2dc173c8
	crc_updateChatParticipantAdd                                          = 0xea4b0e5c
	crc_updateChatParticipantDelete                                       = 0x6e5f8c22
	crc_updateDcOptions                                                   = 0x8e5e9873
	crc_inputMediaUploadedDocument                                        = 0x5b38c6c1
	crc_inputMediaDocument                                                = 0x23ab23d2
	crc_messageMediaDocument                                              = 0x9cb070d7
	crc_inputDocumentEmpty                                                = 0x72f0eaae
	crc_inputDocument                                                     = 0x1abfb575
	crc_inputDocumentFileLocation                                         = 0xbad07584
	crc_documentEmpty                                                     = 0x36f8c871
	crc_document                                                          = 0x9ba29cc1
	crc_help_support                                                      = 0x17c6b5f6
	crc_notifyPeer                                                        = 0x9fd40bd8
	crc_notifyUsers                                                       = 0xb4c83b4c
	crc_notifyChats                                                       = 0xc007cec3
	crc_updateUserBlocked                                                 = 0x80ece81a
	crc_updateNotifySettings                                              = 0xbec268ef
	crc_sendMessageTypingAction                                           = 0x16bf744e
	crc_sendMessageCancelAction                                           = 0xfd5ec8f5
	crc_sendMessageRecordVideoAction                                      = 0xa187d66f
	crc_sendMessageUploadVideoAction                                      = 0xe9763aec
	crc_sendMessageRecordAudioAction                                      = 0xd52f73f7
	crc_sendMessageUploadAudioAction                                      = 0xf351d7ab
	crc_sendMessageUploadPhotoAction                                      = 0xd1d34a26
	crc_sendMessageUploadDocumentAction                                   = 0xaa0cd9e4
	crc_sendMessageGeoLocationAction                                      = 0x176f8ba1
	crc_sendMessageChooseContactAction                                    = 0x628cbc6f
	crc_contacts_found                                                    = 0xb3134d9d
	crc_updateServiceNotification                                         = 0xebe46819
	crc_userStatusRecently                                                = 0xe26f42f1
	crc_userStatusLastWeek                                                = 0x07bf09fc
	crc_userStatusLastMonth                                               = 0x77ebc742
	crc_updatePrivacy                                                     = 0xee3b272a
	crc_inputPrivacyKeyStatusTimestamp                                    = 0x4f96cb18
	crc_privacyKeyStatusTimestamp                                         = 0xbc2eab30
	crc_inputPrivacyValueAllowContacts                                    = 0x0d09e07b
	crc_inputPrivacyValueAllowAll                                         = 0x184b35ce
	crc_inputPrivacyValueAllowUsers                                       = 0x131cc67f
	crc_inputPrivacyValueDisallowContacts                                 = 0x0ba52007
	crc_inputPrivacyValueDisallowAll                                      = 0xd66b66c9
	crc_inputPrivacyValueDisallowUsers                                    = 0x90110467
	crc_privacyValueAllowContacts                                         = 0xfffe1bac
	crc_privacyValueAllowAll                                              = 0x65427b82
	crc_privacyValueAllowUsers                                            = 0x4d5bbe0c
	crc_privacyValueDisallowContacts                                      = 0xf888fa1a
	crc_privacyValueDisallowAll                                           = 0x8b73e763
	crc_privacyValueDisallowUsers                                         = 0x0c7f49b7
	crc_account_privacyRules                                              = 0x50a04e45
	crc_accountDaysTTL                                                    = 0xb8d0afdf
	crc_updateUserPhone                                                   = 0x12b9417b
	crc_documentAttributeImageSize                                        = 0x6c37c15c
	crc_documentAttributeAnimated                                         = 0x11b58939
	crc_documentAttributeSticker                                          = 0x6319d612
	crc_documentAttributeVideo                                            = 0x0ef02ce6
	crc_documentAttributeAudio                                            = 0x9852f9c6
	crc_documentAttributeFilename                                         = 0x15590068
	crc_messages_stickersNotModified                                      = 0xf1749a22
	crc_messages_stickers                                                 = 0xe4599bbd
	crc_stickerPack                                                       = 0x12b299d4
	crc_messages_allStickersNotModified                                   = 0xe86602c3
	crc_messages_allStickers                                              = 0xedfd405f
	crc_updateReadHistoryInbox                                            = 0x9c974fdf
	crc_updateReadHistoryOutbox                                           = 0x2f2f21bf
	crc_messages_affectedMessages                                         = 0x84d19185
	crc_updateWebPage                                                     = 0x7f891213
	crc_webPageEmpty                                                      = 0xeb1477e8
	crc_webPagePending                                                    = 0xc586da1c
	crc_webPage                                                           = 0xfa64e172
	crc_messageMediaWebPage                                               = 0xa32dd600
	crc_authorization                                                     = 0xad01d61d
	crc_account_authorizations                                            = 0x1250abde
	crc_account_password                                                  = 0xad2641f8
	crc_account_passwordSettings                                          = 0x9a5c33e5
	crc_account_passwordInputSettings                                     = 0xc23727c9
	crc_auth_passwordRecovery                                             = 0x137948a5
	crc_inputMediaVenue                                                   = 0xc13d1c11
	crc_messageMediaVenue                                                 = 0x2ec0533f
	crc_receivedNotifyMessage                                             = 0xa384b779
	crc_chatInviteEmpty                                                   = 0x69df3769
	crc_chatInviteExported                                                = 0xfc2e05bc
	crc_chatInviteAlready                                                 = 0x5a686d7c
	crc_chatInvite                                                        = 0xdfc2f58e
	crc_messageActionChatJoinedByLink                                     = 0xf89cf5e8
	crc_updateReadMessagesContents                                        = 0x68c13933
	crc_inputStickerSetEmpty                                              = 0xffb62b95
	crc_inputStickerSetID                                                 = 0x9de7a269
	crc_inputStickerSetShortName                                          = 0x861cc8a0
	crc_stickerSet                                                        = 0xeeb46f27
	crc_messages_stickerSet                                               = 0xb60a24a6
	crc_user                                                              = 0x938458c1
	crc_botCommand                                                        = 0xc27ac8c7
	crc_botInfo                                                           = 0x98e81d3a
	crc_keyboardButton                                                    = 0xa2fa4880
	crc_keyboardButtonRow                                                 = 0x77608b83
	crc_replyKeyboardHide                                                 = 0xa03e5b85
	crc_replyKeyboardForceReply                                           = 0xf4108aa0
	crc_replyKeyboardMarkup                                               = 0x3502758c
	crc_inputPeerUser                                                     = 0x7b8e7de6
	crc_inputUser                                                         = 0xd8292816
	crc_messageEntityUnknown                                              = 0xbb92ba95
	crc_messageEntityMention                                              = 0xfa04579d
	crc_messageEntityHashtag                                              = 0x6f635b0d
	crc_messageEntityBotCommand                                           = 0x6cef8ac7
	crc_messageEntityUrl                                                  = 0x6ed02538
	crc_messageEntityEmail                                                = 0x64e475c2
	crc_messageEntityBold                                                 = 0xbd610bc9
	crc_messageEntityItalic                                               = 0x826f8b60
	crc_messageEntityCode                                                 = 0x28a20571
	crc_messageEntityPre                                                  = 0x73924be0
	crc_messageEntityTextUrl                                              = 0x76a6d327
	crc_updateShortSentMessage                                            = 0x11f1331c
	crc_inputChannelEmpty                                                 = 0xee8c1e86
	crc_inputChannel                                                      = 0xafeb712e
	crc_peerChannel                                                       = 0xbddde532
	crc_inputPeerChannel                                                  = 0x20adaef8
	crc_channel                                                           = 0xd31a961e
	crc_channelForbidden                                                  = 0x289da732
	crc_contacts_resolvedPeer                                             = 0x7f077ad9
	crc_channelFull                                                       = 0x2d895c74
	crc_messageRange                                                      = 0x0ae30253
	crc_messages_channelMessages                                          = 0x99262e37
	crc_messageActionChannelCreate                                        = 0x95d2ac92
	crc_updateChannelTooLong                                              = 0xeb0467fb
	crc_updateChannel                                                     = 0xb6d45656
	crc_updateNewChannelMessage                                           = 0x62ba04d9
	crc_updateReadChannelInbox                                            = 0x330b5424
	crc_updateDeleteChannelMessages                                       = 0xc37521c9
	crc_updateChannelMessageViews                                         = 0x98a12b4b
	crc_updates_channelDifferenceEmpty                                    = 0x3e11affb
	crc_updates_channelDifferenceTooLong                                  = 0xa4bcc6fe
	crc_updates_channelDifference                                         = 0x2064674e
	crc_channelMessagesFilterEmpty                                        = 0x94d42ee7
	crc_channelMessagesFilter                                             = 0xcd77d957
	crc_channelParticipant                                                = 0x15ebac1d
	crc_channelParticipantSelf                                            = 0xa3289a6d
	crc_channelParticipantCreator                                         = 0x808d15a4
	crc_channelParticipantsRecent                                         = 0xde3f3c79
	crc_channelParticipantsAdmins                                         = 0xb4608969
	crc_channelParticipantsKicked                                         = 0xa3b54985
	crc_channels_channelParticipants                                      = 0xf56ee2a8
	crc_channels_channelParticipant                                       = 0xd0d9b163
	crc_chatParticipantCreator                                            = 0xda13538a
	crc_chatParticipantAdmin                                              = 0xe2d6e436
	crc_updateChatParticipantAdmin                                        = 0xb6901959
	crc_messageActionChatMigrateTo                                        = 0x51bdb021
	crc_messageActionChannelMigrateFrom                                   = 0xb055eaee
	crc_channelParticipantsBots                                           = 0xb0d1865b
	crc_help_termsOfService                                               = 0x780a0310
	crc_updateNewStickerSet                                               = 0x688a30aa
	crc_updateStickerSetsOrder                                            = 0x0bb2d201
	crc_updateStickerSets                                                 = 0x43ae3dec
	crc_foundGif                                                          = 0x162ecc1f
	crc_foundGifCached                                                    = 0x9c750409
	crc_inputMediaGifExternal                                             = 0x4843b0fd
	crc_messages_foundGifs                                                = 0x450a1c0a
	crc_messages_savedGifsNotModified                                     = 0xe8025ca2
	crc_messages_savedGifs                                                = 0x2e0709a5
	crc_updateSavedGifs                                                   = 0x9375341e
	crc_inputBotInlineMessageMediaAuto                                    = 0x3380c786
	crc_inputBotInlineMessageText                                         = 0x3dcd7a87
	crc_inputBotInlineResult                                              = 0x88bf9319
	crc_botInlineMessageMediaAuto                                         = 0x764cf810
	crc_botInlineMessageText                                              = 0x8c7f65e2
	crc_botInlineResult                                                   = 0x11965f3a
	crc_messages_botResults                                               = 0x947ca848
	crc_updateBotInlineQuery                                              = 0x54826690
	crc_updateBotInlineSend                                               = 0x0e48f964
	crc_inputMessagesFilterVoice                                          = 0x50f5c392
	crc_inputMessagesFilterMusic                                          = 0x3751b49e
	crc_inputPrivacyKeyChatInvite                                         = 0xbdfb0426
	crc_privacyKeyChatInvite                                              = 0x500e6dfa
	crc_exportedMessageLink                                               = 0x5dab1af4
	crc_messageFwdHeader                                                  = 0xec338270
	crc_updateEditChannelMessage                                          = 0x1b3f4df7
	crc_updateChannelPinnedMessage                                        = 0x98592475
	crc_messageActionPinMessage                                           = 0x94bd38ed
	crc_auth_codeTypeSms                                                  = 0x72a3158c
	crc_auth_codeTypeCall                                                 = 0x741cd3e3
	crc_auth_codeTypeFlashCall                                            = 0x226ccefb
	crc_auth_sentCodeTypeApp                                              = 0x3dbb5986
	crc_auth_sentCodeTypeSms                                              = 0xc000bba2
	crc_auth_sentCodeTypeCall                                             = 0x5353e5a7
	crc_auth_sentCodeTypeFlashCall                                        = 0xab03c6d9
	crc_keyboardButtonUrl                                                 = 0x258aff05
	crc_keyboardButtonCallback                                            = 0x683a5e46
	crc_keyboardButtonRequestPhone                                        = 0xb16a6c29
	crc_keyboardButtonRequestGeoLocation                                  = 0xfc796b3f
	crc_keyboardButtonSwitchInline                                        = 0x0568a748
	crc_replyInlineMarkup                                                 = 0x48a30254
	crc_messages_botCallbackAnswer                                        = 0x36585ea4
	crc_updateBotCallbackQuery                                            = 0xe73547e1
	crc_messages_messageEditData                                          = 0x26b5dde6
	crc_updateEditMessage                                                 = 0xe40370a3
	crc_inputBotInlineMessageMediaGeo                                     = 0xc1b15d65
	crc_inputBotInlineMessageMediaVenue                                   = 0x417bbf11
	crc_inputBotInlineMessageMediaContact                                 = 0xa6edbffd
	crc_botInlineMessageMediaGeo                                          = 0xb722de65
	crc_botInlineMessageMediaVenue                                        = 0x8a86659c
	crc_botInlineMessageMediaContact                                      = 0x18d1cdc2
	crc_inputBotInlineResultPhoto                                         = 0xa8d864a7
	crc_inputBotInlineResultDocument                                      = 0xfff8fdc4
	crc_botInlineMediaResult                                              = 0x17db940b
	crc_inputBotInlineMessageID                                           = 0x890c3d89
	crc_updateInlineBotCallbackQuery                                      = 0xf9d27a5a
	crc_inlineBotSwitchPM                                                 = 0x3c20629f
	crc_messages_peerDialogs                                              = 0x3371c354
	crc_topPeer                                                           = 0xedcdc05b
	crc_topPeerCategoryBotsPM                                             = 0xab661b5b
	crc_topPeerCategoryBotsInline                                         = 0x148677e2
	crc_topPeerCategoryCorrespondents                                     = 0x0637b7ed
	crc_topPeerCategoryGroups                                             = 0xbd17a14a
	crc_topPeerCategoryChannels                                           = 0x161d9628
	crc_topPeerCategoryPeers                                              = 0xfb834291
	crc_contacts_topPeersNotModified                                      = 0xde266ef5
	crc_contacts_topPeers                                                 = 0x70b772a8
	crc_messageEntityMentionName                                          = 0x352dca58
	crc_inputMessageEntityMentionName                                     = 0x208e68c9
	crc_inputMessagesFilterChatPhotos                                     = 0x3a20ecb8
	crc_updateReadChannelOutbox                                           = 0x25d6c9c7
	crc_updateDraftMessage                                                = 0xee2bb969
	crc_draftMessageEmpty                                                 = 0x1b0c841a
	crc_draftMessage                                                      = 0xfd8e711f
	crc_messageActionHistoryClear                                         = 0x9fbab604
	crc_messages_featuredStickersNotModified                              = 0x04ede3cf
	crc_messages_featuredStickers                                         = 0xf89d88e5
	crc_updateReadFeaturedStickers                                        = 0x571d2742
	crc_messages_recentStickersNotModified                                = 0x0b17f890
	crc_messages_recentStickers                                           = 0x22f3afb3
	crc_updateRecentStickers                                              = 0x9a422c20
	crc_messages_archivedStickers                                         = 0x4fcba9c8
	crc_messages_stickerSetInstallResultSuccess                           = 0x38641628
	crc_messages_stickerSetInstallResultArchive                           = 0x35e410a8
	crc_stickerSetCovered                                                 = 0x6410a5d2
	crc_updateConfig                                                      = 0xa229dd06
	crc_updatePtsChanged                                                  = 0x3354678f
	crc_inputMediaPhotoExternal                                           = 0xe5bbfe1a
	crc_inputMediaDocumentExternal                                        = 0xfb52dc99
	crc_stickerSetMultiCovered                                            = 0x3407e51b
	crc_maskCoords                                                        = 0xaed6dbb2
	crc_documentAttributeHasStickers                                      = 0x9801d2f7
	crc_inputStickeredMediaPhoto                                          = 0x4a992157
	crc_inputStickeredMediaDocument                                       = 0x0438865b
	crc_game                                                              = 0xbdf9653b
	crc_inputBotInlineResultGame                                          = 0x4fa417f2
	crc_inputBotInlineMessageGame                                         = 0x4b425864
	crc_messageMediaGame                                                  = 0xfdb19008
	crc_inputMediaGame                                                    = 0xd33f43f3
	crc_inputGameID                                                       = 0x032c3e77
	crc_inputGameShortName                                                = 0xc331e80a
	crc_keyboardButtonGame                                                = 0x50f41ccf
	crc_messageActionGameScore                                            = 0x92a72876
	crc_highScore                                                         = 0x58fffcd0
	crc_messages_highScores                                               = 0x9a3bfd99
	crc_updates_differenceTooLong                                         = 0x4afe8f6d
	crc_updateChannelWebPage                                              = 0x40771900
	crc_messages_chatsSlice                                               = 0x9cd81144
	crc_textEmpty                                                         = 0xdc3d824f
	crc_textPlain                                                         = 0x744694e0
	crc_textBold                                                          = 0x6724abc4
	crc_textItalic                                                        = 0xd912a59c
	crc_textUnderline                                                     = 0xc12622c4
	crc_textStrike                                                        = 0x9bf8bb95
	crc_textFixed                                                         = 0x6c3f19b9
	crc_textUrl                                                           = 0x3c2884c1
	crc_textEmail                                                         = 0xde5a0dd6
	crc_textConcat                                                        = 0x7e6260d7
	crc_pageBlockUnsupported                                              = 0x13567e8a
	crc_pageBlockTitle                                                    = 0x70abc3fd
	crc_pageBlockSubtitle                                                 = 0x8ffa9a1f
	crc_pageBlockAuthorDate                                               = 0xbaafe5e0
	crc_pageBlockHeader                                                   = 0xbfd064ec
	crc_pageBlockSubheader                                                = 0xf12bb6e1
	crc_pageBlockParagraph                                                = 0x467a0766
	crc_pageBlockPreformatted                                             = 0xc070d93e
	crc_pageBlockFooter                                                   = 0x48870999
	crc_pageBlockDivider                                                  = 0xdb20b188
	crc_pageBlockAnchor                                                   = 0xce0d37b0
	crc_pageBlockList                                                     = 0xe4e88011
	crc_pageBlockBlockquote                                               = 0x263d7c26
	crc_pageBlockPullquote                                                = 0x4f4456d3
	crc_pageBlockPhoto                                                    = 0x1759c560
	crc_pageBlockVideo                                                    = 0x7c8fe7b6
	crc_pageBlockCover                                                    = 0x39f23300
	crc_pageBlockEmbed                                                    = 0xa8718dc5
	crc_pageBlockEmbedPost                                                = 0xf259a80b
	crc_pageBlockCollage                                                  = 0x65a0fa4d
	crc_pageBlockSlideshow                                                = 0x031f9590
	crc_webPageNotModified                                                = 0x85849473
	crc_inputPrivacyKeyPhoneCall                                          = 0xfabadc5f
	crc_privacyKeyPhoneCall                                               = 0x3d662b7b
	crc_sendMessageGamePlayAction                                         = 0xdd6a8f48
	crc_phoneCallDiscardReasonMissed                                      = 0x85e42301
	crc_phoneCallDiscardReasonDisconnect                                  = 0xe095c1a0
	crc_phoneCallDiscardReasonHangup                                      = 0x57adc690
	crc_phoneCallDiscardReasonBusy                                        = 0xfaf7e8c9
	crc_updateDialogPinned                                                = 0x6e6fe51c
	crc_updatePinnedDialogs                                               = 0xfa0f3ca2
	crc_dataJSON                                                          = 0x7d748d04
	crc_updateBotWebhookJSON                                              = 0x8317c0c3
	crc_updateBotWebhookJSONQuery                                         = 0x9b9240a6
	crc_labeledPrice                                                      = 0xcb296bf8
	crc_invoice                                                           = 0xc30aa358
	crc_inputMediaInvoice                                                 = 0xf4e096c3
	crc_paymentCharge                                                     = 0xea02c27e
	crc_messageActionPaymentSentMe                                        = 0x8f31b327
	crc_messageMediaInvoice                                               = 0x84551347
	crc_postAddress                                                       = 0x1e8caaeb
	crc_paymentRequestedInfo                                              = 0x909c3f94
	crc_keyboardButtonBuy                                                 = 0xafd93fbb
	crc_messageActionPaymentSent                                          = 0x40699cd0
	crc_paymentSavedCredentialsCard                                       = 0xcdc27a1f
	crc_webDocument                                                       = 0x1c570ed1
	crc_inputWebDocument                                                  = 0x9bed434d
	crc_inputWebFileLocation                                              = 0xc239d686
	crc_upload_webFile                                                    = 0x21e753bc
	crc_payments_paymentForm                                              = 0x3f56aea3
	crc_payments_validatedRequestedInfo                                   = 0xd1451883
	crc_payments_paymentResult                                            = 0x4e5f810d
	crc_payments_paymentReceipt                                           = 0x500911e1
	crc_payments_savedInfo                                                = 0xfb8fe43c
	crc_inputPaymentCredentialsSaved                                      = 0xc10eb2cf
	crc_inputPaymentCredentials                                           = 0x3417d728
	crc_account_tmpPassword                                               = 0xdb64fd34
	crc_shippingOption                                                    = 0xb6213cdf
	crc_updateBotShippingQuery                                            = 0xe0cdc940
	crc_updateBotPrecheckoutQuery                                         = 0x5d2f3aa9
	crc_inputStickerSetItem                                               = 0xffa0a496
	crc_updatePhoneCall                                                   = 0xab0f6b1e
	crc_inputPhoneCall                                                    = 0x1e36fded
	crc_phoneCallEmpty                                                    = 0x5366c915
	crc_phoneCallWaiting                                                  = 0x1b8f4ad1
	crc_phoneCallRequested                                                = 0x87eabb53
	crc_phoneCallAccepted                                                 = 0x997c454a
	crc_phoneCall                                                         = 0x8742ae7f
	crc_phoneCallDiscarded                                                = 0x50ca4de1
	crc_phoneConnection                                                   = 0x9d4c17c0
	crc_phoneCallProtocol                                                 = 0xa2bb35cb
	crc_phone_phoneCall                                                   = 0xec82e140
	crc_inputMessagesFilterPhoneCalls                                     = 0x80c99768
	crc_messageActionPhoneCall                                            = 0x80e11a7f
	crc_inputMessagesFilterRoundVoice                                     = 0x7a7c17a4
	crc_inputMessagesFilterRoundVideo                                     = 0xb549da53
	crc_sendMessageRecordRoundAction                                      = 0x88f27fbc
	crc_sendMessageUploadRoundAction                                      = 0x243e1c66
	crc_upload_fileCdnRedirect                                            = 0xf18cda44
	crc_upload_cdnFileReuploadNeeded                                      = 0xeea8e46e
	crc_upload_cdnFile                                                    = 0xa99fca4f
	crc_cdnPublicKey                                                      = 0xc982eaba
	crc_cdnConfig                                                         = 0x5725e40a
	crc_pageBlockChannel                                                  = 0xef1751b5
	crc_langPackString                                                    = 0xcad181f6
	crc_langPackStringPluralized                                          = 0x6c47ac9f
	crc_langPackStringDeleted                                             = 0x2979eeb2
	crc_langPackDifference                                                = 0xf385c1f6
	crc_langPackLanguage                                                  = 0xeeca5ce3
	crc_updateLangPackTooLong                                             = 0x46560264
	crc_updateLangPack                                                    = 0x56022f4d
	crc_channelParticipantAdmin                                           = 0xccbebbaf
	crc_channelParticipantBanned                                          = 0x1c0facaf
	crc_channelParticipantsBanned                                         = 0x1427a5e1
	crc_channelParticipantsSearch                                         = 0x0656ac4b
	crc_channelAdminLogEventActionChangeTitle                             = 0xe6dfb825
	crc_channelAdminLogEventActionChangeAbout                             = 0x55188a2e
	crc_channelAdminLogEventActionChangeUsername                          = 0x6a4afc38
	crc_channelAdminLogEventActionChangePhoto                             = 0x434bd2af
	crc_channelAdminLogEventActionToggleInvites                           = 0x1b7907ae
	crc_channelAdminLogEventActionToggleSignatures                        = 0x26ae0971
	crc_channelAdminLogEventActionUpdatePinned                            = 0xe9e82c18
	crc_channelAdminLogEventActionEditMessage                             = 0x709b2405
	crc_channelAdminLogEventActionDeleteMessage                           = 0x42e047bb
	crc_channelAdminLogEventActionParticipantJoin                         = 0x183040d3
	crc_channelAdminLogEventActionParticipantLeave                        = 0xf89777f2
	crc_channelAdminLogEventActionParticipantInvite                       = 0xe31c34d8
	crc_channelAdminLogEventActionParticipantToggleBan                    = 0xe6d83d7e
	crc_channelAdminLogEventActionParticipantToggleAdmin                  = 0xd5676710
	crc_channelAdminLogEvent                                              = 0x3b5a3e40
	crc_channels_adminLogResults                                          = 0xed8af74d
	crc_channelAdminLogEventsFilter                                       = 0xea107ae4
	crc_topPeerCategoryPhoneCalls                                         = 0x1e76a78c
	crc_pageBlockAudio                                                    = 0x804361ea
	crc_popularContact                                                    = 0x5ce14175
	crc_messageActionScreenshotTaken                                      = 0x4792929b
	crc_messages_favedStickersNotModified                                 = 0x9e8fa6d3
	crc_messages_favedStickers                                            = 0xf37f2f16
	crc_updateFavedStickers                                               = 0xe511996d
	crc_updateChannelReadMessagesContents                                 = 0x89893b45
	crc_inputMessagesFilterMyMentions                                     = 0xc1f8e69a
	crc_updateContactsReset                                               = 0x7084a7be
	crc_channelAdminLogEventActionChangeStickerSet                        = 0xb1c3caa7
	crc_messageActionCustomAction                                         = 0xfae69f56
	crc_inputPaymentCredentialsApplePay                                   = 0x0aa1c39f
	crc_inputPaymentCredentialsAndroidPay                                 = 0xca05d50e
	crc_inputMessagesFilterGeo                                            = 0xe7026d0d
	crc_inputMessagesFilterContacts                                       = 0xe062db83
	crc_updateChannelAvailableMessages                                    = 0x70db6837
	crc_channelAdminLogEventActionTogglePreHistoryHidden                  = 0x5f5c95f1
	crc_inputMediaGeoLive                                                 = 0xce4e82fd
	crc_messageMediaGeoLive                                               = 0x7c3c2609
	crc_recentMeUrlUnknown                                                = 0x46e1d13d
	crc_recentMeUrlUser                                                   = 0x8dbc3336
	crc_recentMeUrlChat                                                   = 0xa01b22f9
	crc_recentMeUrlChatInvite                                             = 0xeb49081d
	crc_recentMeUrlStickerSet                                             = 0xbc0a57dc
	crc_help_recentMeUrls                                                 = 0x0e0310d7
	crc_channels_channelParticipantsNotModified                           = 0xf0173fe9
	crc_messages_messagesNotModified                                      = 0x74535f21
	crc_inputSingleMedia                                                  = 0x1cc6e91f
	crc_webAuthorization                                                  = 0xcac943f2
	crc_account_webAuthorizations                                         = 0xed56c9fc
	crc_inputMessageID                                                    = 0xa676a322
	crc_inputMessageReplyTo                                               = 0xbad88395
	crc_inputMessagePinned                                                = 0x86872538
	crc_messageEntityPhone                                                = 0x9b69e34b
	crc_messageEntityCashtag                                              = 0x4c4e743f
	crc_messageActionBotAllowed                                           = 0xabe9affe
	crc_inputDialogPeer                                                   = 0xfcaafeb7
	crc_dialogPeer                                                        = 0xe56dbf05
	crc_messages_foundStickerSetsNotModified                              = 0x0d54b65d
	crc_messages_foundStickerSets                                         = 0x5108d648
	crc_fileHash                                                          = 0x6242c773
	crc_webDocumentNoProxy                                                = 0xf9c8bcc6
	crc_inputClientProxy                                                  = 0x75588b3f
	crc_help_proxyDataEmpty                                               = 0xe09e1fb8
	crc_help_proxyDataPromo                                               = 0x2bf7ee23
	crc_help_termsOfServiceUpdateEmpty                                    = 0xe3309f7f
	crc_help_termsOfServiceUpdate                                         = 0x28ecf961
	crc_inputSecureFileUploaded                                           = 0x3334b0f0
	crc_inputSecureFile                                                   = 0x5367e5be
	crc_inputSecureFileLocation                                           = 0xcbc7ee28
	crc_secureFileEmpty                                                   = 0x64199744
	crc_secureFile                                                        = 0xe0277a62
	crc_secureData                                                        = 0x8aeabec3
	crc_securePlainPhone                                                  = 0x7d6099dd
	crc_securePlainEmail                                                  = 0x21ec5a5f
	crc_secureValueTypePersonalDetails                                    = 0x9d2a81e3
	crc_secureValueTypePassport                                           = 0x3dac6a00
	crc_secureValueTypeDriverLicense                                      = 0x06e425c4
	crc_secureValueTypeIdentityCard                                       = 0xa0d0744b
	crc_secureValueTypeInternalPassport                                   = 0x99a48f23
	crc_secureValueTypeAddress                                            = 0xcbe31e26
	crc_secureValueTypeUtilityBill                                        = 0xfc36954e
	crc_secureValueTypeBankStatement                                      = 0x89137c0d
	crc_secureValueTypeRentalAgreement                                    = 0x8b883488
	crc_secureValueTypePassportRegistration                               = 0x99e3806a
	crc_secureValueTypeTemporaryRegistration                              = 0xea02ec33
	crc_secureValueTypePhone                                              = 0xb320aadb
	crc_secureValueTypeEmail                                              = 0x8e3ca7ee
	crc_secureValue                                                       = 0x187fa0ca
	crc_inputSecureValue                                                  = 0xdb21d0a7
	crc_secureValueHash                                                   = 0xed1ecdb0
	crc_secureValueErrorData                                              = 0xe8a40bd9
	crc_secureValueErrorFrontSide                                         = 0x00be3dfa
	crc_secureValueErrorReverseSide                                       = 0x868a2aa5
	crc_secureValueErrorSelfie                                            = 0xe537ced6
	crc_secureValueErrorFile                                              = 0x7a700873
	crc_secureValueErrorFiles                                             = 0x666220e9
	crc_secureCredentialsEncrypted                                        = 0x33f0ea47
	crc_account_authorizationForm                                         = 0xad2e1cd8
	crc_account_sentEmailCode                                             = 0x811f854f
	crc_messageActionSecureValuesSentMe                                   = 0x1b287353
	crc_messageActionSecureValuesSent                                     = 0xd95c6154
	crc_help_deepLinkInfoEmpty                                            = 0x66afa166
	crc_help_deepLinkInfo                                                 = 0x6a4ee832
	crc_savedPhoneContact                                                 = 0x1142bd56
	crc_account_takeout                                                   = 0x4dba4501
	crc_inputTakeoutFileLocation                                          = 0x29be5899
	crc_updateDialogUnreadMark                                            = 0xe16459c3
	crc_messages_dialogsNotModified                                       = 0xf0e3e596
	crc_inputWebFileGeoPointLocation                                      = 0x9f2221c9
	crc_contacts_topPeersDisabled                                         = 0xb52c939d
	crc_inputReportReasonCopyright                                        = 0x9b89f93a
	crc_passwordKdfAlgoUnknown                                            = 0xd45ab096
	crc_securePasswordKdfAlgoUnknown                                      = 0x004a8537
	crc_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000                   = 0xbbf2dda0
	crc_securePasswordKdfAlgoSHA512                                       = 0x86471d92
	crc_secureSecretSettings                                              = 0x1527bcac
	crc_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow = 0x3a912d4a
	crc_inputCheckPasswordEmpty                                           = 0x9880f658
	crc_inputCheckPasswordSRP                                             = 0xd27ff082
	crc_secureValueError                                                  = 0x869d758f
	crc_secureValueErrorTranslationFile                                   = 0xa1144770
	crc_secureValueErrorTranslationFiles                                  = 0x34636dd8
	crc_secureRequiredType                                                = 0x829d99da
	crc_secureRequiredTypeOneOf                                           = 0x027477b4
	crc_help_passportConfigNotModified                                    = 0xbfb9f457
	crc_help_passportConfig                                               = 0xa098d6af
	crc_inputAppEvent                                                     = 0x1d1b1245
	crc_jsonObjectValue                                                   = 0xc0de1bd9
	crc_jsonNull                                                          = 0x3f6d7b68
	crc_jsonBool                                                          = 0xc7345e6a
	crc_jsonNumber                                                        = 0x2be0dfa4
	crc_jsonString                                                        = 0xb71e767a
	crc_jsonArray                                                         = 0xf7444763
	crc_jsonObject                                                        = 0x99c1d49d
	crc_updateUserPinnedMessage                                           = 0x4c43da18
	crc_updateChatPinnedMessage                                           = 0xe10db349
	crc_inputNotifyBroadcasts                                             = 0xb1db7c7e
	crc_notifyBroadcasts                                                  = 0xd612e8ef
	crc_textSubscript                                                     = 0xed6a8504
	crc_textSuperscript                                                   = 0xc7fb5e01
	crc_textMarked                                                        = 0x034b8621
	crc_textPhone                                                         = 0x1ccb966a
	crc_textImage                                                         = 0x081ccf4f
	crc_pageBlockKicker                                                   = 0x1e148390
	crc_pageTableCell                                                     = 0x34566b6a
	crc_pageTableRow                                                      = 0xe0c0c5e5
	crc_pageBlockTable                                                    = 0xbf4dea82
	crc_pageCaption                                                       = 0x6f747657
	crc_pageListItemText                                                  = 0xb92fb6cd
	crc_pageListItemBlocks                                                = 0x25e073fc
	crc_pageListOrderedItemText                                           = 0x5e068047
	crc_pageListOrderedItemBlocks                                         = 0x98dd8936
	crc_pageBlockOrderedList                                              = 0x9a8ae1e1
	crc_pageBlockDetails                                                  = 0x76768bed
	crc_pageRelatedArticle                                                = 0xb390dc08
	crc_pageBlockRelatedArticles                                          = 0x16115a96
	crc_pageBlockMap                                                      = 0xa44f3ef6
	crc_page                                                              = 0xae891bec
	crc_inputPrivacyKeyPhoneP2P                                           = 0xdb9e70d2
	crc_privacyKeyPhoneP2P                                                = 0x39491cc8
	crc_textAnchor                                                        = 0x35553762
	crc_help_supportName                                                  = 0x8c05f1c9
	crc_help_userInfoEmpty                                                = 0xf3ae2eed
	crc_help_userInfo                                                     = 0x01eb3758
	crc_messageActionContactSignUp                                        = 0xf3f25f76
	crc_updateMessagePoll                                                 = 0xaca1657b
	crc_pollAnswer                                                        = 0x6ca9c2e9
	crc_poll                                                              = 0xd5529d06
	crc_pollAnswerVoters                                                  = 0x3b6ddad2
	crc_pollResults                                                       = 0x5755785a
	crc_inputMediaPoll                                                    = 0x06b3765b
	crc_messageMediaPoll                                                  = 0x4bd6e798
	crc_chatOnlines                                                       = 0xf041e250
	crc_statsURL                                                          = 0x47a971e0
	crc_photoStrippedSize                                                 = 0xe0b0bc2e
	crc_chatAdminRights                                                   = 0x5fb224d5
	crc_chatBannedRights                                                  = 0x9f120418
	crc_updateChatDefaultBannedRights                                     = 0x54c01850
	crc_inputWallPaper                                                    = 0xe630b979
	crc_inputWallPaperSlug                                                = 0x72091c80
	crc_channelParticipantsContacts                                       = 0xbb6ae88d
	crc_channelAdminLogEventActionDefaultBannedRights                     = 0x2df5fc0a
	crc_channelAdminLogEventActionStopPoll                                = 0x8f079643
	crc_account_wallPapersNotModified                                     = 0x1c199183
	crc_account_wallPapers                                                = 0x702b65a9
	crc_codeSettings                                                      = 0xdebebe83
	crc_wallPaperSettings                                                 = 0xa12f40b8
	crc_autoDownloadSettings                                              = 0xd246fd47
	crc_account_autoDownloadSettings                                      = 0x63cacf26
	crc_emojiKeyword                                                      = 0xd5b3b9f9
	crc_emojiKeywordDeleted                                               = 0x236df622
	crc_emojiKeywordsDifference                                           = 0x5cc761bd
	crc_emojiURL                                                          = 0xa575739d
	crc_emojiLanguage                                                     = 0xb3fb5361
	crc_inputPrivacyKeyForwards                                           = 0xa4dd4c08
	crc_privacyKeyForwards                                                = 0x69ec56a3
	crc_inputPrivacyKeyProfilePhoto                                       = 0x5719bacc
	crc_privacyKeyProfilePhoto                                            = 0x96151fed
	crc_fileLocationToBeDeprecated                                        = 0xbc7fc6cd
	crc_inputPhotoFileLocation                                            = 0x40181ffe
	crc_inputPhotoLegacyFileLocation                                      = 0xd83466f3
	crc_inputPeerPhotoFileLocation                                        = 0x27d69997
	crc_inputStickerSetThumb                                              = 0x0dbaeae9
	crc_folder                                                            = 0xff544e65
	crc_dialogFolder                                                      = 0x71bd134c
	crc_inputDialogPeerFolder                                             = 0x64600527
	crc_dialogPeerFolder                                                  = 0x514519e2
	crc_inputFolderPeer                                                   = 0xfbd2c296
	crc_folderPeer                                                        = 0xe9baa668
	crc_updateFolderPeers                                                 = 0x19360dc0
	crc_inputUserFromMessage                                              = 0x2d117597
	crc_inputChannelFromMessage                                           = 0x2a286531
	crc_inputPeerUserFromMessage                                          = 0x17bae2e6
	crc_inputPeerChannelFromMessage                                       = 0x9c95f7bb
	crc_inputPrivacyKeyPhoneNumber                                        = 0x0352dafa
	crc_privacyKeyPhoneNumber                                             = 0xd19ae46d
	crc_topPeerCategoryForwardUsers                                       = 0xa8406ca9
	crc_topPeerCategoryForwardChats                                       = 0xfbeec0f0
	crc_channelAdminLogEventActionChangeLinkedChat                        = 0xa26f881b
	crc_messages_searchCounter                                            = 0xe844ebff
	crc_keyboardButtonUrlAuth                                             = 0x10b78d29
	crc_inputKeyboardButtonUrlAuth                                        = 0xd02e7fd4
	crc_urlAuthResultRequest                                              = 0x92d33a0e
	crc_urlAuthResultAccepted                                             = 0x8f8c0e4e
	crc_urlAuthResultDefault                                              = 0xa9d6db1f
	crc_inputPrivacyValueAllowChatParticipants                            = 0x4c81c1ba
	crc_inputPrivacyValueDisallowChatParticipants                         = 0xd82363af
	crc_privacyValueAllowChatParticipants                                 = 0x18be796b
	crc_privacyValueDisallowChatParticipants                              = 0xacae0690
	crc_messageEntityUnderline                                            = 0x9c4e7e8b
	crc_messageEntityStrike                                               = 0xbf0693d4
	crc_messageEntityBlockquote                                           = 0x020df5d0
	crc_updatePeerSettings                                                = 0x6a7e7366
	crc_channelLocationEmpty                                              = 0xbfb5ad8b
	crc_channelLocation                                                   = 0x209b82db
	crc_peerLocated                                                       = 0xca461b5d
	crc_updatePeerLocated                                                 = 0xb4afcfb0
	crc_channelAdminLogEventActionChangeLocation                          = 0x0e6b76ae
	crc_inputReportReasonGeoIrrelevant                                    = 0xdbd4feed
	crc_channelAdminLogEventActionToggleSlowMode                          = 0x53909779
	crc_auth_authorizationSignUpRequired                                  = 0x44747e9a
	crc_payments_paymentVerificationNeeded                                = 0xd8411139
	crc_inputStickerSetAnimatedEmoji                                      = 0x028703c8
	crc_updateNewScheduledMessage                                         = 0x39a51dfb
	crc_updateDeleteScheduledMessages                                     = 0x90866cee
	crc_restrictionReason                                                 = 0xd072acb4
	crc_inputTheme                                                        = 0x3c5693e9
	crc_inputThemeSlug                                                    = 0xf5890df1
	crc_themeDocumentNotModified                                          = 0x483d270c
	crc_theme                                                             = 0xf7d90ce0
	crc_account_themesNotModified                                         = 0xf41eb622
	crc_account_themes                                                    = 0x7f676421
	crc_updateTheme                                                       = 0x8216fba3
	crc_inputPrivacyKeyAddedByPhone                                       = 0xd1219bdd
	crc_privacyKeyAddedByPhone                                            = 0x42ffd42b
	crc_invokeAfterMsg                                                    = 0xcb9f372d
	crc_invokeAfterMsgs                                                   = 0x3dc4b4f0
	crc_auth_sendCode                                                     = 0xa677244f
	crc_auth_signUp                                                       = 0x80eee427
	crc_auth_signIn                                                       = 0xbcd51581
	crc_auth_logOut                                                       = 0x5717da40
	crc_auth_resetAuthorizations                                          = 0x9fab0d1a
	crc_auth_exportAuthorization                                          = 0xe5bfffcd
	crc_auth_importAuthorization                                          = 0xe3ef9613
	crc_auth_bindTempAuthKey                                              = 0xcdd42a05
	crc_account_registerDevice                                            = 0x68976c6f
	crc_account_unregisterDevice                                          = 0x3076c4bf
	crc_account_updateNotifySettings                                      = 0x84be5b93
	crc_account_getNotifySettings                                         = 0x12b3ad31
	crc_account_resetNotifySettings                                       = 0xdb7e1747
	crc_account_updateProfile                                             = 0x78515775
	crc_account_updateStatus                                              = 0x6628562c
	crc_account_getWallPapers                                             = 0xaabb1763
	crc_account_reportPeer                                                = 0xae189d5f
	crc_users_getUsers                                                    = 0x0d91a548
	crc_users_getFullUser                                                 = 0xca30a5b1
	crc_contacts_getContactIDs                                            = 0x2caa4a42
	crc_contacts_getStatuses                                              = 0xc4a353ee
	crc_contacts_getContacts                                              = 0xc023849f
	crc_contacts_importContacts                                           = 0x2c800be5
	crc_contacts_deleteContacts                                           = 0x096a0e00
	crc_contacts_deleteByPhones                                           = 0x1013fd9e
	crc_contacts_block                                                    = 0x332b49fc
	crc_contacts_unblock                                                  = 0xe54100bd
	crc_contacts_getBlocked                                               = 0xf57c350f
	crc_messages_getMessages                                              = 0x63c66506
	crc_messages_getDialogs                                               = 0xa0ee3b73
	crc_messages_getHistory                                               = 0xdcbb8260
	crc_messages_search                                                   = 0x8614ef68
	crc_messages_readHistory                                              = 0x0e306d3a
	crc_messages_deleteHistory                                            = 0x1c015b09
	crc_messages_deleteMessages                                           = 0xe58e95d2
	crc_messages_receivedMessages                                         = 0x05a954c0
	crc_messages_setTyping                                                = 0xa3825e50
	crc_messages_sendMessage                                              = 0x520c3870
	crc_messages_sendMedia                                                = 0x3491eba9
	crc_messages_forwardMessages                                          = 0xd9fee60e
	crc_messages_reportSpam                                               = 0xcf1592db
	crc_messages_getPeerSettings                                          = 0x3672e09c
	crc_messages_report                                                   = 0xbd82b658
	crc_messages_getChats                                                 = 0x3c6aa187
	crc_messages_getFullChat                                              = 0x3b831c66
	crc_messages_editChatTitle                                            = 0xdc452855
	crc_messages_editChatPhoto                                            = 0xca4c79d8
	crc_messages_addChatUser                                              = 0xf9a0aa09
	crc_messages_deleteChatUser                                           = 0xe0611f16
	crc_messages_createChat                                               = 0x09cb126e
	crc_updates_getState                                                  = 0xedd4882a
	crc_updates_getDifference                                             = 0x25939651
	crc_photos_updateProfilePhoto                                         = 0xf0bb5152
	crc_photos_uploadProfilePhoto                                         = 0x4f32c098
	crc_photos_deletePhotos                                               = 0x87cf7f2f
	crc_upload_saveFilePart                                               = 0xb304a621
	crc_upload_getFile                                                    = 0xb15a9afc
	crc_help_getConfig                                                    = 0xc4f9186b
	crc_help_getNearestDc                                                 = 0x1fb33026
	crc_help_getAppUpdate                                                 = 0x522d5a7d
	crc_help_getInviteText                                                = 0x4d392343
	crc_photos_getUserPhotos                                              = 0x91cd32a8
	crc_messages_getDhConfig                                              = 0x26cf8950
	crc_messages_requestEncryption                                        = 0xf64daf43
	crc_messages_acceptEncryption                                         = 0x3dbc0415
	crc_messages_discardEncryption                                        = 0xedd923c5
	crc_messages_setEncryptedTyping                                       = 0x791451ed
	crc_messages_readEncryptedHistory                                     = 0x7f4b690a
	crc_messages_sendEncrypted                                            = 0xa9776773
	crc_messages_sendEncryptedFile                                        = 0x9a901b66
	crc_messages_sendEncryptedService                                     = 0x32d439a4
	crc_messages_receivedQueue                                            = 0x55a5bb66
	crc_messages_reportEncryptedSpam                                      = 0x4b0c8c0f
	crc_upload_saveBigFilePart                                            = 0xde7b673d
	crc_initConnection                                                    = 0x785188b8
	crc_help_getSupport                                                   = 0x9cdf08cd
	crc_messages_readMessageContents                                      = 0x36a73f77
	crc_account_checkUsername                                             = 0x2714d86c
	crc_account_updateUsername                                            = 0x3e0bdd7c
	crc_contacts_search                                                   = 0x11f812d8
	crc_account_getPrivacy                                                = 0xdadbc950
	crc_account_setPrivacy                                                = 0xc9f81ce8
	crc_account_deleteAccount                                             = 0x418d4e0b
	crc_account_getAccountTTL                                             = 0x08fc711d
	crc_account_setAccountTTL                                             = 0x2442485e
	crc_invokeWithLayer                                                   = 0xda9b0d0d
	crc_contacts_resolveUsername                                          = 0xf93ccba3
	crc_account_sendChangePhoneCode                                       = 0x82574ae5
	crc_account_changePhone                                               = 0x70c32edb
	crc_messages_getStickers                                              = 0x043d4f2c
	crc_messages_getAllStickers                                           = 0x1c9618b1
	crc_account_updateDeviceLocked                                        = 0x38df3532
	crc_auth_importBotAuthorization                                       = 0x67a3ff2c
	crc_messages_getWebPagePreview                                        = 0x8b68b0cc
	crc_account_getAuthorizations                                         = 0xe320c158
	crc_account_resetAuthorization                                        = 0xdf77f3bc
	crc_account_getPassword                                               = 0x548a30f5
	crc_account_getPasswordSettings                                       = 0x9cd4eaf9
	crc_account_updatePasswordSettings                                    = 0xa59b102f
	crc_auth_checkPassword                                                = 0xd18b4d16
	crc_auth_requestPasswordRecovery                                      = 0xd897bc66
	crc_auth_recoverPassword                                              = 0x4ea56e92
	crc_invokeWithoutUpdates                                              = 0xbf9459b7
	crc_messages_exportChatInvite                                         = 0x0df7534c
	crc_messages_checkChatInvite                                          = 0x3eadb1bb
	crc_messages_importChatInvite                                         = 0x6c50051c
	crc_messages_getStickerSet                                            = 0x2619a90e
	crc_messages_installStickerSet                                        = 0xc78fe460
	crc_messages_uninstallStickerSet                                      = 0xf96e55de
	crc_messages_startBot                                                 = 0xe6df7378
	crc_help_getAppChangelog                                              = 0x9010ef6f
	crc_messages_getMessagesViews                                         = 0xc4c8a55d
	crc_channels_readHistory                                              = 0xcc104937
	crc_channels_deleteMessages                                           = 0x84c1fd4e
	crc_channels_deleteUserHistory                                        = 0xd10dd71b
	crc_channels_reportSpam                                               = 0xfe087810
	crc_channels_getMessages                                              = 0xad8c9a23
	crc_channels_getParticipants                                          = 0x123e05e9
	crc_channels_getParticipant                                           = 0x546dd7a6
	crc_channels_getChannels                                              = 0x0a7f6bbb
	crc_channels_getFullChannel                                           = 0x08736a09
	crc_channels_createChannel                                            = 0x3d5fb10f
	crc_channels_editAdmin                                                = 0xd33c8902
	crc_channels_editTitle                                                = 0x566decd0
	crc_channels_editPhoto                                                = 0xf12e57c9
	crc_channels_checkUsername                                            = 0x10e6bd2c
	crc_channels_updateUsername                                           = 0x3514b3de
	crc_channels_joinChannel                                              = 0x24b524c5
	crc_channels_leaveChannel                                             = 0xf836aa95
	crc_channels_inviteToChannel                                          = 0x199f3a6c
	crc_channels_deleteChannel                                            = 0xc0111fe3
	crc_updates_getChannelDifference                                      = 0x03173d78
	crc_messages_editChatAdmin                                            = 0xa9e69f2e
	crc_messages_migrateChat                                              = 0x15a3b8e3
	crc_messages_searchGlobal                                             = 0xbf7225a4
	crc_messages_reorderStickerSets                                       = 0x78337739
	crc_messages_getDocumentByHash                                        = 0x338e2464
	crc_messages_searchGifs                                               = 0xbf9a776b
	crc_messages_getSavedGifs                                             = 0x83bf3d52
	crc_messages_saveGif                                                  = 0x327a30cb
	crc_messages_getInlineBotResults                                      = 0x514e999d
	crc_messages_setInlineBotResults                                      = 0xeb5ea206
	crc_messages_sendInlineBotResult                                      = 0x220815b0
	crc_channels_exportMessageLink                                        = 0xceb77163
	crc_channels_toggleSignatures                                         = 0x1f69b606
	crc_auth_resendCode                                                   = 0x3ef1a9bf
	crc_auth_cancelCode                                                   = 0x1f040578
	crc_messages_getMessageEditData                                       = 0xfda68d36
	crc_messages_editMessage                                              = 0x48f71778
	crc_messages_editInlineBotMessage                                     = 0x83557dba
	crc_messages_getBotCallbackAnswer                                     = 0x810a9fec
	crc_messages_setBotCallbackAnswer                                     = 0xd58f130a
	crc_contacts_getTopPeers                                              = 0xd4982db5
	crc_contacts_resetTopPeerRating                                       = 0x1ae373ac
	crc_messages_getPeerDialogs                                           = 0xe470bcfd
	crc_messages_saveDraft                                                = 0xbc39e14b
	crc_messages_getAllDrafts                                             = 0x6a3f8d65
	crc_messages_getFeaturedStickers                                      = 0x2dacca4f
	crc_messages_readFeaturedStickers                                     = 0x5b118126
	crc_messages_getRecentStickers                                        = 0x5ea192c9
	crc_messages_saveRecentSticker                                        = 0x392718f8
	crc_messages_clearRecentStickers                                      = 0x8999602d
	crc_messages_getArchivedStickers                                      = 0x57f17692
	crc_account_sendConfirmPhoneCode                                      = 0x1b3faa88
	crc_account_confirmPhone                                              = 0x5f2178c3
	crc_channels_getAdminedPublicChannels                                 = 0xf8b036af
	crc_messages_getMaskStickers                                          = 0x65b8c79f
	crc_messages_getAttachedStickers                                      = 0xcc5b67cc
	crc_auth_dropTempAuthKeys                                             = 0x8e48a188
	crc_messages_setGameScore                                             = 0x8ef8ecc0
	crc_messages_setInlineGameScore                                       = 0x15ad9f64
	crc_messages_getGameHighScores                                        = 0xe822649d
	crc_messages_getInlineGameHighScores                                  = 0x0f635e1b
	crc_messages_getCommonChats                                           = 0x0d0a48c4
	crc_messages_getAllChats                                              = 0xeba80ff0
	crc_help_setBotUpdatesStatus                                          = 0xec22cfcd
	crc_messages_getWebPage                                               = 0x32ca8f91
	crc_messages_toggleDialogPin                                          = 0xa731e257
	crc_messages_reorderPinnedDialogs                                     = 0x3b1adf37
	crc_messages_getPinnedDialogs                                         = 0xd6b94df2
	crc_bots_sendCustomRequest                                            = 0xaa2769ed
	crc_bots_answerWebhookJSONQuery                                       = 0xe6213f4d
	crc_upload_getWebFile                                                 = 0x24e6818d
	crc_payments_getPaymentForm                                           = 0x99f09745
	crc_payments_getPaymentReceipt                                        = 0xa092a980
	crc_payments_validateRequestedInfo                                    = 0x770a8e74
	crc_payments_sendPaymentForm                                          = 0x2b8879b3
	crc_account_getTmpPassword                                            = 0x449e0b51
	crc_payments_getSavedInfo                                             = 0x227d824b
	crc_payments_clearSavedInfo                                           = 0xd83d70c1
	crc_messages_setBotShippingResults                                    = 0xe5f672fa
	crc_messages_setBotPrecheckoutResults                                 = 0x09c2dd95
	crc_stickers_createStickerSet                                         = 0x9bd86e6a
	crc_stickers_removeStickerFromSet                                     = 0xf7760f51
	crc_stickers_changeStickerPosition                                    = 0xffb6d4ca
	crc_stickers_addStickerToSet                                          = 0x8653febe
	crc_messages_uploadMedia                                              = 0x519bc2b1
	crc_phone_getCallConfig                                               = 0x55451fa9
	crc_phone_requestCall                                                 = 0x42ff96ed
	crc_phone_acceptCall                                                  = 0x3bd2b4a0
	crc_phone_confirmCall                                                 = 0x2efe1722
	crc_phone_receivedCall                                                = 0x17d54f61
	crc_phone_discardCall                                                 = 0xb2cbc1c0
	crc_phone_setCallRating                                               = 0x59ead627
	crc_phone_saveCallDebug                                               = 0x277add7e
	crc_upload_getCdnFile                                                 = 0x2000bcc3
	crc_upload_reuploadCdnFile                                            = 0x9b2754a8
	crc_help_getCdnConfig                                                 = 0x52029342
	crc_langpack_getLangPack                                              = 0xf2f2330a
	crc_langpack_getStrings                                               = 0xefea3803
	crc_langpack_getDifference                                            = 0xcd984aa5
	crc_langpack_getLanguages                                             = 0x42c6978f
	crc_channels_editBanned                                               = 0x72796912
	crc_channels_getAdminLog                                              = 0x33ddf480
	crc_upload_getCdnFileHashes                                           = 0x4da54231
	crc_messages_sendScreenshotNotification                               = 0xc97df020
	crc_channels_setStickers                                              = 0xea8ca4f9
	crc_messages_getFavedStickers                                         = 0x21ce0b0e
	crc_messages_faveSticker                                              = 0xb9ffc55b
	crc_channels_readMessageContents                                      = 0xeab5dc38
	crc_contacts_resetSaved                                               = 0x879537f1
	crc_messages_getUnreadMentions                                        = 0x46578472
	crc_channels_deleteHistory                                            = 0xaf369d42
	crc_help_getRecentMeUrls                                              = 0x3dc0f114
	crc_channels_togglePreHistoryHidden                                   = 0xeabbb94c
	crc_messages_readMentions                                             = 0x0f0189d3
	crc_messages_getRecentLocations                                       = 0xbbc45b09
	crc_messages_sendMultiMedia                                           = 0xcc0110cb
	crc_messages_uploadEncryptedFile                                      = 0x5057c497
	crc_account_getWebAuthorizations                                      = 0x182e6d6f
	crc_account_resetWebAuthorization                                     = 0x2d01b9ef
	crc_account_resetWebAuthorizations                                    = 0x682d2594
	crc_messages_searchStickerSets                                        = 0xc2b7d08b
	crc_upload_getFileHashes                                              = 0xc7025931
	crc_help_getProxyData                                                 = 0x3d7758e1
	crc_help_getTermsOfServiceUpdate                                      = 0x2ca51fd1
	crc_help_acceptTermsOfService                                         = 0xee72f79a
	crc_account_getAllSecureValues                                        = 0xb288bc7d
	crc_account_getSecureValue                                            = 0x73665bc2
	crc_account_saveSecureValue                                           = 0x899fe31d
	crc_account_deleteSecureValue                                         = 0xb880bc4b
	crc_users_setSecureValueErrors                                        = 0x90c894b5
	crc_account_getAuthorizationForm                                      = 0xb86ba8e1
	crc_account_acceptAuthorization                                       = 0xe7027c94
	crc_account_sendVerifyPhoneCode                                       = 0xa5a356f9
	crc_account_verifyPhone                                               = 0x4dd3a7f6
	crc_account_sendVerifyEmailCode                                       = 0x7011509f
	crc_account_verifyEmail                                               = 0xecba39db
	crc_help_getDeepLinkInfo                                              = 0x3fedc75f
	crc_contacts_getSaved                                                 = 0x82f1e39f
	crc_channels_getLeftChannels                                          = 0x8341ecc0
	crc_account_initTakeoutSession                                        = 0xf05b4804
	crc_account_finishTakeoutSession                                      = 0x1d2652ee
	crc_messages_getSplitRanges                                           = 0x1cff7e08
	crc_invokeWithMessagesRange                                           = 0x365275f2
	crc_invokeWithTakeout                                                 = 0xaca9fd2e
	crc_messages_markDialogUnread                                         = 0xc286d98f
	crc_messages_getDialogUnreadMarks                                     = 0x22e24e22
	crc_contacts_toggleTopPeers                                           = 0x8514bdda
	crc_messages_clearAllDrafts                                           = 0x7e58ee9c
	crc_help_getAppConfig                                                 = 0x98914110
	crc_help_saveAppLog                                                   = 0x6f02f748
	crc_help_getPassportConfig                                            = 0xc661ad08
	crc_langpack_getLanguage                                              = 0x6a596502
	crc_messages_updatePinnedMessage                                      = 0xd2aaf7ec
	crc_account_confirmPasswordEmail                                      = 0x8fdf1920
	crc_account_resendPasswordEmail                                       = 0x7a7f2a15
	crc_account_cancelPasswordEmail                                       = 0xc1cbd5b6
	crc_help_getSupportName                                               = 0xd360e72c
	crc_help_getUserInfo                                                  = 0x038a08d3
	crc_help_editUserInfo                                                 = 0x66b91b70
	crc_account_getContactSignUpNotification                              = 0x9f07c728
	crc_account_setContactSignUpNotification                              = 0xcff43f61
	crc_account_getNotifyExceptions                                       = 0x53577479
	crc_messages_sendVote                                                 = 0x10ea6184
	crc_messages_getPollResults                                           = 0x73bb643b
	crc_messages_getOnlines                                               = 0x6e2be050
	crc_messages_getStatsURL                                              = 0x812c2ae6
	crc_messages_editChatAbout                                            = 0xdef60797
	crc_messages_editChatDefaultBannedRights                              = 0xa5866b41
	crc_account_getWallPaper                                              = 0xfc8ddbea
	crc_account_uploadWallPaper                                           = 0xdd853661
	crc_account_saveWallPaper                                             = 0x6c5a5b37
	crc_account_installWallPaper                                          = 0xfeed5769
	crc_account_resetWallPapers                                           = 0xbb3b9804
	crc_account_getAutoDownloadSettings                                   = 0x56da0b3f
	crc_account_saveAutoDownloadSettings                                  = 0x76f36233
	crc_messages_getEmojiKeywords                                         = 0x35a0e062
	crc_messages_getEmojiKeywordsDifference                               = 0x1508b6af
	crc_messages_getEmojiKeywordsLanguages                                = 0x4e9963b2
	crc_messages_getEmojiURL                                              = 0xd5b10c26
	crc_folders_editPeerFolders                                           = 0x6847d0ab
	crc_folders_deleteFolder                                              = 0x1c295881
	crc_messages_getSearchCounters                                        = 0x732eef00
	crc_channels_getGroupsForDiscussion                                   = 0xf5dad378
	crc_channels_setDiscussionGroup                                       = 0x40582bb2
	crc_messages_requestUrlAuth                                           = 0xe33f5613
	crc_messages_acceptUrlAuth                                            = 0xf729ea98
	crc_messages_hidePeerSettingsBar                                      = 0x4facb138
	crc_contacts_addContact                                               = 0xe8f463d0
	crc_contacts_acceptContact                                            = 0xf831a20f
	crc_channels_editCreator                                              = 0x8f38cd1f
	crc_contacts_getLocated                                               = 0x0a356056
	crc_channels_editLocation                                             = 0x58e63f6d
	crc_channels_toggleSlowMode                                           = 0xedd49ef0
	crc_messages_getScheduledHistory                                      = 0xe2c2685b
	crc_messages_getScheduledMessages                                     = 0xbdbb0464
	crc_messages_sendScheduledMessages                                    = 0xbd38850a
	crc_messages_deleteScheduledMessages                                  = 0x59ae2b16
	crc_account_uploadTheme                                               = 0x1c3db333
	crc_account_createTheme                                               = 0x2b7ffd7f
	crc_account_updateTheme                                               = 0x3b8ea202
	crc_account_saveTheme                                                 = 0xf257106c
	crc_account_installTheme                                              = 0x7ae43737
	crc_account_getTheme                                                  = 0x8d9d742b
	crc_account_getThemes                                                 = 0x285946f8
)

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_true struct {
}

type TL_error struct {
	code int32
	text string
}

type TL_null struct {
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerChat struct {
	chat_id int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputPhoneContact struct {
	client_id  int64
	phone      string
	first_name string
	last_name  string
}

type TL_inputFile struct {
	id           int64
	parts        int32
	name         string
	md5_checksum string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	flags       TL // #
	file        TL // InputFile
	stickers    TL // flags_0?Vector<InputDocument>
	ttl_seconds TL // flags_1?int
}

type TL_inputMediaPhoto struct {
	flags       TL // #
	id          TL // InputPhoto
	ttl_seconds TL // flags_0?int
}

type TL_inputMediaGeoPoint struct {
	geo_point TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
	vcard        string
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	file TL // InputFile
}

type TL_inputChatPhoto struct {
	id TL // InputPhoto
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	lat  float64
	long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	id             int64
	access_hash    int64
	file_reference []byte
}

type TL_inputFileLocation struct {
	volume_id      int64
	local_id       int32
	secret         int64
	file_reference []byte
}

type TL_peerUser struct {
	user_id int32
}

type TL_peerChat struct {
	chat_id int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_filePdf struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_userEmpty struct {
	id int32
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	photo_id    int64
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
	dc_id       int32
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	expires int32
}

type TL_userStatusOffline struct {
	was_online int32
}

type TL_chatEmpty struct {
	id int32
}

type TL_chat struct {
	flags                 TL // #
	creator               TL // flags_0?true
	kicked                TL // flags_1?true
	left                  TL // flags_2?true
	deactivated           TL // flags_5?true
	id                    int32
	title                 string
	photo                 TL // ChatPhoto
	participants_count    int32
	date                  int32
	version               int32
	migrated_to           TL // flags_6?InputChannel
	admin_rights          TL // flags_14?ChatAdminRights
	default_banned_rights TL // flags_18?ChatBannedRights
}

type TL_chatForbidden struct {
	id    int32
	title string
}

type TL_chatFull struct {
	flags            TL // #
	can_set_username TL // flags_7?true
	has_scheduled    TL // flags_8?true
	id               int32
	about            string
	participants     TL // ChatParticipants
	chat_photo       TL // flags_2?Photo
	notify_settings  TL // PeerNotifySettings
	exported_invite  TL // ExportedChatInvite
	bot_info         TL // flags_3?Vector<BotInfo>
	pinned_msg_id    TL // flags_6?int
	folder_id        TL // flags_11?int
}

type TL_chatParticipant struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_chatParticipantsForbidden struct {
	flags            TL // #
	chat_id          int32
	self_participant TL // flags_0?ChatParticipant
}

type TL_chatParticipants struct {
	chat_id      int32
	participants []TL // ChatParticipant
	version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
	dc_id       int32
}

type TL_messageEmpty struct {
	id int32
}

type TL_message struct {
	flags              TL // #
	out                TL // flags_1?true
	mentioned          TL // flags_4?true
	media_unread       TL // flags_5?true
	silent             TL // flags_13?true
	post               TL // flags_14?true
	from_scheduled     TL // flags_18?true
	legacy             TL // flags_19?true
	edit_hide          TL // flags_21?true
	id                 int32
	from_id            TL // flags_8?int
	to_id              TL // Peer
	fwd_from           TL // flags_2?MessageFwdHeader
	via_bot_id         TL // flags_11?int
	reply_to_msg_id    TL // flags_3?int
	date               int32
	message            string
	media              TL // flags_9?MessageMedia
	reply_markup       TL // flags_6?ReplyMarkup
	entities           TL // flags_7?Vector<MessageEntity>
	views              TL // flags_10?int
	edit_date          TL // flags_15?int
	post_author        TL // flags_16?string
	grouped_id         TL // flags_17?long
	restriction_reason TL // flags_22?Vector<RestrictionReason>
}

type TL_messageService struct {
	flags           TL // #
	out             TL // flags_1?true
	mentioned       TL // flags_4?true
	media_unread    TL // flags_5?true
	silent          TL // flags_13?true
	post            TL // flags_14?true
	legacy          TL // flags_19?true
	id              int32
	from_id         TL // flags_8?int
	to_id           TL // Peer
	reply_to_msg_id TL // flags_3?int
	date            int32
	action          TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	flags       TL // #
	photo       TL // flags_0?Photo
	ttl_seconds TL // flags_2?int
}

type TL_messageMediaGeo struct {
	geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
	vcard        string
	user_id      int32
}

type TL_messageMediaUnsupported struct {
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	title string
	users []int32
}

type TL_messageActionChatEditTitle struct {
	title string
}

type TL_messageActionChatEditPhoto struct {
	photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	users []int32
}

type TL_messageActionChatDeleteUser struct {
	user_id int32
}

type TL_dialog struct {
	flags                 TL // #
	pinned                TL // flags_2?true
	unread_mark           TL // flags_3?true
	peer                  TL // Peer
	top_message           int32
	read_inbox_max_id     int32
	read_outbox_max_id    int32
	unread_count          int32
	unread_mentions_count int32
	notify_settings       TL // PeerNotifySettings
	pts                   TL // flags_0?int
	draft                 TL // flags_1?DraftMessage
	folder_id             TL // flags_4?int
}

type TL_photoEmpty struct {
	id int64
}

type TL_photo struct {
	flags          TL // #
	has_stickers   TL // flags_0?true
	id             int64
	access_hash    int64
	file_reference []byte
	date           int32
	sizes          []TL // PhotoSize
	dc_id          int32
}

type TL_photoSizeEmpty struct {
	_type string
}

type TL_photoSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	size     int32
}

type TL_photoCachedSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	bytes    []byte
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	long        float64
	lat         float64
	access_hash int64
}

type TL_auth_sentCode struct {
	flags           TL // #
	_type           TL // auth_SentCodeType
	phone_code_hash string
	next_type       TL // flags_1?auth_CodeType
	timeout         TL // flags_2?int
}

type TL_auth_authorization struct {
	flags        TL // #
	tmp_sessions TL // flags_0?int
	user         TL // User
}

type TL_auth_exportedAuthorization struct {
	id    int32
	bytes []byte
}

type TL_inputNotifyPeer struct {
	peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputPeerNotifySettings struct {
	flags         TL // #
	show_previews TL // flags_0?Bool
	silent        TL // flags_1?Bool
	mute_until    TL // flags_2?int
	sound         TL // flags_3?string
}

type TL_peerNotifySettings struct {
	flags         TL // #
	show_previews TL // flags_0?Bool
	silent        TL // flags_1?Bool
	mute_until    TL // flags_2?int
	sound         TL // flags_3?string
}

type TL_peerSettings struct {
	flags                   TL // #
	report_spam             TL // flags_0?true
	add_contact             TL // flags_1?true
	block_contact           TL // flags_2?true
	share_contact           TL // flags_3?true
	need_contacts_exception TL // flags_4?true
	report_geo              TL // flags_5?true
}

type TL_wallPaper struct {
	id          int64
	flags       TL // #
	creator     TL // flags_0?true
	defaults    TL // flags_1?true
	pattern     TL // flags_3?true
	dark        TL // flags_4?true
	access_hash int64
	slug        string
	document    TL // Document
	settings    TL // flags_2?WallPaperSettings
}

type TL_inputReportReasonSpam struct {
}

type TL_inputReportReasonViolence struct {
}

type TL_inputReportReasonPornography struct {
}

type TL_inputReportReasonChildAbuse struct {
}

type TL_inputReportReasonOther struct {
	text string
}

type TL_userFull struct {
	flags                 TL // #
	blocked               TL // flags_0?true
	phone_calls_available TL // flags_4?true
	phone_calls_private   TL // flags_5?true
	can_pin_message       TL // flags_7?true
	has_scheduled         TL // flags_12?true
	user                  TL // User
	about                 TL // flags_1?string
	settings              TL // PeerSettings
	profile_photo         TL // flags_2?Photo
	notify_settings       TL // PeerNotifySettings
	bot_info              TL // flags_3?BotInfo
	pinned_msg_id         TL // flags_6?int
	common_chats_count    int32
	folder_id             TL // flags_11?int
}

type TL_contact struct {
	user_id int32
	mutual  TL // Bool
}

type TL_importedContact struct {
	user_id   int32
	client_id int64
}

type TL_contactBlocked struct {
	user_id int32
	date    int32
}

type TL_contactStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_contacts_contactsNotModified struct {
}

type TL_contacts_contacts struct {
	contacts    []TL // Contact
	saved_count int32
	users       []TL // User
}

type TL_contacts_importedContacts struct {
	imported        []TL // ImportedContact
	popular_invites []TL // PopularContact
	retry_contacts  []int64
	users           []TL // User
}

type TL_contacts_blocked struct {
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_contacts_blockedSlice struct {
	count   int32
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_messages_dialogs struct {
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_dialogsSlice struct {
	count    int32
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messages struct {
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messagesSlice struct {
	flags     TL // #
	inexact   TL // flags_1?true
	count     int32
	next_rate TL   // flags_0?int
	messages  []TL // Message
	chats     []TL // Chat
	users     []TL // User
}

type TL_messages_chats struct {
	chats []TL // Chat
}

type TL_messages_chatFull struct {
	full_chat TL   // ChatFull
	chats     []TL // Chat
	users     []TL // User
}

type TL_messages_affectedHistory struct {
	pts       int32
	pts_count int32
	offset    int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterUrl struct {
}

type TL_inputMessagesFilterGif struct {
}

type TL_updateNewMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateMessageID struct {
	id        int32
	random_id int64
}

type TL_updateDeleteMessages struct {
	messages  []int32
	pts       int32
	pts_count int32
}

type TL_updateUserTyping struct {
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	chat_id int32
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_updateUserName struct {
	user_id    int32
	first_name string
	last_name  string
	username   string
}

type TL_updateUserPhoto struct {
	user_id  int32
	date     int32
	photo    TL // UserProfilePhoto
	previous TL // Bool
}

type TL_updates_state struct {
	pts          int32
	qts          int32
	date         int32
	seq          int32
	unread_count int32
}

type TL_updates_differenceEmpty struct {
	date int32
	seq  int32
}

type TL_updates_difference struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	state                  TL   // updates_State
}

type TL_updates_differenceSlice struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	intermediate_state     TL   // updates_State
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	flags           TL // #
	out             TL // flags_1?true
	mentioned       TL // flags_4?true
	media_unread    TL // flags_5?true
	silent          TL // flags_13?true
	id              int32
	user_id         int32
	message         string
	pts             int32
	pts_count       int32
	date            int32
	fwd_from        TL // flags_2?MessageFwdHeader
	via_bot_id      TL // flags_11?int
	reply_to_msg_id TL // flags_3?int
	entities        TL // flags_7?Vector<MessageEntity>
}

type TL_updateShortChatMessage struct {
	flags           TL // #
	out             TL // flags_1?true
	mentioned       TL // flags_4?true
	media_unread    TL // flags_5?true
	silent          TL // flags_13?true
	id              int32
	from_id         int32
	chat_id         int32
	message         string
	pts             int32
	pts_count       int32
	date            int32
	fwd_from        TL // flags_2?MessageFwdHeader
	via_bot_id      TL // flags_11?int
	reply_to_msg_id TL // flags_3?int
	entities        TL // flags_7?Vector<MessageEntity>
}

type TL_updateShort struct {
	update TL // Update
	date   int32
}

type TL_updatesCombined struct {
	updates   []TL // Update
	users     []TL // User
	chats     []TL // Chat
	date      int32
	seq_start int32
	seq       int32
}

type TL_updates struct {
	updates []TL // Update
	users   []TL // User
	chats   []TL // Chat
	date    int32
	seq     int32
}

type TL_photos_photos struct {
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photosSlice struct {
	count  int32
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photo struct {
	photo TL   // Photo
	users []TL // User
}

type TL_upload_file struct {
	_type TL // storage_FileType
	mtime int32
	bytes []byte
}

type TL_dcOption struct {
	flags      TL // #
	ipv6       TL // flags_0?true
	media_only TL // flags_1?true
	tcpo_only  TL // flags_2?true
	cdn        TL // flags_3?true
	static     TL // flags_4?true
	id         int32
	ip_address string
	port       int32
	secret     TL // flags_10?bytes
}

type TL_config struct {
	flags                      TL // #
	phonecalls_enabled         TL // flags_1?true
	default_p2p_contacts       TL // flags_3?true
	preload_featured_stickers  TL // flags_4?true
	ignore_phone_entities      TL // flags_5?true
	revoke_pm_inbox            TL // flags_6?true
	blocked_mode               TL // flags_8?true
	pfs_enabled                TL // flags_13?true
	date                       int32
	expires                    int32
	test_mode                  TL // Bool
	this_dc                    int32
	dc_options                 []TL // DcOption
	dc_txt_domain_name         string
	chat_size_max              int32
	megagroup_size_max         int32
	forwarded_count_max        int32
	online_update_period_ms    int32
	offline_blur_timeout_ms    int32
	offline_idle_timeout_ms    int32
	online_cloud_timeout_ms    int32
	notify_cloud_delay_ms      int32
	notify_default_delay_ms    int32
	push_chat_period_ms        int32
	push_chat_limit            int32
	saved_gifs_limit           int32
	edit_time_limit            int32
	revoke_time_limit          int32
	revoke_pm_time_limit       int32
	rating_e_decay             int32
	stickers_recent_limit      int32
	stickers_faved_limit       int32
	channels_read_media_period int32
	tmp_sessions               TL // flags_0?int
	pinned_dialogs_count_max   int32
	pinned_infolder_count_max  int32
	call_receive_timeout_ms    int32
	call_ring_timeout_ms       int32
	call_connect_timeout_ms    int32
	call_packet_timeout_ms     int32
	me_url_prefix              string
	autoupdate_url_prefix      TL // flags_7?string
	gif_search_username        TL // flags_9?string
	venue_search_username      TL // flags_10?string
	img_search_username        TL // flags_11?string
	static_maps_provider       TL // flags_12?string
	caption_length_max         int32
	message_length_max         int32
	webfile_dc_id              int32
	suggested_lang_code        TL // flags_2?string
	lang_pack_version          TL // flags_2?int
	base_lang_pack_version     TL // flags_2?int
}

type TL_nearestDc struct {
	country    string
	this_dc    int32
	nearest_dc int32
}

type TL_help_appUpdate struct {
	flags        TL // #
	can_not_skip TL // flags_0?true
	id           int32
	version      string
	text         string
	entities     []TL // MessageEntity
	document     TL   // flags_1?Document
	url          TL   // flags_2?string
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	message string
}

type TL_updateNewEncryptedMessage struct {
	message TL // EncryptedMessage
	qts     int32
}

type TL_updateEncryptedChatTyping struct {
	chat_id int32
}

type TL_updateEncryption struct {
	chat TL // EncryptedChat
	date int32
}

type TL_updateEncryptedMessagesRead struct {
	chat_id  int32
	max_date int32
	date     int32
}

type TL_encryptedChatEmpty struct {
	id int32
}

type TL_encryptedChatWaiting struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
}

type TL_encryptedChatRequested struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	g_a            []byte
}

type TL_encryptedChat struct {
	id              int32
	access_hash     int64
	date            int32
	admin_id        int32
	participant_id  int32
	g_a_or_b        []byte
	key_fingerprint int64
}

type TL_encryptedChatDiscarded struct {
	id int32
}

type TL_inputEncryptedChat struct {
	chat_id     int32
	access_hash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	id              int64
	access_hash     int64
	size            int32
	dc_id           int32
	key_fingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	id              int64
	parts           int32
	md5_checksum    string
	key_fingerprint int32
}

type TL_inputEncryptedFile struct {
	id          int64
	access_hash int64
}

type TL_inputEncryptedFileLocation struct {
	id          int64
	access_hash int64
}

type TL_encryptedMessage struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
	file      TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
}

type TL_messages_dhConfigNotModified struct {
	random []byte
}

type TL_messages_dhConfig struct {
	g       int32
	p       []byte
	version int32
	random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	date int32
}

type TL_messages_sentEncryptedFile struct {
	date int32
	file TL // EncryptedFile
}

type TL_inputFileBig struct {
	id    int64
	parts int32
	name  string
}

type TL_inputEncryptedFileBigUploaded struct {
	id              int64
	parts           int32
	key_fingerprint int32
}

type TL_updateChatParticipantAdd struct {
	chat_id    int32
	user_id    int32
	inviter_id int32
	date       int32
	version    int32
}

type TL_updateChatParticipantDelete struct {
	chat_id int32
	user_id int32
	version int32
}

type TL_updateDcOptions struct {
	dc_options []TL // DcOption
}

type TL_inputMediaUploadedDocument struct {
	flags         TL // #
	nosound_video TL // flags_3?true
	file          TL // InputFile
	thumb         TL // flags_2?InputFile
	mime_type     string
	attributes    []TL // DocumentAttribute
	stickers      TL   // flags_0?Vector<InputDocument>
	ttl_seconds   TL   // flags_1?int
}

type TL_inputMediaDocument struct {
	flags       TL // #
	id          TL // InputDocument
	ttl_seconds TL // flags_0?int
}

type TL_messageMediaDocument struct {
	flags       TL // #
	document    TL // flags_0?Document
	ttl_seconds TL // flags_2?int
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	id             int64
	access_hash    int64
	file_reference []byte
}

type TL_inputDocumentFileLocation struct {
	id             int64
	access_hash    int64
	file_reference []byte
	thumb_size     string
}

type TL_documentEmpty struct {
	id int64
}

type TL_document struct {
	flags          TL // #
	id             int64
	access_hash    int64
	file_reference []byte
	date           int32
	mime_type      string
	size           int32
	thumbs         TL // flags_0?Vector<PhotoSize>
	dc_id          int32
	attributes     []TL // DocumentAttribute
}

type TL_help_support struct {
	phone_number string
	user         TL // User
}

type TL_notifyPeer struct {
	peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_notifyChats struct {
}

type TL_updateUserBlocked struct {
	user_id int32
	blocked TL // Bool
}

type TL_updateNotifySettings struct {
	peer            TL // NotifyPeer
	notify_settings TL // PeerNotifySettings
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
	progress int32
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
	progress int32
}

type TL_sendMessageUploadPhotoAction struct {
	progress int32
}

type TL_sendMessageUploadDocumentAction struct {
	progress int32
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_contacts_found struct {
	my_results []TL // Peer
	results    []TL // Peer
	chats      []TL // Chat
	users      []TL // User
}

type TL_updateServiceNotification struct {
	flags      TL // #
	popup      TL // flags_0?true
	inbox_date TL // flags_1?int
	_type      string
	message    string
	media      TL   // MessageMedia
	entities   []TL // MessageEntity
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_updatePrivacy struct {
	key   TL   // PrivacyKey
	rules []TL // PrivacyRule
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	users []TL // InputUser
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	users []int32
}

type TL_account_privacyRules struct {
	rules []TL // PrivacyRule
	chats []TL // Chat
	users []TL // User
}

type TL_accountDaysTTL struct {
	days int32
}

type TL_updateUserPhone struct {
	user_id int32
	phone   string
}

type TL_documentAttributeImageSize struct {
	w int32
	h int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
	flags       TL // #
	mask        TL // flags_1?true
	alt         string
	stickerset  TL // InputStickerSet
	mask_coords TL // flags_0?MaskCoords
}

type TL_documentAttributeVideo struct {
	flags              TL // #
	round_message      TL // flags_0?true
	supports_streaming TL // flags_1?true
	duration           int32
	w                  int32
	h                  int32
}

type TL_documentAttributeAudio struct {
	flags     TL // #
	voice     TL // flags_10?true
	duration  int32
	title     TL // flags_0?string
	performer TL // flags_1?string
	waveform  TL // flags_2?bytes
}

type TL_documentAttributeFilename struct {
	file_name string
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	hash     int32
	stickers []TL // Document
}

type TL_stickerPack struct {
	emoticon  string
	documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	hash int32
	sets []TL // StickerSet
}

type TL_updateReadHistoryInbox struct {
	flags              TL // #
	folder_id          TL // flags_0?int
	peer               TL // Peer
	max_id             int32
	still_unread_count int32
	pts                int32
	pts_count          int32
}

type TL_updateReadHistoryOutbox struct {
	peer      TL // Peer
	max_id    int32
	pts       int32
	pts_count int32
}

type TL_messages_affectedMessages struct {
	pts       int32
	pts_count int32
}

type TL_updateWebPage struct {
	webpage   TL // WebPage
	pts       int32
	pts_count int32
}

type TL_webPageEmpty struct {
	id int64
}

type TL_webPagePending struct {
	id   int64
	date int32
}

type TL_webPage struct {
	flags        TL // #
	id           int64
	url          string
	display_url  string
	hash         int32
	_type        TL // flags_0?string
	site_name    TL // flags_1?string
	title        TL // flags_2?string
	description  TL // flags_3?string
	photo        TL // flags_4?Photo
	embed_url    TL // flags_5?string
	embed_type   TL // flags_5?string
	embed_width  TL // flags_6?int
	embed_height TL // flags_6?int
	duration     TL // flags_7?int
	author       TL // flags_8?string
	document     TL // flags_9?Document
	documents    TL // flags_11?Vector<Document>
	cached_page  TL // flags_10?Page
}

type TL_messageMediaWebPage struct {
	webpage TL // WebPage
}

type TL_authorization struct {
	flags            TL // #
	current          TL // flags_0?true
	official_app     TL // flags_1?true
	password_pending TL // flags_2?true
	hash             int64
	device_model     string
	platform         string
	system_version   string
	api_id           int32
	app_name         string
	app_version      string
	date_created     int32
	date_active      int32
	ip               string
	country          string
	region           string
}

type TL_account_authorizations struct {
	authorizations []TL // Authorization
}

type TL_account_password struct {
	flags                     TL // #
	has_recovery              TL // flags_0?true
	has_secure_values         TL // flags_1?true
	has_password              TL // flags_2?true
	current_algo              TL // flags_2?PasswordKdfAlgo
	srp_B                     TL // flags_2?bytes
	srp_id                    TL // flags_2?long
	hint                      TL // flags_3?string
	email_unconfirmed_pattern TL // flags_4?string
	new_algo                  TL // PasswordKdfAlgo
	new_secure_algo           TL // SecurePasswordKdfAlgo
	secure_random             []byte
}

type TL_account_passwordSettings struct {
	flags           TL // #
	email           TL // flags_0?string
	secure_settings TL // flags_1?SecureSecretSettings
}

type TL_account_passwordInputSettings struct {
	flags               TL // #
	new_algo            TL // flags_0?PasswordKdfAlgo
	new_password_hash   TL // flags_0?bytes
	hint                TL // flags_0?string
	email               TL // flags_1?string
	new_secure_settings TL // flags_2?SecureSecretSettings
}

type TL_auth_passwordRecovery struct {
	email_pattern string
}

type TL_inputMediaVenue struct {
	geo_point  TL // InputGeoPoint
	title      string
	address    string
	provider   string
	venue_id   string
	venue_type string
}

type TL_messageMediaVenue struct {
	geo        TL // GeoPoint
	title      string
	address    string
	provider   string
	venue_id   string
	venue_type string
}

type TL_receivedNotifyMessage struct {
	id    int32
	flags int32
}

type TL_chatInviteEmpty struct {
}

type TL_chatInviteExported struct {
	link string
}

type TL_chatInviteAlready struct {
	chat TL // Chat
}

type TL_chatInvite struct {
	flags              TL // #
	channel            TL // flags_0?true
	broadcast          TL // flags_1?true
	public             TL // flags_2?true
	megagroup          TL // flags_3?true
	title              string
	photo              TL // Photo
	participants_count int32
	participants       TL // flags_4?Vector<User>
}

type TL_messageActionChatJoinedByLink struct {
	inviter_id int32
}

type TL_updateReadMessagesContents struct {
	messages  []int32
	pts       int32
	pts_count int32
}

type TL_inputStickerSetEmpty struct {
}

type TL_inputStickerSetID struct {
	id          int64
	access_hash int64
}

type TL_inputStickerSetShortName struct {
	short_name string
}

type TL_stickerSet struct {
	flags          TL // #
	archived       TL // flags_1?true
	official       TL // flags_2?true
	masks          TL // flags_3?true
	animated       TL // flags_5?true
	installed_date TL // flags_0?int
	id             int64
	access_hash    int64
	title          string
	short_name     string
	thumb          TL // flags_4?PhotoSize
	thumb_dc_id    TL // flags_4?int
	count          int32
	hash           int32
}

type TL_messages_stickerSet struct {
	set       TL   // StickerSet
	packs     []TL // StickerPack
	documents []TL // Document
}

type TL_user struct {
	flags                  TL // #
	self                   TL // flags_10?true
	contact                TL // flags_11?true
	mutual_contact         TL // flags_12?true
	deleted                TL // flags_13?true
	bot                    TL // flags_14?true
	bot_chat_history       TL // flags_15?true
	bot_nochats            TL // flags_16?true
	verified               TL // flags_17?true
	restricted             TL // flags_18?true
	min                    TL // flags_20?true
	bot_inline_geo         TL // flags_21?true
	support                TL // flags_23?true
	scam                   TL // flags_24?true
	id                     int32
	access_hash            TL // flags_0?long
	first_name             TL // flags_1?string
	last_name              TL // flags_2?string
	username               TL // flags_3?string
	phone                  TL // flags_4?string
	photo                  TL // flags_5?UserProfilePhoto
	status                 TL // flags_6?UserStatus
	bot_info_version       TL // flags_14?int
	restriction_reason     TL // flags_18?Vector<RestrictionReason>
	bot_inline_placeholder TL // flags_19?string
	lang_code              TL // flags_22?string
}

type TL_botCommand struct {
	command     string
	description string
}

type TL_botInfo struct {
	user_id     int32
	description string
	commands    []TL // BotCommand
}

type TL_keyboardButton struct {
	text string
}

type TL_keyboardButtonRow struct {
	buttons []TL // KeyboardButton
}

type TL_replyKeyboardHide struct {
	flags     TL // #
	selective TL // flags_2?true
}

type TL_replyKeyboardForceReply struct {
	flags      TL // #
	single_use TL // flags_1?true
	selective  TL // flags_2?true
}

type TL_replyKeyboardMarkup struct {
	flags      TL   // #
	resize     TL   // flags_0?true
	single_use TL   // flags_1?true
	selective  TL   // flags_2?true
	rows       []TL // KeyboardButtonRow
}

type TL_inputPeerUser struct {
	user_id     int32
	access_hash int64
}

type TL_inputUser struct {
	user_id     int32
	access_hash int64
}

type TL_messageEntityUnknown struct {
	offset int32
	length int32
}

type TL_messageEntityMention struct {
	offset int32
	length int32
}

type TL_messageEntityHashtag struct {
	offset int32
	length int32
}

type TL_messageEntityBotCommand struct {
	offset int32
	length int32
}

type TL_messageEntityUrl struct {
	offset int32
	length int32
}

type TL_messageEntityEmail struct {
	offset int32
	length int32
}

type TL_messageEntityBold struct {
	offset int32
	length int32
}

type TL_messageEntityItalic struct {
	offset int32
	length int32
}

type TL_messageEntityCode struct {
	offset int32
	length int32
}

type TL_messageEntityPre struct {
	offset   int32
	length   int32
	language string
}

type TL_messageEntityTextUrl struct {
	offset int32
	length int32
	url    string
}

type TL_updateShortSentMessage struct {
	flags     TL // #
	out       TL // flags_1?true
	id        int32
	pts       int32
	pts_count int32
	date      int32
	media     TL // flags_9?MessageMedia
	entities  TL // flags_7?Vector<MessageEntity>
}

type TL_inputChannelEmpty struct {
}

type TL_inputChannel struct {
	channel_id  int32
	access_hash int64
}

type TL_peerChannel struct {
	channel_id int32
}

type TL_inputPeerChannel struct {
	channel_id  int32
	access_hash int64
}

type TL_channel struct {
	flags                 TL // #
	creator               TL // flags_0?true
	left                  TL // flags_2?true
	broadcast             TL // flags_5?true
	verified              TL // flags_7?true
	megagroup             TL // flags_8?true
	restricted            TL // flags_9?true
	signatures            TL // flags_11?true
	min                   TL // flags_12?true
	scam                  TL // flags_19?true
	has_link              TL // flags_20?true
	has_geo               TL // flags_21?true
	slowmode_enabled      TL // flags_22?true
	id                    int32
	access_hash           TL // flags_13?long
	title                 string
	username              TL // flags_6?string
	photo                 TL // ChatPhoto
	date                  int32
	version               int32
	restriction_reason    TL // flags_9?Vector<RestrictionReason>
	admin_rights          TL // flags_14?ChatAdminRights
	banned_rights         TL // flags_15?ChatBannedRights
	default_banned_rights TL // flags_18?ChatBannedRights
	participants_count    TL // flags_17?int
}

type TL_channelForbidden struct {
	flags       TL // #
	broadcast   TL // flags_5?true
	megagroup   TL // flags_8?true
	id          int32
	access_hash int64
	title       string
	until_date  TL // flags_16?int
}

type TL_contacts_resolvedPeer struct {
	peer  TL   // Peer
	chats []TL // Chat
	users []TL // User
}

type TL_channelFull struct {
	flags                   TL // #
	can_view_participants   TL // flags_3?true
	can_set_username        TL // flags_6?true
	can_set_stickers        TL // flags_7?true
	hidden_prehistory       TL // flags_10?true
	can_view_stats          TL // flags_12?true
	can_set_location        TL // flags_16?true
	has_scheduled           TL // flags_19?true
	id                      int32
	about                   string
	participants_count      TL // flags_0?int
	admins_count            TL // flags_1?int
	kicked_count            TL // flags_2?int
	banned_count            TL // flags_2?int
	online_count            TL // flags_13?int
	read_inbox_max_id       int32
	read_outbox_max_id      int32
	unread_count            int32
	chat_photo              TL   // Photo
	notify_settings         TL   // PeerNotifySettings
	exported_invite         TL   // ExportedChatInvite
	bot_info                []TL // BotInfo
	migrated_from_chat_id   TL   // flags_4?int
	migrated_from_max_id    TL   // flags_4?int
	pinned_msg_id           TL   // flags_5?int
	stickerset              TL   // flags_8?StickerSet
	available_min_id        TL   // flags_9?int
	folder_id               TL   // flags_11?int
	linked_chat_id          TL   // flags_14?int
	location                TL   // flags_15?ChannelLocation
	slowmode_seconds        TL   // flags_17?int
	slowmode_next_send_date TL   // flags_18?int
	pts                     int32
}

type TL_messageRange struct {
	min_id int32
	max_id int32
}

type TL_messages_channelMessages struct {
	flags    TL // #
	inexact  TL // flags_1?true
	pts      int32
	count    int32
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messageActionChannelCreate struct {
	title string
}

type TL_updateChannelTooLong struct {
	flags      TL // #
	channel_id int32
	pts        TL // flags_0?int
}

type TL_updateChannel struct {
	channel_id int32
}

type TL_updateNewChannelMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateReadChannelInbox struct {
	flags              TL // #
	folder_id          TL // flags_0?int
	channel_id         int32
	max_id             int32
	still_unread_count int32
	pts                int32
}

type TL_updateDeleteChannelMessages struct {
	channel_id int32
	messages   []int32
	pts        int32
	pts_count  int32
}

type TL_updateChannelMessageViews struct {
	channel_id int32
	id         int32
	views      int32
}

type TL_updates_channelDifferenceEmpty struct {
	flags   TL // #
	final   TL // flags_0?true
	pts     int32
	timeout TL // flags_1?int
}

type TL_updates_channelDifferenceTooLong struct {
	flags    TL   // #
	final    TL   // flags_0?true
	timeout  TL   // flags_1?int
	dialog   TL   // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_updates_channelDifference struct {
	flags         TL // #
	final         TL // flags_0?true
	pts           int32
	timeout       TL   // flags_1?int
	new_messages  []TL // Message
	other_updates []TL // Update
	chats         []TL // Chat
	users         []TL // User
}

type TL_channelMessagesFilterEmpty struct {
}

type TL_channelMessagesFilter struct {
	flags                TL   // #
	exclude_new_messages TL   // flags_1?true
	ranges               []TL // MessageRange
}

type TL_channelParticipant struct {
	user_id int32
	date    int32
}

type TL_channelParticipantSelf struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_channelParticipantCreator struct {
	flags   TL // #
	user_id int32
	rank    TL // flags_0?string
}

type TL_channelParticipantsRecent struct {
}

type TL_channelParticipantsAdmins struct {
}

type TL_channelParticipantsKicked struct {
	q string
}

type TL_channels_channelParticipants struct {
	count        int32
	participants []TL // ChannelParticipant
	users        []TL // User
}

type TL_channels_channelParticipant struct {
	participant TL   // ChannelParticipant
	users       []TL // User
}

type TL_chatParticipantCreator struct {
	user_id int32
}

type TL_chatParticipantAdmin struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_updateChatParticipantAdmin struct {
	chat_id  int32
	user_id  int32
	is_admin TL // Bool
	version  int32
}

type TL_messageActionChatMigrateTo struct {
	channel_id int32
}

type TL_messageActionChannelMigrateFrom struct {
	title   string
	chat_id int32
}

type TL_channelParticipantsBots struct {
}

type TL_help_termsOfService struct {
	flags           TL // #
	popup           TL // flags_0?true
	id              TL // DataJSON
	text            string
	entities        []TL // MessageEntity
	min_age_confirm TL   // flags_1?int
}

type TL_updateNewStickerSet struct {
	stickerset TL // messages_StickerSet
}

type TL_updateStickerSetsOrder struct {
	flags TL // #
	masks TL // flags_0?true
	order []int64
}

type TL_updateStickerSets struct {
}

type TL_foundGif struct {
	url          string
	thumb_url    string
	content_url  string
	content_type string
	w            int32
	h            int32
}

type TL_foundGifCached struct {
	url      string
	photo    TL // Photo
	document TL // Document
}

type TL_inputMediaGifExternal struct {
	url string
	q   string
}

type TL_messages_foundGifs struct {
	next_offset int32
	results     []TL // FoundGif
}

type TL_messages_savedGifsNotModified struct {
}

type TL_messages_savedGifs struct {
	hash int32
	gifs []TL // Document
}

type TL_updateSavedGifs struct {
}

type TL_inputBotInlineMessageMediaAuto struct {
	flags        TL // #
	message      string
	entities     TL // flags_1?Vector<MessageEntity>
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageText struct {
	flags        TL // #
	no_webpage   TL // flags_0?true
	message      string
	entities     TL // flags_1?Vector<MessageEntity>
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResult struct {
	flags        TL // #
	id           string
	_type        string
	title        TL // flags_1?string
	description  TL // flags_2?string
	url          TL // flags_3?string
	thumb        TL // flags_4?InputWebDocument
	content      TL // flags_5?InputWebDocument
	send_message TL // InputBotInlineMessage
}

type TL_botInlineMessageMediaAuto struct {
	flags        TL // #
	message      string
	entities     TL // flags_1?Vector<MessageEntity>
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageText struct {
	flags        TL // #
	no_webpage   TL // flags_0?true
	message      string
	entities     TL // flags_1?Vector<MessageEntity>
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineResult struct {
	flags        TL // #
	id           string
	_type        string
	title        TL // flags_1?string
	description  TL // flags_2?string
	url          TL // flags_3?string
	thumb        TL // flags_4?WebDocument
	content      TL // flags_5?WebDocument
	send_message TL // BotInlineMessage
}

type TL_messages_botResults struct {
	flags       TL // #
	gallery     TL // flags_0?true
	query_id    int64
	next_offset TL   // flags_1?string
	switch_pm   TL   // flags_2?InlineBotSwitchPM
	results     []TL // BotInlineResult
	cache_time  int32
	users       []TL // User
}

type TL_updateBotInlineQuery struct {
	flags    TL // #
	query_id int64
	user_id  int32
	query    string
	geo      TL // flags_0?GeoPoint
	offset   string
}

type TL_updateBotInlineSend struct {
	flags   TL // #
	user_id int32
	query   string
	geo     TL // flags_0?GeoPoint
	id      string
	msg_id  TL // flags_1?InputBotInlineMessageID
}

type TL_inputMessagesFilterVoice struct {
}

type TL_inputMessagesFilterMusic struct {
}

type TL_inputPrivacyKeyChatInvite struct {
}

type TL_privacyKeyChatInvite struct {
}

type TL_exportedMessageLink struct {
	link string
	html string
}

type TL_messageFwdHeader struct {
	flags             TL // #
	from_id           TL // flags_0?int
	from_name         TL // flags_5?string
	date              int32
	channel_id        TL // flags_1?int
	channel_post      TL // flags_2?int
	post_author       TL // flags_3?string
	saved_from_peer   TL // flags_4?Peer
	saved_from_msg_id TL // flags_4?int
}

type TL_updateEditChannelMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateChannelPinnedMessage struct {
	channel_id int32
	id         int32
}

type TL_messageActionPinMessage struct {
}

type TL_auth_codeTypeSms struct {
}

type TL_auth_codeTypeCall struct {
}

type TL_auth_codeTypeFlashCall struct {
}

type TL_auth_sentCodeTypeApp struct {
	length int32
}

type TL_auth_sentCodeTypeSms struct {
	length int32
}

type TL_auth_sentCodeTypeCall struct {
	length int32
}

type TL_auth_sentCodeTypeFlashCall struct {
	pattern string
}

type TL_keyboardButtonUrl struct {
	text string
	url  string
}

type TL_keyboardButtonCallback struct {
	text string
	data []byte
}

type TL_keyboardButtonRequestPhone struct {
	text string
}

type TL_keyboardButtonRequestGeoLocation struct {
	text string
}

type TL_keyboardButtonSwitchInline struct {
	flags     TL // #
	same_peer TL // flags_0?true
	text      string
	query     string
}

type TL_replyInlineMarkup struct {
	rows []TL // KeyboardButtonRow
}

type TL_messages_botCallbackAnswer struct {
	flags      TL // #
	alert      TL // flags_1?true
	has_url    TL // flags_3?true
	native_ui  TL // flags_4?true
	message    TL // flags_0?string
	url        TL // flags_2?string
	cache_time int32
}

type TL_updateBotCallbackQuery struct {
	flags           TL // #
	query_id        int64
	user_id         int32
	peer            TL // Peer
	msg_id          int32
	chat_instance   int64
	data            TL // flags_0?bytes
	game_short_name TL // flags_1?string
}

type TL_messages_messageEditData struct {
	flags   TL // #
	caption TL // flags_0?true
}

type TL_updateEditMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_inputBotInlineMessageMediaGeo struct {
	flags        TL // #
	geo_point    TL // InputGeoPoint
	period       int32
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaVenue struct {
	flags        TL // #
	geo_point    TL // InputGeoPoint
	title        string
	address      string
	provider     string
	venue_id     string
	venue_type   string
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaContact struct {
	flags        TL // #
	phone_number string
	first_name   string
	last_name    string
	vcard        string
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaGeo struct {
	flags        TL // #
	geo          TL // GeoPoint
	period       int32
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaVenue struct {
	flags        TL // #
	geo          TL // GeoPoint
	title        string
	address      string
	provider     string
	venue_id     string
	venue_type   string
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaContact struct {
	flags        TL // #
	phone_number string
	first_name   string
	last_name    string
	vcard        string
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResultPhoto struct {
	id           string
	_type        string
	photo        TL // InputPhoto
	send_message TL // InputBotInlineMessage
}

type TL_inputBotInlineResultDocument struct {
	flags        TL // #
	id           string
	_type        string
	title        TL // flags_1?string
	description  TL // flags_2?string
	document     TL // InputDocument
	send_message TL // InputBotInlineMessage
}

type TL_botInlineMediaResult struct {
	flags        TL // #
	id           string
	_type        string
	photo        TL // flags_0?Photo
	document     TL // flags_1?Document
	title        TL // flags_2?string
	description  TL // flags_3?string
	send_message TL // BotInlineMessage
}

type TL_inputBotInlineMessageID struct {
	dc_id       int32
	id          int64
	access_hash int64
}

type TL_updateInlineBotCallbackQuery struct {
	flags           TL // #
	query_id        int64
	user_id         int32
	msg_id          TL // InputBotInlineMessageID
	chat_instance   int64
	data            TL // flags_0?bytes
	game_short_name TL // flags_1?string
}

type TL_inlineBotSwitchPM struct {
	text        string
	start_param string
}

type TL_messages_peerDialogs struct {
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
	state    TL   // updates_State
}

type TL_topPeer struct {
	peer   TL // Peer
	rating float64
}

type TL_topPeerCategoryBotsPM struct {
}

type TL_topPeerCategoryBotsInline struct {
}

type TL_topPeerCategoryCorrespondents struct {
}

type TL_topPeerCategoryGroups struct {
}

type TL_topPeerCategoryChannels struct {
}

type TL_topPeerCategoryPeers struct {
	category TL // TopPeerCategory
	count    int32
	peers    []TL // TopPeer
}

type TL_contacts_topPeersNotModified struct {
}

type TL_contacts_topPeers struct {
	categories []TL // TopPeerCategoryPeers
	chats      []TL // Chat
	users      []TL // User
}

type TL_messageEntityMentionName struct {
	offset  int32
	length  int32
	user_id int32
}

type TL_inputMessageEntityMentionName struct {
	offset  int32
	length  int32
	user_id TL // InputUser
}

type TL_inputMessagesFilterChatPhotos struct {
}

type TL_updateReadChannelOutbox struct {
	channel_id int32
	max_id     int32
}

type TL_updateDraftMessage struct {
	peer  TL // Peer
	draft TL // DraftMessage
}

type TL_draftMessageEmpty struct {
	flags TL // #
	date  TL // flags_0?int
}

type TL_draftMessage struct {
	flags           TL // #
	no_webpage      TL // flags_1?true
	reply_to_msg_id TL // flags_0?int
	message         string
	entities        TL // flags_3?Vector<MessageEntity>
	date            int32
}

type TL_messageActionHistoryClear struct {
}

type TL_messages_featuredStickersNotModified struct {
}

type TL_messages_featuredStickers struct {
	hash   int32
	sets   []TL // StickerSetCovered
	unread []int64
}

type TL_updateReadFeaturedStickers struct {
}

type TL_messages_recentStickersNotModified struct {
}

type TL_messages_recentStickers struct {
	hash     int32
	packs    []TL // StickerPack
	stickers []TL // Document
	dates    []int32
}

type TL_updateRecentStickers struct {
}

type TL_messages_archivedStickers struct {
	count int32
	sets  []TL // StickerSetCovered
}

type TL_messages_stickerSetInstallResultSuccess struct {
}

type TL_messages_stickerSetInstallResultArchive struct {
	sets []TL // StickerSetCovered
}

type TL_stickerSetCovered struct {
	set   TL // StickerSet
	cover TL // Document
}

type TL_updateConfig struct {
}

type TL_updatePtsChanged struct {
}

type TL_inputMediaPhotoExternal struct {
	flags       TL // #
	url         string
	ttl_seconds TL // flags_0?int
}

type TL_inputMediaDocumentExternal struct {
	flags       TL // #
	url         string
	ttl_seconds TL // flags_0?int
}

type TL_stickerSetMultiCovered struct {
	set    TL   // StickerSet
	covers []TL // Document
}

type TL_maskCoords struct {
	n    int32
	x    float64
	y    float64
	zoom float64
}

type TL_documentAttributeHasStickers struct {
}

type TL_inputStickeredMediaPhoto struct {
	id TL // InputPhoto
}

type TL_inputStickeredMediaDocument struct {
	id TL // InputDocument
}

type TL_game struct {
	flags       TL // #
	id          int64
	access_hash int64
	short_name  string
	title       string
	description string
	photo       TL // Photo
	document    TL // flags_0?Document
}

type TL_inputBotInlineResultGame struct {
	id           string
	short_name   string
	send_message TL // InputBotInlineMessage
}

type TL_inputBotInlineMessageGame struct {
	flags        TL // #
	reply_markup TL // flags_2?ReplyMarkup
}

type TL_messageMediaGame struct {
	game TL // Game
}

type TL_inputMediaGame struct {
	id TL // InputGame
}

type TL_inputGameID struct {
	id          int64
	access_hash int64
}

type TL_inputGameShortName struct {
	bot_id     TL // InputUser
	short_name string
}

type TL_keyboardButtonGame struct {
	text string
}

type TL_messageActionGameScore struct {
	game_id int64
	score   int32
}

type TL_highScore struct {
	pos     int32
	user_id int32
	score   int32
}

type TL_messages_highScores struct {
	scores []TL // HighScore
	users  []TL // User
}

type TL_updates_differenceTooLong struct {
	pts int32
}

type TL_updateChannelWebPage struct {
	channel_id int32
	webpage    TL // WebPage
	pts        int32
	pts_count  int32
}

type TL_messages_chatsSlice struct {
	count int32
	chats []TL // Chat
}

type TL_textEmpty struct {
}

type TL_textPlain struct {
	text string
}

type TL_textBold struct {
	text TL // RichText
}

type TL_textItalic struct {
	text TL // RichText
}

type TL_textUnderline struct {
	text TL // RichText
}

type TL_textStrike struct {
	text TL // RichText
}

type TL_textFixed struct {
	text TL // RichText
}

type TL_textUrl struct {
	text       TL // RichText
	url        string
	webpage_id int64
}

type TL_textEmail struct {
	text  TL // RichText
	email string
}

type TL_textConcat struct {
	texts []TL // RichText
}

type TL_pageBlockUnsupported struct {
}

type TL_pageBlockTitle struct {
	text TL // RichText
}

type TL_pageBlockSubtitle struct {
	text TL // RichText
}

type TL_pageBlockAuthorDate struct {
	author         TL // RichText
	published_date int32
}

type TL_pageBlockHeader struct {
	text TL // RichText
}

type TL_pageBlockSubheader struct {
	text TL // RichText
}

type TL_pageBlockParagraph struct {
	text TL // RichText
}

type TL_pageBlockPreformatted struct {
	text     TL // RichText
	language string
}

type TL_pageBlockFooter struct {
	text TL // RichText
}

type TL_pageBlockDivider struct {
}

type TL_pageBlockAnchor struct {
	name string
}

type TL_pageBlockList struct {
	items []TL // PageListItem
}

type TL_pageBlockBlockquote struct {
	text    TL // RichText
	caption TL // RichText
}

type TL_pageBlockPullquote struct {
	text    TL // RichText
	caption TL // RichText
}

type TL_pageBlockPhoto struct {
	flags      TL // #
	photo_id   int64
	caption    TL // PageCaption
	url        TL // flags_0?string
	webpage_id TL // flags_0?long
}

type TL_pageBlockVideo struct {
	flags    TL // #
	autoplay TL // flags_0?true
	loop     TL // flags_1?true
	video_id int64
	caption  TL // PageCaption
}

type TL_pageBlockCover struct {
	cover TL // PageBlock
}

type TL_pageBlockEmbed struct {
	flags           TL // #
	full_width      TL // flags_0?true
	allow_scrolling TL // flags_3?true
	url             TL // flags_1?string
	html            TL // flags_2?string
	poster_photo_id TL // flags_4?long
	w               TL // flags_5?int
	h               TL // flags_5?int
	caption         TL // PageCaption
}

type TL_pageBlockEmbedPost struct {
	url             string
	webpage_id      int64
	author_photo_id int64
	author          string
	date            int32
	blocks          []TL // PageBlock
	caption         TL   // PageCaption
}

type TL_pageBlockCollage struct {
	items   []TL // PageBlock
	caption TL   // PageCaption
}

type TL_pageBlockSlideshow struct {
	items   []TL // PageBlock
	caption TL   // PageCaption
}

type TL_webPageNotModified struct {
}

type TL_inputPrivacyKeyPhoneCall struct {
}

type TL_privacyKeyPhoneCall struct {
}

type TL_sendMessageGamePlayAction struct {
}

type TL_phoneCallDiscardReasonMissed struct {
}

type TL_phoneCallDiscardReasonDisconnect struct {
}

type TL_phoneCallDiscardReasonHangup struct {
}

type TL_phoneCallDiscardReasonBusy struct {
}

type TL_updateDialogPinned struct {
	flags     TL // #
	pinned    TL // flags_0?true
	folder_id TL // flags_1?int
	peer      TL // DialogPeer
}

type TL_updatePinnedDialogs struct {
	flags     TL // #
	folder_id TL // flags_1?int
	order     TL // flags_0?Vector<DialogPeer>
}

type TL_dataJSON struct {
	data string
}

type TL_updateBotWebhookJSON struct {
	data TL // DataJSON
}

type TL_updateBotWebhookJSONQuery struct {
	query_id int64
	data     TL // DataJSON
	timeout  int32
}

type TL_labeledPrice struct {
	label  string
	amount int64
}

type TL_invoice struct {
	flags                      TL // #
	test                       TL // flags_0?true
	name_requested             TL // flags_1?true
	phone_requested            TL // flags_2?true
	email_requested            TL // flags_3?true
	shipping_address_requested TL // flags_4?true
	flexible                   TL // flags_5?true
	phone_to_provider          TL // flags_6?true
	email_to_provider          TL // flags_7?true
	currency                   string
	prices                     []TL // LabeledPrice
}

type TL_inputMediaInvoice struct {
	flags         TL // #
	title         string
	description   string
	photo         TL // flags_0?InputWebDocument
	invoice       TL // Invoice
	payload       []byte
	provider      string
	provider_data TL // DataJSON
	start_param   string
}

type TL_paymentCharge struct {
	id                 string
	provider_charge_id string
}

type TL_messageActionPaymentSentMe struct {
	flags              TL // #
	currency           string
	total_amount       int64
	payload            []byte
	info               TL // flags_0?PaymentRequestedInfo
	shipping_option_id TL // flags_1?string
	charge             TL // PaymentCharge
}

type TL_messageMediaInvoice struct {
	flags                      TL // #
	shipping_address_requested TL // flags_1?true
	test                       TL // flags_3?true
	title                      string
	description                string
	photo                      TL // flags_0?WebDocument
	receipt_msg_id             TL // flags_2?int
	currency                   string
	total_amount               int64
	start_param                string
}

type TL_postAddress struct {
	street_line1 string
	street_line2 string
	city         string
	state        string
	country_iso2 string
	post_code    string
}

type TL_paymentRequestedInfo struct {
	flags            TL // #
	name             TL // flags_0?string
	phone            TL // flags_1?string
	email            TL // flags_2?string
	shipping_address TL // flags_3?PostAddress
}

type TL_keyboardButtonBuy struct {
	text string
}

type TL_messageActionPaymentSent struct {
	currency     string
	total_amount int64
}

type TL_paymentSavedCredentialsCard struct {
	id    string
	title string
}

type TL_webDocument struct {
	url         string
	access_hash int64
	size        int32
	mime_type   string
	attributes  []TL // DocumentAttribute
}

type TL_inputWebDocument struct {
	url        string
	size       int32
	mime_type  string
	attributes []TL // DocumentAttribute
}

type TL_inputWebFileLocation struct {
	url         string
	access_hash int64
}

type TL_upload_webFile struct {
	size      int32
	mime_type string
	file_type TL // storage_FileType
	mtime     int32
	bytes     []byte
}

type TL_payments_paymentForm struct {
	flags                TL // #
	can_save_credentials TL // flags_2?true
	password_missing     TL // flags_3?true
	bot_id               int32
	invoice              TL // Invoice
	provider_id          int32
	url                  string
	native_provider      TL   // flags_4?string
	native_params        TL   // flags_4?DataJSON
	saved_info           TL   // flags_0?PaymentRequestedInfo
	saved_credentials    TL   // flags_1?PaymentSavedCredentials
	users                []TL // User
}

type TL_payments_validatedRequestedInfo struct {
	flags            TL // #
	id               TL // flags_0?string
	shipping_options TL // flags_1?Vector<ShippingOption>
}

type TL_payments_paymentResult struct {
	updates TL // Updates
}

type TL_payments_paymentReceipt struct {
	flags             TL // #
	date              int32
	bot_id            int32
	invoice           TL // Invoice
	provider_id       int32
	info              TL // flags_0?PaymentRequestedInfo
	shipping          TL // flags_1?ShippingOption
	currency          string
	total_amount      int64
	credentials_title string
	users             []TL // User
}

type TL_payments_savedInfo struct {
	flags                 TL // #
	has_saved_credentials TL // flags_1?true
	saved_info            TL // flags_0?PaymentRequestedInfo
}

type TL_inputPaymentCredentialsSaved struct {
	id           string
	tmp_password []byte
}

type TL_inputPaymentCredentials struct {
	flags TL // #
	save  TL // flags_0?true
	data  TL // DataJSON
}

type TL_account_tmpPassword struct {
	tmp_password []byte
	valid_until  int32
}

type TL_shippingOption struct {
	id     string
	title  string
	prices []TL // LabeledPrice
}

type TL_updateBotShippingQuery struct {
	query_id         int64
	user_id          int32
	payload          []byte
	shipping_address TL // PostAddress
}

type TL_updateBotPrecheckoutQuery struct {
	flags              TL // #
	query_id           int64
	user_id            int32
	payload            []byte
	info               TL // flags_0?PaymentRequestedInfo
	shipping_option_id TL // flags_1?string
	currency           string
	total_amount       int64
}

type TL_inputStickerSetItem struct {
	flags       TL // #
	document    TL // InputDocument
	emoji       string
	mask_coords TL // flags_0?MaskCoords
}

type TL_updatePhoneCall struct {
	phone_call TL // PhoneCall
}

type TL_inputPhoneCall struct {
	id          int64
	access_hash int64
}

type TL_phoneCallEmpty struct {
	id int64
}

type TL_phoneCallWaiting struct {
	flags          TL // #
	video          TL // flags_5?true
	id             int64
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	protocol       TL // PhoneCallProtocol
	receive_date   TL // flags_0?int
}

type TL_phoneCallRequested struct {
	flags          TL // #
	video          TL // flags_5?true
	id             int64
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	g_a_hash       []byte
	protocol       TL // PhoneCallProtocol
}

type TL_phoneCallAccepted struct {
	flags          TL // #
	video          TL // flags_5?true
	id             int64
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	g_b            []byte
	protocol       TL // PhoneCallProtocol
}

type TL_phoneCall struct {
	flags           TL // #
	p2p_allowed     TL // flags_5?true
	id              int64
	access_hash     int64
	date            int32
	admin_id        int32
	participant_id  int32
	g_a_or_b        []byte
	key_fingerprint int64
	protocol        TL   // PhoneCallProtocol
	connections     []TL // PhoneConnection
	start_date      int32
}

type TL_phoneCallDiscarded struct {
	flags       TL // #
	need_rating TL // flags_2?true
	need_debug  TL // flags_3?true
	video       TL // flags_5?true
	id          int64
	reason      TL // flags_0?PhoneCallDiscardReason
	duration    TL // flags_1?int
}

type TL_phoneConnection struct {
	id       int64
	ip       string
	ipv6     string
	port     int32
	peer_tag []byte
}

type TL_phoneCallProtocol struct {
	flags         TL // #
	udp_p2p       TL // flags_0?true
	udp_reflector TL // flags_1?true
	min_layer     int32
	max_layer     int32
}

type TL_phone_phoneCall struct {
	phone_call TL   // PhoneCall
	users      []TL // User
}

type TL_inputMessagesFilterPhoneCalls struct {
	flags  TL // #
	missed TL // flags_0?true
}

type TL_messageActionPhoneCall struct {
	flags    TL // #
	video    TL // flags_2?true
	call_id  int64
	reason   TL // flags_0?PhoneCallDiscardReason
	duration TL // flags_1?int
}

type TL_inputMessagesFilterRoundVoice struct {
}

type TL_inputMessagesFilterRoundVideo struct {
}

type TL_sendMessageRecordRoundAction struct {
}

type TL_sendMessageUploadRoundAction struct {
	progress int32
}

type TL_upload_fileCdnRedirect struct {
	dc_id          int32
	file_token     []byte
	encryption_key []byte
	encryption_iv  []byte
	file_hashes    []TL // FileHash
}

type TL_upload_cdnFileReuploadNeeded struct {
	request_token []byte
}

type TL_upload_cdnFile struct {
	bytes []byte
}

type TL_cdnPublicKey struct {
	dc_id      int32
	public_key string
}

type TL_cdnConfig struct {
	public_keys []TL // CdnPublicKey
}

type TL_pageBlockChannel struct {
	channel TL // Chat
}

type TL_langPackString struct {
	key   string
	value string
}

type TL_langPackStringPluralized struct {
	flags       TL // #
	key         string
	zero_value  TL // flags_0?string
	one_value   TL // flags_1?string
	two_value   TL // flags_2?string
	few_value   TL // flags_3?string
	many_value  TL // flags_4?string
	other_value string
}

type TL_langPackStringDeleted struct {
	key string
}

type TL_langPackDifference struct {
	lang_code    string
	from_version int32
	version      int32
	strings      []TL // LangPackString
}

type TL_langPackLanguage struct {
	flags            TL // #
	official         TL // flags_0?true
	rtl              TL // flags_2?true
	beta             TL // flags_3?true
	name             string
	native_name      string
	lang_code        string
	base_lang_code   TL // flags_1?string
	plural_code      string
	strings_count    int32
	translated_count int32
	translations_url string
}

type TL_updateLangPackTooLong struct {
	lang_code string
}

type TL_updateLangPack struct {
	difference TL // LangPackDifference
}

type TL_channelParticipantAdmin struct {
	flags        TL // #
	can_edit     TL // flags_0?true
	self         TL // flags_1?true
	user_id      int32
	inviter_id   TL // flags_1?int
	promoted_by  int32
	date         int32
	admin_rights TL // ChatAdminRights
	rank         TL // flags_2?string
}

type TL_channelParticipantBanned struct {
	flags         TL // #
	left          TL // flags_0?true
	user_id       int32
	kicked_by     int32
	date          int32
	banned_rights TL // ChatBannedRights
}

type TL_channelParticipantsBanned struct {
	q string
}

type TL_channelParticipantsSearch struct {
	q string
}

type TL_channelAdminLogEventActionChangeTitle struct {
	prev_value string
	new_value  string
}

type TL_channelAdminLogEventActionChangeAbout struct {
	prev_value string
	new_value  string
}

type TL_channelAdminLogEventActionChangeUsername struct {
	prev_value string
	new_value  string
}

type TL_channelAdminLogEventActionChangePhoto struct {
	prev_photo TL // Photo
	new_photo  TL // Photo
}

type TL_channelAdminLogEventActionToggleInvites struct {
	new_value TL // Bool
}

type TL_channelAdminLogEventActionToggleSignatures struct {
	new_value TL // Bool
}

type TL_channelAdminLogEventActionUpdatePinned struct {
	message TL // Message
}

type TL_channelAdminLogEventActionEditMessage struct {
	prev_message TL // Message
	new_message  TL // Message
}

type TL_channelAdminLogEventActionDeleteMessage struct {
	message TL // Message
}

type TL_channelAdminLogEventActionParticipantJoin struct {
}

type TL_channelAdminLogEventActionParticipantLeave struct {
}

type TL_channelAdminLogEventActionParticipantInvite struct {
	participant TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleBan struct {
	prev_participant TL // ChannelParticipant
	new_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	prev_participant TL // ChannelParticipant
	new_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEvent struct {
	id      int64
	date    int32
	user_id int32
	action  TL // ChannelAdminLogEventAction
}

type TL_channels_adminLogResults struct {
	events []TL // ChannelAdminLogEvent
	chats  []TL // Chat
	users  []TL // User
}

type TL_channelAdminLogEventsFilter struct {
	flags    TL // #
	join     TL // flags_0?true
	leave    TL // flags_1?true
	invite   TL // flags_2?true
	ban      TL // flags_3?true
	unban    TL // flags_4?true
	kick     TL // flags_5?true
	unkick   TL // flags_6?true
	promote  TL // flags_7?true
	demote   TL // flags_8?true
	info     TL // flags_9?true
	settings TL // flags_10?true
	pinned   TL // flags_11?true
	edit     TL // flags_12?true
	delete   TL // flags_13?true
}

type TL_topPeerCategoryPhoneCalls struct {
}

type TL_pageBlockAudio struct {
	audio_id int64
	caption  TL // PageCaption
}

type TL_popularContact struct {
	client_id int64
	importers int32
}

type TL_messageActionScreenshotTaken struct {
}

type TL_messages_favedStickersNotModified struct {
}

type TL_messages_favedStickers struct {
	hash     int32
	packs    []TL // StickerPack
	stickers []TL // Document
}

type TL_updateFavedStickers struct {
}

type TL_updateChannelReadMessagesContents struct {
	channel_id int32
	messages   []int32
}

type TL_inputMessagesFilterMyMentions struct {
}

type TL_updateContactsReset struct {
}

type TL_channelAdminLogEventActionChangeStickerSet struct {
	prev_stickerset TL // InputStickerSet
	new_stickerset  TL // InputStickerSet
}

type TL_messageActionCustomAction struct {
	message string
}

type TL_inputPaymentCredentialsApplePay struct {
	payment_data TL // DataJSON
}

type TL_inputPaymentCredentialsAndroidPay struct {
	payment_token         TL // DataJSON
	google_transaction_id string
}

type TL_inputMessagesFilterGeo struct {
}

type TL_inputMessagesFilterContacts struct {
}

type TL_updateChannelAvailableMessages struct {
	channel_id       int32
	available_min_id int32
}

type TL_channelAdminLogEventActionTogglePreHistoryHidden struct {
	new_value TL // Bool
}

type TL_inputMediaGeoLive struct {
	flags     TL // #
	stopped   TL // flags_0?true
	geo_point TL // InputGeoPoint
	period    TL // flags_1?int
}

type TL_messageMediaGeoLive struct {
	geo    TL // GeoPoint
	period int32
}

type TL_recentMeUrlUnknown struct {
	url string
}

type TL_recentMeUrlUser struct {
	url     string
	user_id int32
}

type TL_recentMeUrlChat struct {
	url     string
	chat_id int32
}

type TL_recentMeUrlChatInvite struct {
	url         string
	chat_invite TL // ChatInvite
}

type TL_recentMeUrlStickerSet struct {
	url string
	set TL // StickerSetCovered
}

type TL_help_recentMeUrls struct {
	urls  []TL // RecentMeUrl
	chats []TL // Chat
	users []TL // User
}

type TL_channels_channelParticipantsNotModified struct {
}

type TL_messages_messagesNotModified struct {
	count int32
}

type TL_inputSingleMedia struct {
	flags     TL // #
	media     TL // InputMedia
	random_id int64
	message   string
	entities  TL // flags_0?Vector<MessageEntity>
}

type TL_webAuthorization struct {
	hash         int64
	bot_id       int32
	domain       string
	browser      string
	platform     string
	date_created int32
	date_active  int32
	ip           string
	region       string
}

type TL_account_webAuthorizations struct {
	authorizations []TL // WebAuthorization
	users          []TL // User
}

type TL_inputMessageID struct {
	id int32
}

type TL_inputMessageReplyTo struct {
	id int32
}

type TL_inputMessagePinned struct {
}

type TL_messageEntityPhone struct {
	offset int32
	length int32
}

type TL_messageEntityCashtag struct {
	offset int32
	length int32
}

type TL_messageActionBotAllowed struct {
	domain string
}

type TL_inputDialogPeer struct {
	peer TL // InputPeer
}

type TL_dialogPeer struct {
	peer TL // Peer
}

type TL_messages_foundStickerSetsNotModified struct {
}

type TL_messages_foundStickerSets struct {
	hash int32
	sets []TL // StickerSetCovered
}

type TL_fileHash struct {
	offset int32
	limit  int32
	hash   []byte
}

type TL_webDocumentNoProxy struct {
	url        string
	size       int32
	mime_type  string
	attributes []TL // DocumentAttribute
}

type TL_inputClientProxy struct {
	address string
	port    int32
}

type TL_help_proxyDataEmpty struct {
	expires int32
}

type TL_help_proxyDataPromo struct {
	expires int32
	peer    TL   // Peer
	chats   []TL // Chat
	users   []TL // User
}

type TL_help_termsOfServiceUpdateEmpty struct {
	expires int32
}

type TL_help_termsOfServiceUpdate struct {
	expires          int32
	terms_of_service TL // help_TermsOfService
}

type TL_inputSecureFileUploaded struct {
	id           int64
	parts        int32
	md5_checksum string
	file_hash    []byte
	secret       []byte
}

type TL_inputSecureFile struct {
	id          int64
	access_hash int64
}

type TL_inputSecureFileLocation struct {
	id          int64
	access_hash int64
}

type TL_secureFileEmpty struct {
}

type TL_secureFile struct {
	id          int64
	access_hash int64
	size        int32
	dc_id       int32
	date        int32
	file_hash   []byte
	secret      []byte
}

type TL_secureData struct {
	data      []byte
	data_hash []byte
	secret    []byte
}

type TL_securePlainPhone struct {
	phone string
}

type TL_securePlainEmail struct {
	email string
}

type TL_secureValueTypePersonalDetails struct {
}

type TL_secureValueTypePassport struct {
}

type TL_secureValueTypeDriverLicense struct {
}

type TL_secureValueTypeIdentityCard struct {
}

type TL_secureValueTypeInternalPassport struct {
}

type TL_secureValueTypeAddress struct {
}

type TL_secureValueTypeUtilityBill struct {
}

type TL_secureValueTypeBankStatement struct {
}

type TL_secureValueTypeRentalAgreement struct {
}

type TL_secureValueTypePassportRegistration struct {
}

type TL_secureValueTypeTemporaryRegistration struct {
}

type TL_secureValueTypePhone struct {
}

type TL_secureValueTypeEmail struct {
}

type TL_secureValue struct {
	flags        TL // #
	_type        TL // SecureValueType
	data         TL // flags_0?SecureData
	front_side   TL // flags_1?SecureFile
	reverse_side TL // flags_2?SecureFile
	selfie       TL // flags_3?SecureFile
	translation  TL // flags_6?Vector<SecureFile>
	files        TL // flags_4?Vector<SecureFile>
	plain_data   TL // flags_5?SecurePlainData
	hash         []byte
}

type TL_inputSecureValue struct {
	flags        TL // #
	_type        TL // SecureValueType
	data         TL // flags_0?SecureData
	front_side   TL // flags_1?InputSecureFile
	reverse_side TL // flags_2?InputSecureFile
	selfie       TL // flags_3?InputSecureFile
	translation  TL // flags_6?Vector<InputSecureFile>
	files        TL // flags_4?Vector<InputSecureFile>
	plain_data   TL // flags_5?SecurePlainData
}

type TL_secureValueHash struct {
	_type TL // SecureValueType
	hash  []byte
}

type TL_secureValueErrorData struct {
	_type     TL // SecureValueType
	data_hash []byte
	field     string
	text      string
}

type TL_secureValueErrorFrontSide struct {
	_type     TL // SecureValueType
	file_hash []byte
	text      string
}

type TL_secureValueErrorReverseSide struct {
	_type     TL // SecureValueType
	file_hash []byte
	text      string
}

type TL_secureValueErrorSelfie struct {
	_type     TL // SecureValueType
	file_hash []byte
	text      string
}

type TL_secureValueErrorFile struct {
	_type     TL // SecureValueType
	file_hash []byte
	text      string
}

type TL_secureValueErrorFiles struct {
	_type     TL   // SecureValueType
	file_hash []TL // bytes
	text      string
}

type TL_secureCredentialsEncrypted struct {
	data   []byte
	hash   []byte
	secret []byte
}

type TL_account_authorizationForm struct {
	flags              TL   // #
	required_types     []TL // SecureRequiredType
	values             []TL // SecureValue
	errors             []TL // SecureValueError
	users              []TL // User
	privacy_policy_url TL   // flags_0?string
}

type TL_account_sentEmailCode struct {
	email_pattern string
	length        int32
}

type TL_messageActionSecureValuesSentMe struct {
	values      []TL // SecureValue
	credentials TL   // SecureCredentialsEncrypted
}

type TL_messageActionSecureValuesSent struct {
	types []TL // SecureValueType
}

type TL_help_deepLinkInfoEmpty struct {
}

type TL_help_deepLinkInfo struct {
	flags      TL // #
	update_app TL // flags_0?true
	message    string
	entities   TL // flags_1?Vector<MessageEntity>
}

type TL_savedPhoneContact struct {
	phone      string
	first_name string
	last_name  string
	date       int32
}

type TL_account_takeout struct {
	id int64
}

type TL_inputTakeoutFileLocation struct {
}

type TL_updateDialogUnreadMark struct {
	flags  TL // #
	unread TL // flags_0?true
	peer   TL // DialogPeer
}

type TL_messages_dialogsNotModified struct {
	count int32
}

type TL_inputWebFileGeoPointLocation struct {
	geo_point   TL // InputGeoPoint
	access_hash int64
	w           int32
	h           int32
	zoom        int32
	scale       int32
}

type TL_contacts_topPeersDisabled struct {
}

type TL_inputReportReasonCopyright struct {
}

type TL_passwordKdfAlgoUnknown struct {
}

type TL_securePasswordKdfAlgoUnknown struct {
}

type TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000 struct {
	salt []byte
}

type TL_securePasswordKdfAlgoSHA512 struct {
	salt []byte
}

type TL_secureSecretSettings struct {
	secure_algo      TL // SecurePasswordKdfAlgo
	secure_secret    []byte
	secure_secret_id int64
}

type TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow struct {
	salt1 []byte
	salt2 []byte
	g     int32
	p     []byte
}

type TL_inputCheckPasswordEmpty struct {
}

type TL_inputCheckPasswordSRP struct {
	srp_id int64
	A      []byte
	M1     []byte
}

type TL_secureValueError struct {
	_type TL // SecureValueType
	hash  []byte
	text  string
}

type TL_secureValueErrorTranslationFile struct {
	_type     TL // SecureValueType
	file_hash []byte
	text      string
}

type TL_secureValueErrorTranslationFiles struct {
	_type     TL   // SecureValueType
	file_hash []TL // bytes
	text      string
}

type TL_secureRequiredType struct {
	flags                TL // #
	native_names         TL // flags_0?true
	selfie_required      TL // flags_1?true
	translation_required TL // flags_2?true
	_type                TL // SecureValueType
}

type TL_secureRequiredTypeOneOf struct {
	types []TL // SecureRequiredType
}

type TL_help_passportConfigNotModified struct {
}

type TL_help_passportConfig struct {
	hash            int32
	countries_langs TL // DataJSON
}

type TL_inputAppEvent struct {
	time  float64
	_type string
	peer  int64
	data  TL // JSONValue
}

type TL_jsonObjectValue struct {
	key   string
	value TL // JSONValue
}

type TL_jsonNull struct {
}

type TL_jsonBool struct {
	value TL // Bool
}

type TL_jsonNumber struct {
	value float64
}

type TL_jsonString struct {
	value string
}

type TL_jsonArray struct {
	value []TL // JSONValue
}

type TL_jsonObject struct {
	value []TL // JSONObjectValue
}

type TL_updateUserPinnedMessage struct {
	user_id int32
	id      int32
}

type TL_updateChatPinnedMessage struct {
	chat_id int32
	id      int32
	version int32
}

type TL_inputNotifyBroadcasts struct {
}

type TL_notifyBroadcasts struct {
}

type TL_textSubscript struct {
	text TL // RichText
}

type TL_textSuperscript struct {
	text TL // RichText
}

type TL_textMarked struct {
	text TL // RichText
}

type TL_textPhone struct {
	text  TL // RichText
	phone string
}

type TL_textImage struct {
	document_id int64
	w           int32
	h           int32
}

type TL_pageBlockKicker struct {
	text TL // RichText
}

type TL_pageTableCell struct {
	flags         TL // #
	header        TL // flags_0?true
	align_center  TL // flags_3?true
	align_right   TL // flags_4?true
	valign_middle TL // flags_5?true
	valign_bottom TL // flags_6?true
	text          TL // flags_7?RichText
	colspan       TL // flags_1?int
	rowspan       TL // flags_2?int
}

type TL_pageTableRow struct {
	cells []TL // PageTableCell
}

type TL_pageBlockTable struct {
	flags    TL   // #
	bordered TL   // flags_0?true
	striped  TL   // flags_1?true
	title    TL   // RichText
	rows     []TL // PageTableRow
}

type TL_pageCaption struct {
	text   TL // RichText
	credit TL // RichText
}

type TL_pageListItemText struct {
	text TL // RichText
}

type TL_pageListItemBlocks struct {
	blocks []TL // PageBlock
}

type TL_pageListOrderedItemText struct {
	num  string
	text TL // RichText
}

type TL_pageListOrderedItemBlocks struct {
	num    string
	blocks []TL // PageBlock
}

type TL_pageBlockOrderedList struct {
	items []TL // PageListOrderedItem
}

type TL_pageBlockDetails struct {
	flags  TL   // #
	open   TL   // flags_0?true
	blocks []TL // PageBlock
	title  TL   // RichText
}

type TL_pageRelatedArticle struct {
	flags          TL // #
	url            string
	webpage_id     int64
	title          TL // flags_0?string
	description    TL // flags_1?string
	photo_id       TL // flags_2?long
	author         TL // flags_3?string
	published_date TL // flags_4?int
}

type TL_pageBlockRelatedArticles struct {
	title    TL   // RichText
	articles []TL // PageRelatedArticle
}

type TL_pageBlockMap struct {
	geo     TL // GeoPoint
	zoom    int32
	w       int32
	h       int32
	caption TL // PageCaption
}

type TL_page struct {
	flags     TL // #
	part      TL // flags_0?true
	rtl       TL // flags_1?true
	v2        TL // flags_2?true
	url       string
	blocks    []TL // PageBlock
	photos    []TL // Photo
	documents []TL // Document
}

type TL_inputPrivacyKeyPhoneP2P struct {
}

type TL_privacyKeyPhoneP2P struct {
}

type TL_textAnchor struct {
	text TL // RichText
	name string
}

type TL_help_supportName struct {
	name string
}

type TL_help_userInfoEmpty struct {
}

type TL_help_userInfo struct {
	message  string
	entities []TL // MessageEntity
	author   string
	date     int32
}

type TL_messageActionContactSignUp struct {
}

type TL_updateMessagePoll struct {
	flags   TL // #
	poll_id int64
	poll    TL // flags_0?Poll
	results TL // PollResults
}

type TL_pollAnswer struct {
	text   string
	option []byte
}

type TL_poll struct {
	id              int64
	flags           TL // #
	closed          TL // flags_0?true
	public_voters   TL // flags_1?true
	multiple_choice TL // flags_2?true
	quiz            TL // flags_3?true
	question        string
	answers         []TL // PollAnswer
}

type TL_pollAnswerVoters struct {
	flags   TL // #
	chosen  TL // flags_0?true
	correct TL // flags_1?true
	option  []byte
	voters  int32
}

type TL_pollResults struct {
	flags        TL // #
	min          TL // flags_0?true
	results      TL // flags_1?Vector<PollAnswerVoters>
	total_voters TL // flags_2?int
}

type TL_inputMediaPoll struct {
	poll TL // Poll
}

type TL_messageMediaPoll struct {
	poll    TL // Poll
	results TL // PollResults
}

type TL_chatOnlines struct {
	onlines int32
}

type TL_statsURL struct {
	url string
}

type TL_photoStrippedSize struct {
	_type string
	bytes []byte
}

type TL_chatAdminRights struct {
	flags           TL // #
	change_info     TL // flags_0?true
	post_messages   TL // flags_1?true
	edit_messages   TL // flags_2?true
	delete_messages TL // flags_3?true
	ban_users       TL // flags_4?true
	invite_users    TL // flags_5?true
	pin_messages    TL // flags_7?true
	add_admins      TL // flags_9?true
}

type TL_chatBannedRights struct {
	flags         TL // #
	view_messages TL // flags_0?true
	send_messages TL // flags_1?true
	send_media    TL // flags_2?true
	send_stickers TL // flags_3?true
	send_gifs     TL // flags_4?true
	send_games    TL // flags_5?true
	send_inline   TL // flags_6?true
	embed_links   TL // flags_7?true
	send_polls    TL // flags_8?true
	change_info   TL // flags_10?true
	invite_users  TL // flags_15?true
	pin_messages  TL // flags_17?true
	until_date    int32
}

type TL_updateChatDefaultBannedRights struct {
	peer                  TL // Peer
	default_banned_rights TL // ChatBannedRights
	version               int32
}

type TL_inputWallPaper struct {
	id          int64
	access_hash int64
}

type TL_inputWallPaperSlug struct {
	slug string
}

type TL_channelParticipantsContacts struct {
	q string
}

type TL_channelAdminLogEventActionDefaultBannedRights struct {
	prev_banned_rights TL // ChatBannedRights
	new_banned_rights  TL // ChatBannedRights
}

type TL_channelAdminLogEventActionStopPoll struct {
	message TL // Message
}

type TL_account_wallPapersNotModified struct {
}

type TL_account_wallPapers struct {
	hash       int32
	wallpapers []TL // WallPaper
}

type TL_codeSettings struct {
	flags           TL // #
	allow_flashcall TL // flags_0?true
	current_number  TL // flags_1?true
	allow_app_hash  TL // flags_4?true
}

type TL_wallPaperSettings struct {
	flags            TL // #
	blur             TL // flags_1?true
	motion           TL // flags_2?true
	background_color TL // flags_0?int
	intensity        TL // flags_3?int
}

type TL_autoDownloadSettings struct {
	flags                TL // #
	disabled             TL // flags_0?true
	video_preload_large  TL // flags_1?true
	audio_preload_next   TL // flags_2?true
	phonecalls_less_data TL // flags_3?true
	photo_size_max       int32
	video_size_max       int32
	file_size_max        int32
}

type TL_account_autoDownloadSettings struct {
	low    TL // AutoDownloadSettings
	medium TL // AutoDownloadSettings
	high   TL // AutoDownloadSettings
}

type TL_emojiKeyword struct {
	keyword   string
	emoticons []string
}

type TL_emojiKeywordDeleted struct {
	keyword   string
	emoticons []string
}

type TL_emojiKeywordsDifference struct {
	lang_code    string
	from_version int32
	version      int32
	keywords     []TL // EmojiKeyword
}

type TL_emojiURL struct {
	url string
}

type TL_emojiLanguage struct {
	lang_code string
}

type TL_inputPrivacyKeyForwards struct {
}

type TL_privacyKeyForwards struct {
}

type TL_inputPrivacyKeyProfilePhoto struct {
}

type TL_privacyKeyProfilePhoto struct {
}

type TL_fileLocationToBeDeprecated struct {
	volume_id int64
	local_id  int32
}

type TL_inputPhotoFileLocation struct {
	id             int64
	access_hash    int64
	file_reference []byte
	thumb_size     string
}

type TL_inputPhotoLegacyFileLocation struct {
	id             int64
	access_hash    int64
	file_reference []byte
	volume_id      int64
	local_id       int32
	secret         int64
}

type TL_inputPeerPhotoFileLocation struct {
	flags     TL // #
	big       TL // flags_0?true
	peer      TL // InputPeer
	volume_id int64
	local_id  int32
}

type TL_inputStickerSetThumb struct {
	stickerset TL // InputStickerSet
	volume_id  int64
	local_id   int32
}

type TL_folder struct {
	flags                       TL // #
	autofill_new_broadcasts     TL // flags_0?true
	autofill_public_groups      TL // flags_1?true
	autofill_new_correspondents TL // flags_2?true
	id                          int32
	title                       string
	photo                       TL // flags_3?ChatPhoto
}

type TL_dialogFolder struct {
	flags                         TL // #
	pinned                        TL // flags_2?true
	folder                        TL // Folder
	peer                          TL // Peer
	top_message                   int32
	unread_muted_peers_count      int32
	unread_unmuted_peers_count    int32
	unread_muted_messages_count   int32
	unread_unmuted_messages_count int32
}

type TL_inputDialogPeerFolder struct {
	folder_id int32
}

type TL_dialogPeerFolder struct {
	folder_id int32
}

type TL_inputFolderPeer struct {
	peer      TL // InputPeer
	folder_id int32
}

type TL_folderPeer struct {
	peer      TL // Peer
	folder_id int32
}

type TL_updateFolderPeers struct {
	folder_peers []TL // FolderPeer
	pts          int32
	pts_count    int32
}

type TL_inputUserFromMessage struct {
	peer    TL // InputPeer
	msg_id  int32
	user_id int32
}

type TL_inputChannelFromMessage struct {
	peer       TL // InputPeer
	msg_id     int32
	channel_id int32
}

type TL_inputPeerUserFromMessage struct {
	peer    TL // InputPeer
	msg_id  int32
	user_id int32
}

type TL_inputPeerChannelFromMessage struct {
	peer       TL // InputPeer
	msg_id     int32
	channel_id int32
}

type TL_inputPrivacyKeyPhoneNumber struct {
}

type TL_privacyKeyPhoneNumber struct {
}

type TL_topPeerCategoryForwardUsers struct {
}

type TL_topPeerCategoryForwardChats struct {
}

type TL_channelAdminLogEventActionChangeLinkedChat struct {
	prev_value int32
	new_value  int32
}

type TL_messages_searchCounter struct {
	flags   TL // #
	inexact TL // flags_1?true
	filter  TL // MessagesFilter
	count   int32
}

type TL_keyboardButtonUrlAuth struct {
	flags     TL // #
	text      string
	fwd_text  TL // flags_0?string
	url       string
	button_id int32
}

type TL_inputKeyboardButtonUrlAuth struct {
	flags                TL // #
	request_write_access TL // flags_0?true
	text                 string
	fwd_text             TL // flags_1?string
	url                  string
	bot                  TL // InputUser
}

type TL_urlAuthResultRequest struct {
	flags                TL // #
	request_write_access TL // flags_0?true
	bot                  TL // User
	domain               string
}

type TL_urlAuthResultAccepted struct {
	url string
}

type TL_urlAuthResultDefault struct {
}

type TL_inputPrivacyValueAllowChatParticipants struct {
	chats []int32
}

type TL_inputPrivacyValueDisallowChatParticipants struct {
	chats []int32
}

type TL_privacyValueAllowChatParticipants struct {
	chats []int32
}

type TL_privacyValueDisallowChatParticipants struct {
	chats []int32
}

type TL_messageEntityUnderline struct {
	offset int32
	length int32
}

type TL_messageEntityStrike struct {
	offset int32
	length int32
}

type TL_messageEntityBlockquote struct {
	offset int32
	length int32
}

type TL_updatePeerSettings struct {
	peer     TL // Peer
	settings TL // PeerSettings
}

type TL_channelLocationEmpty struct {
}

type TL_channelLocation struct {
	geo_point TL // GeoPoint
	address   string
}

type TL_peerLocated struct {
	peer     TL // Peer
	expires  int32
	distance int32
}

type TL_updatePeerLocated struct {
	peers []TL // PeerLocated
}

type TL_channelAdminLogEventActionChangeLocation struct {
	prev_value TL // ChannelLocation
	new_value  TL // ChannelLocation
}

type TL_inputReportReasonGeoIrrelevant struct {
}

type TL_channelAdminLogEventActionToggleSlowMode struct {
	prev_value int32
	new_value  int32
}

type TL_auth_authorizationSignUpRequired struct {
	flags            TL // #
	terms_of_service TL // flags_0?help_TermsOfService
}

type TL_payments_paymentVerificationNeeded struct {
	url string
}

type TL_inputStickerSetAnimatedEmoji struct {
}

type TL_updateNewScheduledMessage struct {
	message TL // Message
}

type TL_updateDeleteScheduledMessages struct {
	peer     TL // Peer
	messages []int32
}

type TL_restrictionReason struct {
	platform string
	reason   string
	text     string
}

type TL_inputTheme struct {
	id          int64
	access_hash int64
}

type TL_inputThemeSlug struct {
	slug string
}

type TL_themeDocumentNotModified struct {
}

type TL_theme struct {
	flags          TL // #
	creator        TL // flags_0?true
	defaults       TL // flags_1?true
	id             int64
	access_hash    int64
	slug           string
	title          string
	document       TL // flags_2?Document
	installs_count int32
}

type TL_account_themesNotModified struct {
}

type TL_account_themes struct {
	hash   int32
	themes []TL // Theme
}

type TL_updateTheme struct {
	theme TL // Theme
}

type TL_inputPrivacyKeyAddedByPhone struct {
}

type TL_privacyKeyAddedByPhone struct {
}

type TL_invokeAfterMsg struct {
	msg_id int64
	query  TL
}

type TL_invokeAfterMsgs struct {
	msg_ids []int64
	query   TL
}

type TL_auth_sendCode struct {
	phone_number string
	api_id       int32
	api_hash     string
	settings     TL // CodeSettings
}

type TL_auth_signUp struct {
	phone_number    string
	phone_code_hash string
	first_name      string
	last_name       string
}

type TL_auth_signIn struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_exportAuthorization struct {
	dc_id int32
}

type TL_auth_importAuthorization struct {
	id    int32
	bytes []byte
}

type TL_auth_bindTempAuthKey struct {
	perm_auth_key_id  int64
	nonce             int64
	expires_at        int32
	encrypted_message []byte
}

type TL_account_registerDevice struct {
	flags       TL // #
	no_muted    TL // flags_0?true
	token_type  int32
	token       string
	app_sandbox TL // Bool
	secret      []byte
	other_uids  []int32
}

type TL_account_unregisterDevice struct {
	token_type int32
	token      string
	other_uids []int32
}

type TL_account_updateNotifySettings struct {
	peer     TL // InputNotifyPeer
	settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	flags      TL // #
	first_name TL // flags_0?string
	last_name  TL // flags_1?string
	about      TL // flags_2?string
}

type TL_account_updateStatus struct {
	offline TL // Bool
}

type TL_account_getWallPapers struct {
	hash int32
}

type TL_account_reportPeer struct {
	peer   TL // InputPeer
	reason TL // ReportReason
}

type TL_users_getUsers struct {
	id []TL // InputUser
}

type TL_users_getFullUser struct {
	id TL // InputUser
}

type TL_contacts_getContactIDs struct {
	hash int32
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	hash int32
}

type TL_contacts_importContacts struct {
	contacts []TL // InputContact
}

type TL_contacts_deleteContacts struct {
	id []TL // InputUser
}

type TL_contacts_deleteByPhones struct {
	phones []string
}

type TL_contacts_block struct {
	id TL // InputUser
}

type TL_contacts_unblock struct {
	id TL // InputUser
}

type TL_contacts_getBlocked struct {
	offset int32
	limit  int32
}

type TL_messages_getMessages struct {
	id []TL // InputMessage
}

type TL_messages_getDialogs struct {
	flags          TL // #
	exclude_pinned TL // flags_0?true
	folder_id      TL // flags_1?int
	offset_date    int32
	offset_id      int32
	offset_peer    TL // InputPeer
	limit          int32
	hash           int32
}

type TL_messages_getHistory struct {
	peer        TL // InputPeer
	offset_id   int32
	offset_date int32
	add_offset  int32
	limit       int32
	max_id      int32
	min_id      int32
	hash        int32
}

type TL_messages_search struct {
	flags      TL // #
	peer       TL // InputPeer
	q          string
	from_id    TL // flags_0?InputUser
	filter     TL // MessagesFilter
	min_date   int32
	max_date   int32
	offset_id  int32
	add_offset int32
	limit      int32
	max_id     int32
	min_id     int32
	hash       int32
}

type TL_messages_readHistory struct {
	peer   TL // InputPeer
	max_id int32
}

type TL_messages_deleteHistory struct {
	flags      TL // #
	just_clear TL // flags_0?true
	revoke     TL // flags_1?true
	peer       TL // InputPeer
	max_id     int32
}

type TL_messages_deleteMessages struct {
	flags  TL // #
	revoke TL // flags_0?true
	id     []int32
}

type TL_messages_receivedMessages struct {
	max_id int32
}

type TL_messages_setTyping struct {
	peer   TL // InputPeer
	action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	flags           TL // #
	no_webpage      TL // flags_1?true
	silent          TL // flags_5?true
	background      TL // flags_6?true
	clear_draft     TL // flags_7?true
	peer            TL // InputPeer
	reply_to_msg_id TL // flags_0?int
	message         string
	random_id       int64
	reply_markup    TL // flags_2?ReplyMarkup
	entities        TL // flags_3?Vector<MessageEntity>
	schedule_date   TL // flags_10?int
}

type TL_messages_sendMedia struct {
	flags           TL // #
	silent          TL // flags_5?true
	background      TL // flags_6?true
	clear_draft     TL // flags_7?true
	peer            TL // InputPeer
	reply_to_msg_id TL // flags_0?int
	media           TL // InputMedia
	message         string
	random_id       int64
	reply_markup    TL // flags_2?ReplyMarkup
	entities        TL // flags_3?Vector<MessageEntity>
	schedule_date   TL // flags_10?int
}

type TL_messages_forwardMessages struct {
	flags         TL // #
	silent        TL // flags_5?true
	background    TL // flags_6?true
	with_my_score TL // flags_8?true
	grouped       TL // flags_9?true
	from_peer     TL // InputPeer
	id            []int32
	random_id     []int64
	to_peer       TL // InputPeer
	schedule_date TL // flags_10?int
}

type TL_messages_reportSpam struct {
	peer TL // InputPeer
}

type TL_messages_getPeerSettings struct {
	peer TL // InputPeer
}

type TL_messages_report struct {
	peer   TL // InputPeer
	id     []int32
	reason TL // ReportReason
}

type TL_messages_getChats struct {
	id []int32
}

type TL_messages_getFullChat struct {
	chat_id int32
}

type TL_messages_editChatTitle struct {
	chat_id int32
	title   string
}

type TL_messages_editChatPhoto struct {
	chat_id int32
	photo   TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	chat_id   int32
	user_id   TL // InputUser
	fwd_limit int32
}

type TL_messages_deleteChatUser struct {
	chat_id int32
	user_id TL // InputUser
}

type TL_messages_createChat struct {
	users []TL // InputUser
	title string
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	flags           TL // #
	pts             int32
	pts_total_limit TL // flags_0?int
	date            int32
	qts             int32
}

type TL_photos_updateProfilePhoto struct {
	id TL // InputPhoto
}

type TL_photos_uploadProfilePhoto struct {
	file TL // InputFile
}

type TL_photos_deletePhotos struct {
	id []TL // InputPhoto
}

type TL_upload_saveFilePart struct {
	file_id   int64
	file_part int32
	bytes     []byte
}

type TL_upload_getFile struct {
	flags         TL // #
	precise       TL // flags_0?true
	cdn_supported TL // flags_1?true
	location      TL // InputFileLocation
	offset        int32
	limit         int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
	source string
}

type TL_help_getInviteText struct {
}

type TL_photos_getUserPhotos struct {
	user_id TL // InputUser
	offset  int32
	max_id  int64
	limit   int32
}

type TL_messages_getDhConfig struct {
	version       int32
	random_length int32
}

type TL_messages_requestEncryption struct {
	user_id   TL // InputUser
	random_id int32
	g_a       []byte
}

type TL_messages_acceptEncryption struct {
	peer            TL // InputEncryptedChat
	g_b             []byte
	key_fingerprint int64
}

type TL_messages_discardEncryption struct {
	chat_id int32
}

type TL_messages_setEncryptedTyping struct {
	peer   TL // InputEncryptedChat
	typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	peer     TL // InputEncryptedChat
	max_date int32
}

type TL_messages_sendEncrypted struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_sendEncryptedFile struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
	file      TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_receivedQueue struct {
	max_qts int32
}

type TL_messages_reportEncryptedSpam struct {
	peer TL // InputEncryptedChat
}

type TL_upload_saveBigFilePart struct {
	file_id          int64
	file_part        int32
	file_total_parts int32
	bytes            []byte
}

type TL_initConnection struct {
	flags            TL // #
	api_id           int32
	device_model     string
	system_version   string
	app_version      string
	system_lang_code string
	lang_pack        string
	lang_code        string
	proxy            TL // flags_0?InputClientProxy
	query            TL
}

type TL_help_getSupport struct {
}

type TL_messages_readMessageContents struct {
	id []int32
}

type TL_account_checkUsername struct {
	username string
}

type TL_account_updateUsername struct {
	username string
}

type TL_contacts_search struct {
	q     string
	limit int32
}

type TL_account_getPrivacy struct {
	key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	key   TL   // InputPrivacyKey
	rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	ttl TL // AccountDaysTTL
}

type TL_invokeWithLayer struct {
	layer int32
	query TL
}

type TL_contacts_resolveUsername struct {
	username string
}

type TL_account_sendChangePhoneCode struct {
	phone_number string
	settings     TL // CodeSettings
}

type TL_account_changePhone struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_messages_getStickers struct {
	emoticon string
	hash     int32
}

type TL_messages_getAllStickers struct {
	hash int32
}

type TL_account_updateDeviceLocked struct {
	period int32
}

type TL_auth_importBotAuthorization struct {
	flags          int32
	api_id         int32
	api_hash       string
	bot_auth_token string
}

type TL_messages_getWebPagePreview struct {
	flags    TL // #
	message  string
	entities TL // flags_3?Vector<MessageEntity>
}

type TL_account_getAuthorizations struct {
}

type TL_account_resetAuthorization struct {
	hash int64
}

type TL_account_getPassword struct {
}

type TL_account_getPasswordSettings struct {
	password TL // InputCheckPasswordSRP
}

type TL_account_updatePasswordSettings struct {
	password     TL // InputCheckPasswordSRP
	new_settings TL // account_PasswordInputSettings
}

type TL_auth_checkPassword struct {
	password TL // InputCheckPasswordSRP
}

type TL_auth_requestPasswordRecovery struct {
}

type TL_auth_recoverPassword struct {
	code string
}

type TL_invokeWithoutUpdates struct {
	query TL
}

type TL_messages_exportChatInvite struct {
	peer TL // InputPeer
}

type TL_messages_checkChatInvite struct {
	hash string
}

type TL_messages_importChatInvite struct {
	hash string
}

type TL_messages_getStickerSet struct {
	stickerset TL // InputStickerSet
}

type TL_messages_installStickerSet struct {
	stickerset TL // InputStickerSet
	archived   TL // Bool
}

type TL_messages_uninstallStickerSet struct {
	stickerset TL // InputStickerSet
}

type TL_messages_startBot struct {
	bot         TL // InputUser
	peer        TL // InputPeer
	random_id   int64
	start_param string
}

type TL_help_getAppChangelog struct {
	prev_app_version string
}

type TL_messages_getMessagesViews struct {
	peer      TL // InputPeer
	id        []int32
	increment TL // Bool
}

type TL_channels_readHistory struct {
	channel TL // InputChannel
	max_id  int32
}

type TL_channels_deleteMessages struct {
	channel TL // InputChannel
	id      []int32
}

type TL_channels_deleteUserHistory struct {
	channel TL // InputChannel
	user_id TL // InputUser
}

type TL_channels_reportSpam struct {
	channel TL // InputChannel
	user_id TL // InputUser
	id      []int32
}

type TL_channels_getMessages struct {
	channel TL   // InputChannel
	id      []TL // InputMessage
}

type TL_channels_getParticipants struct {
	channel TL // InputChannel
	filter  TL // ChannelParticipantsFilter
	offset  int32
	limit   int32
	hash    int32
}

type TL_channels_getParticipant struct {
	channel TL // InputChannel
	user_id TL // InputUser
}

type TL_channels_getChannels struct {
	id []TL // InputChannel
}

type TL_channels_getFullChannel struct {
	channel TL // InputChannel
}

type TL_channels_createChannel struct {
	flags     TL // #
	broadcast TL // flags_0?true
	megagroup TL // flags_1?true
	title     string
	about     string
	geo_point TL // flags_2?InputGeoPoint
	address   TL // flags_2?string
}

type TL_channels_editAdmin struct {
	channel      TL // InputChannel
	user_id      TL // InputUser
	admin_rights TL // ChatAdminRights
	rank         string
}

type TL_channels_editTitle struct {
	channel TL // InputChannel
	title   string
}

type TL_channels_editPhoto struct {
	channel TL // InputChannel
	photo   TL // InputChatPhoto
}

type TL_channels_checkUsername struct {
	channel  TL // InputChannel
	username string
}

type TL_channels_updateUsername struct {
	channel  TL // InputChannel
	username string
}

type TL_channels_joinChannel struct {
	channel TL // InputChannel
}

type TL_channels_leaveChannel struct {
	channel TL // InputChannel
}

type TL_channels_inviteToChannel struct {
	channel TL   // InputChannel
	users   []TL // InputUser
}

type TL_channels_deleteChannel struct {
	channel TL // InputChannel
}

type TL_updates_getChannelDifference struct {
	flags   TL // #
	force   TL // flags_0?true
	channel TL // InputChannel
	filter  TL // ChannelMessagesFilter
	pts     int32
	limit   int32
}

type TL_messages_editChatAdmin struct {
	chat_id  int32
	user_id  TL // InputUser
	is_admin TL // Bool
}

type TL_messages_migrateChat struct {
	chat_id int32
}

type TL_messages_searchGlobal struct {
	flags       TL // #
	folder_id   TL // flags_0?int
	q           string
	offset_rate int32
	offset_peer TL // InputPeer
	offset_id   int32
	limit       int32
}

type TL_messages_reorderStickerSets struct {
	flags TL // #
	masks TL // flags_0?true
	order []int64
}

type TL_messages_getDocumentByHash struct {
	sha256    []byte
	size      int32
	mime_type string
}

type TL_messages_searchGifs struct {
	q      string
	offset int32
}

type TL_messages_getSavedGifs struct {
	hash int32
}

type TL_messages_saveGif struct {
	id     TL // InputDocument
	unsave TL // Bool
}

type TL_messages_getInlineBotResults struct {
	flags     TL // #
	bot       TL // InputUser
	peer      TL // InputPeer
	geo_point TL // flags_0?InputGeoPoint
	query     string
	offset    string
}

type TL_messages_setInlineBotResults struct {
	flags       TL // #
	gallery     TL // flags_0?true
	private     TL // flags_1?true
	query_id    int64
	results     []TL // InputBotInlineResult
	cache_time  int32
	next_offset TL // flags_2?string
	switch_pm   TL // flags_3?InlineBotSwitchPM
}

type TL_messages_sendInlineBotResult struct {
	flags           TL // #
	silent          TL // flags_5?true
	background      TL // flags_6?true
	clear_draft     TL // flags_7?true
	hide_via        TL // flags_11?true
	peer            TL // InputPeer
	reply_to_msg_id TL // flags_0?int
	random_id       int64
	query_id        int64
	id              string
	schedule_date   TL // flags_10?int
}

type TL_channels_exportMessageLink struct {
	channel TL // InputChannel
	id      int32
	grouped TL // Bool
}

type TL_channels_toggleSignatures struct {
	channel TL // InputChannel
	enabled TL // Bool
}

type TL_auth_resendCode struct {
	phone_number    string
	phone_code_hash string
}

type TL_auth_cancelCode struct {
	phone_number    string
	phone_code_hash string
}

type TL_messages_getMessageEditData struct {
	peer TL // InputPeer
	id   int32
}

type TL_messages_editMessage struct {
	flags         TL // #
	no_webpage    TL // flags_1?true
	peer          TL // InputPeer
	id            int32
	message       TL // flags_11?string
	media         TL // flags_14?InputMedia
	reply_markup  TL // flags_2?ReplyMarkup
	entities      TL // flags_3?Vector<MessageEntity>
	schedule_date TL // flags_15?int
}

type TL_messages_editInlineBotMessage struct {
	flags        TL // #
	no_webpage   TL // flags_1?true
	id           TL // InputBotInlineMessageID
	message      TL // flags_11?string
	media        TL // flags_14?InputMedia
	reply_markup TL // flags_2?ReplyMarkup
	entities     TL // flags_3?Vector<MessageEntity>
}

type TL_messages_getBotCallbackAnswer struct {
	flags  TL // #
	game   TL // flags_1?true
	peer   TL // InputPeer
	msg_id int32
	data   TL // flags_0?bytes
}

type TL_messages_setBotCallbackAnswer struct {
	flags      TL // #
	alert      TL // flags_1?true
	query_id   int64
	message    TL // flags_0?string
	url        TL // flags_2?string
	cache_time int32
}

type TL_contacts_getTopPeers struct {
	flags          TL // #
	correspondents TL // flags_0?true
	bots_pm        TL // flags_1?true
	bots_inline    TL // flags_2?true
	phone_calls    TL // flags_3?true
	forward_users  TL // flags_4?true
	forward_chats  TL // flags_5?true
	groups         TL // flags_10?true
	channels       TL // flags_15?true
	offset         int32
	limit          int32
	hash           int32
}

type TL_contacts_resetTopPeerRating struct {
	category TL // TopPeerCategory
	peer     TL // InputPeer
}

type TL_messages_getPeerDialogs struct {
	peers []TL // InputDialogPeer
}

type TL_messages_saveDraft struct {
	flags           TL // #
	no_webpage      TL // flags_1?true
	reply_to_msg_id TL // flags_0?int
	peer            TL // InputPeer
	message         string
	entities        TL // flags_3?Vector<MessageEntity>
}

type TL_messages_getAllDrafts struct {
}

type TL_messages_getFeaturedStickers struct {
	hash int32
}

type TL_messages_readFeaturedStickers struct {
	id []int64
}

type TL_messages_getRecentStickers struct {
	flags    TL // #
	attached TL // flags_0?true
	hash     int32
}

type TL_messages_saveRecentSticker struct {
	flags    TL // #
	attached TL // flags_0?true
	id       TL // InputDocument
	unsave   TL // Bool
}

type TL_messages_clearRecentStickers struct {
	flags    TL // #
	attached TL // flags_0?true
}

type TL_messages_getArchivedStickers struct {
	flags     TL // #
	masks     TL // flags_0?true
	offset_id int64
	limit     int32
}

type TL_account_sendConfirmPhoneCode struct {
	hash     string
	settings TL // CodeSettings
}

type TL_account_confirmPhone struct {
	phone_code_hash string
	phone_code      string
}

type TL_channels_getAdminedPublicChannels struct {
	flags       TL // #
	by_location TL // flags_0?true
	check_limit TL // flags_1?true
}

type TL_messages_getMaskStickers struct {
	hash int32
}

type TL_messages_getAttachedStickers struct {
	media TL // InputStickeredMedia
}

type TL_auth_dropTempAuthKeys struct {
	except_auth_keys []int64
}

type TL_messages_setGameScore struct {
	flags        TL // #
	edit_message TL // flags_0?true
	force        TL // flags_1?true
	peer         TL // InputPeer
	id           int32
	user_id      TL // InputUser
	score        int32
}

type TL_messages_setInlineGameScore struct {
	flags        TL // #
	edit_message TL // flags_0?true
	force        TL // flags_1?true
	id           TL // InputBotInlineMessageID
	user_id      TL // InputUser
	score        int32
}

type TL_messages_getGameHighScores struct {
	peer    TL // InputPeer
	id      int32
	user_id TL // InputUser
}

type TL_messages_getInlineGameHighScores struct {
	id      TL // InputBotInlineMessageID
	user_id TL // InputUser
}

type TL_messages_getCommonChats struct {
	user_id TL // InputUser
	max_id  int32
	limit   int32
}

type TL_messages_getAllChats struct {
	except_ids []int32
}

type TL_help_setBotUpdatesStatus struct {
	pending_updates_count int32
	message               string
}

type TL_messages_getWebPage struct {
	url  string
	hash int32
}

type TL_messages_toggleDialogPin struct {
	flags  TL // #
	pinned TL // flags_0?true
	peer   TL // InputDialogPeer
}

type TL_messages_reorderPinnedDialogs struct {
	flags     TL // #
	force     TL // flags_0?true
	folder_id int32
	order     []TL // InputDialogPeer
}

type TL_messages_getPinnedDialogs struct {
	folder_id int32
}

type TL_bots_sendCustomRequest struct {
	custom_method string
	params        TL // DataJSON
}

type TL_bots_answerWebhookJSONQuery struct {
	query_id int64
	data     TL // DataJSON
}

type TL_upload_getWebFile struct {
	location TL // InputWebFileLocation
	offset   int32
	limit    int32
}

type TL_payments_getPaymentForm struct {
	msg_id int32
}

type TL_payments_getPaymentReceipt struct {
	msg_id int32
}

type TL_payments_validateRequestedInfo struct {
	flags  TL // #
	save   TL // flags_0?true
	msg_id int32
	info   TL // PaymentRequestedInfo
}

type TL_payments_sendPaymentForm struct {
	flags              TL // #
	msg_id             int32
	requested_info_id  TL // flags_0?string
	shipping_option_id TL // flags_1?string
	credentials        TL // InputPaymentCredentials
}

type TL_account_getTmpPassword struct {
	password TL // InputCheckPasswordSRP
	period   int32
}

type TL_payments_getSavedInfo struct {
}

type TL_payments_clearSavedInfo struct {
	flags       TL // #
	credentials TL // flags_0?true
	info        TL // flags_1?true
}

type TL_messages_setBotShippingResults struct {
	flags            TL // #
	query_id         int64
	error            TL // flags_0?string
	shipping_options TL // flags_1?Vector<ShippingOption>
}

type TL_messages_setBotPrecheckoutResults struct {
	flags    TL // #
	success  TL // flags_1?true
	query_id int64
	error    TL // flags_0?string
}

type TL_stickers_createStickerSet struct {
	flags      TL // #
	masks      TL // flags_0?true
	user_id    TL // InputUser
	title      string
	short_name string
	stickers   []TL // InputStickerSetItem
}

type TL_stickers_removeStickerFromSet struct {
	sticker TL // InputDocument
}

type TL_stickers_changeStickerPosition struct {
	sticker  TL // InputDocument
	position int32
}

type TL_stickers_addStickerToSet struct {
	stickerset TL // InputStickerSet
	sticker    TL // InputStickerSetItem
}

type TL_messages_uploadMedia struct {
	peer  TL // InputPeer
	media TL // InputMedia
}

type TL_phone_getCallConfig struct {
}

type TL_phone_requestCall struct {
	flags     TL // #
	video     TL // flags_0?true
	user_id   TL // InputUser
	random_id int32
	g_a_hash  []byte
	protocol  TL // PhoneCallProtocol
}

type TL_phone_acceptCall struct {
	peer     TL // InputPhoneCall
	g_b      []byte
	protocol TL // PhoneCallProtocol
}

type TL_phone_confirmCall struct {
	peer            TL // InputPhoneCall
	g_a             []byte
	key_fingerprint int64
	protocol        TL // PhoneCallProtocol
}

type TL_phone_receivedCall struct {
	peer TL // InputPhoneCall
}

type TL_phone_discardCall struct {
	flags         TL // #
	video         TL // flags_0?true
	peer          TL // InputPhoneCall
	duration      int32
	reason        TL // PhoneCallDiscardReason
	connection_id int64
}

type TL_phone_setCallRating struct {
	flags           TL // #
	user_initiative TL // flags_0?true
	peer            TL // InputPhoneCall
	rating          int32
	comment         string
}

type TL_phone_saveCallDebug struct {
	peer  TL // InputPhoneCall
	debug TL // DataJSON
}

type TL_upload_getCdnFile struct {
	file_token []byte
	offset     int32
	limit      int32
}

type TL_upload_reuploadCdnFile struct {
	file_token    []byte
	request_token []byte
}

type TL_help_getCdnConfig struct {
}

type TL_langpack_getLangPack struct {
	lang_pack string
	lang_code string
}

type TL_langpack_getStrings struct {
	lang_pack string
	lang_code string
	keys      []string
}

type TL_langpack_getDifference struct {
	lang_pack    string
	lang_code    string
	from_version int32
}

type TL_langpack_getLanguages struct {
	lang_pack string
}

type TL_channels_editBanned struct {
	channel       TL // InputChannel
	user_id       TL // InputUser
	banned_rights TL // ChatBannedRights
}

type TL_channels_getAdminLog struct {
	flags         TL // #
	channel       TL // InputChannel
	q             string
	events_filter TL // flags_0?ChannelAdminLogEventsFilter
	admins        TL // flags_1?Vector<InputUser>
	max_id        int64
	min_id        int64
	limit         int32
}

type TL_upload_getCdnFileHashes struct {
	file_token []byte
	offset     int32
}

type TL_messages_sendScreenshotNotification struct {
	peer            TL // InputPeer
	reply_to_msg_id int32
	random_id       int64
}

type TL_channels_setStickers struct {
	channel    TL // InputChannel
	stickerset TL // InputStickerSet
}

type TL_messages_getFavedStickers struct {
	hash int32
}

type TL_messages_faveSticker struct {
	id     TL // InputDocument
	unfave TL // Bool
}

type TL_channels_readMessageContents struct {
	channel TL // InputChannel
	id      []int32
}

type TL_contacts_resetSaved struct {
}

type TL_messages_getUnreadMentions struct {
	peer       TL // InputPeer
	offset_id  int32
	add_offset int32
	limit      int32
	max_id     int32
	min_id     int32
}

type TL_channels_deleteHistory struct {
	channel TL // InputChannel
	max_id  int32
}

type TL_help_getRecentMeUrls struct {
	referer string
}

type TL_channels_togglePreHistoryHidden struct {
	channel TL // InputChannel
	enabled TL // Bool
}

type TL_messages_readMentions struct {
	peer TL // InputPeer
}

type TL_messages_getRecentLocations struct {
	peer  TL // InputPeer
	limit int32
	hash  int32
}

type TL_messages_sendMultiMedia struct {
	flags           TL   // #
	silent          TL   // flags_5?true
	background      TL   // flags_6?true
	clear_draft     TL   // flags_7?true
	peer            TL   // InputPeer
	reply_to_msg_id TL   // flags_0?int
	multi_media     []TL // InputSingleMedia
	schedule_date   TL   // flags_10?int
}

type TL_messages_uploadEncryptedFile struct {
	peer TL // InputEncryptedChat
	file TL // InputEncryptedFile
}

type TL_account_getWebAuthorizations struct {
}

type TL_account_resetWebAuthorization struct {
	hash int64
}

type TL_account_resetWebAuthorizations struct {
}

type TL_messages_searchStickerSets struct {
	flags            TL // #
	exclude_featured TL // flags_0?true
	q                string
	hash             int32
}

type TL_upload_getFileHashes struct {
	location TL // InputFileLocation
	offset   int32
}

type TL_help_getProxyData struct {
}

type TL_help_getTermsOfServiceUpdate struct {
}

type TL_help_acceptTermsOfService struct {
	id TL // DataJSON
}

type TL_account_getAllSecureValues struct {
}

type TL_account_getSecureValue struct {
	types []TL // SecureValueType
}

type TL_account_saveSecureValue struct {
	value            TL // InputSecureValue
	secure_secret_id int64
}

type TL_account_deleteSecureValue struct {
	types []TL // SecureValueType
}

type TL_users_setSecureValueErrors struct {
	id     TL   // InputUser
	errors []TL // SecureValueError
}

type TL_account_getAuthorizationForm struct {
	bot_id     int32
	scope      string
	public_key string
}

type TL_account_acceptAuthorization struct {
	bot_id       int32
	scope        string
	public_key   string
	value_hashes []TL // SecureValueHash
	credentials  TL   // SecureCredentialsEncrypted
}

type TL_account_sendVerifyPhoneCode struct {
	phone_number string
	settings     TL // CodeSettings
}

type TL_account_verifyPhone struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_account_sendVerifyEmailCode struct {
	email string
}

type TL_account_verifyEmail struct {
	email string
	code  string
}

type TL_help_getDeepLinkInfo struct {
	path string
}

type TL_contacts_getSaved struct {
}

type TL_channels_getLeftChannels struct {
	offset int32
}

type TL_account_initTakeoutSession struct {
	flags              TL // #
	contacts           TL // flags_0?true
	message_users      TL // flags_1?true
	message_chats      TL // flags_2?true
	message_megagroups TL // flags_3?true
	message_channels   TL // flags_4?true
	files              TL // flags_5?true
	file_max_size      TL // flags_5?int
}

type TL_account_finishTakeoutSession struct {
	flags   TL // #
	success TL // flags_0?true
}

type TL_messages_getSplitRanges struct {
}

type TL_invokeWithMessagesRange struct {
	ranges TL // MessageRange
	query  TL
}

type TL_invokeWithTakeout struct {
	takeout_id int64
	query      TL
}

type TL_messages_markDialogUnread struct {
	flags  TL // #
	unread TL // flags_0?true
	peer   TL // InputDialogPeer
}

type TL_messages_getDialogUnreadMarks struct {
}

type TL_contacts_toggleTopPeers struct {
	enabled TL // Bool
}

type TL_messages_clearAllDrafts struct {
}

type TL_help_getAppConfig struct {
}

type TL_help_saveAppLog struct {
	events []TL // InputAppEvent
}

type TL_help_getPassportConfig struct {
	hash int32
}

type TL_langpack_getLanguage struct {
	lang_pack string
	lang_code string
}

type TL_messages_updatePinnedMessage struct {
	flags  TL // #
	silent TL // flags_0?true
	peer   TL // InputPeer
	id     int32
}

type TL_account_confirmPasswordEmail struct {
	code string
}

type TL_account_resendPasswordEmail struct {
}

type TL_account_cancelPasswordEmail struct {
}

type TL_help_getSupportName struct {
}

type TL_help_getUserInfo struct {
	user_id TL // InputUser
}

type TL_help_editUserInfo struct {
	user_id  TL // InputUser
	message  string
	entities []TL // MessageEntity
}

type TL_account_getContactSignUpNotification struct {
}

type TL_account_setContactSignUpNotification struct {
	silent TL // Bool
}

type TL_account_getNotifyExceptions struct {
	flags         TL // #
	compare_sound TL // flags_1?true
	peer          TL // flags_0?InputNotifyPeer
}

type TL_messages_sendVote struct {
	peer    TL // InputPeer
	msg_id  int32
	options []TL // bytes
}

type TL_messages_getPollResults struct {
	peer   TL // InputPeer
	msg_id int32
}

type TL_messages_getOnlines struct {
	peer TL // InputPeer
}

type TL_messages_getStatsURL struct {
	flags  TL // #
	dark   TL // flags_0?true
	peer   TL // InputPeer
	params string
}

type TL_messages_editChatAbout struct {
	peer  TL // InputPeer
	about string
}

type TL_messages_editChatDefaultBannedRights struct {
	peer          TL // InputPeer
	banned_rights TL // ChatBannedRights
}

type TL_account_getWallPaper struct {
	wallpaper TL // InputWallPaper
}

type TL_account_uploadWallPaper struct {
	file      TL // InputFile
	mime_type string
	settings  TL // WallPaperSettings
}

type TL_account_saveWallPaper struct {
	wallpaper TL // InputWallPaper
	unsave    TL // Bool
	settings  TL // WallPaperSettings
}

type TL_account_installWallPaper struct {
	wallpaper TL // InputWallPaper
	settings  TL // WallPaperSettings
}

type TL_account_resetWallPapers struct {
}

type TL_account_getAutoDownloadSettings struct {
}

type TL_account_saveAutoDownloadSettings struct {
	flags    TL // #
	low      TL // flags_0?true
	high     TL // flags_1?true
	settings TL // AutoDownloadSettings
}

type TL_messages_getEmojiKeywords struct {
	lang_code string
}

type TL_messages_getEmojiKeywordsDifference struct {
	lang_code    string
	from_version int32
}

type TL_messages_getEmojiKeywordsLanguages struct {
	lang_codes []string
}

type TL_messages_getEmojiURL struct {
	lang_code string
}

type TL_folders_editPeerFolders struct {
	folder_peers []TL // InputFolderPeer
}

type TL_folders_deleteFolder struct {
	folder_id int32
}

type TL_messages_getSearchCounters struct {
	peer    TL   // InputPeer
	filters []TL // MessagesFilter
}

type TL_channels_getGroupsForDiscussion struct {
}

type TL_channels_setDiscussionGroup struct {
	broadcast TL // InputChannel
	group     TL // InputChannel
}

type TL_messages_requestUrlAuth struct {
	peer      TL // InputPeer
	msg_id    int32
	button_id int32
}

type TL_messages_acceptUrlAuth struct {
	flags         TL // #
	write_allowed TL // flags_0?true
	peer          TL // InputPeer
	msg_id        int32
	button_id     int32
}

type TL_messages_hidePeerSettingsBar struct {
	peer TL // InputPeer
}

type TL_contacts_addContact struct {
	flags                       TL // #
	add_phone_privacy_exception TL // flags_0?true
	id                          TL // InputUser
	first_name                  string
	last_name                   string
	phone                       string
}

type TL_contacts_acceptContact struct {
	id TL // InputUser
}

type TL_channels_editCreator struct {
	channel  TL // InputChannel
	user_id  TL // InputUser
	password TL // InputCheckPasswordSRP
}

type TL_contacts_getLocated struct {
	geo_point TL // InputGeoPoint
}

type TL_channels_editLocation struct {
	channel   TL // InputChannel
	geo_point TL // InputGeoPoint
	address   string
}

type TL_channels_toggleSlowMode struct {
	channel TL // InputChannel
	seconds int32
}

type TL_messages_getScheduledHistory struct {
	peer TL // InputPeer
	hash int32
}

type TL_messages_getScheduledMessages struct {
	peer TL // InputPeer
	id   []int32
}

type TL_messages_sendScheduledMessages struct {
	peer TL // InputPeer
	id   []int32
}

type TL_messages_deleteScheduledMessages struct {
	peer TL // InputPeer
	id   []int32
}

type TL_account_uploadTheme struct {
	flags     TL // #
	file      TL // InputFile
	thumb     TL // flags_0?InputFile
	file_name string
	mime_type string
}

type TL_account_createTheme struct {
	slug     string
	title    string
	document TL // InputDocument
}

type TL_account_updateTheme struct {
	flags    TL // #
	format   string
	theme    TL // InputTheme
	slug     TL // flags_0?string
	title    TL // flags_1?string
	document TL // flags_2?InputDocument
}

type TL_account_saveTheme struct {
	theme  TL // InputTheme
	unsave TL // Bool
}

type TL_account_installTheme struct {
	flags  TL // #
	dark   TL // flags_0?true
	format TL // flags_1?string
	theme  TL // flags_1?InputTheme
}

type TL_account_getTheme struct {
	format      string
	theme       TL // InputTheme
	document_id int64
}

type TL_account_getThemes struct {
	format string
	hash   int32
}

func (e TL_boolFalse) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolFalse)
	return x.buf
}

func (e TL_boolTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolTrue)
	return x.buf
}

func (e TL_true) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_true)
	return x.buf
}

func (e TL_error) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.code)
	x.String(e.text)
	return x.buf
}

func (e TL_null) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_null)
	return x.buf
}

func (e TL_inputPeerEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}

func (e TL_inputPeerSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}

func (e TL_inputPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_inputUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}

func (e TL_inputUserSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}

func (e TL_inputPhoneContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.client_id)
	x.String(e.phone)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_inputFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	x.String(e.md5_checksum)
	return x.buf
}

func (e TL_inputMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}

func (e TL_inputMediaUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.file.Encode())
	x.Bytes(e.stickers.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_inputMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_inputMediaGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.geo_point.Encode())
	return x.buf
}

func (e TL_inputMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.vcard)
	return x.buf
}

func (e TL_inputChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}

func (e TL_inputChatUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_inputChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_inputGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}

func (e TL_inputGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.lat)
	x.Double(e.long)
	return x.buf
}

func (e TL_inputPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}

func (e TL_inputPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	return x.buf
}

func (e TL_inputFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	x.StringBytes(e.file_reference)
	return x.buf
}

func (e TL_peerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_peerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_storage_fileUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileUnknown)
	return x.buf
}

func (e TL_storage_filePartial) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePartial)
	return x.buf
}

func (e TL_storage_fileJpeg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileJpeg)
	return x.buf
}

func (e TL_storage_fileGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileGif)
	return x.buf
}

func (e TL_storage_filePng) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePng)
	return x.buf
}

func (e TL_storage_filePdf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePdf)
	return x.buf
}

func (e TL_storage_fileMp3) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp3)
	return x.buf
}

func (e TL_storage_fileMov) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMov)
	return x.buf
}

func (e TL_storage_fileMp4) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp4)
	return x.buf
}

func (e TL_storage_fileWebp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileWebp)
	return x.buf
}

func (e TL_userEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_userProfilePhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}

func (e TL_userProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.photo_id)
	x.Bytes(e.photo_small.Encode())
	x.Bytes(e.photo_big.Encode())
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_userStatusEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}

func (e TL_userStatusOnline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.expires)
	return x.buf
}

func (e TL_userStatusOffline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.was_online)
	return x.buf
}

func (e TL_chatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_chat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.creator.Encode())
	x.Bytes(e.kicked.Encode())
	x.Bytes(e.left.Encode())
	x.Bytes(e.deactivated.Encode())
	x.Int(e.id)
	x.String(e.title)
	x.Bytes(e.photo.Encode())
	x.Int(e.participants_count)
	x.Int(e.date)
	x.Int(e.version)
	x.Bytes(e.migrated_to.Encode())
	x.Bytes(e.admin_rights.Encode())
	x.Bytes(e.default_banned_rights.Encode())
	return x.buf
}

func (e TL_chatForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.id)
	x.String(e.title)
	return x.buf
}

func (e TL_chatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.can_set_username.Encode())
	x.Bytes(e.has_scheduled.Encode())
	x.Int(e.id)
	x.String(e.about)
	x.Bytes(e.participants.Encode())
	x.Bytes(e.chat_photo.Encode())
	x.Bytes(e.notify_settings.Encode())
	x.Bytes(e.exported_invite.Encode())
	x.Bytes(e.bot_info.Encode())
	x.Bytes(e.pinned_msg_id.Encode())
	x.Bytes(e.folder_id.Encode())
	return x.buf
}

func (e TL_chatParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_chatParticipantsForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	x.Bytes(e.flags.Encode())
	x.Int(e.chat_id)
	x.Bytes(e.self_participant.Encode())
	return x.buf
}

func (e TL_chatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.chat_id)
	x.Vector(e.participants)
	x.Int(e.version)
	return x.buf
}

func (e TL_chatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}

func (e TL_chatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.photo_small.Encode())
	x.Bytes(e.photo_big.Encode())
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_messageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_message) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.out.Encode())
	x.Bytes(e.mentioned.Encode())
	x.Bytes(e.media_unread.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.post.Encode())
	x.Bytes(e.from_scheduled.Encode())
	x.Bytes(e.legacy.Encode())
	x.Bytes(e.edit_hide.Encode())
	x.Int(e.id)
	x.Bytes(e.from_id.Encode())
	x.Bytes(e.to_id.Encode())
	x.Bytes(e.fwd_from.Encode())
	x.Bytes(e.via_bot_id.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Int(e.date)
	x.String(e.message)
	x.Bytes(e.media.Encode())
	x.Bytes(e.reply_markup.Encode())
	x.Bytes(e.entities.Encode())
	x.Bytes(e.views.Encode())
	x.Bytes(e.edit_date.Encode())
	x.Bytes(e.post_author.Encode())
	x.Bytes(e.grouped_id.Encode())
	x.Bytes(e.restriction_reason.Encode())
	return x.buf
}

func (e TL_messageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.out.Encode())
	x.Bytes(e.mentioned.Encode())
	x.Bytes(e.media_unread.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.post.Encode())
	x.Bytes(e.legacy.Encode())
	x.Int(e.id)
	x.Bytes(e.from_id.Encode())
	x.Bytes(e.to_id.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Int(e.date)
	x.Bytes(e.action.Encode())
	return x.buf
}

func (e TL_messageMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}

func (e TL_messageMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.photo.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_messageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.geo.Encode())
	return x.buf
}

func (e TL_messageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.vcard)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_messageMediaUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	return x.buf
}

func (e TL_messageActionEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}

func (e TL_messageActionChatCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.title)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_messageActionChatEditTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.title)
	return x.buf
}

func (e TL_messageActionChatEditPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.photo.Encode())
	return x.buf
}

func (e TL_messageActionChatDeletePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}

func (e TL_messageActionChatAddUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_messageActionChatDeleteUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_dialog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.pinned.Encode())
	x.Bytes(e.unread_mark.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.top_message)
	x.Int(e.read_inbox_max_id)
	x.Int(e.read_outbox_max_id)
	x.Int(e.unread_count)
	x.Int(e.unread_mentions_count)
	x.Bytes(e.notify_settings.Encode())
	x.Bytes(e.pts.Encode())
	x.Bytes(e.draft.Encode())
	x.Bytes(e.folder_id.Encode())
	return x.buf
}

func (e TL_photoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_photo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.has_stickers.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	x.Int(e.date)
	x.Vector(e.sizes)
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_photoSizeEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e._type)
	return x.buf
}

func (e TL_photoSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e._type)
	x.Bytes(e.location.Encode())
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.size)
	return x.buf
}

func (e TL_photoCachedSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e._type)
	x.Bytes(e.location.Encode())
	x.Int(e.w)
	x.Int(e.h)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_geoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}

func (e TL_geoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.long)
	x.Double(e.lat)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_auth_sentCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCode)
	x.Bytes(e.flags.Encode())
	x.Bytes(e._type.Encode())
	x.String(e.phone_code_hash)
	x.Bytes(e.next_type.Encode())
	x.Bytes(e.timeout.Encode())
	return x.buf
}

func (e TL_auth_authorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorization)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.tmp_sessions.Encode())
	x.Bytes(e.user.Encode())
	return x.buf
}

func (e TL_auth_exportedAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportedAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_inputNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_inputNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}

func (e TL_inputNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}

func (e TL_inputPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.show_previews.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.mute_until.Encode())
	x.Bytes(e.sound.Encode())
	return x.buf
}

func (e TL_peerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.show_previews.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.mute_until.Encode())
	x.Bytes(e.sound.Encode())
	return x.buf
}

func (e TL_peerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.report_spam.Encode())
	x.Bytes(e.add_contact.Encode())
	x.Bytes(e.block_contact.Encode())
	x.Bytes(e.share_contact.Encode())
	x.Bytes(e.need_contacts_exception.Encode())
	x.Bytes(e.report_geo.Encode())
	return x.buf
}

func (e TL_wallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Long(e.id)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.creator.Encode())
	x.Bytes(e.defaults.Encode())
	x.Bytes(e.pattern.Encode())
	x.Bytes(e.dark.Encode())
	x.Long(e.access_hash)
	x.String(e.slug)
	x.Bytes(e.document.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_inputReportReasonSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonSpam)
	return x.buf
}

func (e TL_inputReportReasonViolence) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonViolence)
	return x.buf
}

func (e TL_inputReportReasonPornography) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonPornography)
	return x.buf
}

func (e TL_inputReportReasonChildAbuse) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonChildAbuse)
	return x.buf
}

func (e TL_inputReportReasonOther) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonOther)
	x.String(e.text)
	return x.buf
}

func (e TL_userFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.blocked.Encode())
	x.Bytes(e.phone_calls_available.Encode())
	x.Bytes(e.phone_calls_private.Encode())
	x.Bytes(e.can_pin_message.Encode())
	x.Bytes(e.has_scheduled.Encode())
	x.Bytes(e.user.Encode())
	x.Bytes(e.about.Encode())
	x.Bytes(e.settings.Encode())
	x.Bytes(e.profile_photo.Encode())
	x.Bytes(e.notify_settings.Encode())
	x.Bytes(e.bot_info.Encode())
	x.Bytes(e.pinned_msg_id.Encode())
	x.Int(e.common_chats_count)
	x.Bytes(e.folder_id.Encode())
	return x.buf
}

func (e TL_contact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.user_id)
	x.Bytes(e.mutual.Encode())
	return x.buf
}

func (e TL_importedContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.user_id)
	x.Long(e.client_id)
	return x.buf
}

func (e TL_contactBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_contactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.Encode())
	return x.buf
}

func (e TL_contacts_contactsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contactsNotModified)
	return x.buf
}

func (e TL_contacts_contacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contacts)
	x.Vector(e.contacts)
	x.Int(e.saved_count)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_importedContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importedContacts)
	x.Vector(e.imported)
	x.Vector(e.popular_invites)
	x.VectorLong(e.retry_contacts)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blocked)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blockedSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blockedSlice)
	x.Int(e.count)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogs)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsSlice)
	x.Int(e.count)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messages)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messagesSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesSlice)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.inexact.Encode())
	x.Int(e.count)
	x.Bytes(e.next_rate.Encode())
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_chats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chats)
	x.Vector(e.chats)
	return x.buf
}

func (e TL_messages_chatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatFull)
	x.Bytes(e.full_chat.Encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_affectedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedHistory)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.offset)
	return x.buf
}

func (e TL_inputMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}

func (e TL_inputMessagesFilterPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TL_inputMessagesFilterDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}

func (e TL_inputMessagesFilterUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterUrl)
	return x.buf
}

func (e TL_inputMessagesFilterGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGif)
	return x.buf
}

func (e TL_updateNewMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.message.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_updateDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.user_id)
	x.Bytes(e.action.Encode())
	return x.buf
}

func (e TL_updateChatUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Bytes(e.action.Encode())
	return x.buf
}

func (e TL_updateChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.participants.Encode())
	return x.buf
}

func (e TL_updateUserStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.Encode())
	return x.buf
}

func (e TL_updateUserName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.user_id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	return x.buf
}

func (e TL_updateUserPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.user_id)
	x.Int(e.date)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.previous.Encode())
	return x.buf
}

func (e TL_updates_state) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_state)
	x.Int(e.pts)
	x.Int(e.qts)
	x.Int(e.date)
	x.Int(e.seq)
	x.Int(e.unread_count)
	return x.buf
}

func (e TL_updates_differenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceEmpty)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates_difference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_difference)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.state.Encode())
	return x.buf
}

func (e TL_updates_differenceSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceSlice)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.intermediate_state.Encode())
	return x.buf
}

func (e TL_updatesTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}

func (e TL_updateShortMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.out.Encode())
	x.Bytes(e.mentioned.Encode())
	x.Bytes(e.media_unread.Encode())
	x.Bytes(e.silent.Encode())
	x.Int(e.id)
	x.Int(e.user_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.fwd_from.Encode())
	x.Bytes(e.via_bot_id.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_updateShortChatMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.out.Encode())
	x.Bytes(e.mentioned.Encode())
	x.Bytes(e.media_unread.Encode())
	x.Bytes(e.silent.Encode())
	x.Int(e.id)
	x.Int(e.from_id)
	x.Int(e.chat_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.fwd_from.Encode())
	x.Bytes(e.via_bot_id.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_updateShort) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.update.Encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updatesCombined) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq_start)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_photos_photos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photos)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photosSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photosSlice)
	x.Int(e.count)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photo)
	x.Bytes(e.photo.Encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_upload_file) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_file)
	x.Bytes(e._type.Encode())
	x.Int(e.mtime)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_dcOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.ipv6.Encode())
	x.Bytes(e.media_only.Encode())
	x.Bytes(e.tcpo_only.Encode())
	x.Bytes(e.cdn.Encode())
	x.Bytes(e.static.Encode())
	x.Int(e.id)
	x.String(e.ip_address)
	x.Int(e.port)
	x.Bytes(e.secret.Encode())
	return x.buf
}

func (e TL_config) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.phonecalls_enabled.Encode())
	x.Bytes(e.default_p2p_contacts.Encode())
	x.Bytes(e.preload_featured_stickers.Encode())
	x.Bytes(e.ignore_phone_entities.Encode())
	x.Bytes(e.revoke_pm_inbox.Encode())
	x.Bytes(e.blocked_mode.Encode())
	x.Bytes(e.pfs_enabled.Encode())
	x.Int(e.date)
	x.Int(e.expires)
	x.Bytes(e.test_mode.Encode())
	x.Int(e.this_dc)
	x.Vector(e.dc_options)
	x.String(e.dc_txt_domain_name)
	x.Int(e.chat_size_max)
	x.Int(e.megagroup_size_max)
	x.Int(e.forwarded_count_max)
	x.Int(e.online_update_period_ms)
	x.Int(e.offline_blur_timeout_ms)
	x.Int(e.offline_idle_timeout_ms)
	x.Int(e.online_cloud_timeout_ms)
	x.Int(e.notify_cloud_delay_ms)
	x.Int(e.notify_default_delay_ms)
	x.Int(e.push_chat_period_ms)
	x.Int(e.push_chat_limit)
	x.Int(e.saved_gifs_limit)
	x.Int(e.edit_time_limit)
	x.Int(e.revoke_time_limit)
	x.Int(e.revoke_pm_time_limit)
	x.Int(e.rating_e_decay)
	x.Int(e.stickers_recent_limit)
	x.Int(e.stickers_faved_limit)
	x.Int(e.channels_read_media_period)
	x.Bytes(e.tmp_sessions.Encode())
	x.Int(e.pinned_dialogs_count_max)
	x.Int(e.pinned_infolder_count_max)
	x.Int(e.call_receive_timeout_ms)
	x.Int(e.call_ring_timeout_ms)
	x.Int(e.call_connect_timeout_ms)
	x.Int(e.call_packet_timeout_ms)
	x.String(e.me_url_prefix)
	x.Bytes(e.autoupdate_url_prefix.Encode())
	x.Bytes(e.gif_search_username.Encode())
	x.Bytes(e.venue_search_username.Encode())
	x.Bytes(e.img_search_username.Encode())
	x.Bytes(e.static_maps_provider.Encode())
	x.Int(e.caption_length_max)
	x.Int(e.message_length_max)
	x.Int(e.webfile_dc_id)
	x.Bytes(e.suggested_lang_code.Encode())
	x.Bytes(e.lang_pack_version.Encode())
	x.Bytes(e.base_lang_pack_version.Encode())
	return x.buf
}

func (e TL_nearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.country)
	x.Int(e.this_dc)
	x.Int(e.nearest_dc)
	return x.buf
}

func (e TL_help_appUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appUpdate)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.can_not_skip.Encode())
	x.Int(e.id)
	x.String(e.version)
	x.String(e.text)
	x.Vector(e.entities)
	x.Bytes(e.document.Encode())
	x.Bytes(e.url.Encode())
	return x.buf
}

func (e TL_help_noAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_noAppUpdate)
	return x.buf
}

func (e TL_help_inviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_inviteText)
	x.String(e.message)
	return x.buf
}

func (e TL_updateNewEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.message.Encode())
	x.Int(e.qts)
	return x.buf
}

func (e TL_updateEncryptedChatTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_updateEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.chat.Encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updateEncryptedMessagesRead) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.chat_id)
	x.Int(e.max_date)
	x.Int(e.date)
	return x.buf
}

func (e TL_encryptedChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_encryptedChatWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	return x.buf
}

func (e TL_encryptedChatRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_encryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a_or_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_encryptedChatDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.id)
	return x.buf
}

func (e TL_inputEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.chat_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}

func (e TL_encryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.size)
	x.Int(e.dc_id)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}

func (e TL_inputEncryptedFileUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.md5_checksum)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputEncryptedFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_encryptedMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_messages_dhConfigNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfigNotModified)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_dhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfig)
	x.Int(e.g)
	x.StringBytes(e.p)
	x.Int(e.version)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_sentEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedMessage)
	x.Int(e.date)
	return x.buf
}

func (e TL_messages_sentEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedFile)
	x.Int(e.date)
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_inputFileBig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	return x.buf
}

func (e TL_inputEncryptedFileBigUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_updateChatParticipantAdd) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateChatParticipantDelete) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateDcOptions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(e.dc_options)
	return x.buf
}

func (e TL_inputMediaUploadedDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.nosound_video.Encode())
	x.Bytes(e.file.Encode())
	x.Bytes(e.thumb.Encode())
	x.String(e.mime_type)
	x.Vector(e.attributes)
	x.Bytes(e.stickers.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_inputMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_messageMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.document.Encode())
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_inputDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}

func (e TL_inputDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	return x.buf
}

func (e TL_inputDocumentFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	x.String(e.thumb_size)
	return x.buf
}

func (e TL_documentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_document) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Bytes(e.flags.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	x.Int(e.date)
	x.String(e.mime_type)
	x.Int(e.size)
	x.Bytes(e.thumbs.Encode())
	x.Int(e.dc_id)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_help_support) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_support)
	x.String(e.phone_number)
	x.Bytes(e.user.Encode())
	return x.buf
}

func (e TL_notifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_notifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}

func (e TL_notifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}

func (e TL_updateUserBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.user_id)
	x.Bytes(e.blocked.Encode())
	return x.buf
}

func (e TL_updateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.notify_settings.Encode())
	return x.buf
}

func (e TL_sendMessageTypingAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}

func (e TL_sendMessageCancelAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}

func (e TL_sendMessageRecordVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}

func (e TL_sendMessageUploadVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageRecordAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageUploadPhotoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageUploadDocumentAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageGeoLocationAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}

func (e TL_sendMessageChooseContactAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}

func (e TL_contacts_found) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_found)
	x.Vector(e.my_results)
	x.Vector(e.results)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updateServiceNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.popup.Encode())
	x.Bytes(e.inbox_date.Encode())
	x.String(e._type)
	x.String(e.message)
	x.Bytes(e.media.Encode())
	x.Vector(e.entities)
	return x.buf
}

func (e TL_userStatusRecently) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}

func (e TL_userStatusLastWeek) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}

func (e TL_userStatusLastMonth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}

func (e TL_updatePrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.key.Encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_privacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_inputPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}

func (e TL_inputPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_privacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}

func (e TL_privacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}

func (e TL_privacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_privacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}

func (e TL_privacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}

func (e TL_privacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_account_privacyRules) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_privacyRules)
	x.Vector(e.rules)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_accountDaysTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.days)
	return x.buf
}

func (e TL_updateUserPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.user_id)
	x.String(e.phone)
	return x.buf
}

func (e TL_documentAttributeImageSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAnimated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}

func (e TL_documentAttributeSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.mask.Encode())
	x.String(e.alt)
	x.Bytes(e.stickerset.Encode())
	x.Bytes(e.mask_coords.Encode())
	return x.buf
}

func (e TL_documentAttributeVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.round_message.Encode())
	x.Bytes(e.supports_streaming.Encode())
	x.Int(e.duration)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.voice.Encode())
	x.Int(e.duration)
	x.Bytes(e.title.Encode())
	x.Bytes(e.performer.Encode())
	x.Bytes(e.waveform.Encode())
	return x.buf
}

func (e TL_documentAttributeFilename) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.file_name)
	return x.buf
}

func (e TL_messages_stickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickersNotModified)
	return x.buf
}

func (e TL_messages_stickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickers)
	x.Int(e.hash)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_stickerPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.emoticon)
	x.VectorLong(e.documents)
	return x.buf
}

func (e TL_messages_allStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickersNotModified)
	return x.buf
}

func (e TL_messages_allStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickers)
	x.Int(e.hash)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_updateReadHistoryInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryInbox)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.max_id)
	x.Int(e.still_unread_count)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateReadHistoryOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryOutbox)
	x.Bytes(e.peer.Encode())
	x.Int(e.max_id)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_messages_affectedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedMessages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateWebPage)
	x.Bytes(e.webpage.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_webPageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_webPagePending) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPagePending)
	x.Long(e.id)
	x.Int(e.date)
	return x.buf
}

func (e TL_webPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPage)
	x.Bytes(e.flags.Encode())
	x.Long(e.id)
	x.String(e.url)
	x.String(e.display_url)
	x.Int(e.hash)
	x.Bytes(e._type.Encode())
	x.Bytes(e.site_name.Encode())
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.photo.Encode())
	x.Bytes(e.embed_url.Encode())
	x.Bytes(e.embed_type.Encode())
	x.Bytes(e.embed_width.Encode())
	x.Bytes(e.embed_height.Encode())
	x.Bytes(e.duration.Encode())
	x.Bytes(e.author.Encode())
	x.Bytes(e.document.Encode())
	x.Bytes(e.documents.Encode())
	x.Bytes(e.cached_page.Encode())
	return x.buf
}

func (e TL_messageMediaWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaWebPage)
	x.Bytes(e.webpage.Encode())
	return x.buf
}

func (e TL_authorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authorization)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.current.Encode())
	x.Bytes(e.official_app.Encode())
	x.Bytes(e.password_pending.Encode())
	x.Long(e.hash)
	x.String(e.device_model)
	x.String(e.platform)
	x.String(e.system_version)
	x.Int(e.api_id)
	x.String(e.app_name)
	x.String(e.app_version)
	x.Int(e.date_created)
	x.Int(e.date_active)
	x.String(e.ip)
	x.String(e.country)
	x.String(e.region)
	return x.buf
}

func (e TL_account_authorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_authorizations)
	x.Vector(e.authorizations)
	return x.buf
}

func (e TL_account_password) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_password)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.has_recovery.Encode())
	x.Bytes(e.has_secure_values.Encode())
	x.Bytes(e.has_password.Encode())
	x.Bytes(e.current_algo.Encode())
	x.Bytes(e.srp_B.Encode())
	x.Bytes(e.srp_id.Encode())
	x.Bytes(e.hint.Encode())
	x.Bytes(e.email_unconfirmed_pattern.Encode())
	x.Bytes(e.new_algo.Encode())
	x.Bytes(e.new_secure_algo.Encode())
	x.StringBytes(e.secure_random)
	return x.buf
}

func (e TL_account_passwordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.email.Encode())
	x.Bytes(e.secure_settings.Encode())
	return x.buf
}

func (e TL_account_passwordInputSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordInputSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.new_algo.Encode())
	x.Bytes(e.new_password_hash.Encode())
	x.Bytes(e.hint.Encode())
	x.Bytes(e.email.Encode())
	x.Bytes(e.new_secure_settings.Encode())
	return x.buf
}

func (e TL_auth_passwordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_passwordRecovery)
	x.String(e.email_pattern)
	return x.buf
}

func (e TL_inputMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVenue)
	x.Bytes(e.geo_point.Encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.String(e.venue_type)
	return x.buf
}

func (e TL_messageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVenue)
	x.Bytes(e.geo.Encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.String(e.venue_type)
	return x.buf
}

func (e TL_receivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_receivedNotifyMessage)
	x.Int(e.id)
	x.Int(e.flags)
	return x.buf
}

func (e TL_chatInviteEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteEmpty)
	return x.buf
}

func (e TL_chatInviteExported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteExported)
	x.String(e.link)
	return x.buf
}

func (e TL_chatInviteAlready) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteAlready)
	x.Bytes(e.chat.Encode())
	return x.buf
}

func (e TL_chatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInvite)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.channel.Encode())
	x.Bytes(e.broadcast.Encode())
	x.Bytes(e.public.Encode())
	x.Bytes(e.megagroup.Encode())
	x.String(e.title)
	x.Bytes(e.photo.Encode())
	x.Int(e.participants_count)
	x.Bytes(e.participants.Encode())
	return x.buf
}

func (e TL_messageActionChatJoinedByLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatJoinedByLink)
	x.Int(e.inviter_id)
	return x.buf
}

func (e TL_updateReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessagesContents)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_inputStickerSetEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetEmpty)
	return x.buf
}

func (e TL_inputStickerSetID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetID)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputStickerSetShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetShortName)
	x.String(e.short_name)
	return x.buf
}

func (e TL_stickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSet)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.archived.Encode())
	x.Bytes(e.official.Encode())
	x.Bytes(e.masks.Encode())
	x.Bytes(e.animated.Encode())
	x.Bytes(e.installed_date.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	x.String(e.short_name)
	x.Bytes(e.thumb.Encode())
	x.Bytes(e.thumb_dc_id.Encode())
	x.Int(e.count)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_stickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSet)
	x.Bytes(e.set.Encode())
	x.Vector(e.packs)
	x.Vector(e.documents)
	return x.buf
}

func (e TL_user) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_user)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.self.Encode())
	x.Bytes(e.contact.Encode())
	x.Bytes(e.mutual_contact.Encode())
	x.Bytes(e.deleted.Encode())
	x.Bytes(e.bot.Encode())
	x.Bytes(e.bot_chat_history.Encode())
	x.Bytes(e.bot_nochats.Encode())
	x.Bytes(e.verified.Encode())
	x.Bytes(e.restricted.Encode())
	x.Bytes(e.min.Encode())
	x.Bytes(e.bot_inline_geo.Encode())
	x.Bytes(e.support.Encode())
	x.Bytes(e.scam.Encode())
	x.Int(e.id)
	x.Bytes(e.access_hash.Encode())
	x.Bytes(e.first_name.Encode())
	x.Bytes(e.last_name.Encode())
	x.Bytes(e.username.Encode())
	x.Bytes(e.phone.Encode())
	x.Bytes(e.photo.Encode())
	x.Bytes(e.status.Encode())
	x.Bytes(e.bot_info_version.Encode())
	x.Bytes(e.restriction_reason.Encode())
	x.Bytes(e.bot_inline_placeholder.Encode())
	x.Bytes(e.lang_code.Encode())
	return x.buf
}

func (e TL_botCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botCommand)
	x.String(e.command)
	x.String(e.description)
	return x.buf
}

func (e TL_botInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInfo)
	x.Int(e.user_id)
	x.String(e.description)
	x.Vector(e.commands)
	return x.buf
}

func (e TL_keyboardButton) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButton)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonRow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRow)
	x.Vector(e.buttons)
	return x.buf
}

func (e TL_replyKeyboardHide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardHide)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.selective.Encode())
	return x.buf
}

func (e TL_replyKeyboardForceReply) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardForceReply)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.single_use.Encode())
	x.Bytes(e.selective.Encode())
	return x.buf
}

func (e TL_replyKeyboardMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardMarkup)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.resize.Encode())
	x.Bytes(e.single_use.Encode())
	x.Bytes(e.selective.Encode())
	x.Vector(e.rows)
	return x.buf
}

func (e TL_inputPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUser)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUser)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_messageEntityUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnknown)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityMention) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMention)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityHashtag) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityHashtag)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBotCommand)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUrl)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityEmail)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBold)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityItalic)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCode)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityPre) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPre)
	x.Int(e.offset)
	x.Int(e.length)
	x.String(e.language)
	return x.buf
}

func (e TL_messageEntityTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityTextUrl)
	x.Int(e.offset)
	x.Int(e.length)
	x.String(e.url)
	return x.buf
}

func (e TL_updateShortSentMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortSentMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.out.Encode())
	x.Int(e.id)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.media.Encode())
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_inputChannelEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelEmpty)
	return x.buf
}

func (e TL_inputChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannel)
	x.Int(e.channel_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_peerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChannel)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_inputPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannel)
	x.Int(e.channel_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_channel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channel)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.creator.Encode())
	x.Bytes(e.left.Encode())
	x.Bytes(e.broadcast.Encode())
	x.Bytes(e.verified.Encode())
	x.Bytes(e.megagroup.Encode())
	x.Bytes(e.restricted.Encode())
	x.Bytes(e.signatures.Encode())
	x.Bytes(e.min.Encode())
	x.Bytes(e.scam.Encode())
	x.Bytes(e.has_link.Encode())
	x.Bytes(e.has_geo.Encode())
	x.Bytes(e.slowmode_enabled.Encode())
	x.Int(e.id)
	x.Bytes(e.access_hash.Encode())
	x.String(e.title)
	x.Bytes(e.username.Encode())
	x.Bytes(e.photo.Encode())
	x.Int(e.date)
	x.Int(e.version)
	x.Bytes(e.restriction_reason.Encode())
	x.Bytes(e.admin_rights.Encode())
	x.Bytes(e.banned_rights.Encode())
	x.Bytes(e.default_banned_rights.Encode())
	x.Bytes(e.participants_count.Encode())
	return x.buf
}

func (e TL_channelForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelForbidden)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.broadcast.Encode())
	x.Bytes(e.megagroup.Encode())
	x.Int(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	x.Bytes(e.until_date.Encode())
	return x.buf
}

func (e TL_contacts_resolvedPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolvedPeer)
	x.Bytes(e.peer.Encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channelFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelFull)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.can_view_participants.Encode())
	x.Bytes(e.can_set_username.Encode())
	x.Bytes(e.can_set_stickers.Encode())
	x.Bytes(e.hidden_prehistory.Encode())
	x.Bytes(e.can_view_stats.Encode())
	x.Bytes(e.can_set_location.Encode())
	x.Bytes(e.has_scheduled.Encode())
	x.Int(e.id)
	x.String(e.about)
	x.Bytes(e.participants_count.Encode())
	x.Bytes(e.admins_count.Encode())
	x.Bytes(e.kicked_count.Encode())
	x.Bytes(e.banned_count.Encode())
	x.Bytes(e.online_count.Encode())
	x.Int(e.read_inbox_max_id)
	x.Int(e.read_outbox_max_id)
	x.Int(e.unread_count)
	x.Bytes(e.chat_photo.Encode())
	x.Bytes(e.notify_settings.Encode())
	x.Bytes(e.exported_invite.Encode())
	x.Vector(e.bot_info)
	x.Bytes(e.migrated_from_chat_id.Encode())
	x.Bytes(e.migrated_from_max_id.Encode())
	x.Bytes(e.pinned_msg_id.Encode())
	x.Bytes(e.stickerset.Encode())
	x.Bytes(e.available_min_id.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Bytes(e.linked_chat_id.Encode())
	x.Bytes(e.location.Encode())
	x.Bytes(e.slowmode_seconds.Encode())
	x.Bytes(e.slowmode_next_send_date.Encode())
	x.Int(e.pts)
	return x.buf
}

func (e TL_messageRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageRange)
	x.Int(e.min_id)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_channelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_channelMessages)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.inexact.Encode())
	x.Int(e.pts)
	x.Int(e.count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messageActionChannelCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelCreate)
	x.String(e.title)
	return x.buf
}

func (e TL_updateChannelTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelTooLong)
	x.Bytes(e.flags.Encode())
	x.Int(e.channel_id)
	x.Bytes(e.pts.Encode())
	return x.buf
}

func (e TL_updateChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannel)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_updateNewChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewChannelMessage)
	x.Bytes(e.message.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateReadChannelInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelInbox)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Int(e.channel_id)
	x.Int(e.max_id)
	x.Int(e.still_unread_count)
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateDeleteChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteChannelMessages)
	x.Int(e.channel_id)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateChannelMessageViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelMessageViews)
	x.Int(e.channel_id)
	x.Int(e.id)
	x.Int(e.views)
	return x.buf
}

func (e TL_updates_channelDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceEmpty)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.final.Encode())
	x.Int(e.pts)
	x.Bytes(e.timeout.Encode())
	return x.buf
}

func (e TL_updates_channelDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceTooLong)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.final.Encode())
	x.Bytes(e.timeout.Encode())
	x.Bytes(e.dialog.Encode())
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updates_channelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifference)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.final.Encode())
	x.Int(e.pts)
	x.Bytes(e.timeout.Encode())
	x.Vector(e.new_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channelMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilterEmpty)
	return x.buf
}

func (e TL_channelMessagesFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilter)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.exclude_new_messages.Encode())
	x.Vector(e.ranges)
	return x.buf
}

func (e TL_channelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipant)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantSelf)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantCreator)
	x.Bytes(e.flags.Encode())
	x.Int(e.user_id)
	x.Bytes(e.rank.Encode())
	return x.buf
}

func (e TL_channelParticipantsRecent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsRecent)
	return x.buf
}

func (e TL_channelParticipantsAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsAdmins)
	return x.buf
}

func (e TL_channelParticipantsKicked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsKicked)
	x.String(e.q)
	return x.buf
}

func (e TL_channels_channelParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipants)
	x.Int(e.count)
	x.Vector(e.participants)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channels_channelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipant)
	x.Bytes(e.participant.Encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_chatParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantCreator)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_chatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantAdmin)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_updateChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdmin)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Bytes(e.is_admin.Encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_messageActionChatMigrateTo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatMigrateTo)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_messageActionChannelMigrateFrom) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelMigrateFrom)
	x.String(e.title)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_channelParticipantsBots) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBots)
	return x.buf
}

func (e TL_help_termsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfService)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.popup.Encode())
	x.Bytes(e.id.Encode())
	x.String(e.text)
	x.Vector(e.entities)
	x.Bytes(e.min_age_confirm.Encode())
	return x.buf
}

func (e TL_updateNewStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewStickerSet)
	x.Bytes(e.stickerset.Encode())
	return x.buf
}

func (e TL_updateStickerSetsOrder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSetsOrder)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.masks.Encode())
	x.VectorLong(e.order)
	return x.buf
}

func (e TL_updateStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSets)
	return x.buf
}

func (e TL_foundGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGif)
	x.String(e.url)
	x.String(e.thumb_url)
	x.String(e.content_url)
	x.String(e.content_type)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_foundGifCached) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGifCached)
	x.String(e.url)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.document.Encode())
	return x.buf
}

func (e TL_inputMediaGifExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGifExternal)
	x.String(e.url)
	x.String(e.q)
	return x.buf
}

func (e TL_messages_foundGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundGifs)
	x.Int(e.next_offset)
	x.Vector(e.results)
	return x.buf
}

func (e TL_messages_savedGifsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifsNotModified)
	return x.buf
}

func (e TL_messages_savedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifs)
	x.Int(e.hash)
	x.Vector(e.gifs)
	return x.buf
}

func (e TL_updateSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateSavedGifs)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaAuto)
	x.Bytes(e.flags.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_inputBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageText)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_inputBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResult)
	x.Bytes(e.flags.Encode())
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.url.Encode())
	x.Bytes(e.thumb.Encode())
	x.Bytes(e.content.Encode())
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_botInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaAuto)
	x.Bytes(e.flags.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_botInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageText)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_botInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineResult)
	x.Bytes(e.flags.Encode())
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.url.Encode())
	x.Bytes(e.thumb.Encode())
	x.Bytes(e.content.Encode())
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_messages_botResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botResults)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.gallery.Encode())
	x.Long(e.query_id)
	x.Bytes(e.next_offset.Encode())
	x.Bytes(e.switch_pm.Encode())
	x.Vector(e.results)
	x.Int(e.cache_time)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updateBotInlineQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineQuery)
	x.Bytes(e.flags.Encode())
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.String(e.query)
	x.Bytes(e.geo.Encode())
	x.String(e.offset)
	return x.buf
}

func (e TL_updateBotInlineSend) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineSend)
	x.Bytes(e.flags.Encode())
	x.Int(e.user_id)
	x.String(e.query)
	x.Bytes(e.geo.Encode())
	x.String(e.id)
	x.Bytes(e.msg_id.Encode())
	return x.buf
}

func (e TL_inputMessagesFilterVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVoice)
	return x.buf
}

func (e TL_inputMessagesFilterMusic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMusic)
	return x.buf
}

func (e TL_inputPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyChatInvite)
	return x.buf
}

func (e TL_privacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyChatInvite)
	return x.buf
}

func (e TL_exportedMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_exportedMessageLink)
	x.String(e.link)
	x.String(e.html)
	return x.buf
}

func (e TL_messageFwdHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageFwdHeader)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.from_id.Encode())
	x.Bytes(e.from_name.Encode())
	x.Int(e.date)
	x.Bytes(e.channel_id.Encode())
	x.Bytes(e.channel_post.Encode())
	x.Bytes(e.post_author.Encode())
	x.Bytes(e.saved_from_peer.Encode())
	x.Bytes(e.saved_from_msg_id.Encode())
	return x.buf
}

func (e TL_updateEditChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditChannelMessage)
	x.Bytes(e.message.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateChannelPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelPinnedMessage)
	x.Int(e.channel_id)
	x.Int(e.id)
	return x.buf
}

func (e TL_messageActionPinMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPinMessage)
	return x.buf
}

func (e TL_auth_codeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeSms)
	return x.buf
}

func (e TL_auth_codeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeCall)
	return x.buf
}

func (e TL_auth_codeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeFlashCall)
	return x.buf
}

func (e TL_auth_sentCodeTypeApp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeApp)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeSms)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeCall)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeFlashCall)
	x.String(e.pattern)
	return x.buf
}

func (e TL_keyboardButtonUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrl)
	x.String(e.text)
	x.String(e.url)
	return x.buf
}

func (e TL_keyboardButtonCallback) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonCallback)
	x.String(e.text)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_keyboardButtonRequestPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestPhone)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonRequestGeoLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestGeoLocation)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonSwitchInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonSwitchInline)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.same_peer.Encode())
	x.String(e.text)
	x.String(e.query)
	return x.buf
}

func (e TL_replyInlineMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyInlineMarkup)
	x.Vector(e.rows)
	return x.buf
}

func (e TL_messages_botCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botCallbackAnswer)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.alert.Encode())
	x.Bytes(e.has_url.Encode())
	x.Bytes(e.native_ui.Encode())
	x.Bytes(e.message.Encode())
	x.Bytes(e.url.Encode())
	x.Int(e.cache_time)
	return x.buf
}

func (e TL_updateBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotCallbackQuery)
	x.Bytes(e.flags.Encode())
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Long(e.chat_instance)
	x.Bytes(e.data.Encode())
	x.Bytes(e.game_short_name.Encode())
	return x.buf
}

func (e TL_messages_messageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messageEditData)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_updateEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditMessage)
	x.Bytes(e.message.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaGeo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.geo_point.Encode())
	x.Int(e.period)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaVenue)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.geo_point.Encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.String(e.venue_type)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaContact)
	x.Bytes(e.flags.Encode())
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.vcard)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_botInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaGeo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.geo.Encode())
	x.Int(e.period)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_botInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaVenue)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.geo.Encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.String(e.venue_type)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_botInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaContact)
	x.Bytes(e.flags.Encode())
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.vcard)
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_inputBotInlineResultPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultPhoto)
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_inputBotInlineResultDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultDocument)
	x.Bytes(e.flags.Encode())
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.document.Encode())
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_botInlineMediaResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMediaResult)
	x.Bytes(e.flags.Encode())
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.document.Encode())
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_inputBotInlineMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageID)
	x.Int(e.dc_id)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_updateInlineBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateInlineBotCallbackQuery)
	x.Bytes(e.flags.Encode())
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.Bytes(e.msg_id.Encode())
	x.Long(e.chat_instance)
	x.Bytes(e.data.Encode())
	x.Bytes(e.game_short_name.Encode())
	return x.buf
}

func (e TL_inlineBotSwitchPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inlineBotSwitchPM)
	x.String(e.text)
	x.String(e.start_param)
	return x.buf
}

func (e TL_messages_peerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_peerDialogs)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.state.Encode())
	return x.buf
}

func (e TL_topPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeer)
	x.Bytes(e.peer.Encode())
	x.Double(e.rating)
	return x.buf
}

func (e TL_topPeerCategoryBotsPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsPM)
	return x.buf
}

func (e TL_topPeerCategoryBotsInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsInline)
	return x.buf
}

func (e TL_topPeerCategoryCorrespondents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryCorrespondents)
	return x.buf
}

func (e TL_topPeerCategoryGroups) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryGroups)
	return x.buf
}

func (e TL_topPeerCategoryChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryChannels)
	return x.buf
}

func (e TL_topPeerCategoryPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPeers)
	x.Bytes(e.category.Encode())
	x.Int(e.count)
	x.Vector(e.peers)
	return x.buf
}

func (e TL_contacts_topPeersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeersNotModified)
	return x.buf
}

func (e TL_contacts_topPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeers)
	x.Vector(e.categories)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMentionName)
	x.Int(e.offset)
	x.Int(e.length)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageEntityMentionName)
	x.Int(e.offset)
	x.Int(e.length)
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_inputMessagesFilterChatPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterChatPhotos)
	return x.buf
}

func (e TL_updateReadChannelOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelOutbox)
	x.Int(e.channel_id)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_updateDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDraftMessage)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.draft.Encode())
	return x.buf
}

func (e TL_draftMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessageEmpty)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.date.Encode())
	return x.buf
}

func (e TL_draftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_messageActionHistoryClear) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionHistoryClear)
	return x.buf
}

func (e TL_messages_featuredStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickersNotModified)
	return x.buf
}

func (e TL_messages_featuredStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickers)
	x.Int(e.hash)
	x.Vector(e.sets)
	x.VectorLong(e.unread)
	return x.buf
}

func (e TL_updateReadFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadFeaturedStickers)
	return x.buf
}

func (e TL_messages_recentStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickersNotModified)
	return x.buf
}

func (e TL_messages_recentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickers)
	x.Int(e.hash)
	x.Vector(e.packs)
	x.Vector(e.stickers)
	x.VectorInt(e.dates)
	return x.buf
}

func (e TL_updateRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateRecentStickers)
	return x.buf
}

func (e TL_messages_archivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_archivedStickers)
	x.Int(e.count)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultSuccess) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultSuccess)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultArchive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultArchive)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_stickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetCovered)
	x.Bytes(e.set.Encode())
	x.Bytes(e.cover.Encode())
	return x.buf
}

func (e TL_updateConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateConfig)
	return x.buf
}

func (e TL_updatePtsChanged) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePtsChanged)
	return x.buf
}

func (e TL_inputMediaPhotoExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhotoExternal)
	x.Bytes(e.flags.Encode())
	x.String(e.url)
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_inputMediaDocumentExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocumentExternal)
	x.Bytes(e.flags.Encode())
	x.String(e.url)
	x.Bytes(e.ttl_seconds.Encode())
	return x.buf
}

func (e TL_stickerSetMultiCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetMultiCovered)
	x.Bytes(e.set.Encode())
	x.Vector(e.covers)
	return x.buf
}

func (e TL_maskCoords) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_maskCoords)
	x.Int(e.n)
	x.Double(e.x)
	x.Double(e.y)
	x.Double(e.zoom)
	return x.buf
}

func (e TL_documentAttributeHasStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeHasStickers)
	return x.buf
}

func (e TL_inputStickeredMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaPhoto)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_inputStickeredMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaDocument)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_game) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_game)
	x.Bytes(e.flags.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.String(e.short_name)
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.document.Encode())
	return x.buf
}

func (e TL_inputBotInlineResultGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultGame)
	x.String(e.id)
	x.String(e.short_name)
	x.Bytes(e.send_message.Encode())
	return x.buf
}

func (e TL_inputBotInlineMessageGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageGame)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.reply_markup.Encode())
	return x.buf
}

func (e TL_messageMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGame)
	x.Bytes(e.game.Encode())
	return x.buf
}

func (e TL_inputMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGame)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_inputGameID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameID)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputGameShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameShortName)
	x.Bytes(e.bot_id.Encode())
	x.String(e.short_name)
	return x.buf
}

func (e TL_keyboardButtonGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonGame)
	x.String(e.text)
	return x.buf
}

func (e TL_messageActionGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGameScore)
	x.Long(e.game_id)
	x.Int(e.score)
	return x.buf
}

func (e TL_highScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_highScore)
	x.Int(e.pos)
	x.Int(e.user_id)
	x.Int(e.score)
	return x.buf
}

func (e TL_messages_highScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_highScores)
	x.Vector(e.scores)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updates_differenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceTooLong)
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateChannelWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelWebPage)
	x.Int(e.channel_id)
	x.Bytes(e.webpage.Encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_messages_chatsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatsSlice)
	x.Int(e.count)
	x.Vector(e.chats)
	return x.buf
}

func (e TL_textEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmpty)
	return x.buf
}

func (e TL_textPlain) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textPlain)
	x.String(e.text)
	return x.buf
}

func (e TL_textBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textBold)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textItalic)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textUnderline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUnderline)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textStrike) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textStrike)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textFixed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textFixed)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUrl)
	x.Bytes(e.text.Encode())
	x.String(e.url)
	x.Long(e.webpage_id)
	return x.buf
}

func (e TL_textEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmail)
	x.Bytes(e.text.Encode())
	x.String(e.email)
	return x.buf
}

func (e TL_textConcat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textConcat)
	x.Vector(e.texts)
	return x.buf
}

func (e TL_pageBlockUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockUnsupported)
	return x.buf
}

func (e TL_pageBlockTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockTitle)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockSubtitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubtitle)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockAuthorDate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAuthorDate)
	x.Bytes(e.author.Encode())
	x.Int(e.published_date)
	return x.buf
}

func (e TL_pageBlockHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockHeader)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockSubheader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubheader)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockParagraph) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockParagraph)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockPreformatted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPreformatted)
	x.Bytes(e.text.Encode())
	x.String(e.language)
	return x.buf
}

func (e TL_pageBlockFooter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockFooter)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageBlockDivider) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockDivider)
	return x.buf
}

func (e TL_pageBlockAnchor) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAnchor)
	x.String(e.name)
	return x.buf
}

func (e TL_pageBlockList) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockList)
	x.Vector(e.items)
	return x.buf
}

func (e TL_pageBlockBlockquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockBlockquote)
	x.Bytes(e.text.Encode())
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockPullquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPullquote)
	x.Bytes(e.text.Encode())
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPhoto)
	x.Bytes(e.flags.Encode())
	x.Long(e.photo_id)
	x.Bytes(e.caption.Encode())
	x.Bytes(e.url.Encode())
	x.Bytes(e.webpage_id.Encode())
	return x.buf
}

func (e TL_pageBlockVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockVideo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.autoplay.Encode())
	x.Bytes(e.loop.Encode())
	x.Long(e.video_id)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockCover) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCover)
	x.Bytes(e.cover.Encode())
	return x.buf
}

func (e TL_pageBlockEmbed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbed)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.full_width.Encode())
	x.Bytes(e.allow_scrolling.Encode())
	x.Bytes(e.url.Encode())
	x.Bytes(e.html.Encode())
	x.Bytes(e.poster_photo_id.Encode())
	x.Bytes(e.w.Encode())
	x.Bytes(e.h.Encode())
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockEmbedPost) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbedPost)
	x.String(e.url)
	x.Long(e.webpage_id)
	x.Long(e.author_photo_id)
	x.String(e.author)
	x.Int(e.date)
	x.Vector(e.blocks)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockCollage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCollage)
	x.Vector(e.items)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_pageBlockSlideshow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSlideshow)
	x.Vector(e.items)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_webPageNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageNotModified)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneCall)
	return x.buf
}

func (e TL_privacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneCall)
	return x.buf
}

func (e TL_sendMessageGamePlayAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGamePlayAction)
	return x.buf
}

func (e TL_phoneCallDiscardReasonMissed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonMissed)
	return x.buf
}

func (e TL_phoneCallDiscardReasonDisconnect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonDisconnect)
	return x.buf
}

func (e TL_phoneCallDiscardReasonHangup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonHangup)
	return x.buf
}

func (e TL_phoneCallDiscardReasonBusy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonBusy)
	return x.buf
}

func (e TL_updateDialogPinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDialogPinned)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.pinned.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_updatePinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePinnedDialogs)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Bytes(e.order.Encode())
	return x.buf
}

func (e TL_dataJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dataJSON)
	x.String(e.data)
	return x.buf
}

func (e TL_updateBotWebhookJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSON)
	x.Bytes(e.data.Encode())
	return x.buf
}

func (e TL_updateBotWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSONQuery)
	x.Long(e.query_id)
	x.Bytes(e.data.Encode())
	x.Int(e.timeout)
	return x.buf
}

func (e TL_labeledPrice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_labeledPrice)
	x.String(e.label)
	x.Long(e.amount)
	return x.buf
}

func (e TL_invoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invoice)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.test.Encode())
	x.Bytes(e.name_requested.Encode())
	x.Bytes(e.phone_requested.Encode())
	x.Bytes(e.email_requested.Encode())
	x.Bytes(e.shipping_address_requested.Encode())
	x.Bytes(e.flexible.Encode())
	x.Bytes(e.phone_to_provider.Encode())
	x.Bytes(e.email_to_provider.Encode())
	x.String(e.currency)
	x.Vector(e.prices)
	return x.buf
}

func (e TL_inputMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaInvoice)
	x.Bytes(e.flags.Encode())
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.invoice.Encode())
	x.StringBytes(e.payload)
	x.String(e.provider)
	x.Bytes(e.provider_data.Encode())
	x.String(e.start_param)
	return x.buf
}

func (e TL_paymentCharge) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentCharge)
	x.String(e.id)
	x.String(e.provider_charge_id)
	return x.buf
}

func (e TL_messageActionPaymentSentMe) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSentMe)
	x.Bytes(e.flags.Encode())
	x.String(e.currency)
	x.Long(e.total_amount)
	x.StringBytes(e.payload)
	x.Bytes(e.info.Encode())
	x.Bytes(e.shipping_option_id.Encode())
	x.Bytes(e.charge.Encode())
	return x.buf
}

func (e TL_messageMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaInvoice)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.shipping_address_requested.Encode())
	x.Bytes(e.test.Encode())
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.photo.Encode())
	x.Bytes(e.receipt_msg_id.Encode())
	x.String(e.currency)
	x.Long(e.total_amount)
	x.String(e.start_param)
	return x.buf
}

func (e TL_postAddress) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_postAddress)
	x.String(e.street_line1)
	x.String(e.street_line2)
	x.String(e.city)
	x.String(e.state)
	x.String(e.country_iso2)
	x.String(e.post_code)
	return x.buf
}

func (e TL_paymentRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentRequestedInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.name.Encode())
	x.Bytes(e.phone.Encode())
	x.Bytes(e.email.Encode())
	x.Bytes(e.shipping_address.Encode())
	return x.buf
}

func (e TL_keyboardButtonBuy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonBuy)
	x.String(e.text)
	return x.buf
}

func (e TL_messageActionPaymentSent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSent)
	x.String(e.currency)
	x.Long(e.total_amount)
	return x.buf
}

func (e TL_paymentSavedCredentialsCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentSavedCredentialsCard)
	x.String(e.id)
	x.String(e.title)
	return x.buf
}

func (e TL_webDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webDocument)
	x.String(e.url)
	x.Long(e.access_hash)
	x.Int(e.size)
	x.String(e.mime_type)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_inputWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebDocument)
	x.String(e.url)
	x.Int(e.size)
	x.String(e.mime_type)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_inputWebFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebFileLocation)
	x.String(e.url)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_upload_webFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_webFile)
	x.Int(e.size)
	x.String(e.mime_type)
	x.Bytes(e.file_type.Encode())
	x.Int(e.mtime)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_payments_paymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentForm)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.can_save_credentials.Encode())
	x.Bytes(e.password_missing.Encode())
	x.Int(e.bot_id)
	x.Bytes(e.invoice.Encode())
	x.Int(e.provider_id)
	x.String(e.url)
	x.Bytes(e.native_provider.Encode())
	x.Bytes(e.native_params.Encode())
	x.Bytes(e.saved_info.Encode())
	x.Bytes(e.saved_credentials.Encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_payments_validatedRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validatedRequestedInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.shipping_options.Encode())
	return x.buf
}

func (e TL_payments_paymentResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentResult)
	x.Bytes(e.updates.Encode())
	return x.buf
}

func (e TL_payments_paymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentReceipt)
	x.Bytes(e.flags.Encode())
	x.Int(e.date)
	x.Int(e.bot_id)
	x.Bytes(e.invoice.Encode())
	x.Int(e.provider_id)
	x.Bytes(e.info.Encode())
	x.Bytes(e.shipping.Encode())
	x.String(e.currency)
	x.Long(e.total_amount)
	x.String(e.credentials_title)
	x.Vector(e.users)
	return x.buf
}

func (e TL_payments_savedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_savedInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.has_saved_credentials.Encode())
	x.Bytes(e.saved_info.Encode())
	return x.buf
}

func (e TL_inputPaymentCredentialsSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsSaved)
	x.String(e.id)
	x.StringBytes(e.tmp_password)
	return x.buf
}

func (e TL_inputPaymentCredentials) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentials)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.save.Encode())
	x.Bytes(e.data.Encode())
	return x.buf
}

func (e TL_account_tmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_tmpPassword)
	x.StringBytes(e.tmp_password)
	x.Int(e.valid_until)
	return x.buf
}

func (e TL_shippingOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_shippingOption)
	x.String(e.id)
	x.String(e.title)
	x.Vector(e.prices)
	return x.buf
}

func (e TL_updateBotShippingQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotShippingQuery)
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.StringBytes(e.payload)
	x.Bytes(e.shipping_address.Encode())
	return x.buf
}

func (e TL_updateBotPrecheckoutQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotPrecheckoutQuery)
	x.Bytes(e.flags.Encode())
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.StringBytes(e.payload)
	x.Bytes(e.info.Encode())
	x.Bytes(e.shipping_option_id.Encode())
	x.String(e.currency)
	x.Long(e.total_amount)
	return x.buf
}

func (e TL_inputStickerSetItem) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetItem)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.document.Encode())
	x.String(e.emoji)
	x.Bytes(e.mask_coords.Encode())
	return x.buf
}

func (e TL_updatePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePhoneCall)
	x.Bytes(e.phone_call.Encode())
	return x.buf
}

func (e TL_inputPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneCall)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_phoneCallEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_phoneCallWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallWaiting)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.Bytes(e.protocol.Encode())
	x.Bytes(e.receive_date.Encode())
	return x.buf
}

func (e TL_phoneCallRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallRequested)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a_hash)
	x.Bytes(e.protocol.Encode())
	return x.buf
}

func (e TL_phoneCallAccepted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallAccepted)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_b)
	x.Bytes(e.protocol.Encode())
	return x.buf
}

func (e TL_phoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCall)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.p2p_allowed.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a_or_b)
	x.Long(e.key_fingerprint)
	x.Bytes(e.protocol.Encode())
	x.Vector(e.connections)
	x.Int(e.start_date)
	return x.buf
}

func (e TL_phoneCallDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscarded)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.need_rating.Encode())
	x.Bytes(e.need_debug.Encode())
	x.Bytes(e.video.Encode())
	x.Long(e.id)
	x.Bytes(e.reason.Encode())
	x.Bytes(e.duration.Encode())
	return x.buf
}

func (e TL_phoneConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneConnection)
	x.Long(e.id)
	x.String(e.ip)
	x.String(e.ipv6)
	x.Int(e.port)
	x.StringBytes(e.peer_tag)
	return x.buf
}

func (e TL_phoneCallProtocol) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallProtocol)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.udp_p2p.Encode())
	x.Bytes(e.udp_reflector.Encode())
	x.Int(e.min_layer)
	x.Int(e.max_layer)
	return x.buf
}

func (e TL_phone_phoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_phoneCall)
	x.Bytes(e.phone_call.Encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_inputMessagesFilterPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhoneCalls)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.missed.Encode())
	return x.buf
}

func (e TL_messageActionPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPhoneCall)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Long(e.call_id)
	x.Bytes(e.reason.Encode())
	x.Bytes(e.duration.Encode())
	return x.buf
}

func (e TL_inputMessagesFilterRoundVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVoice)
	return x.buf
}

func (e TL_inputMessagesFilterRoundVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVideo)
	return x.buf
}

func (e TL_sendMessageRecordRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordRoundAction)
	return x.buf
}

func (e TL_sendMessageUploadRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadRoundAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_upload_fileCdnRedirect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_fileCdnRedirect)
	x.Int(e.dc_id)
	x.StringBytes(e.file_token)
	x.StringBytes(e.encryption_key)
	x.StringBytes(e.encryption_iv)
	x.Vector(e.file_hashes)
	return x.buf
}

func (e TL_upload_cdnFileReuploadNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_cdnFileReuploadNeeded)
	x.StringBytes(e.request_token)
	return x.buf
}

func (e TL_upload_cdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_cdnFile)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_cdnPublicKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnPublicKey)
	x.Int(e.dc_id)
	x.String(e.public_key)
	return x.buf
}

func (e TL_cdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnConfig)
	x.Vector(e.public_keys)
	return x.buf
}

func (e TL_pageBlockChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockChannel)
	x.Bytes(e.channel.Encode())
	return x.buf
}

func (e TL_langPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackString)
	x.String(e.key)
	x.String(e.value)
	return x.buf
}

func (e TL_langPackStringPluralized) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringPluralized)
	x.Bytes(e.flags.Encode())
	x.String(e.key)
	x.Bytes(e.zero_value.Encode())
	x.Bytes(e.one_value.Encode())
	x.Bytes(e.two_value.Encode())
	x.Bytes(e.few_value.Encode())
	x.Bytes(e.many_value.Encode())
	x.String(e.other_value)
	return x.buf
}

func (e TL_langPackStringDeleted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringDeleted)
	x.String(e.key)
	return x.buf
}

func (e TL_langPackDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackDifference)
	x.String(e.lang_code)
	x.Int(e.from_version)
	x.Int(e.version)
	x.Vector(e.strings)
	return x.buf
}

func (e TL_langPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackLanguage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.official.Encode())
	x.Bytes(e.rtl.Encode())
	x.Bytes(e.beta.Encode())
	x.String(e.name)
	x.String(e.native_name)
	x.String(e.lang_code)
	x.Bytes(e.base_lang_code.Encode())
	x.String(e.plural_code)
	x.Int(e.strings_count)
	x.Int(e.translated_count)
	x.String(e.translations_url)
	return x.buf
}

func (e TL_updateLangPackTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPackTooLong)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_updateLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPack)
	x.Bytes(e.difference.Encode())
	return x.buf
}

func (e TL_channelParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantAdmin)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.can_edit.Encode())
	x.Bytes(e.self.Encode())
	x.Int(e.user_id)
	x.Bytes(e.inviter_id.Encode())
	x.Int(e.promoted_by)
	x.Int(e.date)
	x.Bytes(e.admin_rights.Encode())
	x.Bytes(e.rank.Encode())
	return x.buf
}

func (e TL_channelParticipantBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantBanned)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.left.Encode())
	x.Int(e.user_id)
	x.Int(e.kicked_by)
	x.Int(e.date)
	x.Bytes(e.banned_rights.Encode())
	return x.buf
}

func (e TL_channelParticipantsBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBanned)
	x.String(e.q)
	return x.buf
}

func (e TL_channelParticipantsSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsSearch)
	x.String(e.q)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeTitle)
	x.String(e.prev_value)
	x.String(e.new_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeAbout)
	x.String(e.prev_value)
	x.String(e.new_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeUsername)
	x.String(e.prev_value)
	x.String(e.new_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangePhoto)
	x.Bytes(e.prev_photo.Encode())
	x.Bytes(e.new_photo.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleInvites)
	x.Bytes(e.new_value.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleSignatures)
	x.Bytes(e.new_value.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionUpdatePinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionUpdatePinned)
	x.Bytes(e.message.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionEditMessage)
	x.Bytes(e.prev_message.Encode())
	x.Bytes(e.new_message.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionDeleteMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionDeleteMessage)
	x.Bytes(e.message.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantJoin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantJoin)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantLeave) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantLeave)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantInvite)
	x.Bytes(e.participant.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.prev_participant.Encode())
	x.Bytes(e.new_participant.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.prev_participant.Encode())
	x.Bytes(e.new_participant.Encode())
	return x.buf
}

func (e TL_channelAdminLogEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEvent)
	x.Long(e.id)
	x.Int(e.date)
	x.Int(e.user_id)
	x.Bytes(e.action.Encode())
	return x.buf
}

func (e TL_channels_adminLogResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_adminLogResults)
	x.Vector(e.events)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channelAdminLogEventsFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventsFilter)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.join.Encode())
	x.Bytes(e.leave.Encode())
	x.Bytes(e.invite.Encode())
	x.Bytes(e.ban.Encode())
	x.Bytes(e.unban.Encode())
	x.Bytes(e.kick.Encode())
	x.Bytes(e.unkick.Encode())
	x.Bytes(e.promote.Encode())
	x.Bytes(e.demote.Encode())
	x.Bytes(e.info.Encode())
	x.Bytes(e.settings.Encode())
	x.Bytes(e.pinned.Encode())
	x.Bytes(e.edit.Encode())
	x.Bytes(e.delete.Encode())
	return x.buf
}

func (e TL_topPeerCategoryPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPhoneCalls)
	return x.buf
}

func (e TL_pageBlockAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAudio)
	x.Long(e.audio_id)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_popularContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_popularContact)
	x.Long(e.client_id)
	x.Int(e.importers)
	return x.buf
}

func (e TL_messageActionScreenshotTaken) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionScreenshotTaken)
	return x.buf
}

func (e TL_messages_favedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_favedStickersNotModified)
	return x.buf
}

func (e TL_messages_favedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_favedStickers)
	x.Int(e.hash)
	x.Vector(e.packs)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_updateFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateFavedStickers)
	return x.buf
}

func (e TL_updateChannelReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelReadMessagesContents)
	x.Int(e.channel_id)
	x.VectorInt(e.messages)
	return x.buf
}

func (e TL_inputMessagesFilterMyMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMyMentions)
	return x.buf
}

func (e TL_updateContactsReset) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactsReset)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.prev_stickerset.Encode())
	x.Bytes(e.new_stickerset.Encode())
	return x.buf
}

func (e TL_messageActionCustomAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionCustomAction)
	x.String(e.message)
	return x.buf
}

func (e TL_inputPaymentCredentialsApplePay) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsApplePay)
	x.Bytes(e.payment_data.Encode())
	return x.buf
}

func (e TL_inputPaymentCredentialsAndroidPay) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsAndroidPay)
	x.Bytes(e.payment_token.Encode())
	x.String(e.google_transaction_id)
	return x.buf
}

func (e TL_inputMessagesFilterGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGeo)
	return x.buf
}

func (e TL_inputMessagesFilterContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterContacts)
	return x.buf
}

func (e TL_updateChannelAvailableMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelAvailableMessages)
	x.Int(e.channel_id)
	x.Int(e.available_min_id)
	return x.buf
}

func (e TL_channelAdminLogEventActionTogglePreHistoryHidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionTogglePreHistoryHidden)
	x.Bytes(e.new_value.Encode())
	return x.buf
}

func (e TL_inputMediaGeoLive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoLive)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.stopped.Encode())
	x.Bytes(e.geo_point.Encode())
	x.Bytes(e.period.Encode())
	return x.buf
}

func (e TL_messageMediaGeoLive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeoLive)
	x.Bytes(e.geo.Encode())
	x.Int(e.period)
	return x.buf
}

func (e TL_recentMeUrlUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_recentMeUrlUnknown)
	x.String(e.url)
	return x.buf
}

func (e TL_recentMeUrlUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_recentMeUrlUser)
	x.String(e.url)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_recentMeUrlChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_recentMeUrlChat)
	x.String(e.url)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_recentMeUrlChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_recentMeUrlChatInvite)
	x.String(e.url)
	x.Bytes(e.chat_invite.Encode())
	return x.buf
}

func (e TL_recentMeUrlStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_recentMeUrlStickerSet)
	x.String(e.url)
	x.Bytes(e.set.Encode())
	return x.buf
}

func (e TL_help_recentMeUrls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_recentMeUrls)
	x.Vector(e.urls)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channels_channelParticipantsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipantsNotModified)
	return x.buf
}

func (e TL_messages_messagesNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesNotModified)
	x.Int(e.count)
	return x.buf
}

func (e TL_inputSingleMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputSingleMedia)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.media.Encode())
	x.Long(e.random_id)
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_webAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webAuthorization)
	x.Long(e.hash)
	x.Int(e.bot_id)
	x.String(e.domain)
	x.String(e.browser)
	x.String(e.platform)
	x.Int(e.date_created)
	x.Int(e.date_active)
	x.String(e.ip)
	x.String(e.region)
	return x.buf
}

func (e TL_account_webAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_webAuthorizations)
	x.Vector(e.authorizations)
	x.Vector(e.users)
	return x.buf
}

func (e TL_inputMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageID)
	x.Int(e.id)
	return x.buf
}

func (e TL_inputMessageReplyTo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageReplyTo)
	x.Int(e.id)
	return x.buf
}

func (e TL_inputMessagePinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagePinned)
	return x.buf
}

func (e TL_messageEntityPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPhone)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityCashtag) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCashtag)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageActionBotAllowed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionBotAllowed)
	x.String(e.domain)
	return x.buf
}

func (e TL_inputDialogPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDialogPeer)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_dialogPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialogPeer)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_foundStickerSetsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundStickerSetsNotModified)
	return x.buf
}

func (e TL_messages_foundStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundStickerSets)
	x.Int(e.hash)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_fileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileHash)
	x.Int(e.offset)
	x.Int(e.limit)
	x.StringBytes(e.hash)
	return x.buf
}

func (e TL_webDocumentNoProxy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webDocumentNoProxy)
	x.String(e.url)
	x.Int(e.size)
	x.String(e.mime_type)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_inputClientProxy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputClientProxy)
	x.String(e.address)
	x.Int(e.port)
	return x.buf
}

func (e TL_help_proxyDataEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_proxyDataEmpty)
	x.Int(e.expires)
	return x.buf
}

func (e TL_help_proxyDataPromo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_proxyDataPromo)
	x.Int(e.expires)
	x.Bytes(e.peer.Encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_help_termsOfServiceUpdateEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfServiceUpdateEmpty)
	x.Int(e.expires)
	return x.buf
}

func (e TL_help_termsOfServiceUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfServiceUpdate)
	x.Int(e.expires)
	x.Bytes(e.terms_of_service.Encode())
	return x.buf
}

func (e TL_inputSecureFileUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputSecureFileUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.md5_checksum)
	x.StringBytes(e.file_hash)
	x.StringBytes(e.secret)
	return x.buf
}

func (e TL_inputSecureFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputSecureFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputSecureFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputSecureFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_secureFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureFileEmpty)
	return x.buf
}

func (e TL_secureFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.size)
	x.Int(e.dc_id)
	x.Int(e.date)
	x.StringBytes(e.file_hash)
	x.StringBytes(e.secret)
	return x.buf
}

func (e TL_secureData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureData)
	x.StringBytes(e.data)
	x.StringBytes(e.data_hash)
	x.StringBytes(e.secret)
	return x.buf
}

func (e TL_securePlainPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_securePlainPhone)
	x.String(e.phone)
	return x.buf
}

func (e TL_securePlainEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_securePlainEmail)
	x.String(e.email)
	return x.buf
}

func (e TL_secureValueTypePersonalDetails) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypePersonalDetails)
	return x.buf
}

func (e TL_secureValueTypePassport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypePassport)
	return x.buf
}

func (e TL_secureValueTypeDriverLicense) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeDriverLicense)
	return x.buf
}

func (e TL_secureValueTypeIdentityCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeIdentityCard)
	return x.buf
}

func (e TL_secureValueTypeInternalPassport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeInternalPassport)
	return x.buf
}

func (e TL_secureValueTypeAddress) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeAddress)
	return x.buf
}

func (e TL_secureValueTypeUtilityBill) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeUtilityBill)
	return x.buf
}

func (e TL_secureValueTypeBankStatement) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeBankStatement)
	return x.buf
}

func (e TL_secureValueTypeRentalAgreement) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeRentalAgreement)
	return x.buf
}

func (e TL_secureValueTypePassportRegistration) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypePassportRegistration)
	return x.buf
}

func (e TL_secureValueTypeTemporaryRegistration) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeTemporaryRegistration)
	return x.buf
}

func (e TL_secureValueTypePhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypePhone)
	return x.buf
}

func (e TL_secureValueTypeEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueTypeEmail)
	return x.buf
}

func (e TL_secureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValue)
	x.Bytes(e.flags.Encode())
	x.Bytes(e._type.Encode())
	x.Bytes(e.data.Encode())
	x.Bytes(e.front_side.Encode())
	x.Bytes(e.reverse_side.Encode())
	x.Bytes(e.selfie.Encode())
	x.Bytes(e.translation.Encode())
	x.Bytes(e.files.Encode())
	x.Bytes(e.plain_data.Encode())
	x.StringBytes(e.hash)
	return x.buf
}

func (e TL_inputSecureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputSecureValue)
	x.Bytes(e.flags.Encode())
	x.Bytes(e._type.Encode())
	x.Bytes(e.data.Encode())
	x.Bytes(e.front_side.Encode())
	x.Bytes(e.reverse_side.Encode())
	x.Bytes(e.selfie.Encode())
	x.Bytes(e.translation.Encode())
	x.Bytes(e.files.Encode())
	x.Bytes(e.plain_data.Encode())
	return x.buf
}

func (e TL_secureValueHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueHash)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.hash)
	return x.buf
}

func (e TL_secureValueErrorData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorData)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.data_hash)
	x.String(e.field)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorFrontSide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorFrontSide)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorReverseSide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorReverseSide)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorSelfie) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorSelfie)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorFile)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorFiles) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorFiles)
	x.Bytes(e._type.Encode())
	x.Vector(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureCredentialsEncrypted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureCredentialsEncrypted)
	x.StringBytes(e.data)
	x.StringBytes(e.hash)
	x.StringBytes(e.secret)
	return x.buf
}

func (e TL_account_authorizationForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_authorizationForm)
	x.Bytes(e.flags.Encode())
	x.Vector(e.required_types)
	x.Vector(e.values)
	x.Vector(e.errors)
	x.Vector(e.users)
	x.Bytes(e.privacy_policy_url.Encode())
	return x.buf
}

func (e TL_account_sentEmailCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sentEmailCode)
	x.String(e.email_pattern)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageActionSecureValuesSentMe) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionSecureValuesSentMe)
	x.Vector(e.values)
	x.Bytes(e.credentials.Encode())
	return x.buf
}

func (e TL_messageActionSecureValuesSent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionSecureValuesSent)
	x.Vector(e.types)
	return x.buf
}

func (e TL_help_deepLinkInfoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_deepLinkInfoEmpty)
	return x.buf
}

func (e TL_help_deepLinkInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_deepLinkInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.update_app.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_savedPhoneContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_savedPhoneContact)
	x.String(e.phone)
	x.String(e.first_name)
	x.String(e.last_name)
	x.Int(e.date)
	return x.buf
}

func (e TL_account_takeout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_takeout)
	x.Long(e.id)
	return x.buf
}

func (e TL_inputTakeoutFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputTakeoutFileLocation)
	return x.buf
}

func (e TL_updateDialogUnreadMark) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDialogUnreadMark)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.unread.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_dialogsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsNotModified)
	x.Int(e.count)
	return x.buf
}

func (e TL_inputWebFileGeoPointLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebFileGeoPointLocation)
	x.Bytes(e.geo_point.Encode())
	x.Long(e.access_hash)
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.zoom)
	x.Int(e.scale)
	return x.buf
}

func (e TL_contacts_topPeersDisabled) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeersDisabled)
	return x.buf
}

func (e TL_inputReportReasonCopyright) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonCopyright)
	return x.buf
}

func (e TL_passwordKdfAlgoUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_passwordKdfAlgoUnknown)
	return x.buf
}

func (e TL_securePasswordKdfAlgoUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_securePasswordKdfAlgoUnknown)
	return x.buf
}

func (e TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000)
	x.StringBytes(e.salt)
	return x.buf
}

func (e TL_securePasswordKdfAlgoSHA512) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_securePasswordKdfAlgoSHA512)
	x.StringBytes(e.salt)
	return x.buf
}

func (e TL_secureSecretSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureSecretSettings)
	x.Bytes(e.secure_algo.Encode())
	x.StringBytes(e.secure_secret)
	x.Long(e.secure_secret_id)
	return x.buf
}

func (e TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
	x.StringBytes(e.salt1)
	x.StringBytes(e.salt2)
	x.Int(e.g)
	x.StringBytes(e.p)
	return x.buf
}

func (e TL_inputCheckPasswordEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputCheckPasswordEmpty)
	return x.buf
}

func (e TL_inputCheckPasswordSRP) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputCheckPasswordSRP)
	x.Long(e.srp_id)
	x.StringBytes(e.A)
	x.StringBytes(e.M1)
	return x.buf
}

func (e TL_secureValueError) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueError)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorTranslationFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorTranslationFile)
	x.Bytes(e._type.Encode())
	x.StringBytes(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureValueErrorTranslationFiles) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureValueErrorTranslationFiles)
	x.Bytes(e._type.Encode())
	x.Vector(e.file_hash)
	x.String(e.text)
	return x.buf
}

func (e TL_secureRequiredType) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureRequiredType)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.native_names.Encode())
	x.Bytes(e.selfie_required.Encode())
	x.Bytes(e.translation_required.Encode())
	x.Bytes(e._type.Encode())
	return x.buf
}

func (e TL_secureRequiredTypeOneOf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_secureRequiredTypeOneOf)
	x.Vector(e.types)
	return x.buf
}

func (e TL_help_passportConfigNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_passportConfigNotModified)
	return x.buf
}

func (e TL_help_passportConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_passportConfig)
	x.Int(e.hash)
	x.Bytes(e.countries_langs.Encode())
	return x.buf
}

func (e TL_inputAppEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.time)
	x.String(e._type)
	x.Long(e.peer)
	x.Bytes(e.data.Encode())
	return x.buf
}

func (e TL_jsonObjectValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonObjectValue)
	x.String(e.key)
	x.Bytes(e.value.Encode())
	return x.buf
}

func (e TL_jsonNull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonNull)
	return x.buf
}

func (e TL_jsonBool) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonBool)
	x.Bytes(e.value.Encode())
	return x.buf
}

func (e TL_jsonNumber) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonNumber)
	x.Double(e.value)
	return x.buf
}

func (e TL_jsonString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonString)
	x.String(e.value)
	return x.buf
}

func (e TL_jsonArray) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonArray)
	x.Vector(e.value)
	return x.buf
}

func (e TL_jsonObject) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_jsonObject)
	x.Vector(e.value)
	return x.buf
}

func (e TL_updateUserPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPinnedMessage)
	x.Int(e.user_id)
	x.Int(e.id)
	return x.buf
}

func (e TL_updateChatPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatPinnedMessage)
	x.Int(e.chat_id)
	x.Int(e.id)
	x.Int(e.version)
	return x.buf
}

func (e TL_inputNotifyBroadcasts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyBroadcasts)
	return x.buf
}

func (e TL_notifyBroadcasts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyBroadcasts)
	return x.buf
}

func (e TL_textSubscript) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textSubscript)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textSuperscript) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textSuperscript)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textMarked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textMarked)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_textPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textPhone)
	x.Bytes(e.text.Encode())
	x.String(e.phone)
	return x.buf
}

func (e TL_textImage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textImage)
	x.Long(e.document_id)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_pageBlockKicker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockKicker)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageTableCell) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageTableCell)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.header.Encode())
	x.Bytes(e.align_center.Encode())
	x.Bytes(e.align_right.Encode())
	x.Bytes(e.valign_middle.Encode())
	x.Bytes(e.valign_bottom.Encode())
	x.Bytes(e.text.Encode())
	x.Bytes(e.colspan.Encode())
	x.Bytes(e.rowspan.Encode())
	return x.buf
}

func (e TL_pageTableRow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageTableRow)
	x.Vector(e.cells)
	return x.buf
}

func (e TL_pageBlockTable) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockTable)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.bordered.Encode())
	x.Bytes(e.striped.Encode())
	x.Bytes(e.title.Encode())
	x.Vector(e.rows)
	return x.buf
}

func (e TL_pageCaption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageCaption)
	x.Bytes(e.text.Encode())
	x.Bytes(e.credit.Encode())
	return x.buf
}

func (e TL_pageListItemText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageListItemText)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageListItemBlocks) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageListItemBlocks)
	x.Vector(e.blocks)
	return x.buf
}

func (e TL_pageListOrderedItemText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageListOrderedItemText)
	x.String(e.num)
	x.Bytes(e.text.Encode())
	return x.buf
}

func (e TL_pageListOrderedItemBlocks) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageListOrderedItemBlocks)
	x.String(e.num)
	x.Vector(e.blocks)
	return x.buf
}

func (e TL_pageBlockOrderedList) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockOrderedList)
	x.Vector(e.items)
	return x.buf
}

func (e TL_pageBlockDetails) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockDetails)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.open.Encode())
	x.Vector(e.blocks)
	x.Bytes(e.title.Encode())
	return x.buf
}

func (e TL_pageRelatedArticle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageRelatedArticle)
	x.Bytes(e.flags.Encode())
	x.String(e.url)
	x.Long(e.webpage_id)
	x.Bytes(e.title.Encode())
	x.Bytes(e.description.Encode())
	x.Bytes(e.photo_id.Encode())
	x.Bytes(e.author.Encode())
	x.Bytes(e.published_date.Encode())
	return x.buf
}

func (e TL_pageBlockRelatedArticles) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockRelatedArticles)
	x.Bytes(e.title.Encode())
	x.Vector(e.articles)
	return x.buf
}

func (e TL_pageBlockMap) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockMap)
	x.Bytes(e.geo.Encode())
	x.Int(e.zoom)
	x.Int(e.w)
	x.Int(e.h)
	x.Bytes(e.caption.Encode())
	return x.buf
}

func (e TL_page) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_page)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.part.Encode())
	x.Bytes(e.rtl.Encode())
	x.Bytes(e.v2.Encode())
	x.String(e.url)
	x.Vector(e.blocks)
	x.Vector(e.photos)
	x.Vector(e.documents)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneP2P) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneP2P)
	return x.buf
}

func (e TL_privacyKeyPhoneP2P) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneP2P)
	return x.buf
}

func (e TL_textAnchor) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textAnchor)
	x.Bytes(e.text.Encode())
	x.String(e.name)
	return x.buf
}

func (e TL_help_supportName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_supportName)
	x.String(e.name)
	return x.buf
}

func (e TL_help_userInfoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_userInfoEmpty)
	return x.buf
}

func (e TL_help_userInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_userInfo)
	x.String(e.message)
	x.Vector(e.entities)
	x.String(e.author)
	x.Int(e.date)
	return x.buf
}

func (e TL_messageActionContactSignUp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionContactSignUp)
	return x.buf
}

func (e TL_updateMessagePoll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessagePoll)
	x.Bytes(e.flags.Encode())
	x.Long(e.poll_id)
	x.Bytes(e.poll.Encode())
	x.Bytes(e.results.Encode())
	return x.buf
}

func (e TL_pollAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pollAnswer)
	x.String(e.text)
	x.StringBytes(e.option)
	return x.buf
}

func (e TL_poll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_poll)
	x.Long(e.id)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.closed.Encode())
	x.Bytes(e.public_voters.Encode())
	x.Bytes(e.multiple_choice.Encode())
	x.Bytes(e.quiz.Encode())
	x.String(e.question)
	x.Vector(e.answers)
	return x.buf
}

func (e TL_pollAnswerVoters) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pollAnswerVoters)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.chosen.Encode())
	x.Bytes(e.correct.Encode())
	x.StringBytes(e.option)
	x.Int(e.voters)
	return x.buf
}

func (e TL_pollResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pollResults)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.min.Encode())
	x.Bytes(e.results.Encode())
	x.Bytes(e.total_voters.Encode())
	return x.buf
}

func (e TL_inputMediaPoll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPoll)
	x.Bytes(e.poll.Encode())
	return x.buf
}

func (e TL_messageMediaPoll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPoll)
	x.Bytes(e.poll.Encode())
	x.Bytes(e.results.Encode())
	return x.buf
}

func (e TL_chatOnlines) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatOnlines)
	x.Int(e.onlines)
	return x.buf
}

func (e TL_statsURL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_statsURL)
	x.String(e.url)
	return x.buf
}

func (e TL_photoStrippedSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoStrippedSize)
	x.String(e._type)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_chatAdminRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatAdminRights)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.change_info.Encode())
	x.Bytes(e.post_messages.Encode())
	x.Bytes(e.edit_messages.Encode())
	x.Bytes(e.delete_messages.Encode())
	x.Bytes(e.ban_users.Encode())
	x.Bytes(e.invite_users.Encode())
	x.Bytes(e.pin_messages.Encode())
	x.Bytes(e.add_admins.Encode())
	return x.buf
}

func (e TL_chatBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatBannedRights)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.view_messages.Encode())
	x.Bytes(e.send_messages.Encode())
	x.Bytes(e.send_media.Encode())
	x.Bytes(e.send_stickers.Encode())
	x.Bytes(e.send_gifs.Encode())
	x.Bytes(e.send_games.Encode())
	x.Bytes(e.send_inline.Encode())
	x.Bytes(e.embed_links.Encode())
	x.Bytes(e.send_polls.Encode())
	x.Bytes(e.change_info.Encode())
	x.Bytes(e.invite_users.Encode())
	x.Bytes(e.pin_messages.Encode())
	x.Int(e.until_date)
	return x.buf
}

func (e TL_updateChatDefaultBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatDefaultBannedRights)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.default_banned_rights.Encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_inputWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWallPaper)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputWallPaperSlug) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWallPaperSlug)
	x.String(e.slug)
	return x.buf
}

func (e TL_channelParticipantsContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsContacts)
	x.String(e.q)
	return x.buf
}

func (e TL_channelAdminLogEventActionDefaultBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionDefaultBannedRights)
	x.Bytes(e.prev_banned_rights.Encode())
	x.Bytes(e.new_banned_rights.Encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionStopPoll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionStopPoll)
	x.Bytes(e.message.Encode())
	return x.buf
}

func (e TL_account_wallPapersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_wallPapersNotModified)
	return x.buf
}

func (e TL_account_wallPapers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_wallPapers)
	x.Int(e.hash)
	x.Vector(e.wallpapers)
	return x.buf
}

func (e TL_codeSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_codeSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.allow_flashcall.Encode())
	x.Bytes(e.current_number.Encode())
	x.Bytes(e.allow_app_hash.Encode())
	return x.buf
}

func (e TL_wallPaperSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.blur.Encode())
	x.Bytes(e.motion.Encode())
	x.Bytes(e.background_color.Encode())
	x.Bytes(e.intensity.Encode())
	return x.buf
}

func (e TL_autoDownloadSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_autoDownloadSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.disabled.Encode())
	x.Bytes(e.video_preload_large.Encode())
	x.Bytes(e.audio_preload_next.Encode())
	x.Bytes(e.phonecalls_less_data.Encode())
	x.Int(e.photo_size_max)
	x.Int(e.video_size_max)
	x.Int(e.file_size_max)
	return x.buf
}

func (e TL_account_autoDownloadSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_autoDownloadSettings)
	x.Bytes(e.low.Encode())
	x.Bytes(e.medium.Encode())
	x.Bytes(e.high.Encode())
	return x.buf
}

func (e TL_emojiKeyword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_emojiKeyword)
	x.String(e.keyword)
	x.VectorString(e.emoticons)
	return x.buf
}

func (e TL_emojiKeywordDeleted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_emojiKeywordDeleted)
	x.String(e.keyword)
	x.VectorString(e.emoticons)
	return x.buf
}

func (e TL_emojiKeywordsDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_emojiKeywordsDifference)
	x.String(e.lang_code)
	x.Int(e.from_version)
	x.Int(e.version)
	x.Vector(e.keywords)
	return x.buf
}

func (e TL_emojiURL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_emojiURL)
	x.String(e.url)
	return x.buf
}

func (e TL_emojiLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_emojiLanguage)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_inputPrivacyKeyForwards) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyForwards)
	return x.buf
}

func (e TL_privacyKeyForwards) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyForwards)
	return x.buf
}

func (e TL_inputPrivacyKeyProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyProfilePhoto)
	return x.buf
}

func (e TL_privacyKeyProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyProfilePhoto)
	return x.buf
}

func (e TL_fileLocationToBeDeprecated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationToBeDeprecated)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	return x.buf
}

func (e TL_inputPhotoFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	x.String(e.thumb_size)
	return x.buf
}

func (e TL_inputPhotoLegacyFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoLegacyFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.StringBytes(e.file_reference)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_inputPeerPhotoFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerPhotoFileLocation)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.big.Encode())
	x.Bytes(e.peer.Encode())
	x.Long(e.volume_id)
	x.Int(e.local_id)
	return x.buf
}

func (e TL_inputStickerSetThumb) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetThumb)
	x.Bytes(e.stickerset.Encode())
	x.Long(e.volume_id)
	x.Int(e.local_id)
	return x.buf
}

func (e TL_folder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_folder)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.autofill_new_broadcasts.Encode())
	x.Bytes(e.autofill_public_groups.Encode())
	x.Bytes(e.autofill_new_correspondents.Encode())
	x.Int(e.id)
	x.String(e.title)
	x.Bytes(e.photo.Encode())
	return x.buf
}

func (e TL_dialogFolder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialogFolder)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.pinned.Encode())
	x.Bytes(e.folder.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.top_message)
	x.Int(e.unread_muted_peers_count)
	x.Int(e.unread_unmuted_peers_count)
	x.Int(e.unread_muted_messages_count)
	x.Int(e.unread_unmuted_messages_count)
	return x.buf
}

func (e TL_inputDialogPeerFolder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDialogPeerFolder)
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_dialogPeerFolder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialogPeerFolder)
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_inputFolderPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFolderPeer)
	x.Bytes(e.peer.Encode())
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_folderPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_folderPeer)
	x.Bytes(e.peer.Encode())
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_updateFolderPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateFolderPeers)
	x.Vector(e.folder_peers)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_inputUserFromMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserFromMessage)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputChannelFromMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelFromMessage)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_inputPeerUserFromMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUserFromMessage)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputPeerChannelFromMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannelFromMessage)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneNumber) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneNumber)
	return x.buf
}

func (e TL_privacyKeyPhoneNumber) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneNumber)
	return x.buf
}

func (e TL_topPeerCategoryForwardUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryForwardUsers)
	return x.buf
}

func (e TL_topPeerCategoryForwardChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryForwardChats)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeLinkedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeLinkedChat)
	x.Int(e.prev_value)
	x.Int(e.new_value)
	return x.buf
}

func (e TL_messages_searchCounter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchCounter)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.inexact.Encode())
	x.Bytes(e.filter.Encode())
	x.Int(e.count)
	return x.buf
}

func (e TL_keyboardButtonUrlAuth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrlAuth)
	x.Bytes(e.flags.Encode())
	x.String(e.text)
	x.Bytes(e.fwd_text.Encode())
	x.String(e.url)
	x.Int(e.button_id)
	return x.buf
}

func (e TL_inputKeyboardButtonUrlAuth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputKeyboardButtonUrlAuth)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.request_write_access.Encode())
	x.String(e.text)
	x.Bytes(e.fwd_text.Encode())
	x.String(e.url)
	x.Bytes(e.bot.Encode())
	return x.buf
}

func (e TL_urlAuthResultRequest) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_urlAuthResultRequest)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.request_write_access.Encode())
	x.Bytes(e.bot.Encode())
	x.String(e.domain)
	return x.buf
}

func (e TL_urlAuthResultAccepted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_urlAuthResultAccepted)
	x.String(e.url)
	return x.buf
}

func (e TL_urlAuthResultDefault) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_urlAuthResultDefault)
	return x.buf
}

func (e TL_inputPrivacyValueAllowChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowChatParticipants)
	x.VectorInt(e.chats)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowChatParticipants)
	x.VectorInt(e.chats)
	return x.buf
}

func (e TL_privacyValueAllowChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowChatParticipants)
	x.VectorInt(e.chats)
	return x.buf
}

func (e TL_privacyValueDisallowChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowChatParticipants)
	x.VectorInt(e.chats)
	return x.buf
}

func (e TL_messageEntityUnderline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnderline)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityStrike) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityStrike)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityBlockquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBlockquote)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_updatePeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePeerSettings)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_channelLocationEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelLocationEmpty)
	return x.buf
}

func (e TL_channelLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelLocation)
	x.Bytes(e.geo_point.Encode())
	x.String(e.address)
	return x.buf
}

func (e TL_peerLocated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerLocated)
	x.Bytes(e.peer.Encode())
	x.Int(e.expires)
	x.Int(e.distance)
	return x.buf
}

func (e TL_updatePeerLocated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePeerLocated)
	x.Vector(e.peers)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeLocation)
	x.Bytes(e.prev_value.Encode())
	x.Bytes(e.new_value.Encode())
	return x.buf
}

func (e TL_inputReportReasonGeoIrrelevant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonGeoIrrelevant)
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleSlowMode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleSlowMode)
	x.Int(e.prev_value)
	x.Int(e.new_value)
	return x.buf
}

func (e TL_auth_authorizationSignUpRequired) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorizationSignUpRequired)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.terms_of_service.Encode())
	return x.buf
}

func (e TL_payments_paymentVerificationNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentVerificationNeeded)
	x.String(e.url)
	return x.buf
}

func (e TL_inputStickerSetAnimatedEmoji) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetAnimatedEmoji)
	return x.buf
}

func (e TL_updateNewScheduledMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewScheduledMessage)
	x.Bytes(e.message.Encode())
	return x.buf
}

func (e TL_updateDeleteScheduledMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteScheduledMessages)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.messages)
	return x.buf
}

func (e TL_restrictionReason) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_restrictionReason)
	x.String(e.platform)
	x.String(e.reason)
	x.String(e.text)
	return x.buf
}

func (e TL_inputTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputTheme)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputThemeSlug) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputThemeSlug)
	x.String(e.slug)
	return x.buf
}

func (e TL_themeDocumentNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_themeDocumentNotModified)
	return x.buf
}

func (e TL_theme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_theme)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.creator.Encode())
	x.Bytes(e.defaults.Encode())
	x.Long(e.id)
	x.Long(e.access_hash)
	x.String(e.slug)
	x.String(e.title)
	x.Bytes(e.document.Encode())
	x.Int(e.installs_count)
	return x.buf
}

func (e TL_account_themesNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_themesNotModified)
	return x.buf
}

func (e TL_account_themes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_themes)
	x.Int(e.hash)
	x.Vector(e.themes)
	return x.buf
}

func (e TL_updateTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateTheme)
	x.Bytes(e.theme.Encode())
	return x.buf
}

func (e TL_inputPrivacyKeyAddedByPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyAddedByPhone)
	return x.buf
}

func (e TL_privacyKeyAddedByPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyAddedByPhone)
	return x.buf
}

func (e TL_invokeAfterMsg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.msg_id)
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_invokeAfterMsgs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.msg_ids)
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_auth_sendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCode)
	x.String(e.phone_number)
	x.Int(e.api_id)
	x.String(e.api_hash)
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_auth_signUp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signUp)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_auth_signIn) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signIn)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_auth_logOut) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_logOut)
	return x.buf
}

func (e TL_auth_resetAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resetAuthorizations)
	return x.buf
}

func (e TL_auth_exportAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportAuthorization)
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_auth_importAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_auth_bindTempAuthKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_bindTempAuthKey)
	x.Long(e.perm_auth_key_id)
	x.Long(e.nonce)
	x.Int(e.expires_at)
	x.StringBytes(e.encrypted_message)
	return x.buf
}

func (e TL_account_registerDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_registerDevice)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_muted.Encode())
	x.Int(e.token_type)
	x.String(e.token)
	x.Bytes(e.app_sandbox.Encode())
	x.StringBytes(e.secret)
	x.VectorInt(e.other_uids)
	return x.buf
}

func (e TL_account_unregisterDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_unregisterDevice)
	x.Int(e.token_type)
	x.String(e.token)
	x.VectorInt(e.other_uids)
	return x.buf
}

func (e TL_account_updateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateNotifySettings)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_getNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifySettings)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_account_resetNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetNotifySettings)
	return x.buf
}

func (e TL_account_updateProfile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateProfile)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.first_name.Encode())
	x.Bytes(e.last_name.Encode())
	x.Bytes(e.about.Encode())
	return x.buf
}

func (e TL_account_updateStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateStatus)
	x.Bytes(e.offline.Encode())
	return x.buf
}

func (e TL_account_getWallPapers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPapers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_account_reportPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_reportPeer)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.reason.Encode())
	return x.buf
}

func (e TL_users_getUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getUsers)
	x.Vector(e.id)
	return x.buf
}

func (e TL_users_getFullUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getFullUser)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_contacts_getContactIDs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContactIDs)
	x.Int(e.hash)
	return x.buf
}

func (e TL_contacts_getStatuses) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getStatuses)
	return x.buf
}

func (e TL_contacts_getContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContacts)
	x.Int(e.hash)
	return x.buf
}

func (e TL_contacts_importContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importContacts)
	x.Vector(e.contacts)
	return x.buf
}

func (e TL_contacts_deleteContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContacts)
	x.Vector(e.id)
	return x.buf
}

func (e TL_contacts_deleteByPhones) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteByPhones)
	x.VectorString(e.phones)
	return x.buf
}

func (e TL_contacts_block) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_block)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_contacts_unblock) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_unblock)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_contacts_getBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getBlocked)
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessages)
	x.Vector(e.id)
	return x.buf
}

func (e TL_messages_getDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogs)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.exclude_pinned.Encode())
	x.Bytes(e.folder_id.Encode())
	x.Int(e.offset_date)
	x.Int(e.offset_id)
	x.Bytes(e.offset_peer.Encode())
	x.Int(e.limit)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_getHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getHistory)
	x.Bytes(e.peer.Encode())
	x.Int(e.offset_id)
	x.Int(e.offset_date)
	x.Int(e.add_offset)
	x.Int(e.limit)
	x.Int(e.max_id)
	x.Int(e.min_id)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_search) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_search)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.peer.Encode())
	x.String(e.q)
	x.Bytes(e.from_id.Encode())
	x.Bytes(e.filter.Encode())
	x.Int(e.min_date)
	x.Int(e.max_date)
	x.Int(e.offset_id)
	x.Int(e.add_offset)
	x.Int(e.limit)
	x.Int(e.max_id)
	x.Int(e.min_id)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_readHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readHistory)
	x.Bytes(e.peer.Encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_deleteHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteHistory)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.just_clear.Encode())
	x.Bytes(e.revoke.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_deleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteMessages)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.revoke.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_receivedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedMessages)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_setTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setTyping)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.action.Encode())
	return x.buf
}

func (e TL_messages_sendMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.background.Encode())
	x.Bytes(e.clear_draft.Encode())
	x.Bytes(e.peer.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.String(e.message)
	x.Long(e.random_id)
	x.Bytes(e.reply_markup.Encode())
	x.Bytes(e.entities.Encode())
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_messages_sendMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMedia)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.background.Encode())
	x.Bytes(e.clear_draft.Encode())
	x.Bytes(e.peer.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Bytes(e.media.Encode())
	x.String(e.message)
	x.Long(e.random_id)
	x.Bytes(e.reply_markup.Encode())
	x.Bytes(e.entities.Encode())
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_messages_forwardMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessages)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.background.Encode())
	x.Bytes(e.with_my_score.Encode())
	x.Bytes(e.grouped.Encode())
	x.Bytes(e.from_peer.Encode())
	x.VectorInt(e.id)
	x.VectorLong(e.random_id)
	x.Bytes(e.to_peer.Encode())
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_messages_reportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportSpam)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_getPeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerSettings)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_report) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_report)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.id)
	x.Bytes(e.reason.Encode())
	return x.buf
}

func (e TL_messages_getChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getChats)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getFullChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFullChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_editChatTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatTitle)
	x.Int(e.chat_id)
	x.String(e.title)
	return x.buf
}

func (e TL_messages_editChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatPhoto)
	x.Int(e.chat_id)
	x.Bytes(e.photo.Encode())
	return x.buf
}

func (e TL_messages_addChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_addChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.Encode())
	x.Int(e.fwd_limit)
	return x.buf
}

func (e TL_messages_deleteChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_messages_createChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_createChat)
	x.Vector(e.users)
	x.String(e.title)
	return x.buf
}

func (e TL_updates_getState) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getState)
	return x.buf
}

func (e TL_updates_getDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getDifference)
	x.Bytes(e.flags.Encode())
	x.Int(e.pts)
	x.Bytes(e.pts_total_limit.Encode())
	x.Int(e.date)
	x.Int(e.qts)
	return x.buf
}

func (e TL_photos_updateProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_updateProfilePhoto)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_photos_uploadProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_uploadProfilePhoto)
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_photos_deletePhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_deletePhotos)
	x.Vector(e.id)
	return x.buf
}

func (e TL_upload_saveFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_upload_getFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFile)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.precise.Encode())
	x.Bytes(e.cdn_supported.Encode())
	x.Bytes(e.location.Encode())
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_help_getConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getConfig)
	return x.buf
}

func (e TL_help_getNearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getNearestDc)
	return x.buf
}

func (e TL_help_getAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppUpdate)
	x.String(e.source)
	return x.buf
}

func (e TL_help_getInviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getInviteText)
	return x.buf
}

func (e TL_photos_getUserPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_getUserPhotos)
	x.Bytes(e.user_id.Encode())
	x.Int(e.offset)
	x.Long(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getDhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDhConfig)
	x.Int(e.version)
	x.Int(e.random_length)
	return x.buf
}

func (e TL_messages_requestEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestEncryption)
	x.Bytes(e.user_id.Encode())
	x.Int(e.random_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_messages_acceptEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptEncryption)
	x.Bytes(e.peer.Encode())
	x.StringBytes(e.g_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_messages_discardEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_discardEncryption)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_setEncryptedTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setEncryptedTyping)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.typing.Encode())
	return x.buf
}

func (e TL_messages_readEncryptedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readEncryptedHistory)
	x.Bytes(e.peer.Encode())
	x.Int(e.max_date)
	return x.buf
}

func (e TL_messages_sendEncrypted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncrypted)
	x.Bytes(e.peer.Encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_sendEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedFile)
	x.Bytes(e.peer.Encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_messages_sendEncryptedService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedService)
	x.Bytes(e.peer.Encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_receivedQueue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedQueue)
	x.Int(e.max_qts)
	return x.buf
}

func (e TL_messages_reportEncryptedSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportEncryptedSpam)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_upload_saveBigFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveBigFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.Int(e.file_total_parts)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_initConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Bytes(e.flags.Encode())
	x.Int(e.api_id)
	x.String(e.device_model)
	x.String(e.system_version)
	x.String(e.app_version)
	x.String(e.system_lang_code)
	x.String(e.lang_pack)
	x.String(e.lang_code)
	x.Bytes(e.proxy.Encode())
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_help_getSupport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupport)
	return x.buf
}

func (e TL_messages_readMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMessageContents)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_account_checkUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_checkUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_updateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_contacts_search) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_search)
	x.String(e.q)
	x.Int(e.limit)
	return x.buf
}

func (e TL_account_getPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPrivacy)
	x.Bytes(e.key.Encode())
	return x.buf
}

func (e TL_account_setPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setPrivacy)
	x.Bytes(e.key.Encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_account_deleteAccount) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteAccount)
	x.String(e.reason)
	return x.buf
}

func (e TL_account_getAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAccountTTL)
	return x.buf
}

func (e TL_account_setAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setAccountTTL)
	x.Bytes(e.ttl.Encode())
	return x.buf
}

func (e TL_invokeWithLayer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.layer)
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_contacts_resolveUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolveUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_sendChangePhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendChangePhoneCode)
	x.String(e.phone_number)
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_changePhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_changePhone)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_messages_getStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickers)
	x.String(e.emoticon)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_getAllStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_account_updateDeviceLocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateDeviceLocked)
	x.Int(e.period)
	return x.buf
}

func (e TL_auth_importBotAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importBotAuthorization)
	x.Int(e.flags)
	x.Int(e.api_id)
	x.String(e.api_hash)
	x.String(e.bot_auth_token)
	return x.buf
}

func (e TL_messages_getWebPagePreview) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPagePreview)
	x.Bytes(e.flags.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_account_getAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAuthorizations)
	return x.buf
}

func (e TL_account_resetAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetAuthorization)
	x.Long(e.hash)
	return x.buf
}

func (e TL_account_getPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPassword)
	return x.buf
}

func (e TL_account_getPasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPasswordSettings)
	x.Bytes(e.password.Encode())
	return x.buf
}

func (e TL_account_updatePasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updatePasswordSettings)
	x.Bytes(e.password.Encode())
	x.Bytes(e.new_settings.Encode())
	return x.buf
}

func (e TL_auth_checkPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPassword)
	x.Bytes(e.password.Encode())
	return x.buf
}

func (e TL_auth_requestPasswordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_requestPasswordRecovery)
	return x.buf
}

func (e TL_auth_recoverPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_recoverPassword)
	x.String(e.code)
	return x.buf
}

func (e TL_invokeWithoutUpdates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithoutUpdates)
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_messages_exportChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_exportChatInvite)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_checkChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_checkChatInvite)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_importChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_importChatInvite)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_getStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickerSet)
	x.Bytes(e.stickerset.Encode())
	return x.buf
}

func (e TL_messages_installStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_installStickerSet)
	x.Bytes(e.stickerset.Encode())
	x.Bytes(e.archived.Encode())
	return x.buf
}

func (e TL_messages_uninstallStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uninstallStickerSet)
	x.Bytes(e.stickerset.Encode())
	return x.buf
}

func (e TL_messages_startBot) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_startBot)
	x.Bytes(e.bot.Encode())
	x.Bytes(e.peer.Encode())
	x.Long(e.random_id)
	x.String(e.start_param)
	return x.buf
}

func (e TL_help_getAppChangelog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppChangelog)
	x.String(e.prev_app_version)
	return x.buf
}

func (e TL_messages_getMessagesViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessagesViews)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.id)
	x.Bytes(e.increment.Encode())
	return x.buf
}

func (e TL_channels_readHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readHistory)
	x.Bytes(e.channel.Encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_channels_deleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteMessages)
	x.Bytes(e.channel.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_channels_deleteUserHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteUserHistory)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_channels_reportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_reportSpam)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_channels_getMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getMessages)
	x.Bytes(e.channel.Encode())
	x.Vector(e.id)
	return x.buf
}

func (e TL_channels_getParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipants)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.filter.Encode())
	x.Int(e.offset)
	x.Int(e.limit)
	x.Int(e.hash)
	return x.buf
}

func (e TL_channels_getParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipant)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_channels_getChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getChannels)
	x.Vector(e.id)
	return x.buf
}

func (e TL_channels_getFullChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getFullChannel)
	x.Bytes(e.channel.Encode())
	return x.buf
}

func (e TL_channels_createChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_createChannel)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.broadcast.Encode())
	x.Bytes(e.megagroup.Encode())
	x.String(e.title)
	x.String(e.about)
	x.Bytes(e.geo_point.Encode())
	x.Bytes(e.address.Encode())
	return x.buf
}

func (e TL_channels_editAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAdmin)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	x.Bytes(e.admin_rights.Encode())
	x.String(e.rank)
	return x.buf
}

func (e TL_channels_editTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editTitle)
	x.Bytes(e.channel.Encode())
	x.String(e.title)
	return x.buf
}

func (e TL_channels_editPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editPhoto)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.photo.Encode())
	return x.buf
}

func (e TL_channels_checkUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_checkUsername)
	x.Bytes(e.channel.Encode())
	x.String(e.username)
	return x.buf
}

func (e TL_channels_updateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updateUsername)
	x.Bytes(e.channel.Encode())
	x.String(e.username)
	return x.buf
}

func (e TL_channels_joinChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_joinChannel)
	x.Bytes(e.channel.Encode())
	return x.buf
}

func (e TL_channels_leaveChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_leaveChannel)
	x.Bytes(e.channel.Encode())
	return x.buf
}

func (e TL_channels_inviteToChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_inviteToChannel)
	x.Bytes(e.channel.Encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_channels_deleteChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteChannel)
	x.Bytes(e.channel.Encode())
	return x.buf
}

func (e TL_updates_getChannelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getChannelDifference)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.force.Encode())
	x.Bytes(e.channel.Encode())
	x.Bytes(e.filter.Encode())
	x.Int(e.pts)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_editChatAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatAdmin)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.Encode())
	x.Bytes(e.is_admin.Encode())
	return x.buf
}

func (e TL_messages_migrateChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_migrateChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_searchGlobal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGlobal)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.folder_id.Encode())
	x.String(e.q)
	x.Int(e.offset_rate)
	x.Bytes(e.offset_peer.Encode())
	x.Int(e.offset_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_reorderStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderStickerSets)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.masks.Encode())
	x.VectorLong(e.order)
	return x.buf
}

func (e TL_messages_getDocumentByHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDocumentByHash)
	x.StringBytes(e.sha256)
	x.Int(e.size)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_messages_searchGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGifs)
	x.String(e.q)
	x.Int(e.offset)
	return x.buf
}

func (e TL_messages_getSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSavedGifs)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_saveGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveGif)
	x.Bytes(e.id.Encode())
	x.Bytes(e.unsave.Encode())
	return x.buf
}

func (e TL_messages_getInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineBotResults)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.bot.Encode())
	x.Bytes(e.peer.Encode())
	x.Bytes(e.geo_point.Encode())
	x.String(e.query)
	x.String(e.offset)
	return x.buf
}

func (e TL_messages_setInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineBotResults)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.gallery.Encode())
	x.Bytes(e.private.Encode())
	x.Long(e.query_id)
	x.Vector(e.results)
	x.Int(e.cache_time)
	x.Bytes(e.next_offset.Encode())
	x.Bytes(e.switch_pm.Encode())
	return x.buf
}

func (e TL_messages_sendInlineBotResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendInlineBotResult)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.background.Encode())
	x.Bytes(e.clear_draft.Encode())
	x.Bytes(e.hide_via.Encode())
	x.Bytes(e.peer.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Long(e.random_id)
	x.Long(e.query_id)
	x.String(e.id)
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_channels_exportMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportMessageLink)
	x.Bytes(e.channel.Encode())
	x.Int(e.id)
	x.Bytes(e.grouped.Encode())
	return x.buf
}

func (e TL_channels_toggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleSignatures)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.enabled.Encode())
	return x.buf
}

func (e TL_auth_resendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resendCode)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_auth_cancelCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_cancelCode)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_messages_getMessageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessageEditData)
	x.Bytes(e.peer.Encode())
	x.Int(e.id)
	return x.buf
}

func (e TL_messages_editMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.id)
	x.Bytes(e.message.Encode())
	x.Bytes(e.media.Encode())
	x.Bytes(e.reply_markup.Encode())
	x.Bytes(e.entities.Encode())
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_messages_editInlineBotMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editInlineBotMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.message.Encode())
	x.Bytes(e.media.Encode())
	x.Bytes(e.reply_markup.Encode())
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_messages_getBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getBotCallbackAnswer)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.game.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Bytes(e.data.Encode())
	return x.buf
}

func (e TL_messages_setBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotCallbackAnswer)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.alert.Encode())
	x.Long(e.query_id)
	x.Bytes(e.message.Encode())
	x.Bytes(e.url.Encode())
	x.Int(e.cache_time)
	return x.buf
}

func (e TL_contacts_getTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getTopPeers)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.correspondents.Encode())
	x.Bytes(e.bots_pm.Encode())
	x.Bytes(e.bots_inline.Encode())
	x.Bytes(e.phone_calls.Encode())
	x.Bytes(e.forward_users.Encode())
	x.Bytes(e.forward_chats.Encode())
	x.Bytes(e.groups.Encode())
	x.Bytes(e.channels.Encode())
	x.Int(e.offset)
	x.Int(e.limit)
	x.Int(e.hash)
	return x.buf
}

func (e TL_contacts_resetTopPeerRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetTopPeerRating)
	x.Bytes(e.category.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_getPeerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerDialogs)
	x.Vector(e.peers)
	return x.buf
}

func (e TL_messages_saveDraft) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveDraft)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.no_webpage.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Bytes(e.peer.Encode())
	x.String(e.message)
	x.Bytes(e.entities.Encode())
	return x.buf
}

func (e TL_messages_getAllDrafts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllDrafts)
	return x.buf
}

func (e TL_messages_getFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFeaturedStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_readFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readFeaturedStickers)
	x.VectorLong(e.id)
	return x.buf
}

func (e TL_messages_getRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getRecentStickers)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.attached.Encode())
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_saveRecentSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveRecentSticker)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.attached.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.unsave.Encode())
	return x.buf
}

func (e TL_messages_clearRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_clearRecentStickers)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.attached.Encode())
	return x.buf
}

func (e TL_messages_getArchivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getArchivedStickers)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.masks.Encode())
	x.Long(e.offset_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_account_sendConfirmPhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendConfirmPhoneCode)
	x.String(e.hash)
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_confirmPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_confirmPhone)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_channels_getAdminedPublicChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getAdminedPublicChannels)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.by_location.Encode())
	x.Bytes(e.check_limit.Encode())
	return x.buf
}

func (e TL_messages_getMaskStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMaskStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_getAttachedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAttachedStickers)
	x.Bytes(e.media.Encode())
	return x.buf
}

func (e TL_auth_dropTempAuthKeys) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_dropTempAuthKeys)
	x.VectorLong(e.except_auth_keys)
	return x.buf
}

func (e TL_messages_setGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setGameScore)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.edit_message.Encode())
	x.Bytes(e.force.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.id)
	x.Bytes(e.user_id.Encode())
	x.Int(e.score)
	return x.buf
}

func (e TL_messages_setInlineGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineGameScore)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.edit_message.Encode())
	x.Bytes(e.force.Encode())
	x.Bytes(e.id.Encode())
	x.Bytes(e.user_id.Encode())
	x.Int(e.score)
	return x.buf
}

func (e TL_messages_getGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getGameHighScores)
	x.Bytes(e.peer.Encode())
	x.Int(e.id)
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_messages_getInlineGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineGameHighScores)
	x.Bytes(e.id.Encode())
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_messages_getCommonChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getCommonChats)
	x.Bytes(e.user_id.Encode())
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getAllChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllChats)
	x.VectorInt(e.except_ids)
	return x.buf
}

func (e TL_help_setBotUpdatesStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_setBotUpdatesStatus)
	x.Int(e.pending_updates_count)
	x.String(e.message)
	return x.buf
}

func (e TL_messages_getWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPage)
	x.String(e.url)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_toggleDialogPin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleDialogPin)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.pinned.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_reorderPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderPinnedDialogs)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.force.Encode())
	x.Int(e.folder_id)
	x.Vector(e.order)
	return x.buf
}

func (e TL_messages_getPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPinnedDialogs)
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_bots_sendCustomRequest) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_sendCustomRequest)
	x.String(e.custom_method)
	x.Bytes(e.params.Encode())
	return x.buf
}

func (e TL_bots_answerWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_answerWebhookJSONQuery)
	x.Long(e.query_id)
	x.Bytes(e.data.Encode())
	return x.buf
}

func (e TL_upload_getWebFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getWebFile)
	x.Bytes(e.location.Encode())
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_payments_getPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentForm)
	x.Int(e.msg_id)
	return x.buf
}

func (e TL_payments_getPaymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentReceipt)
	x.Int(e.msg_id)
	return x.buf
}

func (e TL_payments_validateRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validateRequestedInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.save.Encode())
	x.Int(e.msg_id)
	x.Bytes(e.info.Encode())
	return x.buf
}

func (e TL_payments_sendPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_sendPaymentForm)
	x.Bytes(e.flags.Encode())
	x.Int(e.msg_id)
	x.Bytes(e.requested_info_id.Encode())
	x.Bytes(e.shipping_option_id.Encode())
	x.Bytes(e.credentials.Encode())
	return x.buf
}

func (e TL_account_getTmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getTmpPassword)
	x.Bytes(e.password.Encode())
	x.Int(e.period)
	return x.buf
}

func (e TL_payments_getSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getSavedInfo)
	return x.buf
}

func (e TL_payments_clearSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_clearSavedInfo)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.credentials.Encode())
	x.Bytes(e.info.Encode())
	return x.buf
}

func (e TL_messages_setBotShippingResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotShippingResults)
	x.Bytes(e.flags.Encode())
	x.Long(e.query_id)
	x.Bytes(e.error.Encode())
	x.Bytes(e.shipping_options.Encode())
	return x.buf
}

func (e TL_messages_setBotPrecheckoutResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotPrecheckoutResults)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.success.Encode())
	x.Long(e.query_id)
	x.Bytes(e.error.Encode())
	return x.buf
}

func (e TL_stickers_createStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_createStickerSet)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.masks.Encode())
	x.Bytes(e.user_id.Encode())
	x.String(e.title)
	x.String(e.short_name)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_stickers_removeStickerFromSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_removeStickerFromSet)
	x.Bytes(e.sticker.Encode())
	return x.buf
}

func (e TL_stickers_changeStickerPosition) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_changeStickerPosition)
	x.Bytes(e.sticker.Encode())
	x.Int(e.position)
	return x.buf
}

func (e TL_stickers_addStickerToSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_addStickerToSet)
	x.Bytes(e.stickerset.Encode())
	x.Bytes(e.sticker.Encode())
	return x.buf
}

func (e TL_messages_uploadMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uploadMedia)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.media.Encode())
	return x.buf
}

func (e TL_phone_getCallConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_getCallConfig)
	return x.buf
}

func (e TL_phone_requestCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_requestCall)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Bytes(e.user_id.Encode())
	x.Int(e.random_id)
	x.StringBytes(e.g_a_hash)
	x.Bytes(e.protocol.Encode())
	return x.buf
}

func (e TL_phone_acceptCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_acceptCall)
	x.Bytes(e.peer.Encode())
	x.StringBytes(e.g_b)
	x.Bytes(e.protocol.Encode())
	return x.buf
}

func (e TL_phone_confirmCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_confirmCall)
	x.Bytes(e.peer.Encode())
	x.StringBytes(e.g_a)
	x.Long(e.key_fingerprint)
	x.Bytes(e.protocol.Encode())
	return x.buf
}

func (e TL_phone_receivedCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_receivedCall)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_phone_discardCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_discardCall)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.video.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.duration)
	x.Bytes(e.reason.Encode())
	x.Long(e.connection_id)
	return x.buf
}

func (e TL_phone_setCallRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_setCallRating)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.user_initiative.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.rating)
	x.String(e.comment)
	return x.buf
}

func (e TL_phone_saveCallDebug) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_saveCallDebug)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.debug.Encode())
	return x.buf
}

func (e TL_upload_getCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getCdnFile)
	x.StringBytes(e.file_token)
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_upload_reuploadCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_reuploadCdnFile)
	x.StringBytes(e.file_token)
	x.StringBytes(e.request_token)
	return x.buf
}

func (e TL_help_getCdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getCdnConfig)
	return x.buf
}

func (e TL_langpack_getLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getLangPack)
	x.String(e.lang_pack)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_langpack_getStrings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getStrings)
	x.String(e.lang_pack)
	x.String(e.lang_code)
	x.VectorString(e.keys)
	return x.buf
}

func (e TL_langpack_getDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getDifference)
	x.String(e.lang_pack)
	x.String(e.lang_code)
	x.Int(e.from_version)
	return x.buf
}

func (e TL_langpack_getLanguages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getLanguages)
	x.String(e.lang_pack)
	return x.buf
}

func (e TL_channels_editBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editBanned)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	x.Bytes(e.banned_rights.Encode())
	return x.buf
}

func (e TL_channels_getAdminLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getAdminLog)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.channel.Encode())
	x.String(e.q)
	x.Bytes(e.events_filter.Encode())
	x.Bytes(e.admins.Encode())
	x.Long(e.max_id)
	x.Long(e.min_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_upload_getCdnFileHashes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getCdnFileHashes)
	x.StringBytes(e.file_token)
	x.Int(e.offset)
	return x.buf
}

func (e TL_messages_sendScreenshotNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendScreenshotNotification)
	x.Bytes(e.peer.Encode())
	x.Int(e.reply_to_msg_id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_channels_setStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_setStickers)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.stickerset.Encode())
	return x.buf
}

func (e TL_messages_getFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFavedStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_faveSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_faveSticker)
	x.Bytes(e.id.Encode())
	x.Bytes(e.unfave.Encode())
	return x.buf
}

func (e TL_channels_readMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readMessageContents)
	x.Bytes(e.channel.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_contacts_resetSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetSaved)
	return x.buf
}

func (e TL_messages_getUnreadMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getUnreadMentions)
	x.Bytes(e.peer.Encode())
	x.Int(e.offset_id)
	x.Int(e.add_offset)
	x.Int(e.limit)
	x.Int(e.max_id)
	x.Int(e.min_id)
	return x.buf
}

func (e TL_channels_deleteHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteHistory)
	x.Bytes(e.channel.Encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_help_getRecentMeUrls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getRecentMeUrls)
	x.String(e.referer)
	return x.buf
}

func (e TL_channels_togglePreHistoryHidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_togglePreHistoryHidden)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.enabled.Encode())
	return x.buf
}

func (e TL_messages_readMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMentions)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_getRecentLocations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getRecentLocations)
	x.Bytes(e.peer.Encode())
	x.Int(e.limit)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_sendMultiMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMultiMedia)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.background.Encode())
	x.Bytes(e.clear_draft.Encode())
	x.Bytes(e.peer.Encode())
	x.Bytes(e.reply_to_msg_id.Encode())
	x.Vector(e.multi_media)
	x.Bytes(e.schedule_date.Encode())
	return x.buf
}

func (e TL_messages_uploadEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uploadEncryptedFile)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.file.Encode())
	return x.buf
}

func (e TL_account_getWebAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWebAuthorizations)
	return x.buf
}

func (e TL_account_resetWebAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetWebAuthorization)
	x.Long(e.hash)
	return x.buf
}

func (e TL_account_resetWebAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetWebAuthorizations)
	return x.buf
}

func (e TL_messages_searchStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchStickerSets)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.exclude_featured.Encode())
	x.String(e.q)
	x.Int(e.hash)
	return x.buf
}

func (e TL_upload_getFileHashes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFileHashes)
	x.Bytes(e.location.Encode())
	x.Int(e.offset)
	return x.buf
}

func (e TL_help_getProxyData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getProxyData)
	return x.buf
}

func (e TL_help_getTermsOfServiceUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getTermsOfServiceUpdate)
	return x.buf
}

func (e TL_help_acceptTermsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_acceptTermsOfService)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_account_getAllSecureValues) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAllSecureValues)
	return x.buf
}

func (e TL_account_getSecureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getSecureValue)
	x.Vector(e.types)
	return x.buf
}

func (e TL_account_saveSecureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_saveSecureValue)
	x.Bytes(e.value.Encode())
	x.Long(e.secure_secret_id)
	return x.buf
}

func (e TL_account_deleteSecureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteSecureValue)
	x.Vector(e.types)
	return x.buf
}

func (e TL_users_setSecureValueErrors) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_setSecureValueErrors)
	x.Bytes(e.id.Encode())
	x.Vector(e.errors)
	return x.buf
}

func (e TL_account_getAuthorizationForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAuthorizationForm)
	x.Int(e.bot_id)
	x.String(e.scope)
	x.String(e.public_key)
	return x.buf
}

func (e TL_account_acceptAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_acceptAuthorization)
	x.Int(e.bot_id)
	x.String(e.scope)
	x.String(e.public_key)
	x.Vector(e.value_hashes)
	x.Bytes(e.credentials.Encode())
	return x.buf
}

func (e TL_account_sendVerifyPhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendVerifyPhoneCode)
	x.String(e.phone_number)
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_verifyPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_verifyPhone)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_account_sendVerifyEmailCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendVerifyEmailCode)
	x.String(e.email)
	return x.buf
}

func (e TL_account_verifyEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_verifyEmail)
	x.String(e.email)
	x.String(e.code)
	return x.buf
}

func (e TL_help_getDeepLinkInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getDeepLinkInfo)
	x.String(e.path)
	return x.buf
}

func (e TL_contacts_getSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getSaved)
	return x.buf
}

func (e TL_channels_getLeftChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getLeftChannels)
	x.Int(e.offset)
	return x.buf
}

func (e TL_account_initTakeoutSession) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_initTakeoutSession)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.contacts.Encode())
	x.Bytes(e.message_users.Encode())
	x.Bytes(e.message_chats.Encode())
	x.Bytes(e.message_megagroups.Encode())
	x.Bytes(e.message_channels.Encode())
	x.Bytes(e.files.Encode())
	x.Bytes(e.file_max_size.Encode())
	return x.buf
}

func (e TL_account_finishTakeoutSession) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_finishTakeoutSession)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.success.Encode())
	return x.buf
}

func (e TL_messages_getSplitRanges) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSplitRanges)
	return x.buf
}

func (e TL_invokeWithMessagesRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithMessagesRange)
	x.Bytes(e.ranges.Encode())
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_invokeWithTakeout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithTakeout)
	x.Long(e.takeout_id)
	x.Bytes(e.query.Encode())
	return x.buf
}

func (e TL_messages_markDialogUnread) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_markDialogUnread)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.unread.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_getDialogUnreadMarks) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogUnreadMarks)
	return x.buf
}

func (e TL_contacts_toggleTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_toggleTopPeers)
	x.Bytes(e.enabled.Encode())
	return x.buf
}

func (e TL_messages_clearAllDrafts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_clearAllDrafts)
	return x.buf
}

func (e TL_help_getAppConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppConfig)
	return x.buf
}

func (e TL_help_saveAppLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_saveAppLog)
	x.Vector(e.events)
	return x.buf
}

func (e TL_help_getPassportConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getPassportConfig)
	x.Int(e.hash)
	return x.buf
}

func (e TL_langpack_getLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getLanguage)
	x.String(e.lang_pack)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_messages_updatePinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_updatePinnedMessage)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.silent.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.id)
	return x.buf
}

func (e TL_account_confirmPasswordEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_confirmPasswordEmail)
	x.String(e.code)
	return x.buf
}

func (e TL_account_resendPasswordEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resendPasswordEmail)
	return x.buf
}

func (e TL_account_cancelPasswordEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_cancelPasswordEmail)
	return x.buf
}

func (e TL_help_getSupportName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupportName)
	return x.buf
}

func (e TL_help_getUserInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getUserInfo)
	x.Bytes(e.user_id.Encode())
	return x.buf
}

func (e TL_help_editUserInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_editUserInfo)
	x.Bytes(e.user_id.Encode())
	x.String(e.message)
	x.Vector(e.entities)
	return x.buf
}

func (e TL_account_getContactSignUpNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getContactSignUpNotification)
	return x.buf
}

func (e TL_account_setContactSignUpNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setContactSignUpNotification)
	x.Bytes(e.silent.Encode())
	return x.buf
}

func (e TL_account_getNotifyExceptions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifyExceptions)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.compare_sound.Encode())
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_sendVote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendVote)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Vector(e.options)
	return x.buf
}

func (e TL_messages_getPollResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPollResults)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	return x.buf
}

func (e TL_messages_getOnlines) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getOnlines)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_messages_getStatsURL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStatsURL)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.dark.Encode())
	x.Bytes(e.peer.Encode())
	x.String(e.params)
	return x.buf
}

func (e TL_messages_editChatAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatAbout)
	x.Bytes(e.peer.Encode())
	x.String(e.about)
	return x.buf
}

func (e TL_messages_editChatDefaultBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatDefaultBannedRights)
	x.Bytes(e.peer.Encode())
	x.Bytes(e.banned_rights.Encode())
	return x.buf
}

func (e TL_account_getWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPaper)
	x.Bytes(e.wallpaper.Encode())
	return x.buf
}

func (e TL_account_uploadWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_uploadWallPaper)
	x.Bytes(e.file.Encode())
	x.String(e.mime_type)
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_saveWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_saveWallPaper)
	x.Bytes(e.wallpaper.Encode())
	x.Bytes(e.unsave.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_installWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_installWallPaper)
	x.Bytes(e.wallpaper.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_account_resetWallPapers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetWallPapers)
	return x.buf
}

func (e TL_account_getAutoDownloadSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAutoDownloadSettings)
	return x.buf
}

func (e TL_account_saveAutoDownloadSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_saveAutoDownloadSettings)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.low.Encode())
	x.Bytes(e.high.Encode())
	x.Bytes(e.settings.Encode())
	return x.buf
}

func (e TL_messages_getEmojiKeywords) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getEmojiKeywords)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_messages_getEmojiKeywordsDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getEmojiKeywordsDifference)
	x.String(e.lang_code)
	x.Int(e.from_version)
	return x.buf
}

func (e TL_messages_getEmojiKeywordsLanguages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getEmojiKeywordsLanguages)
	x.VectorString(e.lang_codes)
	return x.buf
}

func (e TL_messages_getEmojiURL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getEmojiURL)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_folders_editPeerFolders) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_folders_editPeerFolders)
	x.Vector(e.folder_peers)
	return x.buf
}

func (e TL_folders_deleteFolder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_folders_deleteFolder)
	x.Int(e.folder_id)
	return x.buf
}

func (e TL_messages_getSearchCounters) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSearchCounters)
	x.Bytes(e.peer.Encode())
	x.Vector(e.filters)
	return x.buf
}

func (e TL_channels_getGroupsForDiscussion) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getGroupsForDiscussion)
	return x.buf
}

func (e TL_channels_setDiscussionGroup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_setDiscussionGroup)
	x.Bytes(e.broadcast.Encode())
	x.Bytes(e.group.Encode())
	return x.buf
}

func (e TL_messages_requestUrlAuth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestUrlAuth)
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.button_id)
	return x.buf
}

func (e TL_messages_acceptUrlAuth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptUrlAuth)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.write_allowed.Encode())
	x.Bytes(e.peer.Encode())
	x.Int(e.msg_id)
	x.Int(e.button_id)
	return x.buf
}

func (e TL_messages_hidePeerSettingsBar) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_hidePeerSettingsBar)
	x.Bytes(e.peer.Encode())
	return x.buf
}

func (e TL_contacts_addContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_addContact)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.add_phone_privacy_exception.Encode())
	x.Bytes(e.id.Encode())
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.phone)
	return x.buf
}

func (e TL_contacts_acceptContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_acceptContact)
	x.Bytes(e.id.Encode())
	return x.buf
}

func (e TL_channels_editCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editCreator)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.user_id.Encode())
	x.Bytes(e.password.Encode())
	return x.buf
}

func (e TL_contacts_getLocated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getLocated)
	x.Bytes(e.geo_point.Encode())
	return x.buf
}

func (e TL_channels_editLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editLocation)
	x.Bytes(e.channel.Encode())
	x.Bytes(e.geo_point.Encode())
	x.String(e.address)
	return x.buf
}

func (e TL_channels_toggleSlowMode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleSlowMode)
	x.Bytes(e.channel.Encode())
	x.Int(e.seconds)
	return x.buf
}

func (e TL_messages_getScheduledHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getScheduledHistory)
	x.Bytes(e.peer.Encode())
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_getScheduledMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getScheduledMessages)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_sendScheduledMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendScheduledMessages)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_deleteScheduledMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteScheduledMessages)
	x.Bytes(e.peer.Encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_account_uploadTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_uploadTheme)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.file.Encode())
	x.Bytes(e.thumb.Encode())
	x.String(e.file_name)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_account_createTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_createTheme)
	x.String(e.slug)
	x.String(e.title)
	x.Bytes(e.document.Encode())
	return x.buf
}

func (e TL_account_updateTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateTheme)
	x.Bytes(e.flags.Encode())
	x.String(e.format)
	x.Bytes(e.theme.Encode())
	x.Bytes(e.slug.Encode())
	x.Bytes(e.title.Encode())
	x.Bytes(e.document.Encode())
	return x.buf
}

func (e TL_account_saveTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_saveTheme)
	x.Bytes(e.theme.Encode())
	x.Bytes(e.unsave.Encode())
	return x.buf
}

func (e TL_account_installTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_installTheme)
	x.Bytes(e.flags.Encode())
	x.Bytes(e.dark.Encode())
	x.Bytes(e.format.Encode())
	x.Bytes(e.theme.Encode())
	return x.buf
}

func (e TL_account_getTheme) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getTheme)
	x.String(e.format)
	x.Bytes(e.theme.Encode())
	x.Long(e.document_id)
	return x.buf
}

func (e TL_account_getThemes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getThemes)
	x.String(e.format)
	x.Int(e.hash)
	return x.buf
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		print("boolFalse")
		r = TL_boolFalse{}

	case crc_boolTrue:
		print("boolTrue")
		r = TL_boolTrue{}

	case crc_true:
		print("true")
		r = TL_true{}

	case crc_error:
		print("error")
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case crc_null:
		print("null")
		r = TL_null{}

	case crc_inputPeerEmpty:
		print("inputPeerEmpty")
		r = TL_inputPeerEmpty{}

	case crc_inputPeerSelf:
		print("inputPeerSelf")
		r = TL_inputPeerSelf{}

	case crc_inputPeerChat:
		print("inputPeerChat")
		r = TL_inputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		print("inputUserEmpty")
		r = TL_inputUserEmpty{}

	case crc_inputUserSelf:
		print("inputUserSelf")
		r = TL_inputUserSelf{}

	case crc_inputPhoneContact:
		print("inputPhoneContact")
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		print("inputFile")
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		print("inputMediaEmpty")
		r = TL_inputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		print("inputMediaUploadedPhoto")
		r = TL_inputMediaUploadedPhoto{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaPhoto:
		print("inputMediaPhoto")
		r = TL_inputMediaPhoto{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaGeoPoint:
		print("inputMediaGeoPoint")
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case crc_inputMediaContact:
		print("inputMediaContact")
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputChatPhotoEmpty:
		print("inputChatPhotoEmpty")
		r = TL_inputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		print("inputChatUploadedPhoto")
		r = TL_inputChatUploadedPhoto{
			m.Object(),
		}

	case crc_inputChatPhoto:
		print("inputChatPhoto")
		r = TL_inputChatPhoto{
			m.Object(),
		}

	case crc_inputGeoPointEmpty:
		print("inputGeoPointEmpty")
		r = TL_inputGeoPointEmpty{}

	case crc_inputGeoPoint:
		print("inputGeoPoint")
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		print("inputPhotoEmpty")
		r = TL_inputPhotoEmpty{}

	case crc_inputPhoto:
		print("inputPhoto")
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_inputFileLocation:
		print("inputFileLocation")
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_peerUser:
		print("peerUser")
		r = TL_peerUser{
			m.Int(),
		}

	case crc_peerChat:
		print("peerChat")
		r = TL_peerChat{
			m.Int(),
		}

	case crc_storage_fileUnknown:
		print("storage_fileUnknown")
		r = TL_storage_fileUnknown{}

	case crc_storage_filePartial:
		print("storage_filePartial")
		r = TL_storage_filePartial{}

	case crc_storage_fileJpeg:
		print("storage_fileJpeg")
		r = TL_storage_fileJpeg{}

	case crc_storage_fileGif:
		print("storage_fileGif")
		r = TL_storage_fileGif{}

	case crc_storage_filePng:
		print("storage_filePng")
		r = TL_storage_filePng{}

	case crc_storage_filePdf:
		print("storage_filePdf")
		r = TL_storage_filePdf{}

	case crc_storage_fileMp3:
		print("storage_fileMp3")
		r = TL_storage_fileMp3{}

	case crc_storage_fileMov:
		print("storage_fileMov")
		r = TL_storage_fileMov{}

	case crc_storage_fileMp4:
		print("storage_fileMp4")
		r = TL_storage_fileMp4{}

	case crc_storage_fileWebp:
		print("storage_fileWebp")
		r = TL_storage_fileWebp{}

	case crc_userEmpty:
		print("userEmpty")
		r = TL_userEmpty{
			m.Int(),
		}

	case crc_userProfilePhotoEmpty:
		print("userProfilePhotoEmpty")
		r = TL_userProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		print("userProfilePhoto")
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_userStatusEmpty:
		print("userStatusEmpty")
		r = TL_userStatusEmpty{}

	case crc_userStatusOnline:
		print("userStatusOnline")
		r = TL_userStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		print("userStatusOffline")
		r = TL_userStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		print("chatEmpty")
		r = TL_chatEmpty{
			m.Int(),
		}

	case crc_chat:
		print("chat")
		r = TL_chat{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_chatForbidden:
		print("chatForbidden")
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
		}

	case crc_chatFull:
		print("chatFull")
		r = TL_chatFull{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_chatParticipant:
		print("chatParticipant")
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		print("chatParticipantsForbidden")
		r = TL_chatParticipantsForbidden{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_chatParticipants:
		print("chatParticipants")
		r = TL_chatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		print("chatPhotoEmpty")
		r = TL_chatPhotoEmpty{}

	case crc_chatPhoto:
		print("chatPhoto")
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messageEmpty:
		print("messageEmpty")
		r = TL_messageEmpty{
			m.Int(),
		}

	case crc_message:
		print("message")
		r = TL_message{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageService:
		print("messageService")
		r = TL_messageService{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messageMediaEmpty:
		print("messageMediaEmpty")
		r = TL_messageMediaEmpty{}

	case crc_messageMediaPhoto:
		print("messageMediaPhoto")
		r = TL_messageMediaPhoto{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaGeo:
		print("messageMediaGeo")
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case crc_messageMediaContact:
		print("messageMediaContact")
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		print("messageMediaUnsupported")
		r = TL_messageMediaUnsupported{}

	case crc_messageActionEmpty:
		print("messageActionEmpty")
		r = TL_messageActionEmpty{}

	case crc_messageActionChatCreate:
		print("messageActionChatCreate")
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		print("messageActionChatEditTitle")
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		print("messageActionChatEditPhoto")
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case crc_messageActionChatDeletePhoto:
		print("messageActionChatDeletePhoto")
		r = TL_messageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		print("messageActionChatAddUser")
		r = TL_messageActionChatAddUser{
			m.VectorInt(),
		}

	case crc_messageActionChatDeleteUser:
		print("messageActionChatDeleteUser")
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		print("dialog")
		r = TL_dialog{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_photoEmpty:
		print("photoEmpty")
		r = TL_photoEmpty{
			m.Long(),
		}

	case crc_photo:
		print("photo")
		r = TL_photo{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_photoSizeEmpty:
		print("photoSizeEmpty")
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		print("photoSize")
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		print("photoCachedSize")
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_geoPointEmpty:
		print("geoPointEmpty")
		r = TL_geoPointEmpty{}

	case crc_geoPoint:
		print("geoPoint")
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
			m.Long(),
		}

	case crc_auth_sentCode:
		print("auth_sentCode")
		r = TL_auth_sentCode{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_auth_authorization:
		print("auth_authorization")
		r = TL_auth_authorization{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_auth_exportedAuthorization:
		print("auth_exportedAuthorization")
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		print("inputNotifyPeer")
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case crc_inputNotifyUsers:
		print("inputNotifyUsers")
		r = TL_inputNotifyUsers{}

	case crc_inputNotifyChats:
		print("inputNotifyChats")
		r = TL_inputNotifyChats{}

	case crc_inputPeerNotifySettings:
		print("inputPeerNotifySettings")
		r = TL_inputPeerNotifySettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_peerNotifySettings:
		print("peerNotifySettings")
		r = TL_peerNotifySettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_peerSettings:
		print("peerSettings")
		r = TL_peerSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_wallPaper:
		print("wallPaper")
		r = TL_wallPaper{
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputReportReasonSpam:
		print("inputReportReasonSpam")
		r = TL_inputReportReasonSpam{}

	case crc_inputReportReasonViolence:
		print("inputReportReasonViolence")
		r = TL_inputReportReasonViolence{}

	case crc_inputReportReasonPornography:
		print("inputReportReasonPornography")
		r = TL_inputReportReasonPornography{}

	case crc_inputReportReasonChildAbuse:
		print("inputReportReasonChildAbuse")
		r = TL_inputReportReasonChildAbuse{}

	case crc_inputReportReasonOther:
		print("inputReportReasonOther")
		r = TL_inputReportReasonOther{
			m.String(),
		}

	case crc_userFull:
		print("userFull")
		r = TL_userFull{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_contact:
		print("contact")
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case crc_importedContact:
		print("importedContact")
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		print("contactBlocked")
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		print("contactStatus")
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_contactsNotModified:
		print("contacts_contactsNotModified")
		r = TL_contacts_contactsNotModified{}

	case crc_contacts_contacts:
		print("contacts_contacts")
		r = TL_contacts_contacts{
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_importedContacts:
		print("contacts_importedContacts")
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crc_contacts_blocked:
		print("contacts_blocked")
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_blockedSlice:
		print("contacts_blockedSlice")
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogs:
		print("messages_dialogs")
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogsSlice:
		print("messages_dialogsSlice")
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messages:
		print("messages_messages")
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messagesSlice:
		print("messages_messagesSlice")
		r = TL_messages_messagesSlice{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chats:
		print("messages_chats")
		r = TL_messages_chats{
			m.Vector(),
		}

	case crc_messages_chatFull:
		print("messages_chatFull")
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_affectedHistory:
		print("messages_affectedHistory")
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		print("inputMessagesFilterEmpty")
		r = TL_inputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		print("inputMessagesFilterPhotos")
		r = TL_inputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		print("inputMessagesFilterVideo")
		r = TL_inputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		print("inputMessagesFilterPhotoVideo")
		r = TL_inputMessagesFilterPhotoVideo{}

	case crc_inputMessagesFilterDocument:
		print("inputMessagesFilterDocument")
		r = TL_inputMessagesFilterDocument{}

	case crc_inputMessagesFilterUrl:
		print("inputMessagesFilterUrl")
		r = TL_inputMessagesFilterUrl{}

	case crc_inputMessagesFilterGif:
		print("inputMessagesFilterGif")
		r = TL_inputMessagesFilterGif{}

	case crc_updateNewMessage:
		print("updateNewMessage")
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateMessageID:
		print("updateMessageID")
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateDeleteMessages:
		print("updateDeleteMessages")
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateUserTyping:
		print("updateUserTyping")
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crc_updateChatUserTyping:
		print("updateChatUserTyping")
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChatParticipants:
		print("updateChatParticipants")
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case crc_updateUserStatus:
		print("updateUserStatus")
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crc_updateUserName:
		print("updateUserName")
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		print("updateUserPhoto")
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updates_state:
		print("updates_state")
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceEmpty:
		print("updates_differenceEmpty")
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updates_difference:
		print("updates_difference")
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updates_differenceSlice:
		print("updates_differenceSlice")
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updatesTooLong:
		print("updatesTooLong")
		r = TL_updatesTooLong{}

	case crc_updateShortMessage:
		print("updateShortMessage")
		r = TL_updateShortMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_updateShortChatMessage:
		print("updateShortChatMessage")
		r = TL_updateShortChatMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_updateShort:
		print("updateShort")
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case crc_updatesCombined:
		print("updatesCombined")
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		print("updates")
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_photos:
		print("photos_photos")
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photosSlice:
		print("photos_photosSlice")
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photo:
		print("photos_photo")
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case crc_upload_file:
		print("upload_file")
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		print("dcOption")
		r = TL_dcOption{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
		}

	case crc_config:
		print("config")
		r = TL_config{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_nearestDc:
		print("nearestDc")
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_help_appUpdate:
		print("help_appUpdate")
		r = TL_help_appUpdate{
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.Object(),
			m.Object(),
		}

	case crc_help_noAppUpdate:
		print("help_noAppUpdate")
		r = TL_help_noAppUpdate{}

	case crc_help_inviteText:
		print("help_inviteText")
		r = TL_help_inviteText{
			m.String(),
		}

	case crc_updateNewEncryptedMessage:
		print("updateNewEncryptedMessage")
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		print("updateEncryptedChatTyping")
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		print("updateEncryption")
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		print("updateEncryptedMessagesRead")
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		print("encryptedChatEmpty")
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		print("encryptedChatWaiting")
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		print("encryptedChatRequested")
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		print("encryptedChat")
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		print("encryptedChatDiscarded")
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		print("inputEncryptedChat")
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		print("encryptedFileEmpty")
		r = TL_encryptedFileEmpty{}

	case crc_encryptedFile:
		print("encryptedFile")
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		print("inputEncryptedFileEmpty")
		r = TL_inputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		print("inputEncryptedFileUploaded")
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		print("inputEncryptedFile")
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		print("inputEncryptedFileLocation")
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		print("encryptedMessage")
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_encryptedMessageService:
		print("encryptedMessageService")
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_dhConfigNotModified:
		print("messages_dhConfigNotModified")
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messages_dhConfig:
		print("messages_dhConfig")
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_sentEncryptedMessage:
		print("messages_sentEncryptedMessage")
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case crc_messages_sentEncryptedFile:
		print("messages_sentEncryptedFile")
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crc_inputFileBig:
		print("inputFileBig")
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		print("inputEncryptedFileBigUploaded")
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantAdd:
		print("updateChatParticipantAdd")
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		print("updateChatParticipantDelete")
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		print("updateDcOptions")
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case crc_inputMediaUploadedDocument:
		print("inputMediaUploadedDocument")
		r = TL_inputMediaUploadedDocument{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaDocument:
		print("inputMediaDocument")
		r = TL_inputMediaDocument{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaDocument:
		print("messageMediaDocument")
		r = TL_messageMediaDocument{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputDocumentEmpty:
		print("inputDocumentEmpty")
		r = TL_inputDocumentEmpty{}

	case crc_inputDocument:
		print("inputDocument")
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_inputDocumentFileLocation:
		print("inputDocumentFileLocation")
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case crc_documentEmpty:
		print("documentEmpty")
		r = TL_documentEmpty{
			m.Long(),
		}

	case crc_document:
		print("document")
		r = TL_document{
			m.Object(),
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_help_support:
		print("help_support")
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case crc_notifyPeer:
		print("notifyPeer")
		r = TL_notifyPeer{
			m.Object(),
		}

	case crc_notifyUsers:
		print("notifyUsers")
		r = TL_notifyUsers{}

	case crc_notifyChats:
		print("notifyChats")
		r = TL_notifyChats{}

	case crc_updateUserBlocked:
		print("updateUserBlocked")
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crc_updateNotifySettings:
		print("updateNotifySettings")
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_sendMessageTypingAction:
		print("sendMessageTypingAction")
		r = TL_sendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		print("sendMessageCancelAction")
		r = TL_sendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		print("sendMessageRecordVideoAction")
		r = TL_sendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		print("sendMessageUploadVideoAction")
		r = TL_sendMessageUploadVideoAction{
			m.Int(),
		}

	case crc_sendMessageRecordAudioAction:
		print("sendMessageRecordAudioAction")
		r = TL_sendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		print("sendMessageUploadAudioAction")
		r = TL_sendMessageUploadAudioAction{
			m.Int(),
		}

	case crc_sendMessageUploadPhotoAction:
		print("sendMessageUploadPhotoAction")
		r = TL_sendMessageUploadPhotoAction{
			m.Int(),
		}

	case crc_sendMessageUploadDocumentAction:
		print("sendMessageUploadDocumentAction")
		r = TL_sendMessageUploadDocumentAction{
			m.Int(),
		}

	case crc_sendMessageGeoLocationAction:
		print("sendMessageGeoLocationAction")
		r = TL_sendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		print("sendMessageChooseContactAction")
		r = TL_sendMessageChooseContactAction{}

	case crc_contacts_found:
		print("contacts_found")
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updateServiceNotification:
		print("updateServiceNotification")
		r = TL_updateServiceNotification{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case crc_userStatusRecently:
		print("userStatusRecently")
		r = TL_userStatusRecently{}

	case crc_userStatusLastWeek:
		print("userStatusLastWeek")
		r = TL_userStatusLastWeek{}

	case crc_userStatusLastMonth:
		print("userStatusLastMonth")
		r = TL_userStatusLastMonth{}

	case crc_updatePrivacy:
		print("updatePrivacy")
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		print("inputPrivacyKeyStatusTimestamp")
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		print("privacyKeyStatusTimestamp")
		r = TL_privacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		print("inputPrivacyValueAllowContacts")
		r = TL_inputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		print("inputPrivacyValueAllowAll")
		r = TL_inputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		print("inputPrivacyValueAllowUsers")
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crc_inputPrivacyValueDisallowContacts:
		print("inputPrivacyValueDisallowContacts")
		r = TL_inputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		print("inputPrivacyValueDisallowAll")
		r = TL_inputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		print("inputPrivacyValueDisallowUsers")
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crc_privacyValueAllowContacts:
		print("privacyValueAllowContacts")
		r = TL_privacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		print("privacyValueAllowAll")
		r = TL_privacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		print("privacyValueAllowUsers")
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		print("privacyValueDisallowContacts")
		r = TL_privacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		print("privacyValueDisallowAll")
		r = TL_privacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		print("privacyValueDisallowUsers")
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_account_privacyRules:
		print("account_privacyRules")
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_accountDaysTTL:
		print("accountDaysTTL")
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case crc_updateUserPhone:
		print("updateUserPhone")
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		print("documentAttributeImageSize")
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		print("documentAttributeAnimated")
		r = TL_documentAttributeAnimated{}

	case crc_documentAttributeSticker:
		print("documentAttributeSticker")
		r = TL_documentAttributeSticker{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_documentAttributeVideo:
		print("documentAttributeVideo")
		r = TL_documentAttributeVideo{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		print("documentAttributeAudio")
		r = TL_documentAttributeAudio{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_documentAttributeFilename:
		print("documentAttributeFilename")
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case crc_messages_stickersNotModified:
		print("messages_stickersNotModified")
		r = TL_messages_stickersNotModified{}

	case crc_messages_stickers:
		print("messages_stickers")
		r = TL_messages_stickers{
			m.Int(),
			m.Vector(),
		}

	case crc_stickerPack:
		print("stickerPack")
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messages_allStickersNotModified:
		print("messages_allStickersNotModified")
		r = TL_messages_allStickersNotModified{}

	case crc_messages_allStickers:
		print("messages_allStickers")
		r = TL_messages_allStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_updateReadHistoryInbox:
		print("updateReadHistoryInbox")
		r = TL_updateReadHistoryInbox{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadHistoryOutbox:
		print("updateReadHistoryOutbox")
		r = TL_updateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_affectedMessages:
		print("messages_affectedMessages")
		r = TL_messages_affectedMessages{
			m.Int(),
			m.Int(),
		}

	case crc_updateWebPage:
		print("updateWebPage")
		r = TL_updateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_webPageEmpty:
		print("webPageEmpty")
		r = TL_webPageEmpty{
			m.Long(),
		}

	case crc_webPagePending:
		print("webPagePending")
		r = TL_webPagePending{
			m.Long(),
			m.Int(),
		}

	case crc_webPage:
		print("webPage")
		r = TL_webPage{
			m.Object(),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaWebPage:
		print("messageMediaWebPage")
		r = TL_messageMediaWebPage{
			m.Object(),
		}

	case crc_authorization:
		print("authorization")
		r = TL_authorization{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_account_authorizations:
		print("account_authorizations")
		r = TL_account_authorizations{
			m.Vector(),
		}

	case crc_account_password:
		print("account_password")
		r = TL_account_password{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.StringBytes(),
		}

	case crc_account_passwordSettings:
		print("account_passwordSettings")
		r = TL_account_passwordSettings{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_passwordInputSettings:
		print("account_passwordInputSettings")
		r = TL_account_passwordInputSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_auth_passwordRecovery:
		print("auth_passwordRecovery")
		r = TL_auth_passwordRecovery{
			m.String(),
		}

	case crc_inputMediaVenue:
		print("inputMediaVenue")
		r = TL_inputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messageMediaVenue:
		print("messageMediaVenue")
		r = TL_messageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_receivedNotifyMessage:
		print("receivedNotifyMessage")
		r = TL_receivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crc_chatInviteEmpty:
		print("chatInviteEmpty")
		r = TL_chatInviteEmpty{}

	case crc_chatInviteExported:
		print("chatInviteExported")
		r = TL_chatInviteExported{
			m.String(),
		}

	case crc_chatInviteAlready:
		print("chatInviteAlready")
		r = TL_chatInviteAlready{
			m.Object(),
		}

	case crc_chatInvite:
		print("chatInvite")
		r = TL_chatInvite{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messageActionChatJoinedByLink:
		print("messageActionChatJoinedByLink")
		r = TL_messageActionChatJoinedByLink{
			m.Int(),
		}

	case crc_updateReadMessagesContents:
		print("updateReadMessagesContents")
		r = TL_updateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_inputStickerSetEmpty:
		print("inputStickerSetEmpty")
		r = TL_inputStickerSetEmpty{}

	case crc_inputStickerSetID:
		print("inputStickerSetID")
		r = TL_inputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crc_inputStickerSetShortName:
		print("inputStickerSetShortName")
		r = TL_inputStickerSetShortName{
			m.String(),
		}

	case crc_stickerSet:
		print("stickerSet")
		r = TL_stickerSet{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_stickerSet:
		print("messages_stickerSet")
		r = TL_messages_stickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_user:
		print("user")
		r = TL_user{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_botCommand:
		print("botCommand")
		r = TL_botCommand{
			m.String(),
			m.String(),
		}

	case crc_botInfo:
		print("botInfo")
		r = TL_botInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_keyboardButton:
		print("keyboardButton")
		r = TL_keyboardButton{
			m.String(),
		}

	case crc_keyboardButtonRow:
		print("keyboardButtonRow")
		r = TL_keyboardButtonRow{
			m.Vector(),
		}

	case crc_replyKeyboardHide:
		print("replyKeyboardHide")
		r = TL_replyKeyboardHide{
			m.Object(),
			m.Object(),
		}

	case crc_replyKeyboardForceReply:
		print("replyKeyboardForceReply")
		r = TL_replyKeyboardForceReply{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_replyKeyboardMarkup:
		print("replyKeyboardMarkup")
		r = TL_replyKeyboardMarkup{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_inputPeerUser:
		print("inputPeerUser")
		r = TL_inputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crc_inputUser:
		print("inputUser")
		r = TL_inputUser{
			m.Int(),
			m.Long(),
		}

	case crc_messageEntityUnknown:
		print("messageEntityUnknown")
		r = TL_messageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityMention:
		print("messageEntityMention")
		r = TL_messageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityHashtag:
		print("messageEntityHashtag")
		r = TL_messageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBotCommand:
		print("messageEntityBotCommand")
		r = TL_messageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityUrl:
		print("messageEntityUrl")
		r = TL_messageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityEmail:
		print("messageEntityEmail")
		r = TL_messageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBold:
		print("messageEntityBold")
		r = TL_messageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityItalic:
		print("messageEntityItalic")
		r = TL_messageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCode:
		print("messageEntityCode")
		r = TL_messageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityPre:
		print("messageEntityPre")
		r = TL_messageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_messageEntityTextUrl:
		print("messageEntityTextUrl")
		r = TL_messageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateShortSentMessage:
		print("updateShortSentMessage")
		r = TL_updateShortSentMessage{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_inputChannelEmpty:
		print("inputChannelEmpty")
		r = TL_inputChannelEmpty{}

	case crc_inputChannel:
		print("inputChannel")
		r = TL_inputChannel{
			m.Int(),
			m.Long(),
		}

	case crc_peerChannel:
		print("peerChannel")
		r = TL_peerChannel{
			m.Int(),
		}

	case crc_inputPeerChannel:
		print("inputPeerChannel")
		r = TL_inputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crc_channel:
		print("channel")
		r = TL_channel{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channelForbidden:
		print("channelForbidden")
		r = TL_channelForbidden{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.String(),
			m.Object(),
		}

	case crc_contacts_resolvedPeer:
		print("contacts_resolvedPeer")
		r = TL_contacts_resolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelFull:
		print("channelFull")
		r = TL_channelFull{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messageRange:
		print("messageRange")
		r = TL_messageRange{
			m.Int(),
			m.Int(),
		}

	case crc_messages_channelMessages:
		print("messages_channelMessages")
		r = TL_messages_channelMessages{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageActionChannelCreate:
		print("messageActionChannelCreate")
		r = TL_messageActionChannelCreate{
			m.String(),
		}

	case crc_updateChannelTooLong:
		print("updateChannelTooLong")
		r = TL_updateChannelTooLong{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChannel:
		print("updateChannel")
		r = TL_updateChannel{
			m.Int(),
		}

	case crc_updateNewChannelMessage:
		print("updateNewChannelMessage")
		r = TL_updateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadChannelInbox:
		print("updateReadChannelInbox")
		r = TL_updateReadChannelInbox{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDeleteChannelMessages:
		print("updateDeleteChannelMessages")
		r = TL_updateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelMessageViews:
		print("updateChannelMessageViews")
		r = TL_updateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_channelDifferenceEmpty:
		print("updates_channelDifferenceEmpty")
		r = TL_updates_channelDifferenceEmpty{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_updates_channelDifferenceTooLong:
		print("updates_channelDifferenceTooLong")
		r = TL_updates_channelDifferenceTooLong{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updates_channelDifference:
		print("updates_channelDifference")
		r = TL_updates_channelDifference{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelMessagesFilterEmpty:
		print("channelMessagesFilterEmpty")
		r = TL_channelMessagesFilterEmpty{}

	case crc_channelMessagesFilter:
		print("channelMessagesFilter")
		r = TL_channelMessagesFilter{
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_channelParticipant:
		print("channelParticipant")
		r = TL_channelParticipant{
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantSelf:
		print("channelParticipantSelf")
		r = TL_channelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantCreator:
		print("channelParticipantCreator")
		r = TL_channelParticipantCreator{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantsRecent:
		print("channelParticipantsRecent")
		r = TL_channelParticipantsRecent{}

	case crc_channelParticipantsAdmins:
		print("channelParticipantsAdmins")
		r = TL_channelParticipantsAdmins{}

	case crc_channelParticipantsKicked:
		print("channelParticipantsKicked")
		r = TL_channelParticipantsKicked{
			m.String(),
		}

	case crc_channels_channelParticipants:
		print("channels_channelParticipants")
		r = TL_channels_channelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channels_channelParticipant:
		print("channels_channelParticipant")
		r = TL_channels_channelParticipant{
			m.Object(),
			m.Vector(),
		}

	case crc_chatParticipantCreator:
		print("chatParticipantCreator")
		r = TL_chatParticipantCreator{
			m.Int(),
		}

	case crc_chatParticipantAdmin:
		print("chatParticipantAdmin")
		r = TL_chatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantAdmin:
		print("updateChatParticipantAdmin")
		r = TL_updateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messageActionChatMigrateTo:
		print("messageActionChatMigrateTo")
		r = TL_messageActionChatMigrateTo{
			m.Int(),
		}

	case crc_messageActionChannelMigrateFrom:
		print("messageActionChannelMigrateFrom")
		r = TL_messageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crc_channelParticipantsBots:
		print("channelParticipantsBots")
		r = TL_channelParticipantsBots{}

	case crc_help_termsOfService:
		print("help_termsOfService")
		r = TL_help_termsOfService{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
			m.Object(),
		}

	case crc_updateNewStickerSet:
		print("updateNewStickerSet")
		r = TL_updateNewStickerSet{
			m.Object(),
		}

	case crc_updateStickerSetsOrder:
		print("updateStickerSetsOrder")
		r = TL_updateStickerSetsOrder{
			m.Object(),
			m.Object(),
			m.VectorLong(),
		}

	case crc_updateStickerSets:
		print("updateStickerSets")
		r = TL_updateStickerSets{}

	case crc_foundGif:
		print("foundGif")
		r = TL_foundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_foundGifCached:
		print("foundGifCached")
		r = TL_foundGifCached{
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaGifExternal:
		print("inputMediaGifExternal")
		r = TL_inputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case crc_messages_foundGifs:
		print("messages_foundGifs")
		r = TL_messages_foundGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_savedGifsNotModified:
		print("messages_savedGifsNotModified")
		r = TL_messages_savedGifsNotModified{}

	case crc_messages_savedGifs:
		print("messages_savedGifs")
		r = TL_messages_savedGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_updateSavedGifs:
		print("updateSavedGifs")
		r = TL_updateSavedGifs{}

	case crc_inputBotInlineMessageMediaAuto:
		print("inputBotInlineMessageMediaAuto")
		r = TL_inputBotInlineMessageMediaAuto{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineMessageText:
		print("inputBotInlineMessageText")
		r = TL_inputBotInlineMessageText{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResult:
		print("inputBotInlineResult")
		r = TL_inputBotInlineResult{
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineMessageMediaAuto:
		print("botInlineMessageMediaAuto")
		r = TL_botInlineMessageMediaAuto{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineMessageText:
		print("botInlineMessageText")
		r = TL_botInlineMessageText{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineResult:
		print("botInlineResult")
		r = TL_botInlineResult{
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_botResults:
		print("messages_botResults")
		r = TL_messages_botResults{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crc_updateBotInlineQuery:
		print("updateBotInlineQuery")
		r = TL_updateBotInlineQuery{
			m.Object(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_updateBotInlineSend:
		print("updateBotInlineSend")
		r = TL_updateBotInlineSend{
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_inputMessagesFilterVoice:
		print("inputMessagesFilterVoice")
		r = TL_inputMessagesFilterVoice{}

	case crc_inputMessagesFilterMusic:
		print("inputMessagesFilterMusic")
		r = TL_inputMessagesFilterMusic{}

	case crc_inputPrivacyKeyChatInvite:
		print("inputPrivacyKeyChatInvite")
		r = TL_inputPrivacyKeyChatInvite{}

	case crc_privacyKeyChatInvite:
		print("privacyKeyChatInvite")
		r = TL_privacyKeyChatInvite{}

	case crc_exportedMessageLink:
		print("exportedMessageLink")
		r = TL_exportedMessageLink{
			m.String(),
			m.String(),
		}

	case crc_messageFwdHeader:
		print("messageFwdHeader")
		r = TL_messageFwdHeader{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_updateEditChannelMessage:
		print("updateEditChannelMessage")
		r = TL_updateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelPinnedMessage:
		print("updateChannelPinnedMessage")
		r = TL_updateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_messageActionPinMessage:
		print("messageActionPinMessage")
		r = TL_messageActionPinMessage{}

	case crc_auth_codeTypeSms:
		print("auth_codeTypeSms")
		r = TL_auth_codeTypeSms{}

	case crc_auth_codeTypeCall:
		print("auth_codeTypeCall")
		r = TL_auth_codeTypeCall{}

	case crc_auth_codeTypeFlashCall:
		print("auth_codeTypeFlashCall")
		r = TL_auth_codeTypeFlashCall{}

	case crc_auth_sentCodeTypeApp:
		print("auth_sentCodeTypeApp")
		r = TL_auth_sentCodeTypeApp{
			m.Int(),
		}

	case crc_auth_sentCodeTypeSms:
		print("auth_sentCodeTypeSms")
		r = TL_auth_sentCodeTypeSms{
			m.Int(),
		}

	case crc_auth_sentCodeTypeCall:
		print("auth_sentCodeTypeCall")
		r = TL_auth_sentCodeTypeCall{
			m.Int(),
		}

	case crc_auth_sentCodeTypeFlashCall:
		print("auth_sentCodeTypeFlashCall")
		r = TL_auth_sentCodeTypeFlashCall{
			m.String(),
		}

	case crc_keyboardButtonUrl:
		print("keyboardButtonUrl")
		r = TL_keyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crc_keyboardButtonCallback:
		print("keyboardButtonCallback")
		r = TL_keyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crc_keyboardButtonRequestPhone:
		print("keyboardButtonRequestPhone")
		r = TL_keyboardButtonRequestPhone{
			m.String(),
		}

	case crc_keyboardButtonRequestGeoLocation:
		print("keyboardButtonRequestGeoLocation")
		r = TL_keyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crc_keyboardButtonSwitchInline:
		print("keyboardButtonSwitchInline")
		r = TL_keyboardButtonSwitchInline{
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_replyInlineMarkup:
		print("replyInlineMarkup")
		r = TL_replyInlineMarkup{
			m.Vector(),
		}

	case crc_messages_botCallbackAnswer:
		print("messages_botCallbackAnswer")
		r = TL_messages_botCallbackAnswer{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_updateBotCallbackQuery:
		print("updateBotCallbackQuery")
		r = TL_updateBotCallbackQuery{
			m.Object(),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_messageEditData:
		print("messages_messageEditData")
		r = TL_messages_messageEditData{
			m.Object(),
			m.Object(),
		}

	case crc_updateEditMessage:
		print("updateEditMessage")
		r = TL_updateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputBotInlineMessageMediaGeo:
		print("inputBotInlineMessageMediaGeo")
		r = TL_inputBotInlineMessageMediaGeo{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_inputBotInlineMessageMediaVenue:
		print("inputBotInlineMessageMediaVenue")
		r = TL_inputBotInlineMessageMediaVenue{
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_inputBotInlineMessageMediaContact:
		print("inputBotInlineMessageMediaContact")
		r = TL_inputBotInlineMessageMediaContact{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_botInlineMessageMediaGeo:
		print("botInlineMessageMediaGeo")
		r = TL_botInlineMessageMediaGeo{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_botInlineMessageMediaVenue:
		print("botInlineMessageMediaVenue")
		r = TL_botInlineMessageMediaVenue{
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_botInlineMessageMediaContact:
		print("botInlineMessageMediaContact")
		r = TL_botInlineMessageMediaContact{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_inputBotInlineResultPhoto:
		print("inputBotInlineResultPhoto")
		r = TL_inputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResultDocument:
		print("inputBotInlineResultDocument")
		r = TL_inputBotInlineResultDocument{
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineMediaResult:
		print("botInlineMediaResult")
		r = TL_botInlineMediaResult{
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineMessageID:
		print("inputBotInlineMessageID")
		r = TL_inputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crc_updateInlineBotCallbackQuery:
		print("updateInlineBotCallbackQuery")
		r = TL_updateInlineBotCallbackQuery{
			m.Object(),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_inlineBotSwitchPM:
		print("inlineBotSwitchPM")
		r = TL_inlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crc_messages_peerDialogs:
		print("messages_peerDialogs")
		r = TL_messages_peerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_topPeer:
		print("topPeer")
		r = TL_topPeer{
			m.Object(),
			m.Double(),
		}

	case crc_topPeerCategoryBotsPM:
		print("topPeerCategoryBotsPM")
		r = TL_topPeerCategoryBotsPM{}

	case crc_topPeerCategoryBotsInline:
		print("topPeerCategoryBotsInline")
		r = TL_topPeerCategoryBotsInline{}

	case crc_topPeerCategoryCorrespondents:
		print("topPeerCategoryCorrespondents")
		r = TL_topPeerCategoryCorrespondents{}

	case crc_topPeerCategoryGroups:
		print("topPeerCategoryGroups")
		r = TL_topPeerCategoryGroups{}

	case crc_topPeerCategoryChannels:
		print("topPeerCategoryChannels")
		r = TL_topPeerCategoryChannels{}

	case crc_topPeerCategoryPeers:
		print("topPeerCategoryPeers")
		r = TL_topPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_topPeersNotModified:
		print("contacts_topPeersNotModified")
		r = TL_contacts_topPeersNotModified{}

	case crc_contacts_topPeers:
		print("contacts_topPeers")
		r = TL_contacts_topPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageEntityMentionName:
		print("messageEntityMentionName")
		r = TL_messageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessageEntityMentionName:
		print("inputMessageEntityMentionName")
		r = TL_inputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_inputMessagesFilterChatPhotos:
		print("inputMessagesFilterChatPhotos")
		r = TL_inputMessagesFilterChatPhotos{}

	case crc_updateReadChannelOutbox:
		print("updateReadChannelOutbox")
		r = TL_updateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDraftMessage:
		print("updateDraftMessage")
		r = TL_updateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case crc_draftMessageEmpty:
		print("draftMessageEmpty")
		r = TL_draftMessageEmpty{
			m.Object(),
			m.Object(),
		}

	case crc_draftMessage:
		print("draftMessage")
		r = TL_draftMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case crc_messageActionHistoryClear:
		print("messageActionHistoryClear")
		r = TL_messageActionHistoryClear{}

	case crc_messages_featuredStickersNotModified:
		print("messages_featuredStickersNotModified")
		r = TL_messages_featuredStickersNotModified{}

	case crc_messages_featuredStickers:
		print("messages_featuredStickers")
		r = TL_messages_featuredStickers{
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case crc_updateReadFeaturedStickers:
		print("updateReadFeaturedStickers")
		r = TL_updateReadFeaturedStickers{}

	case crc_messages_recentStickersNotModified:
		print("messages_recentStickersNotModified")
		r = TL_messages_recentStickersNotModified{}

	case crc_messages_recentStickers:
		print("messages_recentStickers")
		r = TL_messages_recentStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.VectorInt(),
		}

	case crc_updateRecentStickers:
		print("updateRecentStickers")
		r = TL_updateRecentStickers{}

	case crc_messages_archivedStickers:
		print("messages_archivedStickers")
		r = TL_messages_archivedStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_stickerSetInstallResultSuccess:
		print("messages_stickerSetInstallResultSuccess")
		r = TL_messages_stickerSetInstallResultSuccess{}

	case crc_messages_stickerSetInstallResultArchive:
		print("messages_stickerSetInstallResultArchive")
		r = TL_messages_stickerSetInstallResultArchive{
			m.Vector(),
		}

	case crc_stickerSetCovered:
		print("stickerSetCovered")
		r = TL_stickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case crc_updateConfig:
		print("updateConfig")
		r = TL_updateConfig{}

	case crc_updatePtsChanged:
		print("updatePtsChanged")
		r = TL_updatePtsChanged{}

	case crc_inputMediaPhotoExternal:
		print("inputMediaPhotoExternal")
		r = TL_inputMediaPhotoExternal{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_inputMediaDocumentExternal:
		print("inputMediaDocumentExternal")
		r = TL_inputMediaDocumentExternal{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_stickerSetMultiCovered:
		print("stickerSetMultiCovered")
		r = TL_stickerSetMultiCovered{
			m.Object(),
			m.Vector(),
		}

	case crc_maskCoords:
		print("maskCoords")
		r = TL_maskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_documentAttributeHasStickers:
		print("documentAttributeHasStickers")
		r = TL_documentAttributeHasStickers{}

	case crc_inputStickeredMediaPhoto:
		print("inputStickeredMediaPhoto")
		r = TL_inputStickeredMediaPhoto{
			m.Object(),
		}

	case crc_inputStickeredMediaDocument:
		print("inputStickeredMediaDocument")
		r = TL_inputStickeredMediaDocument{
			m.Object(),
		}

	case crc_game:
		print("game")
		r = TL_game{
			m.Object(),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResultGame:
		print("inputBotInlineResultGame")
		r = TL_inputBotInlineResultGame{
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_inputBotInlineMessageGame:
		print("inputBotInlineMessageGame")
		r = TL_inputBotInlineMessageGame{
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaGame:
		print("messageMediaGame")
		r = TL_messageMediaGame{
			m.Object(),
		}

	case crc_inputMediaGame:
		print("inputMediaGame")
		r = TL_inputMediaGame{
			m.Object(),
		}

	case crc_inputGameID:
		print("inputGameID")
		r = TL_inputGameID{
			m.Long(),
			m.Long(),
		}

	case crc_inputGameShortName:
		print("inputGameShortName")
		r = TL_inputGameShortName{
			m.Object(),
			m.String(),
		}

	case crc_keyboardButtonGame:
		print("keyboardButtonGame")
		r = TL_keyboardButtonGame{
			m.String(),
		}

	case crc_messageActionGameScore:
		print("messageActionGameScore")
		r = TL_messageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case crc_highScore:
		print("highScore")
		r = TL_highScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_highScores:
		print("messages_highScores")
		r = TL_messages_highScores{
			m.Vector(),
			m.Vector(),
		}

	case crc_updates_differenceTooLong:
		print("updates_differenceTooLong")
		r = TL_updates_differenceTooLong{
			m.Int(),
		}

	case crc_updateChannelWebPage:
		print("updateChannelWebPage")
		r = TL_updateChannelWebPage{
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_chatsSlice:
		print("messages_chatsSlice")
		r = TL_messages_chatsSlice{
			m.Int(),
			m.Vector(),
		}

	case crc_textEmpty:
		print("textEmpty")
		r = TL_textEmpty{}

	case crc_textPlain:
		print("textPlain")
		r = TL_textPlain{
			m.String(),
		}

	case crc_textBold:
		print("textBold")
		r = TL_textBold{
			m.Object(),
		}

	case crc_textItalic:
		print("textItalic")
		r = TL_textItalic{
			m.Object(),
		}

	case crc_textUnderline:
		print("textUnderline")
		r = TL_textUnderline{
			m.Object(),
		}

	case crc_textStrike:
		print("textStrike")
		r = TL_textStrike{
			m.Object(),
		}

	case crc_textFixed:
		print("textFixed")
		r = TL_textFixed{
			m.Object(),
		}

	case crc_textUrl:
		print("textUrl")
		r = TL_textUrl{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_textEmail:
		print("textEmail")
		r = TL_textEmail{
			m.Object(),
			m.String(),
		}

	case crc_textConcat:
		print("textConcat")
		r = TL_textConcat{
			m.Vector(),
		}

	case crc_pageBlockUnsupported:
		print("pageBlockUnsupported")
		r = TL_pageBlockUnsupported{}

	case crc_pageBlockTitle:
		print("pageBlockTitle")
		r = TL_pageBlockTitle{
			m.Object(),
		}

	case crc_pageBlockSubtitle:
		print("pageBlockSubtitle")
		r = TL_pageBlockSubtitle{
			m.Object(),
		}

	case crc_pageBlockAuthorDate:
		print("pageBlockAuthorDate")
		r = TL_pageBlockAuthorDate{
			m.Object(),
			m.Int(),
		}

	case crc_pageBlockHeader:
		print("pageBlockHeader")
		r = TL_pageBlockHeader{
			m.Object(),
		}

	case crc_pageBlockSubheader:
		print("pageBlockSubheader")
		r = TL_pageBlockSubheader{
			m.Object(),
		}

	case crc_pageBlockParagraph:
		print("pageBlockParagraph")
		r = TL_pageBlockParagraph{
			m.Object(),
		}

	case crc_pageBlockPreformatted:
		print("pageBlockPreformatted")
		r = TL_pageBlockPreformatted{
			m.Object(),
			m.String(),
		}

	case crc_pageBlockFooter:
		print("pageBlockFooter")
		r = TL_pageBlockFooter{
			m.Object(),
		}

	case crc_pageBlockDivider:
		print("pageBlockDivider")
		r = TL_pageBlockDivider{}

	case crc_pageBlockAnchor:
		print("pageBlockAnchor")
		r = TL_pageBlockAnchor{
			m.String(),
		}

	case crc_pageBlockList:
		print("pageBlockList")
		r = TL_pageBlockList{
			m.Vector(),
		}

	case crc_pageBlockBlockquote:
		print("pageBlockBlockquote")
		r = TL_pageBlockBlockquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPullquote:
		print("pageBlockPullquote")
		r = TL_pageBlockPullquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPhoto:
		print("pageBlockPhoto")
		r = TL_pageBlockPhoto{
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockVideo:
		print("pageBlockVideo")
		r = TL_pageBlockVideo{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
		}

	case crc_pageBlockCover:
		print("pageBlockCover")
		r = TL_pageBlockCover{
			m.Object(),
		}

	case crc_pageBlockEmbed:
		print("pageBlockEmbed")
		r = TL_pageBlockEmbed{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockEmbedPost:
		print("pageBlockEmbedPost")
		r = TL_pageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			m.Vector(),
			m.Object(),
		}

	case crc_pageBlockCollage:
		print("pageBlockCollage")
		r = TL_pageBlockCollage{
			m.Vector(),
			m.Object(),
		}

	case crc_pageBlockSlideshow:
		print("pageBlockSlideshow")
		r = TL_pageBlockSlideshow{
			m.Vector(),
			m.Object(),
		}

	case crc_webPageNotModified:
		print("webPageNotModified")
		r = TL_webPageNotModified{}

	case crc_inputPrivacyKeyPhoneCall:
		print("inputPrivacyKeyPhoneCall")
		r = TL_inputPrivacyKeyPhoneCall{}

	case crc_privacyKeyPhoneCall:
		print("privacyKeyPhoneCall")
		r = TL_privacyKeyPhoneCall{}

	case crc_sendMessageGamePlayAction:
		print("sendMessageGamePlayAction")
		r = TL_sendMessageGamePlayAction{}

	case crc_phoneCallDiscardReasonMissed:
		print("phoneCallDiscardReasonMissed")
		r = TL_phoneCallDiscardReasonMissed{}

	case crc_phoneCallDiscardReasonDisconnect:
		print("phoneCallDiscardReasonDisconnect")
		r = TL_phoneCallDiscardReasonDisconnect{}

	case crc_phoneCallDiscardReasonHangup:
		print("phoneCallDiscardReasonHangup")
		r = TL_phoneCallDiscardReasonHangup{}

	case crc_phoneCallDiscardReasonBusy:
		print("phoneCallDiscardReasonBusy")
		r = TL_phoneCallDiscardReasonBusy{}

	case crc_updateDialogPinned:
		print("updateDialogPinned")
		r = TL_updateDialogPinned{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_updatePinnedDialogs:
		print("updatePinnedDialogs")
		r = TL_updatePinnedDialogs{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_dataJSON:
		print("dataJSON")
		r = TL_dataJSON{
			m.String(),
		}

	case crc_updateBotWebhookJSON:
		print("updateBotWebhookJSON")
		r = TL_updateBotWebhookJSON{
			m.Object(),
		}

	case crc_updateBotWebhookJSONQuery:
		print("updateBotWebhookJSONQuery")
		r = TL_updateBotWebhookJSONQuery{
			m.Long(),
			m.Object(),
			m.Int(),
		}

	case crc_labeledPrice:
		print("labeledPrice")
		r = TL_labeledPrice{
			m.String(),
			m.Long(),
		}

	case crc_invoice:
		print("invoice")
		r = TL_invoice{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crc_inputMediaInvoice:
		print("inputMediaInvoice")
		r = TL_inputMediaInvoice{
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_paymentCharge:
		print("paymentCharge")
		r = TL_paymentCharge{
			m.String(),
			m.String(),
		}

	case crc_messageActionPaymentSentMe:
		print("messageActionPaymentSentMe")
		r = TL_messageActionPaymentSentMe{
			m.Object(),
			m.String(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaInvoice:
		print("messageMediaInvoice")
		r = TL_messageMediaInvoice{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_postAddress:
		print("postAddress")
		r = TL_postAddress{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_paymentRequestedInfo:
		print("paymentRequestedInfo")
		r = TL_paymentRequestedInfo{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_keyboardButtonBuy:
		print("keyboardButtonBuy")
		r = TL_keyboardButtonBuy{
			m.String(),
		}

	case crc_messageActionPaymentSent:
		print("messageActionPaymentSent")
		r = TL_messageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case crc_paymentSavedCredentialsCard:
		print("paymentSavedCredentialsCard")
		r = TL_paymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case crc_webDocument:
		print("webDocument")
		r = TL_webDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_inputWebDocument:
		print("inputWebDocument")
		r = TL_inputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_inputWebFileLocation:
		print("inputWebFileLocation")
		r = TL_inputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case crc_upload_webFile:
		print("upload_webFile")
		r = TL_upload_webFile{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_payments_paymentForm:
		print("payments_paymentForm")
		r = TL_payments_paymentForm{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_payments_validatedRequestedInfo:
		print("payments_validatedRequestedInfo")
		r = TL_payments_validatedRequestedInfo{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_payments_paymentResult:
		print("payments_paymentResult")
		r = TL_payments_paymentResult{
			m.Object(),
		}

	case crc_payments_paymentReceipt:
		print("payments_paymentReceipt")
		r = TL_payments_paymentReceipt{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Long(),
			m.String(),
			m.Vector(),
		}

	case crc_payments_savedInfo:
		print("payments_savedInfo")
		r = TL_payments_savedInfo{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputPaymentCredentialsSaved:
		print("inputPaymentCredentialsSaved")
		r = TL_inputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case crc_inputPaymentCredentials:
		print("inputPaymentCredentials")
		r = TL_inputPaymentCredentials{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_tmpPassword:
		print("account_tmpPassword")
		r = TL_account_tmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_shippingOption:
		print("shippingOption")
		r = TL_shippingOption{
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_updateBotShippingQuery:
		print("updateBotShippingQuery")
		r = TL_updateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_updateBotPrecheckoutQuery:
		print("updateBotPrecheckoutQuery")
		r = TL_updateBotPrecheckoutQuery{
			m.Object(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_inputStickerSetItem:
		print("inputStickerSetItem")
		r = TL_inputStickerSetItem{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_updatePhoneCall:
		print("updatePhoneCall")
		r = TL_updatePhoneCall{
			m.Object(),
		}

	case crc_inputPhoneCall:
		print("inputPhoneCall")
		r = TL_inputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case crc_phoneCallEmpty:
		print("phoneCallEmpty")
		r = TL_phoneCallEmpty{
			m.Long(),
		}

	case crc_phoneCallWaiting:
		print("phoneCallWaiting")
		r = TL_phoneCallWaiting{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_phoneCallRequested:
		print("phoneCallRequested")
		r = TL_phoneCallRequested{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phoneCallAccepted:
		print("phoneCallAccepted")
		r = TL_phoneCallAccepted{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phoneCall:
		print("phoneCall")
		r = TL_phoneCall{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
			m.Vector(),
			m.Int(),
		}

	case crc_phoneCallDiscarded:
		print("phoneCallDiscarded")
		r = TL_phoneCallDiscarded{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_phoneConnection:
		print("phoneConnection")
		r = TL_phoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_phoneCallProtocol:
		print("phoneCallProtocol")
		r = TL_phoneCallProtocol{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_phone_phoneCall:
		print("phone_phoneCall")
		r = TL_phone_phoneCall{
			m.Object(),
			m.Vector(),
		}

	case crc_inputMessagesFilterPhoneCalls:
		print("inputMessagesFilterPhoneCalls")
		r = TL_inputMessagesFilterPhoneCalls{
			m.Object(),
			m.Object(),
		}

	case crc_messageActionPhoneCall:
		print("messageActionPhoneCall")
		r = TL_messageActionPhoneCall{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMessagesFilterRoundVoice:
		print("inputMessagesFilterRoundVoice")
		r = TL_inputMessagesFilterRoundVoice{}

	case crc_inputMessagesFilterRoundVideo:
		print("inputMessagesFilterRoundVideo")
		r = TL_inputMessagesFilterRoundVideo{}

	case crc_sendMessageRecordRoundAction:
		print("sendMessageRecordRoundAction")
		r = TL_sendMessageRecordRoundAction{}

	case crc_sendMessageUploadRoundAction:
		print("sendMessageUploadRoundAction")
		r = TL_sendMessageUploadRoundAction{
			m.Int(),
		}

	case crc_upload_fileCdnRedirect:
		print("upload_fileCdnRedirect")
		r = TL_upload_fileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			m.Vector(),
		}

	case crc_upload_cdnFileReuploadNeeded:
		print("upload_cdnFileReuploadNeeded")
		r = TL_upload_cdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case crc_upload_cdnFile:
		print("upload_cdnFile")
		r = TL_upload_cdnFile{
			m.StringBytes(),
		}

	case crc_cdnPublicKey:
		print("cdnPublicKey")
		r = TL_cdnPublicKey{
			m.Int(),
			m.String(),
		}

	case crc_cdnConfig:
		print("cdnConfig")
		r = TL_cdnConfig{
			m.Vector(),
		}

	case crc_pageBlockChannel:
		print("pageBlockChannel")
		r = TL_pageBlockChannel{
			m.Object(),
		}

	case crc_langPackString:
		print("langPackString")
		r = TL_langPackString{
			m.String(),
			m.String(),
		}

	case crc_langPackStringPluralized:
		print("langPackStringPluralized")
		r = TL_langPackStringPluralized{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crc_langPackStringDeleted:
		print("langPackStringDeleted")
		r = TL_langPackStringDeleted{
			m.String(),
		}

	case crc_langPackDifference:
		print("langPackDifference")
		r = TL_langPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_langPackLanguage:
		print("langPackLanguage")
		r = TL_langPackLanguage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateLangPackTooLong:
		print("updateLangPackTooLong")
		r = TL_updateLangPackTooLong{
			m.String(),
		}

	case crc_updateLangPack:
		print("updateLangPack")
		r = TL_updateLangPack{
			m.Object(),
		}

	case crc_channelParticipantAdmin:
		print("channelParticipantAdmin")
		r = TL_channelParticipantAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_channelParticipantBanned:
		print("channelParticipantBanned")
		r = TL_channelParticipantBanned{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantsBanned:
		print("channelParticipantsBanned")
		r = TL_channelParticipantsBanned{
			m.String(),
		}

	case crc_channelParticipantsSearch:
		print("channelParticipantsSearch")
		r = TL_channelParticipantsSearch{
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeTitle:
		print("channelAdminLogEventActionChangeTitle")
		r = TL_channelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeAbout:
		print("channelAdminLogEventActionChangeAbout")
		r = TL_channelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeUsername:
		print("channelAdminLogEventActionChangeUsername")
		r = TL_channelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangePhoto:
		print("channelAdminLogEventActionChangePhoto")
		r = TL_channelAdminLogEventActionChangePhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleInvites:
		print("channelAdminLogEventActionToggleInvites")
		r = TL_channelAdminLogEventActionToggleInvites{
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleSignatures:
		print("channelAdminLogEventActionToggleSignatures")
		r = TL_channelAdminLogEventActionToggleSignatures{
			m.Object(),
		}

	case crc_channelAdminLogEventActionUpdatePinned:
		print("channelAdminLogEventActionUpdatePinned")
		r = TL_channelAdminLogEventActionUpdatePinned{
			m.Object(),
		}

	case crc_channelAdminLogEventActionEditMessage:
		print("channelAdminLogEventActionEditMessage")
		r = TL_channelAdminLogEventActionEditMessage{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionDeleteMessage:
		print("channelAdminLogEventActionDeleteMessage")
		r = TL_channelAdminLogEventActionDeleteMessage{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantJoin:
		print("channelAdminLogEventActionParticipantJoin")
		r = TL_channelAdminLogEventActionParticipantJoin{}

	case crc_channelAdminLogEventActionParticipantLeave:
		print("channelAdminLogEventActionParticipantLeave")
		r = TL_channelAdminLogEventActionParticipantLeave{}

	case crc_channelAdminLogEventActionParticipantInvite:
		print("channelAdminLogEventActionParticipantInvite")
		r = TL_channelAdminLogEventActionParticipantInvite{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleBan:
		print("channelAdminLogEventActionParticipantToggleBan")
		r = TL_channelAdminLogEventActionParticipantToggleBan{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleAdmin:
		print("channelAdminLogEventActionParticipantToggleAdmin")
		r = TL_channelAdminLogEventActionParticipantToggleAdmin{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEvent:
		print("channelAdminLogEvent")
		r = TL_channelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channels_adminLogResults:
		print("channels_adminLogResults")
		r = TL_channels_adminLogResults{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelAdminLogEventsFilter:
		print("channelAdminLogEventsFilter")
		r = TL_channelAdminLogEventsFilter{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_topPeerCategoryPhoneCalls:
		print("topPeerCategoryPhoneCalls")
		r = TL_topPeerCategoryPhoneCalls{}

	case crc_pageBlockAudio:
		print("pageBlockAudio")
		r = TL_pageBlockAudio{
			m.Long(),
			m.Object(),
		}

	case crc_popularContact:
		print("popularContact")
		r = TL_popularContact{
			m.Long(),
			m.Int(),
		}

	case crc_messageActionScreenshotTaken:
		print("messageActionScreenshotTaken")
		r = TL_messageActionScreenshotTaken{}

	case crc_messages_favedStickersNotModified:
		print("messages_favedStickersNotModified")
		r = TL_messages_favedStickersNotModified{}

	case crc_messages_favedStickers:
		print("messages_favedStickers")
		r = TL_messages_favedStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updateFavedStickers:
		print("updateFavedStickers")
		r = TL_updateFavedStickers{}

	case crc_updateChannelReadMessagesContents:
		print("updateChannelReadMessagesContents")
		r = TL_updateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case crc_inputMessagesFilterMyMentions:
		print("inputMessagesFilterMyMentions")
		r = TL_inputMessagesFilterMyMentions{}

	case crc_updateContactsReset:
		print("updateContactsReset")
		r = TL_updateContactsReset{}

	case crc_channelAdminLogEventActionChangeStickerSet:
		print("channelAdminLogEventActionChangeStickerSet")
		r = TL_channelAdminLogEventActionChangeStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_messageActionCustomAction:
		print("messageActionCustomAction")
		r = TL_messageActionCustomAction{
			m.String(),
		}

	case crc_inputPaymentCredentialsApplePay:
		print("inputPaymentCredentialsApplePay")
		r = TL_inputPaymentCredentialsApplePay{
			m.Object(),
		}

	case crc_inputPaymentCredentialsAndroidPay:
		print("inputPaymentCredentialsAndroidPay")
		r = TL_inputPaymentCredentialsAndroidPay{
			m.Object(),
			m.String(),
		}

	case crc_inputMessagesFilterGeo:
		print("inputMessagesFilterGeo")
		r = TL_inputMessagesFilterGeo{}

	case crc_inputMessagesFilterContacts:
		print("inputMessagesFilterContacts")
		r = TL_inputMessagesFilterContacts{}

	case crc_updateChannelAvailableMessages:
		print("updateChannelAvailableMessages")
		r = TL_updateChannelAvailableMessages{
			m.Int(),
			m.Int(),
		}

	case crc_channelAdminLogEventActionTogglePreHistoryHidden:
		print("channelAdminLogEventActionTogglePreHistoryHidden")
		r = TL_channelAdminLogEventActionTogglePreHistoryHidden{
			m.Object(),
		}

	case crc_inputMediaGeoLive:
		print("inputMediaGeoLive")
		r = TL_inputMediaGeoLive{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messageMediaGeoLive:
		print("messageMediaGeoLive")
		r = TL_messageMediaGeoLive{
			m.Object(),
			m.Int(),
		}

	case crc_recentMeUrlUnknown:
		print("recentMeUrlUnknown")
		r = TL_recentMeUrlUnknown{
			m.String(),
		}

	case crc_recentMeUrlUser:
		print("recentMeUrlUser")
		r = TL_recentMeUrlUser{
			m.String(),
			m.Int(),
		}

	case crc_recentMeUrlChat:
		print("recentMeUrlChat")
		r = TL_recentMeUrlChat{
			m.String(),
			m.Int(),
		}

	case crc_recentMeUrlChatInvite:
		print("recentMeUrlChatInvite")
		r = TL_recentMeUrlChatInvite{
			m.String(),
			m.Object(),
		}

	case crc_recentMeUrlStickerSet:
		print("recentMeUrlStickerSet")
		r = TL_recentMeUrlStickerSet{
			m.String(),
			m.Object(),
		}

	case crc_help_recentMeUrls:
		print("help_recentMeUrls")
		r = TL_help_recentMeUrls{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channels_channelParticipantsNotModified:
		print("channels_channelParticipantsNotModified")
		r = TL_channels_channelParticipantsNotModified{}

	case crc_messages_messagesNotModified:
		print("messages_messagesNotModified")
		r = TL_messages_messagesNotModified{
			m.Int(),
		}

	case crc_inputSingleMedia:
		print("inputSingleMedia")
		r = TL_inputSingleMedia{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
			m.Object(),
		}

	case crc_webAuthorization:
		print("webAuthorization")
		r = TL_webAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_account_webAuthorizations:
		print("account_webAuthorizations")
		r = TL_account_webAuthorizations{
			m.Vector(),
			m.Vector(),
		}

	case crc_inputMessageID:
		print("inputMessageID")
		r = TL_inputMessageID{
			m.Int(),
		}

	case crc_inputMessageReplyTo:
		print("inputMessageReplyTo")
		r = TL_inputMessageReplyTo{
			m.Int(),
		}

	case crc_inputMessagePinned:
		print("inputMessagePinned")
		r = TL_inputMessagePinned{}

	case crc_messageEntityPhone:
		print("messageEntityPhone")
		r = TL_messageEntityPhone{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCashtag:
		print("messageEntityCashtag")
		r = TL_messageEntityCashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageActionBotAllowed:
		print("messageActionBotAllowed")
		r = TL_messageActionBotAllowed{
			m.String(),
		}

	case crc_inputDialogPeer:
		print("inputDialogPeer")
		r = TL_inputDialogPeer{
			m.Object(),
		}

	case crc_dialogPeer:
		print("dialogPeer")
		r = TL_dialogPeer{
			m.Object(),
		}

	case crc_messages_foundStickerSetsNotModified:
		print("messages_foundStickerSetsNotModified")
		r = TL_messages_foundStickerSetsNotModified{}

	case crc_messages_foundStickerSets:
		print("messages_foundStickerSets")
		r = TL_messages_foundStickerSets{
			m.Int(),
			m.Vector(),
		}

	case crc_fileHash:
		print("fileHash")
		r = TL_fileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_webDocumentNoProxy:
		print("webDocumentNoProxy")
		r = TL_webDocumentNoProxy{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_inputClientProxy:
		print("inputClientProxy")
		r = TL_inputClientProxy{
			m.String(),
			m.Int(),
		}

	case crc_help_proxyDataEmpty:
		print("help_proxyDataEmpty")
		r = TL_help_proxyDataEmpty{
			m.Int(),
		}

	case crc_help_proxyDataPromo:
		print("help_proxyDataPromo")
		r = TL_help_proxyDataPromo{
			m.Int(),
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_help_termsOfServiceUpdateEmpty:
		print("help_termsOfServiceUpdateEmpty")
		r = TL_help_termsOfServiceUpdateEmpty{
			m.Int(),
		}

	case crc_help_termsOfServiceUpdate:
		print("help_termsOfServiceUpdate")
		r = TL_help_termsOfServiceUpdate{
			m.Int(),
			m.Object(),
		}

	case crc_inputSecureFileUploaded:
		print("inputSecureFileUploaded")
		r = TL_inputSecureFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_inputSecureFile:
		print("inputSecureFile")
		r = TL_inputSecureFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputSecureFileLocation:
		print("inputSecureFileLocation")
		r = TL_inputSecureFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_secureFileEmpty:
		print("secureFileEmpty")
		r = TL_secureFileEmpty{}

	case crc_secureFile:
		print("secureFile")
		r = TL_secureFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_secureData:
		print("secureData")
		r = TL_secureData{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_securePlainPhone:
		print("securePlainPhone")
		r = TL_securePlainPhone{
			m.String(),
		}

	case crc_securePlainEmail:
		print("securePlainEmail")
		r = TL_securePlainEmail{
			m.String(),
		}

	case crc_secureValueTypePersonalDetails:
		print("secureValueTypePersonalDetails")
		r = TL_secureValueTypePersonalDetails{}

	case crc_secureValueTypePassport:
		print("secureValueTypePassport")
		r = TL_secureValueTypePassport{}

	case crc_secureValueTypeDriverLicense:
		print("secureValueTypeDriverLicense")
		r = TL_secureValueTypeDriverLicense{}

	case crc_secureValueTypeIdentityCard:
		print("secureValueTypeIdentityCard")
		r = TL_secureValueTypeIdentityCard{}

	case crc_secureValueTypeInternalPassport:
		print("secureValueTypeInternalPassport")
		r = TL_secureValueTypeInternalPassport{}

	case crc_secureValueTypeAddress:
		print("secureValueTypeAddress")
		r = TL_secureValueTypeAddress{}

	case crc_secureValueTypeUtilityBill:
		print("secureValueTypeUtilityBill")
		r = TL_secureValueTypeUtilityBill{}

	case crc_secureValueTypeBankStatement:
		print("secureValueTypeBankStatement")
		r = TL_secureValueTypeBankStatement{}

	case crc_secureValueTypeRentalAgreement:
		print("secureValueTypeRentalAgreement")
		r = TL_secureValueTypeRentalAgreement{}

	case crc_secureValueTypePassportRegistration:
		print("secureValueTypePassportRegistration")
		r = TL_secureValueTypePassportRegistration{}

	case crc_secureValueTypeTemporaryRegistration:
		print("secureValueTypeTemporaryRegistration")
		r = TL_secureValueTypeTemporaryRegistration{}

	case crc_secureValueTypePhone:
		print("secureValueTypePhone")
		r = TL_secureValueTypePhone{}

	case crc_secureValueTypeEmail:
		print("secureValueTypeEmail")
		r = TL_secureValueTypeEmail{}

	case crc_secureValue:
		print("secureValue")
		r = TL_secureValue{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.StringBytes(),
		}

	case crc_inputSecureValue:
		print("inputSecureValue")
		r = TL_inputSecureValue{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_secureValueHash:
		print("secureValueHash")
		r = TL_secureValueHash{
			m.Object(),
			m.StringBytes(),
		}

	case crc_secureValueErrorData:
		print("secureValueErrorData")
		r = TL_secureValueErrorData{
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case crc_secureValueErrorFrontSide:
		print("secureValueErrorFrontSide")
		r = TL_secureValueErrorFrontSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorReverseSide:
		print("secureValueErrorReverseSide")
		r = TL_secureValueErrorReverseSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorSelfie:
		print("secureValueErrorSelfie")
		r = TL_secureValueErrorSelfie{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorFile:
		print("secureValueErrorFile")
		r = TL_secureValueErrorFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorFiles:
		print("secureValueErrorFiles")
		r = TL_secureValueErrorFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case crc_secureCredentialsEncrypted:
		print("secureCredentialsEncrypted")
		r = TL_secureCredentialsEncrypted{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_account_authorizationForm:
		print("account_authorizationForm")
		r = TL_account_authorizationForm{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_account_sentEmailCode:
		print("account_sentEmailCode")
		r = TL_account_sentEmailCode{
			m.String(),
			m.Int(),
		}

	case crc_messageActionSecureValuesSentMe:
		print("messageActionSecureValuesSentMe")
		r = TL_messageActionSecureValuesSentMe{
			m.Vector(),
			m.Object(),
		}

	case crc_messageActionSecureValuesSent:
		print("messageActionSecureValuesSent")
		r = TL_messageActionSecureValuesSent{
			m.Vector(),
		}

	case crc_help_deepLinkInfoEmpty:
		print("help_deepLinkInfoEmpty")
		r = TL_help_deepLinkInfoEmpty{}

	case crc_help_deepLinkInfo:
		print("help_deepLinkInfo")
		r = TL_help_deepLinkInfo{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_savedPhoneContact:
		print("savedPhoneContact")
		r = TL_savedPhoneContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_account_takeout:
		print("account_takeout")
		r = TL_account_takeout{
			m.Long(),
		}

	case crc_inputTakeoutFileLocation:
		print("inputTakeoutFileLocation")
		r = TL_inputTakeoutFileLocation{}

	case crc_updateDialogUnreadMark:
		print("updateDialogUnreadMark")
		r = TL_updateDialogUnreadMark{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_dialogsNotModified:
		print("messages_dialogsNotModified")
		r = TL_messages_dialogsNotModified{
			m.Int(),
		}

	case crc_inputWebFileGeoPointLocation:
		print("inputWebFileGeoPointLocation")
		r = TL_inputWebFileGeoPointLocation{
			m.Object(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_contacts_topPeersDisabled:
		print("contacts_topPeersDisabled")
		r = TL_contacts_topPeersDisabled{}

	case crc_inputReportReasonCopyright:
		print("inputReportReasonCopyright")
		r = TL_inputReportReasonCopyright{}

	case crc_passwordKdfAlgoUnknown:
		print("passwordKdfAlgoUnknown")
		r = TL_passwordKdfAlgoUnknown{}

	case crc_securePasswordKdfAlgoUnknown:
		print("securePasswordKdfAlgoUnknown")
		r = TL_securePasswordKdfAlgoUnknown{}

	case crc_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000:
		print("securePasswordKdfAlgoPBKDF2HMACSHA512iter100000")
		r = TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000{
			m.StringBytes(),
		}

	case crc_securePasswordKdfAlgoSHA512:
		print("securePasswordKdfAlgoSHA512")
		r = TL_securePasswordKdfAlgoSHA512{
			m.StringBytes(),
		}

	case crc_secureSecretSettings:
		print("secureSecretSettings")
		r = TL_secureSecretSettings{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow:
		print("passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow")
		r = TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{
			m.StringBytes(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputCheckPasswordEmpty:
		print("inputCheckPasswordEmpty")
		r = TL_inputCheckPasswordEmpty{}

	case crc_inputCheckPasswordSRP:
		print("inputCheckPasswordSRP")
		r = TL_inputCheckPasswordSRP{
			m.Long(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_secureValueError:
		print("secureValueError")
		r = TL_secureValueError{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorTranslationFile:
		print("secureValueErrorTranslationFile")
		r = TL_secureValueErrorTranslationFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crc_secureValueErrorTranslationFiles:
		print("secureValueErrorTranslationFiles")
		r = TL_secureValueErrorTranslationFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case crc_secureRequiredType:
		print("secureRequiredType")
		r = TL_secureRequiredType{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_secureRequiredTypeOneOf:
		print("secureRequiredTypeOneOf")
		r = TL_secureRequiredTypeOneOf{
			m.Vector(),
		}

	case crc_help_passportConfigNotModified:
		print("help_passportConfigNotModified")
		r = TL_help_passportConfigNotModified{}

	case crc_help_passportConfig:
		print("help_passportConfig")
		r = TL_help_passportConfig{
			m.Int(),
			m.Object(),
		}

	case crc_inputAppEvent:
		print("inputAppEvent")
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.Object(),
		}

	case crc_jsonObjectValue:
		print("jsonObjectValue")
		r = TL_jsonObjectValue{
			m.String(),
			m.Object(),
		}

	case crc_jsonNull:
		print("jsonNull")
		r = TL_jsonNull{}

	case crc_jsonBool:
		print("jsonBool")
		r = TL_jsonBool{
			m.Object(),
		}

	case crc_jsonNumber:
		print("jsonNumber")
		r = TL_jsonNumber{
			m.Double(),
		}

	case crc_jsonString:
		print("jsonString")
		r = TL_jsonString{
			m.String(),
		}

	case crc_jsonArray:
		print("jsonArray")
		r = TL_jsonArray{
			m.Vector(),
		}

	case crc_jsonObject:
		print("jsonObject")
		r = TL_jsonObject{
			m.Vector(),
		}

	case crc_updateUserPinnedMessage:
		print("updateUserPinnedMessage")
		r = TL_updateUserPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_updateChatPinnedMessage:
		print("updateChatPinnedMessage")
		r = TL_updateChatPinnedMessage{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputNotifyBroadcasts:
		print("inputNotifyBroadcasts")
		r = TL_inputNotifyBroadcasts{}

	case crc_notifyBroadcasts:
		print("notifyBroadcasts")
		r = TL_notifyBroadcasts{}

	case crc_textSubscript:
		print("textSubscript")
		r = TL_textSubscript{
			m.Object(),
		}

	case crc_textSuperscript:
		print("textSuperscript")
		r = TL_textSuperscript{
			m.Object(),
		}

	case crc_textMarked:
		print("textMarked")
		r = TL_textMarked{
			m.Object(),
		}

	case crc_textPhone:
		print("textPhone")
		r = TL_textPhone{
			m.Object(),
			m.String(),
		}

	case crc_textImage:
		print("textImage")
		r = TL_textImage{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_pageBlockKicker:
		print("pageBlockKicker")
		r = TL_pageBlockKicker{
			m.Object(),
		}

	case crc_pageTableCell:
		print("pageTableCell")
		r = TL_pageTableCell{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_pageTableRow:
		print("pageTableRow")
		r = TL_pageTableRow{
			m.Vector(),
		}

	case crc_pageBlockTable:
		print("pageBlockTable")
		r = TL_pageBlockTable{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_pageCaption:
		print("pageCaption")
		r = TL_pageCaption{
			m.Object(),
			m.Object(),
		}

	case crc_pageListItemText:
		print("pageListItemText")
		r = TL_pageListItemText{
			m.Object(),
		}

	case crc_pageListItemBlocks:
		print("pageListItemBlocks")
		r = TL_pageListItemBlocks{
			m.Vector(),
		}

	case crc_pageListOrderedItemText:
		print("pageListOrderedItemText")
		r = TL_pageListOrderedItemText{
			m.String(),
			m.Object(),
		}

	case crc_pageListOrderedItemBlocks:
		print("pageListOrderedItemBlocks")
		r = TL_pageListOrderedItemBlocks{
			m.String(),
			m.Vector(),
		}

	case crc_pageBlockOrderedList:
		print("pageBlockOrderedList")
		r = TL_pageBlockOrderedList{
			m.Vector(),
		}

	case crc_pageBlockDetails:
		print("pageBlockDetails")
		r = TL_pageBlockDetails{
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Object(),
		}

	case crc_pageRelatedArticle:
		print("pageRelatedArticle")
		r = TL_pageRelatedArticle{
			m.Object(),
			m.String(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockRelatedArticles:
		print("pageBlockRelatedArticles")
		r = TL_pageBlockRelatedArticles{
			m.Object(),
			m.Vector(),
		}

	case crc_pageBlockMap:
		print("pageBlockMap")
		r = TL_pageBlockMap{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_page:
		print("page")
		r = TL_page{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyPhoneP2P:
		print("inputPrivacyKeyPhoneP2P")
		r = TL_inputPrivacyKeyPhoneP2P{}

	case crc_privacyKeyPhoneP2P:
		print("privacyKeyPhoneP2P")
		r = TL_privacyKeyPhoneP2P{}

	case crc_textAnchor:
		print("textAnchor")
		r = TL_textAnchor{
			m.Object(),
			m.String(),
		}

	case crc_help_supportName:
		print("help_supportName")
		r = TL_help_supportName{
			m.String(),
		}

	case crc_help_userInfoEmpty:
		print("help_userInfoEmpty")
		r = TL_help_userInfoEmpty{}

	case crc_help_userInfo:
		print("help_userInfo")
		r = TL_help_userInfo{
			m.String(),
			m.Vector(),
			m.String(),
			m.Int(),
		}

	case crc_messageActionContactSignUp:
		print("messageActionContactSignUp")
		r = TL_messageActionContactSignUp{}

	case crc_updateMessagePoll:
		print("updateMessagePoll")
		r = TL_updateMessagePoll{
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_pollAnswer:
		print("pollAnswer")
		r = TL_pollAnswer{
			m.String(),
			m.StringBytes(),
		}

	case crc_poll:
		print("poll")
		r = TL_poll{
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crc_pollAnswerVoters:
		print("pollAnswerVoters")
		r = TL_pollAnswerVoters{
			m.Object(),
			m.Object(),
			m.Object(),
			m.StringBytes(),
			m.Int(),
		}

	case crc_pollResults:
		print("pollResults")
		r = TL_pollResults{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaPoll:
		print("inputMediaPoll")
		r = TL_inputMediaPoll{
			m.Object(),
		}

	case crc_messageMediaPoll:
		print("messageMediaPoll")
		r = TL_messageMediaPoll{
			m.Object(),
			m.Object(),
		}

	case crc_chatOnlines:
		print("chatOnlines")
		r = TL_chatOnlines{
			m.Int(),
		}

	case crc_statsURL:
		print("statsURL")
		r = TL_statsURL{
			m.String(),
		}

	case crc_photoStrippedSize:
		print("photoStrippedSize")
		r = TL_photoStrippedSize{
			m.String(),
			m.StringBytes(),
		}

	case crc_chatAdminRights:
		print("chatAdminRights")
		r = TL_chatAdminRights{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_chatBannedRights:
		print("chatBannedRights")
		r = TL_chatBannedRights{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_updateChatDefaultBannedRights:
		print("updateChatDefaultBannedRights")
		r = TL_updateChatDefaultBannedRights{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_inputWallPaper:
		print("inputWallPaper")
		r = TL_inputWallPaper{
			m.Long(),
			m.Long(),
		}

	case crc_inputWallPaperSlug:
		print("inputWallPaperSlug")
		r = TL_inputWallPaperSlug{
			m.String(),
		}

	case crc_channelParticipantsContacts:
		print("channelParticipantsContacts")
		r = TL_channelParticipantsContacts{
			m.String(),
		}

	case crc_channelAdminLogEventActionDefaultBannedRights:
		print("channelAdminLogEventActionDefaultBannedRights")
		r = TL_channelAdminLogEventActionDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionStopPoll:
		print("channelAdminLogEventActionStopPoll")
		r = TL_channelAdminLogEventActionStopPoll{
			m.Object(),
		}

	case crc_account_wallPapersNotModified:
		print("account_wallPapersNotModified")
		r = TL_account_wallPapersNotModified{}

	case crc_account_wallPapers:
		print("account_wallPapers")
		r = TL_account_wallPapers{
			m.Int(),
			m.Vector(),
		}

	case crc_codeSettings:
		print("codeSettings")
		r = TL_codeSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_wallPaperSettings:
		print("wallPaperSettings")
		r = TL_wallPaperSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_autoDownloadSettings:
		print("autoDownloadSettings")
		r = TL_autoDownloadSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_account_autoDownloadSettings:
		print("account_autoDownloadSettings")
		r = TL_account_autoDownloadSettings{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_emojiKeyword:
		print("emojiKeyword")
		r = TL_emojiKeyword{
			m.String(),
			m.VectorString(),
		}

	case crc_emojiKeywordDeleted:
		print("emojiKeywordDeleted")
		r = TL_emojiKeywordDeleted{
			m.String(),
			m.VectorString(),
		}

	case crc_emojiKeywordsDifference:
		print("emojiKeywordsDifference")
		r = TL_emojiKeywordsDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_emojiURL:
		print("emojiURL")
		r = TL_emojiURL{
			m.String(),
		}

	case crc_emojiLanguage:
		print("emojiLanguage")
		r = TL_emojiLanguage{
			m.String(),
		}

	case crc_inputPrivacyKeyForwards:
		print("inputPrivacyKeyForwards")
		r = TL_inputPrivacyKeyForwards{}

	case crc_privacyKeyForwards:
		print("privacyKeyForwards")
		r = TL_privacyKeyForwards{}

	case crc_inputPrivacyKeyProfilePhoto:
		print("inputPrivacyKeyProfilePhoto")
		r = TL_inputPrivacyKeyProfilePhoto{}

	case crc_privacyKeyProfilePhoto:
		print("privacyKeyProfilePhoto")
		r = TL_privacyKeyProfilePhoto{}

	case crc_fileLocationToBeDeprecated:
		print("fileLocationToBeDeprecated")
		r = TL_fileLocationToBeDeprecated{
			m.Long(),
			m.Int(),
		}

	case crc_inputPhotoFileLocation:
		print("inputPhotoFileLocation")
		r = TL_inputPhotoFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case crc_inputPhotoLegacyFileLocation:
		print("inputPhotoLegacyFileLocation")
		r = TL_inputPhotoLegacyFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputPeerPhotoFileLocation:
		print("inputPeerPhotoFileLocation")
		r = TL_inputPeerPhotoFileLocation{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case crc_inputStickerSetThumb:
		print("inputStickerSetThumb")
		r = TL_inputStickerSetThumb{
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case crc_folder:
		print("folder")
		r = TL_folder{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crc_dialogFolder:
		print("dialogFolder")
		r = TL_dialogFolder{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputDialogPeerFolder:
		print("inputDialogPeerFolder")
		r = TL_inputDialogPeerFolder{
			m.Int(),
		}

	case crc_dialogPeerFolder:
		print("dialogPeerFolder")
		r = TL_dialogPeerFolder{
			m.Int(),
		}

	case crc_inputFolderPeer:
		print("inputFolderPeer")
		r = TL_inputFolderPeer{
			m.Object(),
			m.Int(),
		}

	case crc_folderPeer:
		print("folderPeer")
		r = TL_folderPeer{
			m.Object(),
			m.Int(),
		}

	case crc_updateFolderPeers:
		print("updateFolderPeers")
		r = TL_updateFolderPeers{
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_inputUserFromMessage:
		print("inputUserFromMessage")
		r = TL_inputUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputChannelFromMessage:
		print("inputChannelFromMessage")
		r = TL_inputChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputPeerUserFromMessage:
		print("inputPeerUserFromMessage")
		r = TL_inputPeerUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputPeerChannelFromMessage:
		print("inputPeerChannelFromMessage")
		r = TL_inputPeerChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputPrivacyKeyPhoneNumber:
		print("inputPrivacyKeyPhoneNumber")
		r = TL_inputPrivacyKeyPhoneNumber{}

	case crc_privacyKeyPhoneNumber:
		print("privacyKeyPhoneNumber")
		r = TL_privacyKeyPhoneNumber{}

	case crc_topPeerCategoryForwardUsers:
		print("topPeerCategoryForwardUsers")
		r = TL_topPeerCategoryForwardUsers{}

	case crc_topPeerCategoryForwardChats:
		print("topPeerCategoryForwardChats")
		r = TL_topPeerCategoryForwardChats{}

	case crc_channelAdminLogEventActionChangeLinkedChat:
		print("channelAdminLogEventActionChangeLinkedChat")
		r = TL_channelAdminLogEventActionChangeLinkedChat{
			m.Int(),
			m.Int(),
		}

	case crc_messages_searchCounter:
		print("messages_searchCounter")
		r = TL_messages_searchCounter{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_keyboardButtonUrlAuth:
		print("keyboardButtonUrlAuth")
		r = TL_keyboardButtonUrlAuth{
			m.Object(),
			m.String(),
			m.Object(),
			m.String(),
			m.Int(),
		}

	case crc_inputKeyboardButtonUrlAuth:
		print("inputKeyboardButtonUrlAuth")
		r = TL_inputKeyboardButtonUrlAuth{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_urlAuthResultRequest:
		print("urlAuthResultRequest")
		r = TL_urlAuthResultRequest{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crc_urlAuthResultAccepted:
		print("urlAuthResultAccepted")
		r = TL_urlAuthResultAccepted{
			m.String(),
		}

	case crc_urlAuthResultDefault:
		print("urlAuthResultDefault")
		r = TL_urlAuthResultDefault{}

	case crc_inputPrivacyValueAllowChatParticipants:
		print("inputPrivacyValueAllowChatParticipants")
		r = TL_inputPrivacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case crc_inputPrivacyValueDisallowChatParticipants:
		print("inputPrivacyValueDisallowChatParticipants")
		r = TL_inputPrivacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case crc_privacyValueAllowChatParticipants:
		print("privacyValueAllowChatParticipants")
		r = TL_privacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowChatParticipants:
		print("privacyValueDisallowChatParticipants")
		r = TL_privacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case crc_messageEntityUnderline:
		print("messageEntityUnderline")
		r = TL_messageEntityUnderline{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityStrike:
		print("messageEntityStrike")
		r = TL_messageEntityStrike{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBlockquote:
		print("messageEntityBlockquote")
		r = TL_messageEntityBlockquote{
			m.Int(),
			m.Int(),
		}

	case crc_updatePeerSettings:
		print("updatePeerSettings")
		r = TL_updatePeerSettings{
			m.Object(),
			m.Object(),
		}

	case crc_channelLocationEmpty:
		print("channelLocationEmpty")
		r = TL_channelLocationEmpty{}

	case crc_channelLocation:
		print("channelLocation")
		r = TL_channelLocation{
			m.Object(),
			m.String(),
		}

	case crc_peerLocated:
		print("peerLocated")
		r = TL_peerLocated{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updatePeerLocated:
		print("updatePeerLocated")
		r = TL_updatePeerLocated{
			m.Vector(),
		}

	case crc_channelAdminLogEventActionChangeLocation:
		print("channelAdminLogEventActionChangeLocation")
		r = TL_channelAdminLogEventActionChangeLocation{
			m.Object(),
			m.Object(),
		}

	case crc_inputReportReasonGeoIrrelevant:
		print("inputReportReasonGeoIrrelevant")
		r = TL_inputReportReasonGeoIrrelevant{}

	case crc_channelAdminLogEventActionToggleSlowMode:
		print("channelAdminLogEventActionToggleSlowMode")
		r = TL_channelAdminLogEventActionToggleSlowMode{
			m.Int(),
			m.Int(),
		}

	case crc_auth_authorizationSignUpRequired:
		print("auth_authorizationSignUpRequired")
		r = TL_auth_authorizationSignUpRequired{
			m.Object(),
			m.Object(),
		}

	case crc_payments_paymentVerificationNeeded:
		print("payments_paymentVerificationNeeded")
		r = TL_payments_paymentVerificationNeeded{
			m.String(),
		}

	case crc_inputStickerSetAnimatedEmoji:
		print("inputStickerSetAnimatedEmoji")
		r = TL_inputStickerSetAnimatedEmoji{}

	case crc_updateNewScheduledMessage:
		print("updateNewScheduledMessage")
		r = TL_updateNewScheduledMessage{
			m.Object(),
		}

	case crc_updateDeleteScheduledMessages:
		print("updateDeleteScheduledMessages")
		r = TL_updateDeleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_restrictionReason:
		print("restrictionReason")
		r = TL_restrictionReason{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputTheme:
		print("inputTheme")
		r = TL_inputTheme{
			m.Long(),
			m.Long(),
		}

	case crc_inputThemeSlug:
		print("inputThemeSlug")
		r = TL_inputThemeSlug{
			m.String(),
		}

	case crc_themeDocumentNotModified:
		print("themeDocumentNotModified")
		r = TL_themeDocumentNotModified{}

	case crc_theme:
		print("theme")
		r = TL_theme{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case crc_account_themesNotModified:
		print("account_themesNotModified")
		r = TL_account_themesNotModified{}

	case crc_account_themes:
		print("account_themes")
		r = TL_account_themes{
			m.Int(),
			m.Vector(),
		}

	case crc_updateTheme:
		print("updateTheme")
		r = TL_updateTheme{
			m.Object(),
		}

	case crc_inputPrivacyKeyAddedByPhone:
		print("inputPrivacyKeyAddedByPhone")
		r = TL_inputPrivacyKeyAddedByPhone{}

	case crc_privacyKeyAddedByPhone:
		print("privacyKeyAddedByPhone")
		r = TL_privacyKeyAddedByPhone{}

	case crc_invokeAfterMsg:
		print("invokeAfterMsg")
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crc_invokeAfterMsgs:
		print("invokeAfterMsgs")
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crc_auth_sendCode:
		print("auth_sendCode")
		r = TL_auth_sendCode{
			m.String(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crc_auth_signUp:
		print("auth_signUp")
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_signIn:
		print("auth_signIn")
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_logOut:
		print("auth_logOut")
		r = TL_auth_logOut{}

	case crc_auth_resetAuthorizations:
		print("auth_resetAuthorizations")
		r = TL_auth_resetAuthorizations{}

	case crc_auth_exportAuthorization:
		print("auth_exportAuthorization")
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case crc_auth_importAuthorization:
		print("auth_importAuthorization")
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_auth_bindTempAuthKey:
		print("auth_bindTempAuthKey")
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_account_registerDevice:
		print("account_registerDevice")
		r = TL_account_registerDevice{
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
			m.StringBytes(),
			m.VectorInt(),
		}

	case crc_account_unregisterDevice:
		print("account_unregisterDevice")
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
			m.VectorInt(),
		}

	case crc_account_updateNotifySettings:
		print("account_updateNotifySettings")
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_account_getNotifySettings:
		print("account_getNotifySettings")
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case crc_account_resetNotifySettings:
		print("account_resetNotifySettings")
		r = TL_account_resetNotifySettings{}

	case crc_account_updateProfile:
		print("account_updateProfile")
		r = TL_account_updateProfile{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_updateStatus:
		print("account_updateStatus")
		r = TL_account_updateStatus{
			m.Object(),
		}

	case crc_account_getWallPapers:
		print("account_getWallPapers")
		r = TL_account_getWallPapers{
			m.Int(),
		}

	case crc_account_reportPeer:
		print("account_reportPeer")
		r = TL_account_reportPeer{
			m.Object(),
			m.Object(),
		}

	case crc_users_getUsers:
		print("users_getUsers")
		r = TL_users_getUsers{
			m.Vector(),
		}

	case crc_users_getFullUser:
		print("users_getFullUser")
		r = TL_users_getFullUser{
			m.Object(),
		}

	case crc_contacts_getContactIDs:
		print("contacts_getContactIDs")
		r = TL_contacts_getContactIDs{
			m.Int(),
		}

	case crc_contacts_getStatuses:
		print("contacts_getStatuses")
		r = TL_contacts_getStatuses{}

	case crc_contacts_getContacts:
		print("contacts_getContacts")
		r = TL_contacts_getContacts{
			m.Int(),
		}

	case crc_contacts_importContacts:
		print("contacts_importContacts")
		r = TL_contacts_importContacts{
			m.Vector(),
		}

	case crc_contacts_deleteContacts:
		print("contacts_deleteContacts")
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case crc_contacts_deleteByPhones:
		print("contacts_deleteByPhones")
		r = TL_contacts_deleteByPhones{
			m.VectorString(),
		}

	case crc_contacts_block:
		print("contacts_block")
		r = TL_contacts_block{
			m.Object(),
		}

	case crc_contacts_unblock:
		print("contacts_unblock")
		r = TL_contacts_unblock{
			m.Object(),
		}

	case crc_contacts_getBlocked:
		print("contacts_getBlocked")
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_messages_getMessages:
		print("messages_getMessages")
		r = TL_messages_getMessages{
			m.Vector(),
		}

	case crc_messages_getDialogs:
		print("messages_getDialogs")
		r = TL_messages_getDialogs{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_getHistory:
		print("messages_getHistory")
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_search:
		print("messages_search")
		r = TL_messages_search{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_readHistory:
		print("messages_readHistory")
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteHistory:
		print("messages_deleteHistory")
		r = TL_messages_deleteHistory{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteMessages:
		print("messages_deleteMessages")
		r = TL_messages_deleteMessages{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crc_messages_receivedMessages:
		print("messages_receivedMessages")
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case crc_messages_setTyping:
		print("messages_setTyping")
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMessage:
		print("messages_sendMessage")
		r = TL_messages_sendMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMedia:
		print("messages_sendMedia")
		r = TL_messages_sendMedia{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_forwardMessages:
		print("messages_forwardMessages")
		r = TL_messages_forwardMessages{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.VectorInt(),
			m.VectorLong(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_reportSpam:
		print("messages_reportSpam")
		r = TL_messages_reportSpam{
			m.Object(),
		}

	case crc_messages_getPeerSettings:
		print("messages_getPeerSettings")
		r = TL_messages_getPeerSettings{
			m.Object(),
		}

	case crc_messages_report:
		print("messages_report")
		r = TL_messages_report{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crc_messages_getChats:
		print("messages_getChats")
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case crc_messages_getFullChat:
		print("messages_getFullChat")
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case crc_messages_editChatTitle:
		print("messages_editChatTitle")
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messages_editChatPhoto:
		print("messages_editChatPhoto")
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crc_messages_addChatUser:
		print("messages_addChatUser")
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteChatUser:
		print("messages_deleteChatUser")
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crc_messages_createChat:
		print("messages_createChat")
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case crc_updates_getState:
		print("updates_getState")
		r = TL_updates_getState{}

	case crc_updates_getDifference:
		print("updates_getDifference")
		r = TL_updates_getDifference{
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_updateProfilePhoto:
		print("photos_updateProfilePhoto")
		r = TL_photos_updateProfilePhoto{
			m.Object(),
		}

	case crc_photos_uploadProfilePhoto:
		print("photos_uploadProfilePhoto")
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
		}

	case crc_photos_deletePhotos:
		print("photos_deletePhotos")
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case crc_upload_saveFilePart:
		print("upload_saveFilePart")
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_upload_getFile:
		print("upload_getFile")
		r = TL_upload_getFile{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getConfig:
		print("help_getConfig")
		r = TL_help_getConfig{}

	case crc_help_getNearestDc:
		print("help_getNearestDc")
		r = TL_help_getNearestDc{}

	case crc_help_getAppUpdate:
		print("help_getAppUpdate")
		r = TL_help_getAppUpdate{
			m.String(),
		}

	case crc_help_getInviteText:
		print("help_getInviteText")
		r = TL_help_getInviteText{}

	case crc_photos_getUserPhotos:
		print("photos_getUserPhotos")
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crc_messages_getDhConfig:
		print("messages_getDhConfig")
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messages_requestEncryption:
		print("messages_requestEncryption")
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_acceptEncryption:
		print("messages_acceptEncryption")
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messages_discardEncryption:
		print("messages_discardEncryption")
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case crc_messages_setEncryptedTyping:
		print("messages_setEncryptedTyping")
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readEncryptedHistory:
		print("messages_readEncryptedHistory")
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_sendEncrypted:
		print("messages_sendEncrypted")
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_sendEncryptedFile:
		print("messages_sendEncryptedFile")
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_messages_sendEncryptedService:
		print("messages_sendEncryptedService")
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_receivedQueue:
		print("messages_receivedQueue")
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case crc_messages_reportEncryptedSpam:
		print("messages_reportEncryptedSpam")
		r = TL_messages_reportEncryptedSpam{
			m.Object(),
		}

	case crc_upload_saveBigFilePart:
		print("upload_saveBigFilePart")
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		print("initConnection")
		r = TL_initConnection{
			m.Object(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_help_getSupport:
		print("help_getSupport")
		r = TL_help_getSupport{}

	case crc_messages_readMessageContents:
		print("messages_readMessageContents")
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case crc_account_checkUsername:
		print("account_checkUsername")
		r = TL_account_checkUsername{
			m.String(),
		}

	case crc_account_updateUsername:
		print("account_updateUsername")
		r = TL_account_updateUsername{
			m.String(),
		}

	case crc_contacts_search:
		print("contacts_search")
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case crc_account_getPrivacy:
		print("account_getPrivacy")
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case crc_account_setPrivacy:
		print("account_setPrivacy")
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_account_deleteAccount:
		print("account_deleteAccount")
		r = TL_account_deleteAccount{
			m.String(),
		}

	case crc_account_getAccountTTL:
		print("account_getAccountTTL")
		r = TL_account_getAccountTTL{}

	case crc_account_setAccountTTL:
		print("account_setAccountTTL")
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case crc_invokeWithLayer:
		print("invokeWithLayer")
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_resolveUsername:
		print("contacts_resolveUsername")
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case crc_account_sendChangePhoneCode:
		print("account_sendChangePhoneCode")
		r = TL_account_sendChangePhoneCode{
			m.String(),
			m.Object(),
		}

	case crc_account_changePhone:
		print("account_changePhone")
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messages_getStickers:
		print("messages_getStickers")
		r = TL_messages_getStickers{
			m.String(),
			m.Int(),
		}

	case crc_messages_getAllStickers:
		print("messages_getAllStickers")
		r = TL_messages_getAllStickers{
			m.Int(),
		}

	case crc_account_updateDeviceLocked:
		print("account_updateDeviceLocked")
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	case crc_auth_importBotAuthorization:
		print("auth_importBotAuthorization")
		r = TL_auth_importBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_messages_getWebPagePreview:
		print("messages_getWebPagePreview")
		r = TL_messages_getWebPagePreview{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_account_getAuthorizations:
		print("account_getAuthorizations")
		r = TL_account_getAuthorizations{}

	case crc_account_resetAuthorization:
		print("account_resetAuthorization")
		r = TL_account_resetAuthorization{
			m.Long(),
		}

	case crc_account_getPassword:
		print("account_getPassword")
		r = TL_account_getPassword{}

	case crc_account_getPasswordSettings:
		print("account_getPasswordSettings")
		r = TL_account_getPasswordSettings{
			m.Object(),
		}

	case crc_account_updatePasswordSettings:
		print("account_updatePasswordSettings")
		r = TL_account_updatePasswordSettings{
			m.Object(),
			m.Object(),
		}

	case crc_auth_checkPassword:
		print("auth_checkPassword")
		r = TL_auth_checkPassword{
			m.Object(),
		}

	case crc_auth_requestPasswordRecovery:
		print("auth_requestPasswordRecovery")
		r = TL_auth_requestPasswordRecovery{}

	case crc_auth_recoverPassword:
		print("auth_recoverPassword")
		r = TL_auth_recoverPassword{
			m.String(),
		}

	case crc_invokeWithoutUpdates:
		print("invokeWithoutUpdates")
		r = TL_invokeWithoutUpdates{
			m.Object(),
		}

	case crc_messages_exportChatInvite:
		print("messages_exportChatInvite")
		r = TL_messages_exportChatInvite{
			m.Object(),
		}

	case crc_messages_checkChatInvite:
		print("messages_checkChatInvite")
		r = TL_messages_checkChatInvite{
			m.String(),
		}

	case crc_messages_importChatInvite:
		print("messages_importChatInvite")
		r = TL_messages_importChatInvite{
			m.String(),
		}

	case crc_messages_getStickerSet:
		print("messages_getStickerSet")
		r = TL_messages_getStickerSet{
			m.Object(),
		}

	case crc_messages_installStickerSet:
		print("messages_installStickerSet")
		r = TL_messages_installStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_uninstallStickerSet:
		print("messages_uninstallStickerSet")
		r = TL_messages_uninstallStickerSet{
			m.Object(),
		}

	case crc_messages_startBot:
		print("messages_startBot")
		r = TL_messages_startBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case crc_help_getAppChangelog:
		print("help_getAppChangelog")
		r = TL_help_getAppChangelog{
			m.String(),
		}

	case crc_messages_getMessagesViews:
		print("messages_getMessagesViews")
		r = TL_messages_getMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crc_channels_readHistory:
		print("channels_readHistory")
		r = TL_channels_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_channels_deleteMessages:
		print("channels_deleteMessages")
		r = TL_channels_deleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_deleteUserHistory:
		print("channels_deleteUserHistory")
		r = TL_channels_deleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case crc_channels_reportSpam:
		print("channels_reportSpam")
		r = TL_channels_reportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getMessages:
		print("channels_getMessages")
		r = TL_channels_getMessages{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_getParticipants:
		print("channels_getParticipants")
		r = TL_channels_getParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_getParticipant:
		print("channels_getParticipant")
		r = TL_channels_getParticipant{
			m.Object(),
			m.Object(),
		}

	case crc_channels_getChannels:
		print("channels_getChannels")
		r = TL_channels_getChannels{
			m.Vector(),
		}

	case crc_channels_getFullChannel:
		print("channels_getFullChannel")
		r = TL_channels_getFullChannel{
			m.Object(),
		}

	case crc_channels_createChannel:
		print("channels_createChannel")
		r = TL_channels_createChannel{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_editAdmin:
		print("channels_editAdmin")
		r = TL_channels_editAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crc_channels_editTitle:
		print("channels_editTitle")
		r = TL_channels_editTitle{
			m.Object(),
			m.String(),
		}

	case crc_channels_editPhoto:
		print("channels_editPhoto")
		r = TL_channels_editPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channels_checkUsername:
		print("channels_checkUsername")
		r = TL_channels_checkUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_updateUsername:
		print("channels_updateUsername")
		r = TL_channels_updateUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_joinChannel:
		print("channels_joinChannel")
		r = TL_channels_joinChannel{
			m.Object(),
		}

	case crc_channels_leaveChannel:
		print("channels_leaveChannel")
		r = TL_channels_leaveChannel{
			m.Object(),
		}

	case crc_channels_inviteToChannel:
		print("channels_inviteToChannel")
		r = TL_channels_inviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_deleteChannel:
		print("channels_deleteChannel")
		r = TL_channels_deleteChannel{
			m.Object(),
		}

	case crc_updates_getChannelDifference:
		print("updates_getChannelDifference")
		r = TL_updates_getChannelDifference{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_editChatAdmin:
		print("messages_editChatAdmin")
		r = TL_messages_editChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_migrateChat:
		print("messages_migrateChat")
		r = TL_messages_migrateChat{
			m.Int(),
		}

	case crc_messages_searchGlobal:
		print("messages_searchGlobal")
		r = TL_messages_searchGlobal{
			m.Object(),
			m.Object(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_reorderStickerSets:
		print("messages_reorderStickerSets")
		r = TL_messages_reorderStickerSets{
			m.Object(),
			m.Object(),
			m.VectorLong(),
		}

	case crc_messages_getDocumentByHash:
		print("messages_getDocumentByHash")
		r = TL_messages_getDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crc_messages_searchGifs:
		print("messages_searchGifs")
		r = TL_messages_searchGifs{
			m.String(),
			m.Int(),
		}

	case crc_messages_getSavedGifs:
		print("messages_getSavedGifs")
		r = TL_messages_getSavedGifs{
			m.Int(),
		}

	case crc_messages_saveGif:
		print("messages_saveGif")
		r = TL_messages_saveGif{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getInlineBotResults:
		print("messages_getInlineBotResults")
		r = TL_messages_getInlineBotResults{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_messages_setInlineBotResults:
		print("messages_setInlineBotResults")
		r = TL_messages_setInlineBotResults{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Vector(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendInlineBotResult:
		print("messages_sendInlineBotResult")
		r = TL_messages_sendInlineBotResult{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Object(),
		}

	case crc_channels_exportMessageLink:
		print("channels_exportMessageLink")
		r = TL_channels_exportMessageLink{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_channels_toggleSignatures:
		print("channels_toggleSignatures")
		r = TL_channels_toggleSignatures{
			m.Object(),
			m.Object(),
		}

	case crc_auth_resendCode:
		print("auth_resendCode")
		r = TL_auth_resendCode{
			m.String(),
			m.String(),
		}

	case crc_auth_cancelCode:
		print("auth_cancelCode")
		r = TL_auth_cancelCode{
			m.String(),
			m.String(),
		}

	case crc_messages_getMessageEditData:
		print("messages_getMessageEditData")
		r = TL_messages_getMessageEditData{
			m.Object(),
			m.Int(),
		}

	case crc_messages_editMessage:
		print("messages_editMessage")
		r = TL_messages_editMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_editInlineBotMessage:
		print("messages_editInlineBotMessage")
		r = TL_messages_editInlineBotMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_getBotCallbackAnswer:
		print("messages_getBotCallbackAnswer")
		r = TL_messages_getBotCallbackAnswer{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_setBotCallbackAnswer:
		print("messages_setBotCallbackAnswer")
		r = TL_messages_setBotCallbackAnswer{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_contacts_getTopPeers:
		print("contacts_getTopPeers")
		r = TL_contacts_getTopPeers{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_contacts_resetTopPeerRating:
		print("contacts_resetTopPeerRating")
		r = TL_contacts_resetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getPeerDialogs:
		print("messages_getPeerDialogs")
		r = TL_messages_getPeerDialogs{
			m.Vector(),
		}

	case crc_messages_saveDraft:
		print("messages_saveDraft")
		r = TL_messages_saveDraft{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_messages_getAllDrafts:
		print("messages_getAllDrafts")
		r = TL_messages_getAllDrafts{}

	case crc_messages_getFeaturedStickers:
		print("messages_getFeaturedStickers")
		r = TL_messages_getFeaturedStickers{
			m.Int(),
		}

	case crc_messages_readFeaturedStickers:
		print("messages_readFeaturedStickers")
		r = TL_messages_readFeaturedStickers{
			m.VectorLong(),
		}

	case crc_messages_getRecentStickers:
		print("messages_getRecentStickers")
		r = TL_messages_getRecentStickers{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_saveRecentSticker:
		print("messages_saveRecentSticker")
		r = TL_messages_saveRecentSticker{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_clearRecentStickers:
		print("messages_clearRecentStickers")
		r = TL_messages_clearRecentStickers{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getArchivedStickers:
		print("messages_getArchivedStickers")
		r = TL_messages_getArchivedStickers{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case crc_account_sendConfirmPhoneCode:
		print("account_sendConfirmPhoneCode")
		r = TL_account_sendConfirmPhoneCode{
			m.String(),
			m.Object(),
		}

	case crc_account_confirmPhone:
		print("account_confirmPhone")
		r = TL_account_confirmPhone{
			m.String(),
			m.String(),
		}

	case crc_channels_getAdminedPublicChannels:
		print("channels_getAdminedPublicChannels")
		r = TL_channels_getAdminedPublicChannels{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_getMaskStickers:
		print("messages_getMaskStickers")
		r = TL_messages_getMaskStickers{
			m.Int(),
		}

	case crc_messages_getAttachedStickers:
		print("messages_getAttachedStickers")
		r = TL_messages_getAttachedStickers{
			m.Object(),
		}

	case crc_auth_dropTempAuthKeys:
		print("auth_dropTempAuthKeys")
		r = TL_auth_dropTempAuthKeys{
			m.VectorLong(),
		}

	case crc_messages_setGameScore:
		print("messages_setGameScore")
		r = TL_messages_setGameScore{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_setInlineGameScore:
		print("messages_setInlineGameScore")
		r = TL_messages_setInlineGameScore{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getGameHighScores:
		print("messages_getGameHighScores")
		r = TL_messages_getGameHighScores{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_getInlineGameHighScores:
		print("messages_getInlineGameHighScores")
		r = TL_messages_getInlineGameHighScores{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getCommonChats:
		print("messages_getCommonChats")
		r = TL_messages_getCommonChats{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_getAllChats:
		print("messages_getAllChats")
		r = TL_messages_getAllChats{
			m.VectorInt(),
		}

	case crc_help_setBotUpdatesStatus:
		print("help_setBotUpdatesStatus")
		r = TL_help_setBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case crc_messages_getWebPage:
		print("messages_getWebPage")
		r = TL_messages_getWebPage{
			m.String(),
			m.Int(),
		}

	case crc_messages_toggleDialogPin:
		print("messages_toggleDialogPin")
		r = TL_messages_toggleDialogPin{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_reorderPinnedDialogs:
		print("messages_reorderPinnedDialogs")
		r = TL_messages_reorderPinnedDialogs{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_messages_getPinnedDialogs:
		print("messages_getPinnedDialogs")
		r = TL_messages_getPinnedDialogs{
			m.Int(),
		}

	case crc_bots_sendCustomRequest:
		print("bots_sendCustomRequest")
		r = TL_bots_sendCustomRequest{
			m.String(),
			m.Object(),
		}

	case crc_bots_answerWebhookJSONQuery:
		print("bots_answerWebhookJSONQuery")
		r = TL_bots_answerWebhookJSONQuery{
			m.Long(),
			m.Object(),
		}

	case crc_upload_getWebFile:
		print("upload_getWebFile")
		r = TL_upload_getWebFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_payments_getPaymentForm:
		print("payments_getPaymentForm")
		r = TL_payments_getPaymentForm{
			m.Int(),
		}

	case crc_payments_getPaymentReceipt:
		print("payments_getPaymentReceipt")
		r = TL_payments_getPaymentReceipt{
			m.Int(),
		}

	case crc_payments_validateRequestedInfo:
		print("payments_validateRequestedInfo")
		r = TL_payments_validateRequestedInfo{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_payments_sendPaymentForm:
		print("payments_sendPaymentForm")
		r = TL_payments_sendPaymentForm{
			m.Object(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_getTmpPassword:
		print("account_getTmpPassword")
		r = TL_account_getTmpPassword{
			m.Object(),
			m.Int(),
		}

	case crc_payments_getSavedInfo:
		print("payments_getSavedInfo")
		r = TL_payments_getSavedInfo{}

	case crc_payments_clearSavedInfo:
		print("payments_clearSavedInfo")
		r = TL_payments_clearSavedInfo{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_setBotShippingResults:
		print("messages_setBotShippingResults")
		r = TL_messages_setBotShippingResults{
			m.Object(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_setBotPrecheckoutResults:
		print("messages_setBotPrecheckoutResults")
		r = TL_messages_setBotPrecheckoutResults{
			m.Object(),
			m.Object(),
			m.Long(),
			m.Object(),
		}

	case crc_stickers_createStickerSet:
		print("stickers_createStickerSet")
		r = TL_stickers_createStickerSet{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_stickers_removeStickerFromSet:
		print("stickers_removeStickerFromSet")
		r = TL_stickers_removeStickerFromSet{
			m.Object(),
		}

	case crc_stickers_changeStickerPosition:
		print("stickers_changeStickerPosition")
		r = TL_stickers_changeStickerPosition{
			m.Object(),
			m.Int(),
		}

	case crc_stickers_addStickerToSet:
		print("stickers_addStickerToSet")
		r = TL_stickers_addStickerToSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_uploadMedia:
		print("messages_uploadMedia")
		r = TL_messages_uploadMedia{
			m.Object(),
			m.Object(),
		}

	case crc_phone_getCallConfig:
		print("phone_getCallConfig")
		r = TL_phone_getCallConfig{}

	case crc_phone_requestCall:
		print("phone_requestCall")
		r = TL_phone_requestCall{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_acceptCall:
		print("phone_acceptCall")
		r = TL_phone_acceptCall{
			m.Object(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_confirmCall:
		print("phone_confirmCall")
		r = TL_phone_confirmCall{
			m.Object(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
		}

	case crc_phone_receivedCall:
		print("phone_receivedCall")
		r = TL_phone_receivedCall{
			m.Object(),
		}

	case crc_phone_discardCall:
		print("phone_discardCall")
		r = TL_phone_discardCall{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Object(),
			m.Long(),
		}

	case crc_phone_setCallRating:
		print("phone_setCallRating")
		r = TL_phone_setCallRating{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.String(),
		}

	case crc_phone_saveCallDebug:
		print("phone_saveCallDebug")
		r = TL_phone_saveCallDebug{
			m.Object(),
			m.Object(),
		}

	case crc_upload_getCdnFile:
		print("upload_getCdnFile")
		r = TL_upload_getCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case crc_upload_reuploadCdnFile:
		print("upload_reuploadCdnFile")
		r = TL_upload_reuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_help_getCdnConfig:
		print("help_getCdnConfig")
		r = TL_help_getCdnConfig{}

	case crc_langpack_getLangPack:
		print("langpack_getLangPack")
		r = TL_langpack_getLangPack{
			m.String(),
			m.String(),
		}

	case crc_langpack_getStrings:
		print("langpack_getStrings")
		r = TL_langpack_getStrings{
			m.String(),
			m.String(),
			m.VectorString(),
		}

	case crc_langpack_getDifference:
		print("langpack_getDifference")
		r = TL_langpack_getDifference{
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_langpack_getLanguages:
		print("langpack_getLanguages")
		r = TL_langpack_getLanguages{
			m.String(),
		}

	case crc_channels_editBanned:
		print("channels_editBanned")
		r = TL_channels_editBanned{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_getAdminLog:
		print("channels_getAdminLog")
		r = TL_channels_getAdminLog{
			m.Object(),
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_upload_getCdnFileHashes:
		print("upload_getCdnFileHashes")
		r = TL_upload_getCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messages_sendScreenshotNotification:
		print("messages_sendScreenshotNotification")
		r = TL_messages_sendScreenshotNotification{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_channels_setStickers:
		print("channels_setStickers")
		r = TL_channels_setStickers{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getFavedStickers:
		print("messages_getFavedStickers")
		r = TL_messages_getFavedStickers{
			m.Int(),
		}

	case crc_messages_faveSticker:
		print("messages_faveSticker")
		r = TL_messages_faveSticker{
			m.Object(),
			m.Object(),
		}

	case crc_channels_readMessageContents:
		print("channels_readMessageContents")
		r = TL_channels_readMessageContents{
			m.Object(),
			m.VectorInt(),
		}

	case crc_contacts_resetSaved:
		print("contacts_resetSaved")
		r = TL_contacts_resetSaved{}

	case crc_messages_getUnreadMentions:
		print("messages_getUnreadMentions")
		r = TL_messages_getUnreadMentions{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_deleteHistory:
		print("channels_deleteHistory")
		r = TL_channels_deleteHistory{
			m.Object(),
			m.Int(),
		}

	case crc_help_getRecentMeUrls:
		print("help_getRecentMeUrls")
		r = TL_help_getRecentMeUrls{
			m.String(),
		}

	case crc_channels_togglePreHistoryHidden:
		print("channels_togglePreHistoryHidden")
		r = TL_channels_togglePreHistoryHidden{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readMentions:
		print("messages_readMentions")
		r = TL_messages_readMentions{
			m.Object(),
		}

	case crc_messages_getRecentLocations:
		print("messages_getRecentLocations")
		r = TL_messages_getRecentLocations{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_sendMultiMedia:
		print("messages_sendMultiMedia")
		r = TL_messages_sendMultiMedia{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Object(),
		}

	case crc_messages_uploadEncryptedFile:
		print("messages_uploadEncryptedFile")
		r = TL_messages_uploadEncryptedFile{
			m.Object(),
			m.Object(),
		}

	case crc_account_getWebAuthorizations:
		print("account_getWebAuthorizations")
		r = TL_account_getWebAuthorizations{}

	case crc_account_resetWebAuthorization:
		print("account_resetWebAuthorization")
		r = TL_account_resetWebAuthorization{
			m.Long(),
		}

	case crc_account_resetWebAuthorizations:
		print("account_resetWebAuthorizations")
		r = TL_account_resetWebAuthorizations{}

	case crc_messages_searchStickerSets:
		print("messages_searchStickerSets")
		r = TL_messages_searchStickerSets{
			m.Object(),
			m.Object(),
			m.String(),
			m.Int(),
		}

	case crc_upload_getFileHashes:
		print("upload_getFileHashes")
		r = TL_upload_getFileHashes{
			m.Object(),
			m.Int(),
		}

	case crc_help_getProxyData:
		print("help_getProxyData")
		r = TL_help_getProxyData{}

	case crc_help_getTermsOfServiceUpdate:
		print("help_getTermsOfServiceUpdate")
		r = TL_help_getTermsOfServiceUpdate{}

	case crc_help_acceptTermsOfService:
		print("help_acceptTermsOfService")
		r = TL_help_acceptTermsOfService{
			m.Object(),
		}

	case crc_account_getAllSecureValues:
		print("account_getAllSecureValues")
		r = TL_account_getAllSecureValues{}

	case crc_account_getSecureValue:
		print("account_getSecureValue")
		r = TL_account_getSecureValue{
			m.Vector(),
		}

	case crc_account_saveSecureValue:
		print("account_saveSecureValue")
		r = TL_account_saveSecureValue{
			m.Object(),
			m.Long(),
		}

	case crc_account_deleteSecureValue:
		print("account_deleteSecureValue")
		r = TL_account_deleteSecureValue{
			m.Vector(),
		}

	case crc_users_setSecureValueErrors:
		print("users_setSecureValueErrors")
		r = TL_users_setSecureValueErrors{
			m.Object(),
			m.Vector(),
		}

	case crc_account_getAuthorizationForm:
		print("account_getAuthorizationForm")
		r = TL_account_getAuthorizationForm{
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_account_acceptAuthorization:
		print("account_acceptAuthorization")
		r = TL_account_acceptAuthorization{
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.Object(),
		}

	case crc_account_sendVerifyPhoneCode:
		print("account_sendVerifyPhoneCode")
		r = TL_account_sendVerifyPhoneCode{
			m.String(),
			m.Object(),
		}

	case crc_account_verifyPhone:
		print("account_verifyPhone")
		r = TL_account_verifyPhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_account_sendVerifyEmailCode:
		print("account_sendVerifyEmailCode")
		r = TL_account_sendVerifyEmailCode{
			m.String(),
		}

	case crc_account_verifyEmail:
		print("account_verifyEmail")
		r = TL_account_verifyEmail{
			m.String(),
			m.String(),
		}

	case crc_help_getDeepLinkInfo:
		print("help_getDeepLinkInfo")
		r = TL_help_getDeepLinkInfo{
			m.String(),
		}

	case crc_contacts_getSaved:
		print("contacts_getSaved")
		r = TL_contacts_getSaved{}

	case crc_channels_getLeftChannels:
		print("channels_getLeftChannels")
		r = TL_channels_getLeftChannels{
			m.Int(),
		}

	case crc_account_initTakeoutSession:
		print("account_initTakeoutSession")
		r = TL_account_initTakeoutSession{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_finishTakeoutSession:
		print("account_finishTakeoutSession")
		r = TL_account_finishTakeoutSession{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getSplitRanges:
		print("messages_getSplitRanges")
		r = TL_messages_getSplitRanges{}

	case crc_invokeWithMessagesRange:
		print("invokeWithMessagesRange")
		r = TL_invokeWithMessagesRange{
			m.Object(),
			m.Object(),
		}

	case crc_invokeWithTakeout:
		print("invokeWithTakeout")
		r = TL_invokeWithTakeout{
			m.Long(),
			m.Object(),
		}

	case crc_messages_markDialogUnread:
		print("messages_markDialogUnread")
		r = TL_messages_markDialogUnread{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_getDialogUnreadMarks:
		print("messages_getDialogUnreadMarks")
		r = TL_messages_getDialogUnreadMarks{}

	case crc_contacts_toggleTopPeers:
		print("contacts_toggleTopPeers")
		r = TL_contacts_toggleTopPeers{
			m.Object(),
		}

	case crc_messages_clearAllDrafts:
		print("messages_clearAllDrafts")
		r = TL_messages_clearAllDrafts{}

	case crc_help_getAppConfig:
		print("help_getAppConfig")
		r = TL_help_getAppConfig{}

	case crc_help_saveAppLog:
		print("help_saveAppLog")
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case crc_help_getPassportConfig:
		print("help_getPassportConfig")
		r = TL_help_getPassportConfig{
			m.Int(),
		}

	case crc_langpack_getLanguage:
		print("langpack_getLanguage")
		r = TL_langpack_getLanguage{
			m.String(),
			m.String(),
		}

	case crc_messages_updatePinnedMessage:
		print("messages_updatePinnedMessage")
		r = TL_messages_updatePinnedMessage{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_account_confirmPasswordEmail:
		print("account_confirmPasswordEmail")
		r = TL_account_confirmPasswordEmail{
			m.String(),
		}

	case crc_account_resendPasswordEmail:
		print("account_resendPasswordEmail")
		r = TL_account_resendPasswordEmail{}

	case crc_account_cancelPasswordEmail:
		print("account_cancelPasswordEmail")
		r = TL_account_cancelPasswordEmail{}

	case crc_help_getSupportName:
		print("help_getSupportName")
		r = TL_help_getSupportName{}

	case crc_help_getUserInfo:
		print("help_getUserInfo")
		r = TL_help_getUserInfo{
			m.Object(),
		}

	case crc_help_editUserInfo:
		print("help_editUserInfo")
		r = TL_help_editUserInfo{
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crc_account_getContactSignUpNotification:
		print("account_getContactSignUpNotification")
		r = TL_account_getContactSignUpNotification{}

	case crc_account_setContactSignUpNotification:
		print("account_setContactSignUpNotification")
		r = TL_account_setContactSignUpNotification{
			m.Object(),
		}

	case crc_account_getNotifyExceptions:
		print("account_getNotifyExceptions")
		r = TL_account_getNotifyExceptions{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendVote:
		print("messages_sendVote")
		r = TL_messages_sendVote{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_messages_getPollResults:
		print("messages_getPollResults")
		r = TL_messages_getPollResults{
			m.Object(),
			m.Int(),
		}

	case crc_messages_getOnlines:
		print("messages_getOnlines")
		r = TL_messages_getOnlines{
			m.Object(),
		}

	case crc_messages_getStatsURL:
		print("messages_getStatsURL")
		r = TL_messages_getStatsURL{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crc_messages_editChatAbout:
		print("messages_editChatAbout")
		r = TL_messages_editChatAbout{
			m.Object(),
			m.String(),
		}

	case crc_messages_editChatDefaultBannedRights:
		print("messages_editChatDefaultBannedRights")
		r = TL_messages_editChatDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case crc_account_getWallPaper:
		print("account_getWallPaper")
		r = TL_account_getWallPaper{
			m.Object(),
		}

	case crc_account_uploadWallPaper:
		print("account_uploadWallPaper")
		r = TL_account_uploadWallPaper{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crc_account_saveWallPaper:
		print("account_saveWallPaper")
		r = TL_account_saveWallPaper{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_installWallPaper:
		print("account_installWallPaper")
		r = TL_account_installWallPaper{
			m.Object(),
			m.Object(),
		}

	case crc_account_resetWallPapers:
		print("account_resetWallPapers")
		r = TL_account_resetWallPapers{}

	case crc_account_getAutoDownloadSettings:
		print("account_getAutoDownloadSettings")
		r = TL_account_getAutoDownloadSettings{}

	case crc_account_saveAutoDownloadSettings:
		print("account_saveAutoDownloadSettings")
		r = TL_account_saveAutoDownloadSettings{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_getEmojiKeywords:
		print("messages_getEmojiKeywords")
		r = TL_messages_getEmojiKeywords{
			m.String(),
		}

	case crc_messages_getEmojiKeywordsDifference:
		print("messages_getEmojiKeywordsDifference")
		r = TL_messages_getEmojiKeywordsDifference{
			m.String(),
			m.Int(),
		}

	case crc_messages_getEmojiKeywordsLanguages:
		print("messages_getEmojiKeywordsLanguages")
		r = TL_messages_getEmojiKeywordsLanguages{
			m.VectorString(),
		}

	case crc_messages_getEmojiURL:
		print("messages_getEmojiURL")
		r = TL_messages_getEmojiURL{
			m.String(),
		}

	case crc_folders_editPeerFolders:
		print("folders_editPeerFolders")
		r = TL_folders_editPeerFolders{
			m.Vector(),
		}

	case crc_folders_deleteFolder:
		print("folders_deleteFolder")
		r = TL_folders_deleteFolder{
			m.Int(),
		}

	case crc_messages_getSearchCounters:
		print("messages_getSearchCounters")
		r = TL_messages_getSearchCounters{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_getGroupsForDiscussion:
		print("channels_getGroupsForDiscussion")
		r = TL_channels_getGroupsForDiscussion{}

	case crc_channels_setDiscussionGroup:
		print("channels_setDiscussionGroup")
		r = TL_channels_setDiscussionGroup{
			m.Object(),
			m.Object(),
		}

	case crc_messages_requestUrlAuth:
		print("messages_requestUrlAuth")
		r = TL_messages_requestUrlAuth{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_acceptUrlAuth:
		print("messages_acceptUrlAuth")
		r = TL_messages_acceptUrlAuth{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_hidePeerSettingsBar:
		print("messages_hidePeerSettingsBar")
		r = TL_messages_hidePeerSettingsBar{
			m.Object(),
		}

	case crc_contacts_addContact:
		print("contacts_addContact")
		r = TL_contacts_addContact{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_contacts_acceptContact:
		print("contacts_acceptContact")
		r = TL_contacts_acceptContact{
			m.Object(),
		}

	case crc_channels_editCreator:
		print("channels_editCreator")
		r = TL_channels_editCreator{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_contacts_getLocated:
		print("contacts_getLocated")
		r = TL_contacts_getLocated{
			m.Object(),
		}

	case crc_channels_editLocation:
		print("channels_editLocation")
		r = TL_channels_editLocation{
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crc_channels_toggleSlowMode:
		print("channels_toggleSlowMode")
		r = TL_channels_toggleSlowMode{
			m.Object(),
			m.Int(),
		}

	case crc_messages_getScheduledHistory:
		print("messages_getScheduledHistory")
		r = TL_messages_getScheduledHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_getScheduledMessages:
		print("messages_getScheduledMessages")
		r = TL_messages_getScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_messages_sendScheduledMessages:
		print("messages_sendScheduledMessages")
		r = TL_messages_sendScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_messages_deleteScheduledMessages:
		print("messages_deleteScheduledMessages")
		r = TL_messages_deleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_account_uploadTheme:
		print("account_uploadTheme")
		r = TL_account_uploadTheme{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_account_createTheme:
		print("account_createTheme")
		r = TL_account_createTheme{
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_account_updateTheme:
		print("account_updateTheme")
		r = TL_account_updateTheme{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_saveTheme:
		print("account_saveTheme")
		r = TL_account_saveTheme{
			m.Object(),
			m.Object(),
		}

	case crc_account_installTheme:
		print("account_installTheme")
		r = TL_account_installTheme{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_account_getTheme:
		print("account_getTheme")
		r = TL_account_getTheme{
			m.String(),
			m.Object(),
			m.Long(),
		}

	case crc_account_getThemes:
		print("account_getThemes")
		r = TL_account_getThemes{
			m.String(),
			m.Int(),
		}

	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil
	}
	if m.err != nil {
		return nil
	}
	return
}
