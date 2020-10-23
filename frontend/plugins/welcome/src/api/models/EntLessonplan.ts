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
    EntLessonplanEdges,
    EntLessonplanEdgesFromJSON,
    EntLessonplanEdgesFromJSONTyped,
    EntLessonplanEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLessonplan
 */
export interface EntLessonplan {
    /**
     * Room holds the value of the "Room" field.
     * @type {string}
     * @memberof EntLessonplan
     */
    room?: string;
    /**
     * 
     * @type {EntLessonplanEdges}
     * @memberof EntLessonplan
     */
    edges?: EntLessonplanEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntLessonplan
     */
    id?: number;
}

export function EntLessonplanFromJSON(json: any): EntLessonplan {
    return EntLessonplanFromJSONTyped(json, false);
}

export function EntLessonplanFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLessonplan {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'room': !exists(json, 'Room') ? undefined : json['Room'],
        'edges': !exists(json, 'edges') ? undefined : EntLessonplanEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntLessonplanToJSON(value?: EntLessonplan | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Room': value.room,
        'edges': EntLessonplanEdgesToJSON(value.edges),
        'id': value.id,
    };
}


