import { Badge } from "@/components/ui/badge";
import { Tooltip, TooltipContent, TooltipTrigger, } from "@/components/ui/tooltip";
import { openExternal } from "@/lib/utils";
import { formatRelativeTime } from "@/lib/relative-time";
interface HeaderProps {
<<<<<<< HEAD
  version: string;
  hasUpdate: boolean;
  releaseDate?: string | null;
}
export function Header({ version, hasUpdate, releaseDate }: HeaderProps) {
  return (<div className="hero-gradient text-center space-y-4">
    <div className="flex items-center justify-center gap-4">
      <img
        src="/icon.svg"
        alt="SpotiFLACNeo"
        className="w-14 h-14 cursor-pointer drop-shadow-md transition-transform hover:scale-105"
        onClick={() => window.location.reload()}
      />
      <div className="text-left">
        <h1
          className="font-minecraft text-3xl md:text-4xl tracking-wide cursor-pointer leading-none flex items-center gap-2.5"
          onClick={() => window.location.reload()}
        >
          <span>SpotiFLAC</span>
          <span className="neo-badge" aria-label="Neo edition">Neo</span>
        </h1>
        <p className="text-xs text-muted-foreground mt-1 font-medium">
          neo edition · русская локализация · Apple UI
        </p>
      </div>
      <div className="relative self-start">
        <Tooltip>
          <TooltipTrigger asChild>
            <Badge variant="default" asChild>
              <button
                type="button"
                onClick={() => openExternal("https://github.com/krwg/SpotiFLACNeo/releases")}
                className="cursor-pointer rounded-full px-3 py-1 hover:opacity-90 transition-opacity shadow-sm"
              >
                v{version}
              </button>
            </Badge>
          </TooltipTrigger>
          {hasUpdate && releaseDate && (<TooltipContent>
              <p>Обновление · {formatRelativeTime(releaseDate)}</p>
            </TooltipContent>)}
        </Tooltip>
        {hasUpdate && (<span className="absolute -top-1 -right-1 flex h-3 w-3">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
            <span className="relative inline-flex rounded-full h-3 w-3 bg-green-500"></span>
          </span>)}
      </div>
    </div>
    <p className="text-muted-foreground text-sm md:text-base max-w-xl mx-auto leading-relaxed">
      Скачивай треки Spotify в настоящем FLAC из Tidal, Qobuz и Amazon Music — без аккаунтов.
      <span className="block mt-1 text-xs opacity-80">
        Скоро: Яндекс Музыка, VK Музыка, Apple Music
      </span>
    </p>
  </div>);
=======
    version: string;
    hasUpdate: boolean;
    releaseDate?: string | null;
}
export function Header({ version, hasUpdate, releaseDate }: HeaderProps) {
    return (<div className="relative">
      <div className="text-center space-y-2">
        <div className="flex items-center justify-center gap-3">
          <img src="/icon.svg" alt="SpotiFLAC" className="w-12 h-12 cursor-pointer" onClick={() => window.location.reload()}/>
          <h1 className="text-4xl font-bold cursor-pointer" onClick={() => window.location.reload()}>
            SpotiFLAC
          </h1>
          <div className="relative">
            <Tooltip>
              <TooltipTrigger asChild>
                <Badge variant="default" asChild>
                  <button type="button" onClick={() => openExternal("https://github.com/spotbye/SpotiFLAC/releases")} className="cursor-pointer hover:opacity-80 transition-opacity">
                    v{version}
                  </button>
                </Badge>
              </TooltipTrigger>
              {hasUpdate && releaseDate && (<TooltipContent>
                  <p>{formatRelativeTime(releaseDate)}</p>
                </TooltipContent>)}
            </Tooltip>
            {hasUpdate && (<span className="absolute -top-1 -right-1 flex h-3 w-3">
                <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
                <span className="relative inline-flex rounded-full h-3 w-3 bg-green-500"></span>
              </span>)}
          </div>
        </div>
        <p className="text-muted-foreground">
          Get Spotify tracks in true FLAC from Tidal, Qobuz & Amazon Music — no account required.
        </p>
      </div>
    </div>);
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}
