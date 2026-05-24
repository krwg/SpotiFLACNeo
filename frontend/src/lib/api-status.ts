<<<<<<< HEAD
import { CheckAPIStatus } from "../../wailsjs/go/main/App";
import { CHECK_TIMEOUT_MS, withTimeout } from "@/lib/async-timeout";
export type ApiCheckStatus = "checking" | "online" | "offline" | "idle";
=======
import { CheckAPIStatus, CheckCustomTidalAPI } from "../../wailsjs/go/main/App";
import { CHECK_TIMEOUT_MS, withTimeout } from "@/lib/async-timeout";
import { getSettings, hasConfiguredCustomTidalApi } from "@/lib/settings";

export type ApiCheckStatus = "checking" | "online" | "offline" | "idle";

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export interface ApiSource {
    id: string;
    type: string;
    name: string;
    url: string;
}
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
interface SpotiFLACNextSource {
    id: string;
    name: string;
    statusKey?: string;
    statusPrefix?: string;
}
<<<<<<< HEAD
type SpotiFLACNextStatusResponse = Partial<Record<string, string>>;
=======

type SpotiFLACNextStatusResponse = Partial<Record<string, string>>;
type ApiStatusTargetReport = {
    target?: string;
    label?: string;
    online?: boolean;
    message?: string;
};
type ApiStatusReport = {
    type?: string;
    online?: boolean;
    require_all?: boolean;
    details?: ApiStatusTargetReport[];
};

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export const API_SOURCES: ApiSource[] = [
    { id: "tidal", type: "tidal", name: "Tidal", url: "" },
    { id: "qobuz", type: "qobuz", name: "Qobuz", url: "" },
    { id: "amazon", type: "amazon", name: "Amazon Music", url: "" },
<<<<<<< HEAD
    { id: "musicbrainz", type: "musicbrainz", name: "MusicBrainz", url: "https://musicbrainz.org" },
];
=======
];

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export const SPOTIFLAC_NEXT_SOURCES: SpotiFLACNextSource[] = [
    { id: "tidal", name: "Tidal", statusKey: "tidal" },
    { id: "qobuz", name: "Qobuz", statusPrefix: "qobuz_" },
    { id: "amazon", name: "Amazon Music", statusPrefix: "amazon_" },
    { id: "deezer", name: "Deezer", statusPrefix: "deezer_" },
    { id: "apple", name: "Apple Music", statusKey: "apple" },
];
<<<<<<< HEAD
const SPOTIFLAC_NEXT_STATUS_URL = "https://gist.githubusercontent.com/afkarxyz/6e57cd362cbd67f889e3a91a76254a5e/raw";
const SPOTIFLAC_NEXT_MAX_ATTEMPTS = 3;
const SPOTIFLAC_NEXT_RETRY_DELAY_MS = 1200;
=======

const SPOTIFLAC_STATUS_URL = "https://gist.githubusercontent.com/afkarxyz/6e57cd362cbd67f889e3a91a76254a5e/raw";
const SPOTIFLAC_CURRENT_AMAZON_STATUS_KEY = "amazon_a";
const SPOTIFLAC_STATUS_MAX_ATTEMPTS = 3;
const SPOTIFLAC_STATUS_RETRY_DELAY_MS = 1200;
const CheckAPIStatusReport = (apiType: string, apiURL: string): Promise<ApiStatusReport> => (window as any)["go"]["main"]["App"]["CheckAPIStatusReport"](apiType, apiURL);
const LogStatusConsole = (level: string, message: string): Promise<void> => (window as any)["go"]["main"]["App"]["LogStatusConsole"](level, message);

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
type ApiStatusState = {
    checkingSources: Record<string, boolean>;
    statuses: Record<string, ApiCheckStatus>;
    nextStatuses: Record<string, ApiCheckStatus>;
};
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
let apiStatusState: ApiStatusState = {
    checkingSources: {},
    statuses: {},
    nextStatuses: {},
};
<<<<<<< HEAD
let activeCheckNextOnly: Promise<void> | null = null;
const activeSourceChecks = new Map<string, Promise<void>>();
const listeners = new Set<() => void>();
=======

let activeCheckCurrentOnly: Promise<void> | null = null;
let activeCheckNextOnly: Promise<void> | null = null;
let activeStatusPayloadFetch: Promise<SpotiFLACNextStatusResponse> | null = null;

