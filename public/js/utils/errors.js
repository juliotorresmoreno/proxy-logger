class ErrorUser extends Error {

}
class ErrorUsernameNotFound extends ErrorUser {
    constructor() {
        super('Username not found!.');
    }
}

class ErrorPasswordNotFound extends ErrorUser {
    constructor() {
        super('Password not found!.');
    }
}

/**
 * @typedef {import('./errors').ErrorCode} ErrorCode
 * @typedef {import('./errors').ResponseError} ResponseError
 */

class ErrorHTTP extends Error {
    /**
     * @type {ErrorCode}
     */
    statusCode = 500;
}

class ErrorUnauthorized extends ErrorHTTP {
    constructor() {
        super('Unauthorized!.');
        this.statusCode = 401;
    }
}

class BadRequestError extends ErrorHTTP {
    /**
     * 
     * @param {string | undefined} message 
     */
    constructor(message) {
        super(message || 'Bad Request');
        this.statusCode = 400;
    }
}

class ErrorInternalServerError extends ErrorHTTP {
    constructor() {
        super('Internal Server Error');
        this.statusCode = 500;
    }
}

/**
 * @param {Response} response 
 */
async function parseError(response) {
    if (!response.ok) {
        if (response.status === 401)
            throw new ErrorUnauthorized();
        if (response.status === 400) {
            /**
             * @type {ResponseError}
             */
            const badRequestError = response.json();
            throw new BadRequestError(badRequestError.message);
        }
    }
}