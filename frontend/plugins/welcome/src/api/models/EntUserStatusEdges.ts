/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserStatusEdges
 */
export interface EntUserStatusEdges {
    /**
     * UserstatusUser holds the value of the UserstatusUser edge.
     * @type {Array<EntUser>}
     * @memberof EntUserStatusEdges
     */
    userstatusUser?: Array<EntUser>;
}

export function EntUserStatusEdgesFromJSON(json: any): EntUserStatusEdges {
    return EntUserStatusEdgesFromJSONTyped(json, false);
}

export function EntUserStatusEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserStatusEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userstatusUser': !exists(json, 'userstatusUser') ? undefined : ((json['userstatusUser'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntUserStatusEdgesToJSON(value?: EntUserStatusEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userstatusUser': value.userstatusUser === undefined ? undefined : ((value.userstatusUser as Array<any>).map(EntUserToJSON)),
    };
}