const activeSourceChecks = new Map<string, Promise<void>>();
const listeners = new Set<() => void>();

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
function emitApiStatusChange() {
    for (const listener of listeners) {
        listener();
    }
}
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
function setApiStatusState(updater: (current: ApiStatusState) => ApiStatusState) {
    apiStatusState = updater(apiStatusState);
    emitApiStatusChange();
}
<<<<<<< HEAD
async function checkSourceStatus(source: ApiSource): Promise<ApiCheckStatus> {
    try {
        const isOnline = await withTimeout(CheckAPIStatus(source.type, source.url), CHECK_TIMEOUT_MS, `API status check timed out after 10 seconds for ${source.name}`);
        return isOnline ? "online" : "offline";
    }
    catch {
        return "offline";
    }
}
function anyNextVariantUp(values: Array<string | undefined>): ApiCheckStatus {
    return values.some((value) => value === "up") ? "online" : "offline";
}
=======

function delay(ms: number): Promise<void> {
    return new Promise((resolve) => window.setTimeout(resolve, ms));
}
function sendStatusConsole(level: "info" | "warning" | "error", message: string): void {
    try {
        void LogStatusConsole(level, message);
    }
    catch {
        return;
    }
}
function logStatusInfo(message: string): void {
    sendStatusConsole("info", message);
}
function logStatusWarning(message: string): void {
    sendStatusConsole("warning", message);
}
function logStatusError(message: string): void {
    sendStatusConsole("error", message);
}
function truncateStatusMessage(message?: string, maxLen = 180): string {
    const trimmed = (message || "").trim();
    if (trimmed.length <= maxLen) {
        return trimmed;
    }
    return trimmed.slice(0, maxLen) + "...";
}
function logQobuzStatusReport(report: ApiStatusReport): void {
    const details = Array.isArray(report.details) ? report.details : [];
    if (details.length === 0) {
        logStatusWarning("[Status][Qobuz] No provider details were returned.");
        return;
    }
    const onlineCount = details.filter((detail) => detail.online === true).length;
    logStatusInfo(`[Status][Qobuz] Provider check completed: ${onlineCount}/${details.length} providers online.`);
    for (const detail of details) {
        const label = detail.label || detail.target || "Unknown provider";
        const suffix = detail.message ? ` - ${truncateStatusMessage(detail.message)}` : "";
        if (detail.online) {
            logStatusInfo(`[Status][Qobuz] ${label}: online${suffix}`);
        }
        else {
            logStatusWarning(`[Status][Qobuz] ${label}: offline${suffix}`);
        }
    }
    if (report.online) {
        logStatusInfo(`[Status][Qobuz] SpotiFLAC Qobuz is online (${onlineCount}/${details.length} providers online).`);
    }
    else {
        logStatusWarning(`[Status][Qobuz] SpotiFLAC Qobuz marked maintenance because all ${details.length} providers are offline.`);
    }
}

