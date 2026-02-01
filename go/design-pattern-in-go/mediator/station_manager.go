package main

// StationManager は具象Mediator（駅長）
// プラットフォームの状態を管理し、列車間の調整を行う
type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		trainQueue:     make([]Train, 0),
	}
}

// CanArrive はプラットフォームが空いているか確認し、到着を許可または待機させる
func (s *StationManager) CanArrive(t Train) bool {
	if s.isPlatformFree {
		// （本筋ではないので実装はしないが厳密にはロック取ったほうがよさそう）
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

// NotifyAboutDeparture は列車の出発を通知し、待機中の列車に到着を許可する
func (s *StationManager) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrain.PermitArrival()
	}
}
