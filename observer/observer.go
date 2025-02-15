package observer

import (
	"log"
	//"math"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CHANNEL_ID = -1002348931205
)

type Observer struct {
	followerChannelsList []int64
	sourceChannel        int64
	mu                   *sync.Mutex
}

/* Конструктор */
func NewObserver() *Observer {
	return &Observer{
		followerChannelsList: make([]int64, 0, 16),
		mu:                   &sync.Mutex{},
	}
}

func (obs *Observer) AddChannelToList(ID int64) {
	obs.mu.Lock()
	defer obs.mu.Unlock()
	obs.followerChannelsList = append(obs.followerChannelsList, ID)
}

func (obs *Observer) RemoveChannelToList(ID int64) {
	obs.mu.Lock()
	defer obs.mu.Unlock()
	buffer := make([]int64, 0, 16)
	for _, val := range obs.followerChannelsList {
		if val != ID {
			buffer = append(buffer, val)
		}
	}
	obs.followerChannelsList = buffer
}

// -1002348931205 - ID первого канала
func (obs *Observer) PushMessageToRouting(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	obs.mu.Lock()
	defer obs.mu.Unlock()

	if update.ChannelPost != nil {
		log.Printf("Новое сообщение в канале %v, текст сообщения: %v", update.ChannelPost.SenderChat, update.ChannelPost.Text)

		/*if update.ChannelPost.SenderChat.ID != CHANNEL_ID {
			//msg := tgbotapi.NewForward(CHANNEL_ID, update.ChannelPost.SenderChat.ID, update.ChannelPost.MessageID)
			msg := tgbotapi.NewMessageToChannel("Admin", update.ChannelPost.Text)
			a := update.ChannelPost.Animation
			b := update.ChannelPost.Audio

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Ошибка при пересылке: %v", err)
			}
		}*/




		/*// Собираем подпись (если есть текст)
		var caption string
		if update.ChannelPost.Caption != "" {
			caption = update.ChannelPost.Caption
		} else if update.ChannelPost.Text != "" {
			caption = update.ChannelPost.Text
		}

		// Если есть несколько фото — отправляем как MediaGroup
		if len(update.ChannelPost.Photo) > 0 {
			media := []tgbotapi.MediaGroupConfig{}

			//Добавляем все фото в группу
			for i, _ := range update.ChannelPost.Photo {
				//mediaType := "photo"
				if i == len(update.ChannelPost.Photo)-1 { // Только последнее фото получает подпись
					//media = append(media, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(photo.FileID)).WithCaption(caption))
				} else {
					//media = append(media, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(photo.FileID)))
				}
			}

			_, err := bot.SendMediaGroup(media)
			if err != nil {
				log.Printf("Ошибка при отправке медиа-группы: %v", err)
			}
		} else if update.ChannelPost.Video != nil { // Если есть видео
			msg := tgbotapi.NewVideo(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Video.FileID))
			msg.Caption = caption
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка при отправке видео: %v", err)
			}
		} else if update.ChannelPost.Audio != nil { // Если есть аудио
			msg := tgbotapi.NewAudio(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Audio.FileID))
			msg.Caption = caption
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка при отправке аудио: %v", err)
			}
		} else if update.ChannelPost.Document != nil { // Если есть документ
			msg := tgbotapi.NewDocument(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Document.FileID))
			msg.Caption = caption
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка при отправке документа: %v", err)
			}
		} else if update.ChannelPost.Text != "" { // Если это просто текст
			msg := tgbotapi.NewMessage(CHANNEL_ID, update.ChannelPost.Text)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка при отправке текста: %v", err)
			}
		}*/



		// Извлекаем текст публикации
		text := update.ChannelPost.Text

		// Если есть медиа, отправляем их в новый канал
		if len(update.ChannelPost.Photo) > 0 {
			var mediaGroup []interface{}
			for i, photo := range update.ChannelPost.Photo {
				inputMediaPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(photo.FileID))
				// Добавляем текст только к первому фото
				if i == 0 {
					inputMediaPhoto.Caption = update.ChannelPost.Caption
				}
				mediaGroup = append(mediaGroup, inputMediaPhoto)
			}
			// s := int(math.Sqrt(float64(len(update.ChannelPost.Photo))))
			// for i := 0; i < s; s += s {
			// 	inputMediaPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(update.ChannelPost.Photo[i].FileID))
			// 	// Добавляем текст только к первому фото
			// 	if i == 0 {
			// 		inputMediaPhoto.Caption = update.ChannelPost.Caption
			// 	}
			// 	mediaGroup = append(mediaGroup, inputMediaPhoto)
			// }
			log.Println(">>>>>>>>>>>>>>>", len(update.ChannelPost.Photo))
			// msg := tgbotapi.NewMediaGroup(CHANNEL_ID, mediaGroup)
			// if _, err := bot.Send(msg); err != nil {
			// 	log.Println("Ошибка при отправке альбома:", err)
			// }
			msg := tgbotapi.NewMediaGroup(CHANNEL_ID, []interface{}{tgbotapi.FileID(update.ChannelPost.Photo[0].FileID)})
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке альбома:", err)
			}
		} else if update.ChannelPost.Video != nil {
			msg := tgbotapi.NewVideo(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Video.FileID))
			msg.Caption = text // Добавляем текст как подпись к видео
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке видео:", err)
			}
		} else if update.ChannelPost.Audio != nil {
			msg := tgbotapi.NewAudio(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Audio.FileID))
			msg.Caption = text // Добавляем текст как подпись к аудио
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке аудио:", err)
			}
		} else if update.ChannelPost.Document != nil {
			msg := tgbotapi.NewDocument(CHANNEL_ID, tgbotapi.FileID(update.ChannelPost.Document.FileID))
			msg.Caption = text // Добавляем текст как подпись к документу
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке документа:", err)
			}
		} else {
			// Если нет медиа, отправляем только текст
			msg := tgbotapi.NewMessage(CHANNEL_ID, text)
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке текста:", err)
			}
		}
	}

	if update.Message != nil {
		log.Printf("Новое сообщение в группе-источнике %v, текст сообщения: %v", update.Message.SenderChat, update.Message.Text)

		/*for _, chanID := range obs.followerChannelsList {
			msg := tgbotapi.NewForward(chanID, update.ChannelPost.From.ID, update.ChannelPost.MessageID)
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Ошибка при пересылке: %v", err)
			}
		}*/
	}
}
