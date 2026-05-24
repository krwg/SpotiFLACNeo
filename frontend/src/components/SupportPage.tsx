import { useState } from "react";
import { CircleCheck, Copy } from "lucide-react";
import { Button } from "@/components/ui/button";
import { openExternal } from "@/lib/utils";
import KofiLogo from "@/assets/ko-fi.gif";
import KofiSvg from "@/assets/kofi_symbol.svg";
import PatreonLogo from "@/assets/patreon.svg";
import PatreonSymbol from "@/assets/patreon_symbol.svg";
import UsdtBarcode from "@/assets/usdt.jpg";

export function SupportPage() {
    const [copiedUsdt, setCopiedUsdt] = useState(false);
    const [copiedEmail, setCopiedEmail] = useState(false);
    return (<div className="flex flex-col space-y-3">
      <div className="flex items-center justify-between shrink-0">
        <h2 className="text-2xl font-bold tracking-tight">Support Me</h2>
      </div>

      <div className="flex flex-col items-center justify-center p-4">
        <div className="grid w-full max-w-5xl overflow-hidden rounded-xl border bg-card shadow-sm md:grid-cols-3">
          <div className="flex min-h-[22rem] flex-col items-center justify-between space-y-6 border-b p-6 md:border-b-0 md:border-r">
            <div className="flex flex-col items-center space-y-4">
              <div className="h-32 flex items-center justify-center w-full relative">
                <img src={KofiLogo} className="w-72 absolute pointer-events-none" alt="Ko-fi"/>
              </div>
              <h4 className="font-semibold text-foreground">Support via Ko-fi</h4>
              <p className="text-sm text-muted-foreground text-center px-4">
                Buy me a coffee to help keep development going.
              </p>
            </div>
            <Button className="h-10 w-full gap-2 bg-[#72a4f2] text-sm font-semibold text-white hover:bg-[#5f8cd6]" onClick={() => openExternal("https://ko-fi.com/afkarxyz")}>
              <img src={KofiSvg} className="h-6 w-6 shrink-0" alt="" aria-hidden="true"/>
              Support me on Ko-fi
            </Button>
          </div>

          <div className="flex min-h-[22rem] flex-col items-center justify-between space-y-6 border-b p-6 md:border-b-0 md:border-r">
            <div className="flex flex-col items-center space-y-4 w-full">
              <div className="h-32 flex items-center justify-center w-full px-4">
                <img src={PatreonLogo} className="w-56 max-w-full brightness-0 dark:brightness-100" alt="Patreon"/>
              </div>
              <h4 className="font-semibold text-foreground">Support via Patreon</h4>
              <p className="text-sm text-muted-foreground text-center px-4">
                Join on Patreon to help fund the project and follow updates.
              </p>
            </div>
            <Button className="h-10 w-full gap-2 bg-[#ff424d] text-sm font-semibold text-white hover:bg-[#e63945]" onClick={() => openExternal("https://www.patreon.com/cw/afkarxyz")}>
              <img src={PatreonSymbol} className="h-5 w-5 shrink-0" alt="" aria-hidden="true"/>
              Support me on Patreon
            </Button>
          </div>

          <div className="flex min-h-[22rem] flex-col items-center justify-between space-y-6 p-6">
            <div className="flex flex-col items-center space-y-4 w-full">
              <div className="h-32 flex items-center justify-center">
                <div className="rounded-xl border bg-white p-2 shadow-sm">
                  <img src={UsdtBarcode} className="h-24 w-24 object-contain" alt="USDT Barcode"/>
                </div>
              </div>
              <h4 className="font-semibold text-foreground">USDT (TRC20)</h4>
              <p className="text-sm text-muted-foreground text-center px-4">
                Prefer crypto? Use the QR code or wallet address below.
              </p>
            </div>
            <div className="flex h-10 w-full items-center justify-between gap-2 rounded-lg border bg-muted/50 py-1.5 pl-3 pr-1.5">
              <code className="truncate text-xs font-mono text-muted-foreground" title="THnzAAwZgp2Sq5CAXLP2njQDhTvgZG9EWs">
                THnzAAwZgp2Sq5CAXLP2njQDhTvgZG9EWs
              </code>
              <Button variant="ghost" size="icon" className="h-7 w-7 shrink-0 hover:bg-background" onClick={() => {
            navigator.clipboard.writeText("THnzAAwZgp2Sq5CAXLP2njQDhTvgZG9EWs");
            setCopiedUsdt(true);
            setTimeout(() => setCopiedUsdt(false), 500);
        }}>
                {copiedUsdt ? <CircleCheck className="h-3.5 w-3.5 text-green-500"/> : <Copy className="h-3.5 w-3.5"/>}
              </Button>
            </div>
          </div>
        </div>

        <div className="mt-4 w-full max-w-5xl rounded-xl border bg-muted/30 px-4 py-3 text-center text-sm text-muted-foreground">
          If you have any questions or need help with donating, feel free to reach out via{" "}
          <button type="button" className="font-medium text-foreground underline-offset-4 hover:underline" onClick={() => openExternal("https://t.me/afkarxyz")}>
            Telegram
          </button>{" "}
          or{" "}
          <button type="button" className="font-medium text-foreground underline-offset-4 hover:underline" onClick={() => {
            navigator.clipboard.writeText("hi@afkarxyz.fyi");
            setCopiedEmail(true);
            setTimeout(() => setCopiedEmail(false), 500);
        }}>
            {copiedEmail ? "copied" : "hi@afkarxyz.fyi"}
          </button>
          .
        </div>
      </div>
    </div>);
}
