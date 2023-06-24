

/**
 * Quiz API
 * Quiz App backend 
 *
 * The version of the OpenAPI document: v1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * The id of the user
     * @type {string}
     * @memberof User
     */
    'id'?: string;
    /**
     * The email of the user
     * @type {string}
     * @memberof User
     */
    'email'?: string;
    /**
     * The firstname of the user
     * @type {string}
     * @memberof User
     */
    'firstname'?: string;
    /**
     * The lastname of the user
     * @type {string}
     * @memberof User
     */
    'lastname'?: string;
    /**
     * If the user is active or not
     * @type {boolean}
     * @memberof User
     */
    'active'?: boolean;
    /**
     * The role of the user
     * @type {string}
     * @memberof User
     */
    'role'?: UserRoleEnum;
}

export const UserRoleEnum = {
    NoRole: 'NO_ROLE',
    Student: 'STUDENT',
    Teacher: 'TEACHER',
    Admin: 'ADMIN'
} as const;

export type UserRoleEnum = typeof UserRoleEnum[keyof typeof UserRoleEnum];


