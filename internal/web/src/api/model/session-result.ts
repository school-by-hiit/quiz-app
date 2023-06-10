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
 * An object describing the result of the quiz
 * @export
 * @interface SessionResult
 */
export interface SessionResult {
  /**
   * The number of good answer in the quiz
   * @type {number}
   * @memberof SessionResult
   */
  goodAnswer?: number | null;
  /**
   * The total number of answer in the quiz
   * @type {number}
   * @memberof SessionResult
   */
  totalAnswer?: number | null;
}
