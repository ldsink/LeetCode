package main

type AuthenticationManager struct {
	tokens map[string]int
	ttl    int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{ttl: timeToLive, tokens: make(map[string]int)}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	ttl, ok := this.tokens[tokenId]
	if !ok {
		return
	}
	if ttl > currentTime {
		this.tokens[tokenId] = currentTime + this.ttl
	} else {
		delete(this.tokens, tokenId)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var count int
	for _, ttl := range this.tokens {
		if ttl > currentTime {
			count++
		}
	}
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
