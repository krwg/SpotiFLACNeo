import { Button } from "@/components/ui/button";
<<<<<<< HEAD
import { SearchCheck, CheckCircle2, XCircle, Loader2 } from "lucide-react";
import { TidalIcon, QobuzIcon, AmazonIcon, MusicBrainzIcon, AppleMusicIcon, DeezerIcon } from "./PlatformIcons";
import { useApiStatus } from "@/hooks/useApiStatus";
import { SPOTIFLAC_NEXT_SOURCES } from "@/lib/api-status";
function renderStatusIcon(status: "checking" | "online" | "offline" | "idle") {
=======
import { PlugZap, CheckCircle2, Loader2, Wrench } from "lucide-react";
import { TidalIcon, QobuzIcon, AmazonIcon, AppleMusicIcon, DeezerIcon } from "./PlatformIcons";
import { useApiStatus } from "@/hooks/useApiStatus";
import { SPOTIFLAC_NEXT_SOURCES } from "@/lib/api-status";
function renderStatusIndicator(status: "checking" | "online" | "offline" | "idle") {
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    if (status === "online") {
        return <CheckCircle2 className="h-5 w-5 text-emerald-500"/>;
    }
    if (status === "offline") {
<<<<<<< HEAD
        return <XCircle className="h-5 w-5 text-destructive"/>;
=======
        return <Wrench className="h-4 w-4 text-amber-600 dark:text-amber-400"/>;
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    }
    return null;
}
function renderPlatformIcon(type: string) {
    if (type === "tidal") {
        return <TidalIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
    }
    if (type === "amazon") {
        return <AmazonIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
    }
<<<<<<< HEAD
    if (type === "musicbrainz") {
        return <MusicBrainzIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
    }
=======
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
    if (type === "deezer") {
        return <DeezerIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
    }
    if (type === "apple") {
        return <AppleMusicIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
    }
    return <QobuzIcon className="w-5 h-5 shrink-0 text-muted-foreground"/>;
}
export function ApiStatusTab() {
<<<<<<< HEAD
    const { sources, statuses, nextStatuses, checkingSources, checkOne } = useApiStatus();
    return (<div className="space-y-6">
      <div className="space-y-4">
        <h3 className="text-sm font-semibold tracking-tight">Службы SpotiFLACNeo</h3>
=======
    const { sources, statuses, nextStatuses, checkingSources, checkAllCurrent, checkAllNext } = useApiStatus();
    const isCheckingCurrent = sources.some((source) => checkingSources[source.id] === true);
    const isCheckingNext = SPOTIFLAC_NEXT_SOURCES.some((source) => nextStatuses[source.id] === "checking");
    return (<div className="space-y-6">
      <div className="space-y-4">
        <div className="flex items-center justify-between gap-3">
          <h3 className="text-sm font-semibold tracking-tight">SpotiFLAC</h3>
          <Button variant="outline" size="sm" onClick={() => void checkAllCurrent()} disabled={isCheckingCurrent} className="gap-2">
            {isCheckingCurrent ? <Loader2 className="h-4 w-4 animate-spin"/> : <PlugZap className="h-4 w-4"/>}
            Check
          </Button>
        </div>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {sources.map((source) => {
            const status = statuses[source.id] || "idle";
<<<<<<< HEAD
            const isChecking = checkingSources[source.id] === true;
=======
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
            return (<div key={source.id} className="space-y-4 p-4 border rounded-lg bg-card text-card-foreground shadow-sm">
                <div className="flex items-center justify-between gap-3">
                  <div className="flex items-center gap-3">
                    {renderPlatformIcon(source.type)}
                    <p className="font-medium leading-none">{source.name}</p>
                  </div>
<<<<<<< HEAD
                  <div className="flex items-center">{renderStatusIcon(status)}</div>
                </div>
                <Button variant="outline" size="sm" onClick={() => void checkOne(source.id)} disabled={isChecking} className="w-full gap-2">
                  {isChecking ? <Loader2 className="h-4 w-4 animate-spin"/> : <SearchCheck className="h-4 w-4"/>}
                  Проверить
                </Button>
=======
                  <div className="flex items-center">{renderStatusIndicator(status)}</div>
                </div>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
              </div>);
        })}
        </div>
      </div>

      <div className="border-t"/>

      <div className="space-y-4">
<<<<<<< HEAD
        <h3 className="text-sm font-semibold tracking-tight">Будущие службы SpotiFLACNeo</h3>
=======
        <div className="flex items-center justify-between gap-3">
          <h3 className="text-sm font-semibold tracking-tight">SpotiFLAC Next</h3>
          <Button variant="outline" size="sm" onClick={() => void checkAllNext()} disabled={isCheckingNext} className="gap-2">
            {isCheckingNext ? <Loader2 className="h-4 w-4 animate-spin"/> : <PlugZap className="h-4 w-4"/>}
            Check
          </Button>
        </div>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5">
          {SPOTIFLAC_NEXT_SOURCES.map((source) => {
            const status = nextStatuses[source.id] || "idle";
            return (<div key={source.id} className="flex items-center justify-between p-4 border rounded-lg bg-card text-card-foreground shadow-sm">
              <div className="flex items-center gap-3">
                {renderPlatformIcon(source.id)}
                <p className="font-medium leading-none">{source.name}</p>
              </div>
<<<<<<< HEAD
              <div className="flex items-center">{renderStatusIcon(status)}</div>
=======
              <div className="flex items-center">{renderStatusIndicator(status)}</div>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
            </div>);
        })}
        </div>
      </div>
    </div>);
}
