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
    EntRent,
    EntRentFromJSON,
    EntRentFromJSONTyped,
    EntRentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomageEdges
 */
export interface EntRoomageEdges {
    /**
     * RoomageRent holds the value of the RoomageRent edge.
     * @type {Array<EntRent>}
     * @memberof EntRoomageEdges
     */
    roomageRent?: Array<EntRent>;
}

export function EntRoomageEdgesFromJSON(json: any): EntRoomageEdges {
    return EntRoomageEdgesFromJSONTyped(json, false);
}

export function EntRoomageEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomageEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomageRent': !exists(json, 'roomageRent') ? undefined : ((json['roomageRent'] as Array<any>).map(EntRentFromJSON)),
    };
}

export function EntRoomageEdgesToJSON(value?: EntRoomageEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomageRent': value.roomageRent === undefined ? undefined : ((value.roomageRent as Array<any>).map(EntRentToJSON)),
    };
}