function anyNextVariantUp(values: Array<string | undefined>): ApiCheckStatus {
    return values.some((value) => value === "up") ? "online" : "offline";
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
function getNextSourceValues(payload: SpotiFLACNextStatusResponse, source: SpotiFLACNextSource): string[] {
    if (source.statusKey) {
        const value = payload[source.statusKey];
        return typeof value === "string" ? [value] : [];
    }
    if (!source.statusPrefix) {
        return [];
    }
    const values: string[] = [];
    for (const [key, value] of Object.entries(payload)) {
        if (key.startsWith(source.statusPrefix) && typeof value === "string") {
            values.push(value);
        }
    }
    return values;
}
<<<<<<< HEAD
function delay(ms: number): Promise<void> {
    return new Promise((resolve) => window.setTimeout(resolve, ms));
}
=======

function getCurrentAmazonStatus(payload: SpotiFLACNextStatusResponse): ApiCheckStatus {
    return payload[SPOTIFLAC_CURRENT_AMAZON_STATUS_KEY] === "up" ? "online" : "offline";
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
function getSafeNextStatusesFallback(currentStatuses: Record<string, ApiCheckStatus>): Record<string, ApiCheckStatus> {
    return SPOTIFLAC_NEXT_SOURCES.reduce<Record<string, ApiCheckStatus>>((acc, source) => {
        const current = currentStatuses[source.id];
        acc[source.id] = current === "online" || current === "offline" ? current : "idle";
        return acc;
    }, {});
}
<<<<<<< HEAD
async function fetchSpotiFLACNextStatusesOnce(): Promise<Record<string, ApiCheckStatus>> {
    const response = await withTimeout(fetch(SPOTIFLAC_NEXT_STATUS_URL, {
        method: "GET",
        cache: "no-store",
        headers: {
            Accept: "application/json",
        },
    }), CHECK_TIMEOUT_MS, "SpotiFLAC Next status check timed out after 10 seconds");
    if (!response.ok) {
        throw new Error(`SpotiFLAC Next status returned ${response.status}`);
    }
    const payload = (await response.json()) as SpotiFLACNextStatusResponse;
    return SPOTIFLAC_NEXT_SOURCES.reduce<Record<string, ApiCheckStatus>>((acc, source) => {
        acc[source.id] = anyNextVariantUp(getNextSourceValues(payload, source));
        return acc;
    }, {});
}
async function checkSpotiFLACNextStatuses(): Promise<Record<string, ApiCheckStatus>> {
    let lastError: unknown = null;
    for (let attempt = 1; attempt <= SPOTIFLAC_NEXT_MAX_ATTEMPTS; attempt++) {
        try {
            return await fetchSpotiFLACNextStatusesOnce();
        }
        catch (error) {
            lastError = error;
            if (attempt < SPOTIFLAC_NEXT_MAX_ATTEMPTS) {
                await delay(SPOTIFLAC_NEXT_RETRY_DELAY_MS * attempt);
            }
        }
    }
    throw lastError instanceof Error ? lastError : new Error("SpotiFLAC Next status check failed");
}
export function getApiStatusState(): ApiStatusState {
    return apiStatusState;
}
export function subscribeApiStatus(listener: () => void): () => void {
    listeners.add(listener);
    return () => {
        listeners.delete(listener);
    };
}
=======

function hasCurrentResults(): boolean {
    return API_SOURCES.some((source) => {
        const status = apiStatusState.statuses[source.id];
        return status === "online" || status === "offline";
    });
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
function hasSpotiFLACNextResults(): boolean {
    return SPOTIFLAC_NEXT_SOURCES.some((source) => {
        const status = apiStatusState.nextStatuses[source.id];
        return status === "online" || status === "offline";
    });
}
<<<<<<< HEAD
=======

async function fetchSpotiFLACStatusPayloadOnce(): Promise<SpotiFLACNextStatusResponse> {
    const response = await withTimeout(fetch(SPOTIFLAC_STATUS_URL, {
        method: "GET",
        cache: "no-store",
        headers: {
            Accept: "application/json",
        },
    }), CHECK_TIMEOUT_MS, "SpotiFLAC status check timed out after 10 seconds");

    if (!response.ok) {
        throw new Error(`SpotiFLAC status returned ${response.status}`);
    }

    return (await response.json()) as SpotiFLACNextStatusResponse;
}

async function fetchSpotiFLACStatusPayload(): Promise<SpotiFLACNextStatusResponse> {
    if (activeStatusPayloadFetch) {
        return activeStatusPayloadFetch;
    }

    activeStatusPayloadFetch = (async () => {
        let lastError: unknown = null;
        for (let attempt = 1; attempt <= SPOTIFLAC_STATUS_MAX_ATTEMPTS; attempt++) {
            try {
                return await fetchSpotiFLACStatusPayloadOnce();
            }
            catch (error) {
                lastError = error;
                if (attempt < SPOTIFLAC_STATUS_MAX_ATTEMPTS) {
                    await delay(SPOTIFLAC_STATUS_RETRY_DELAY_MS * attempt);
                }
            }
        }
        throw lastError instanceof Error ? lastError : new Error("SpotiFLAC status check failed");
    })();

    try {
        return await activeStatusPayloadFetch;
    }
    finally {
        activeStatusPayloadFetch = null;
    }
}

async function checkSourceStatus(source: ApiSource): Promise<ApiCheckStatus> {
    try {
        if (source.id === "tidal") {
            const customTidalApi = getSettings().customTidalApi;
            if (!hasConfiguredCustomTidalApi(customTidalApi)) {
                logStatusWarning("[Status][Tidal] Marked maintenance because no custom Tidal instance is configured.");
                return "offline";
            }
            const isOnline = await withTimeout(CheckCustomTidalAPI(customTidalApi), CHECK_TIMEOUT_MS, `API status check timed out after 10 seconds for ${source.name}`);
            return isOnline ? "online" : "offline";
        }

        if (source.id === "amazon") {
            const payload = await fetchSpotiFLACStatusPayload();
            return getCurrentAmazonStatus(payload);
        }

        if (source.id === "qobuz") {
            logStatusInfo("[Status][Qobuz] Checking current SpotiFLAC providers...");
            const report = await withTimeout(CheckAPIStatusReport(source.type, source.url), CHECK_TIMEOUT_MS, `API status report timed out after 10 seconds for ${source.name}`);
            logQobuzStatusReport(report);
            return report.online ? "online" : "offline";
        }

        const isOnline = await withTimeout(CheckAPIStatus(source.type, source.url), CHECK_TIMEOUT_MS, `API status check timed out after 10 seconds for ${source.name}`);
        return isOnline ? "online" : "offline";
    }
    catch (error) {
        if (source.id === "qobuz") {
            logStatusError(`[Status][Qobuz] Provider check failed: ${error instanceof Error ? error.message : String(error)}`);
        }
        return "offline";
    }
}

async function checkSpotiFLACNextStatuses(): Promise<Record<string, ApiCheckStatus>> {
    const payload = await fetchSpotiFLACStatusPayload();
    return SPOTIFLAC_NEXT_SOURCES.reduce<Record<string, ApiCheckStatus>>((acc, source) => {
        acc[source.id] = anyNextVariantUp(getNextSourceValues(payload, source));
        return acc;
    }, {});
}

export function getApiStatusState(): ApiStatusState {
    return apiStatusState;
}

export function subscribeApiStatus(listener: () => void): () => void {
    listeners.add(listener);
    return () => {
        listeners.delete(listener);
    };
}

export async function checkCurrentApiStatusesOnly(): Promise<void> {
    if (activeCheckCurrentOnly) {
        return activeCheckCurrentOnly;
    }

    activeCheckCurrentOnly = (async () => {
        await Promise.all(API_SOURCES.map((source) => checkApiStatus(source.id)));
    })();

    try {
        await activeCheckCurrentOnly;
    }
    finally {
        activeCheckCurrentOnly = null;
    }
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export async function checkSpotiFLACNextStatusesOnly(): Promise<void> {
    if (activeCheckNextOnly) {
        return activeCheckNextOnly;
    }
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    activeCheckNextOnly = (async () => {
        const checkingNextStatuses = Object.fromEntries(SPOTIFLAC_NEXT_SOURCES.map((source) => [source.id, "checking" as ApiCheckStatus]));
        setApiStatusState((current) => ({
            ...current,
            nextStatuses: {
                ...current.nextStatuses,
                ...checkingNextStatuses,
            },
        }));
<<<<<<< HEAD
        try {
            setApiStatusState((current) => ({
                ...current,
                nextStatuses: { ...current.nextStatuses },
            }));
=======

        try {
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
            const nextStatuses = await checkSpotiFLACNextStatuses();
            setApiStatusState((current) => ({
                ...current,
                nextStatuses: {
                    ...current.nextStatuses,
                    ...nextStatuses,
                },
            }));
        }
        catch {
            setApiStatusState((current) => ({
                ...current,
                nextStatuses: getSafeNextStatusesFallback(current.nextStatuses),
            }));
        }
<<<<<<< HEAD
        finally {
            activeCheckNextOnly = null;
        }
    })();
    return activeCheckNextOnly;
}
export function ensureSpotiFLACNextStatusCheckStarted(): void {
=======
    })();

    try {
        await activeCheckNextOnly;
    }
    finally {
        activeCheckNextOnly = null;
    }
}

export function ensureApiStatusCheckStarted(): void {
    if (!activeCheckCurrentOnly && !hasCurrentResults()) {
        void checkCurrentApiStatusesOnly();
    }
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    if (!activeCheckNextOnly && !hasSpotiFLACNextResults()) {
        void checkSpotiFLACNextStatusesOnly();
    }
}
<<<<<<< HEAD
=======

export function ensureSpotiFLACNextStatusCheckStarted(): void {
    ensureApiStatusCheckStarted();
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export async function checkApiStatus(sourceId: string): Promise<void> {
    const source = API_SOURCES.find((item) => item.id === sourceId);
    if (!source) {
        return;
    }
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    const activeCheck = activeSourceChecks.get(sourceId);
    if (activeCheck) {
        return activeCheck;
    }
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    const task = (async () => {
        setApiStatusState((current) => ({
            ...current,
            checkingSources: {
                ...current.checkingSources,
                [sourceId]: true,
            },
            statuses: {
                ...current.statuses,
                [sourceId]: "checking",
            },
        }));
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
        try {
            const status = await checkSourceStatus(source);
            setApiStatusState((current) => ({
                ...current,
                statuses: {
                    ...current.statuses,
                    [sourceId]: status,
                },
            }));
        }
        finally {
            setApiStatusState((current) => ({
                ...current,
                checkingSources: {
                    ...current.checkingSources,
                    [sourceId]: false,
                },
            }));
            activeSourceChecks.delete(sourceId);
        }
    })();
<<<<<<< HEAD
=======

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    activeSourceChecks.set(sourceId, task);
    return task;
}
