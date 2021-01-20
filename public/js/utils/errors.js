
class UsernameNotFound extends Error {
    constructor() {
        super('Username not found!.');
    }
}

class PasswordNotFound extends Error {
    constructor() {
        super('Password not found!.');
    }
}

class Unauthorized extends Error {
    constructor() {
        super('Unauthorized!.');
    }
}