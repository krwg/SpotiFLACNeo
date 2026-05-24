import { useEffect, useState } from "react";
<<<<<<< HEAD
import { API_SOURCES, checkApiStatus, getApiStatusState, subscribeApiStatus, } from "@/lib/api-status";
=======
import { API_SOURCES, checkApiStatus, checkCurrentApiStatusesOnly, checkSpotiFLACNextStatusesOnly, getApiStatusState, subscribeApiStatus, } from "@/lib/api-status";
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
export function useApiStatus() {
    const [state, setState] = useState(getApiStatusState);
    useEffect(() => {
        return subscribeApiStatus(() => {
            setState(getApiStatusState());
        });
    }, []);
    return {
        ...state,
        sources: API_SOURCES,
        checkOne: (sourceId: string) => checkApiStatus(sourceId),
<<<<<<< HEAD
=======
        checkAllCurrent: () => checkCurrentApiStatusesOnly(),
        checkAllNext: () => checkSpotiFLACNextStatusesOnly(),
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    };
}
