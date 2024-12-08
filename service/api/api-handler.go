package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// User routes
	//
	// -- Do Login -- //
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))
	//
	// -- Set my UserName -- //
	rt.router.PUT("/profiles/:user/username", rt.wrap(rt.setMyUserName, true))
	//
	// -- Set my Photo -- //
	rt.router.PUT("/profiles/:user/photo", rt.wrap(rt.setMyPhoto, true))
	//
	// -- Delete User -- //
	rt.router.DELETE("/profiles/:user", rt.wrap(rt.deleteUser, true))
	//
	// -- Seacr User -- //
	rt.router.GET("/profiles", rt.wrap(rt.searchUsers, true))

	// Group routes
	//
	// -- Create Group -- //
	rt.router.POST("/profiles/:user/groups", rt.wrap(rt.createGroup, true))
	//
	// -- Leave Group -- //
	rt.router.DELETE("/profiles/:user/groups/:group", rt.wrap(rt.leaveGroup, true))
	//
	// -- Add to Group -- //
	rt.router.PUT("/profiles/:user/groups/:group", rt.wrap(rt.addToGroup, true))
	//
	// -- Set new name to the group -- //
	rt.router.PUT("/profiles/:user/groups/:group/name", rt.wrap(rt.setGroupName, true))
	//
	// -- Set new photo group -- //
	rt.router.PUT("/profiles/:user/groups/:group/photo", rt.wrap(rt.setGroupPhoto, true))

	// Conversation routes
	//
	// -- Create Conversation -- //
	rt.router.PUT("/profiles/:user/conversations/:receiver", rt.wrap(rt.createConversation, true))
	//
	// -- Get My Conversations -- //
	rt.router.GET("/profiles/:user/conversations", rt.wrap(rt.getMyConversations, true))
	//
	// -- Get Conversation -- //
	rt.router.GET("/profiles/:user/conversations/:conversation", rt.wrap(rt.getConversation, true))

	// Message routes
	//
	// -- Send Message -- //
	rt.router.POST("/profiles/:user/conversations/:conversation/messages", rt.wrap(rt.sendMessage, true))
	//
	// -- Fortward Message -- //
	rt.router.POST("/profiles/:user/conversations/:conversation/messages/:message", rt.wrap(rt.forwardMessage, true))
	//
	// -- Delete Message -- //
	rt.router.DELETE("/profiles/:user/conversations/:conversation/messages/:message", rt.wrap(rt.deleteMessage, true))

	// Comment routes
	//
	// -- Comment Message -- //
	rt.router.PUT("/profiles/:user/conversations/:receiver/messages/:message/comments", rt.wrap(rt.commentMessage, true))
	//
	// -- Uncomment Message -- //
	rt.router.DELETE("/profiles/:user/conversations/:conversation/messages/:message/comments/:comment", rt.wrap(rt.uncommentMessage, true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
