/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelCloudComplianceScanTrigger
 */
export interface ModelCloudComplianceScanTrigger {
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelCloudComplianceScanTrigger
     */
    benchmark_types: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudComplianceScanTrigger
     */
    node_id: string;
}

/**
 * Check if a given object implements the ModelCloudComplianceScanTrigger interface.
 */
export function instanceOfModelCloudComplianceScanTrigger(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "benchmark_types" in value;
    isInstance = isInstance && "node_id" in value;

    return isInstance;
}

export function ModelCloudComplianceScanTriggerFromJSON(json: any): ModelCloudComplianceScanTrigger {
    return ModelCloudComplianceScanTriggerFromJSONTyped(json, false);
}

export function ModelCloudComplianceScanTriggerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCloudComplianceScanTrigger {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'benchmark_types': json['benchmark_types'],
        'node_id': json['node_id'],
    };
}

export function ModelCloudComplianceScanTriggerToJSON(value?: ModelCloudComplianceScanTrigger | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'benchmark_types': value.benchmark_types,
        'node_id': value.node_id,
    };
}

