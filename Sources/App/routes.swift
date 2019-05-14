import Vapor
import Authentication


/// Register your application's routes here.
public func routes(_ router: Router) throws {

    let userController = UserController()
    router.get("register", use: userController.renderRegister)
    router.post("register", use: userController.register)
    router.get("login", use: userController.renderLogin)
    
    let authSessionRouter = router.grouped(User.authSessionsMiddleware())
    authSessionRouter.post("login", use: userController.login)
    
    let protectedRouter = authSessionRouter.grouped(RedirectMiddleware<User>(path: "/login"))
    protectedRouter.get("profile", use: userController.renderProfile)

    router.get("logout", use: userController.logout)
    
    router.get("") { req -> Future<View> in
        return try req.view().render("home")
    }
}
