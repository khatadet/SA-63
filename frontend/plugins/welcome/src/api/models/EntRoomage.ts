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
    EntRoomageEdges,
    EntRoomageEdgesFromJSON,
    EntRoomageEdgesFromJSONTyped,
    EntRoomageEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomage
 */
export interface EntRoomage {
    /**
     * RoomAge holds the value of the "RoomAge" field.
     * @type {number}
     * @memberof EntRoomage
     */
    roomAge?: number;
    /**
     * Text holds the value of the "Text" field.
     * @type {string}
     * @memberof EntRoomage
     */
    text?: string;
    /**
     * 
     * @type {EntRoomageEdges}
     * @memberof EntRoomage
     */
    edges?: EntRoomageEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRoomage
     */
    id?: number;
}

export function EntRoomageFromJSON(json: any): EntRoomage {
    return EntRoomageFromJSONTyped(json, false);
}

export function EntRoomageFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomage {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomAge': !exists(json, 'RoomAge') ? undefined : json['RoomAge'],
        'text': !exists(json, 'Text') ? undefined : json['Text'],
        'edges': !exists(json, 'edges') ? undefined : EntRoomageEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntRoomageToJSON(value?: EntRoomage | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'RoomAge': value.roomAge,
        'Text': value.text,
        'edges': EntRoomageEdgesToJSON(value.edges),
        'id': value.id,
    };
}

