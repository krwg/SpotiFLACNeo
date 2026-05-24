import { useRef, useState, type RefObject } from "react";
import { HomeIcon } from "@/components/ui/home";
import { HistoryIcon } from "@/components/ui/history-icon";
import { SettingsIcon } from "@/components/ui/settings";
import { ActivityIcon, type ActivityIconHandle } from "@/components/ui/activity";
import { TerminalIcon } from "@/components/ui/terminal";
import { FileMusicIcon, type FileMusicIconHandle } from "@/components/ui/file-music";
import { FilePenIcon, type FilePenIconHandle } from "@/components/ui/file-pen";
<<<<<<< HEAD
import { CoffeeIcon } from "@/components/ui/coffee";
import { BadgeAlertIcon } from "@/components/ui/badge-alert";
import { GithubIcon } from "@/components/ui/github";
import { BlocksIcon } from "@/components/ui/blocks-icon";
import { AudioLinesIcon, type AudioLinesIconHandle } from "@/components/ui/audio-lines";
=======
import { BugReportIcon } from "@/components/ui/bug-report-icon";
import { CoffeeIcon } from "@/components/ui/coffee";
import { BlocksIcon } from "@/components/ui/blocks-icon";
import { AudioLinesIcon, type AudioLinesIconHandle } from "@/components/ui/audio-lines";
import { ToolCaseIcon } from "@/components/ui/tool-case";
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Checkbox } from "@/components/ui/checkbox";
import { Tooltip, TooltipContent, TooltipTrigger, } from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";
import { openExternal } from "@/lib/utils";
<<<<<<< HEAD
export type PageType = "main" | "settings" | "debug" | "audio-analysis" | "audio-converter" | "audio-resampler" | "file-manager" | "about" | "history";
=======
export type PageType = "main" | "settings" | "debug" | "audio-analysis" | "audio-converter" | "audio-resampler" | "file-manager" | "projects" | "support" | "history";
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
interface SidebarProps {
    currentPage: PageType;
    onPageChange: (page: PageType) => void;
}
interface AnimatedIconHandle {
    startAnimation: () => void;
    stopAnimation: () => void;
}
export function Sidebar({ currentPage, onPageChange }: SidebarProps) {
    const [isIssuesDialogOpen, setIsIssuesDialogOpen] = useState(false);
    const [hasIssueAgreement, setHasIssueAgreement] = useState(false);
    const analyzerIconRef = useRef<ActivityIconHandle>(null);
    const resamplerIconRef = useRef<AudioLinesIconHandle>(null);
    const converterIconRef = useRef<FileMusicIconHandle>(null);
    const fileManagerIconRef = useRef<FilePenIconHandle>(null);
    const handleIssuesDialogChange = (open: boolean) => {
        setIsIssuesDialogOpen(open);
        if (!open) {
            setHasIssueAgreement(false);
        }
    };
    const handleOpenIssues = () => {
<<<<<<< HEAD
        openExternal("https://github.com/krwg/SpotiFLACNeo/issues");
=======
        openExternal("https://github.com/spotbye/SpotiFLAC/issues");
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
        handleIssuesDialogChange(false);
    };
    const getAnimatedItemHandlers = <T extends AnimatedIconHandle>(iconRef: RefObject<T | null>) => ({
        onMouseEnter: () => iconRef.current?.startAnimation(),
        onMouseLeave: () => iconRef.current?.stopAnimation(),
        onFocus: () => iconRef.current?.startAnimation(),
        onBlur: () => iconRef.current?.stopAnimation(),
    });
<<<<<<< HEAD
    return (<div className="fixed left-0 top-0 h-full w-14 glass-panel border-r border-border/50 flex flex-col items-center py-14 z-30">
        <div className="flex flex-col gap-2 flex-1">
            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant={currentPage === "main" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${currentPage === "main" ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("main")}>
                        <HomeIcon size={20} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>Главная</p>
                </TooltipContent>
            </Tooltip>

            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant={currentPage === "history" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${currentPage === "history" ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("history")}>
                        <HistoryIcon size={20} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>История</p>
                </TooltipContent>
            </Tooltip>

            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant={currentPage === "settings" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${currentPage === "settings" ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("settings")}>
                        <SettingsIcon size={20} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>Настройки</p>
                </TooltipContent>
            </Tooltip>

            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant={currentPage === "debug" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${currentPage === "debug" ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("debug")}>
                        <TerminalIcon size={20} loop={true} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>Логи Отладки</p>
                </TooltipContent>
            </Tooltip>

            <DropdownMenu>
                <Tooltip delayDuration={0}>
                    <DropdownMenuTrigger asChild>
                        <TooltipTrigger asChild>
                            <Button variant={["audio-analysis", "audio-converter", "audio-resampler", "file-manager"].includes(currentPage) ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${["audio-analysis", "audio-converter", "audio-resampler", "file-manager"].includes(currentPage) ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`}>
                                <BlocksIcon size={20} loop={true} />
                            </Button>
                        </TooltipTrigger>
                    </DropdownMenuTrigger>
                    <TooltipContent side="right">
                        <p>Инструменты</p>
                    </TooltipContent>
                </Tooltip>
                <DropdownMenuContent side="right" sideOffset={14} className="min-w-[200px] ml-2">
                    <DropdownMenuItem onClick={() => onPageChange("audio-analysis")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(analyzerIconRef)}>
                        <ActivityIcon ref={analyzerIconRef} size={16} />
                        <span>Анализатор Качества Звука</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => onPageChange("audio-resampler")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(resamplerIconRef)}>
                        <AudioLinesIcon ref={resamplerIconRef} size={16} />
                        <span>Ресемплер Звука</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => onPageChange("audio-converter")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(converterIconRef)}>
                        <FileMusicIcon ref={converterIconRef} size={16} />
                        <span>Аудиоконвертер</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => onPageChange("file-manager")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(fileManagerIconRef)}>
                        <FilePenIcon ref={fileManagerIconRef} size={16} />
                        <span>Файловый Менеджер</span>
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>

        <div className="mt-auto flex flex-col gap-2">
            <Dialog open={isIssuesDialogOpen} onOpenChange={handleIssuesDialogChange}>
                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="h-10 w-10 hover:bg-primary/10 hover:text-primary" onClick={() => setIsIssuesDialogOpen(true)}>
                            <GithubIcon size={20} />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Сообщить об ошибке / Предложить идею</p>
                    </TooltipContent>
                </Tooltip>
                <DialogContent className="max-w-xl">
                    <DialogHeader>
                        <DialogTitle>Перед созданием Issue на GitHub</DialogTitle>
                        <DialogDescription />
                    </DialogHeader>

                    <div className="space-y-4 text-sm">
                        <div className="rounded-lg border border-amber-500/40 bg-amber-500/10 p-4">
                            <p className="font-semibold text-amber-900 dark:text-amber-200">Внимание</p>
                            <p className="mt-1 text-amber-950/90 dark:text-amber-100/90">
                                Сначала поищите решение в существующих issue и обязательно используйте шаблоны при создании нового отчета или запроса.
                            </p>
                        </div>

                        <label className="flex cursor-pointer items-center gap-3 rounded-lg border p-4">
                            <Checkbox className="shrink-0" checked={hasIssueAgreement} onCheckedChange={(checked) => setHasIssueAgreement(checked === true)} />
                            <span className="leading-5 text-foreground/90">
                                Я понимаю, что должен использовать шаблон и избегать создания дубликатов.
                            </span>
                        </label>
                    </div>

                    <DialogFooter className="sm:justify-between gap-2">
                        <Button variant="outline" onClick={() => handleIssuesDialogChange(false)}>
                            Отмена
                        </Button>
                        <Button disabled={!hasIssueAgreement} onClick={handleOpenIssues}>
                            Открыть Issues
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>

            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant={currentPage === "about" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 rounded-xl ${currentPage === "about" ? "bg-primary/15 text-primary shadow-sm hover:bg-primary/25" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("about")}>
                        <BadgeAlertIcon size={20} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>О программе</p>
                </TooltipContent>
            </Tooltip>

            <Tooltip delayDuration={0}>
                <TooltipTrigger asChild>
                    <Button variant="ghost" size="icon" className="h-10 w-10 hover:bg-primary/10 hover:text-primary" onClick={() => openExternal("https://dalink.to/fallnix")}>
                        <CoffeeIcon size={20} loop={true} />
                    </Button>
                </TooltipTrigger>
                <TooltipContent side="right">
                    <p>Поддержать на DonationAlerts</p>
                </TooltipContent>
            </Tooltip>
        </div>
    </div>);
=======
    return (<div className="fixed left-0 top-0 h-full w-14 bg-card border-r border-border flex flex-col items-center py-14 z-30">
            <div className="flex flex-col gap-2 flex-1">
                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "main" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "main" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("main")}>
                            <HomeIcon size={20}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Home</p>
                    </TooltipContent>
                </Tooltip>

                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "history" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "history" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("history")}>
                            <HistoryIcon size={20}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>History</p>
                    </TooltipContent>
                </Tooltip>

                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "settings" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "settings" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("settings")}>
                            <SettingsIcon size={20}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Settings</p>
                    </TooltipContent>
                </Tooltip>

                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "debug" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "debug" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("debug")}>
                            <TerminalIcon size={20} loop={true}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Debug Logs</p>
                    </TooltipContent>
                </Tooltip>

                <DropdownMenu>
                    <Tooltip delayDuration={0}>
                        <DropdownMenuTrigger asChild>
                            <TooltipTrigger asChild>
                                <Button variant={["audio-analysis", "audio-converter", "audio-resampler", "file-manager"].includes(currentPage) ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${["audio-analysis", "audio-converter", "audio-resampler", "file-manager"].includes(currentPage) ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`}>
                                    <ToolCaseIcon size={20}/>
                                </Button>
                            </TooltipTrigger>
                        </DropdownMenuTrigger>
                        <TooltipContent side="right">
                            <p>Tools</p>
                        </TooltipContent>
                    </Tooltip>
                    <DropdownMenuContent side="right" sideOffset={14} className="min-w-[200px] ml-2">
                        <DropdownMenuItem onClick={() => onPageChange("audio-analysis")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(analyzerIconRef)}>
                            <ActivityIcon ref={analyzerIconRef} size={16}/>
                            <span>Audio Quality Analyzer</span>
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => onPageChange("audio-resampler")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(resamplerIconRef)}>
                            <AudioLinesIcon ref={resamplerIconRef} size={16}/>
                            <span>Audio Resampler</span>
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => onPageChange("audio-converter")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(converterIconRef)}>
                            <FileMusicIcon ref={converterIconRef} size={16}/>
                            <span>Audio Converter</span>
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => onPageChange("file-manager")} className="gap-3 cursor-pointer py-2 px-3" {...getAnimatedItemHandlers(fileManagerIconRef)}>
                            <FilePenIcon ref={fileManagerIconRef} size={16}/>
                            <span>File Manager</span>
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>

            <div className="mt-auto flex flex-col gap-2">
                <Dialog open={isIssuesDialogOpen} onOpenChange={handleIssuesDialogChange}>
                    <Tooltip delayDuration={0}>
                        <TooltipTrigger asChild>
                            <Button variant="ghost" size="icon" className="h-10 w-10 hover:bg-primary/10 hover:text-primary" onClick={() => setIsIssuesDialogOpen(true)}>
                                <BugReportIcon size={20} loop={true}/>
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent side="right">
                            <p>Report Bugs or Request Features</p>
                        </TooltipContent>
                    </Tooltip>
                    <DialogContent className="max-w-xl">
                        <DialogHeader>
                            <DialogTitle>Before Opening GitHub Issues</DialogTitle>
                            <DialogDescription />
                        </DialogHeader>

                        <div className="space-y-4 text-sm">
                            <div className="rounded-lg border border-amber-500/40 bg-amber-500/10 p-4">
                                <p className="font-semibold text-amber-900 dark:text-amber-200">Important</p>
                                <p className="mt-1 text-amber-950/90 dark:text-amber-100/90">
                                    Search existing issues first and use the issue template when opening a new report or request.
                                </p>
                            </div>

                            <label className="flex cursor-pointer items-center gap-3 rounded-lg border p-4">
                                <Checkbox className="shrink-0" checked={hasIssueAgreement} onCheckedChange={(checked) => setHasIssueAgreement(checked === true)}/>
                                <span className="leading-5 text-foreground/90">
                                    I understand that I should use the issue template and avoid duplicate issues.
                                </span>
                            </label>
                        </div>

                        <DialogFooter className="sm:justify-between gap-2">
                            <Button variant="outline" onClick={() => handleIssuesDialogChange(false)}>
                                Cancel
                            </Button>
                            <Button disabled={!hasIssueAgreement} onClick={handleOpenIssues}>
                                Open Issues
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>

                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "projects" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "projects" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("projects")}>
                            <BlocksIcon size={20} loop={true}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Other Projects</p>
                    </TooltipContent>
                </Tooltip>

                <Tooltip delayDuration={0}>
                    <TooltipTrigger asChild>
                        <Button variant={currentPage === "support" ? "secondary" : "ghost"} size="icon" className={`h-10 w-10 ${currentPage === "support" ? "bg-primary/10 text-primary hover:bg-primary/20" : "hover:bg-primary/10 hover:text-primary"}`} onClick={() => onPageChange("support")}>
                            <CoffeeIcon size={20} loop={true}/>
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right">
                        <p>Support Me</p>
                    </TooltipContent>
                </Tooltip>
            </div>
        </div>);
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}
