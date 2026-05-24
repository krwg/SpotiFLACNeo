import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue, } from "@/components/ui/select";
import { Search, ArrowUpDown, XCircle } from "lucide-react";
interface SearchAndSortProps {
    searchQuery: string;
    sortBy: string;
    onSearchChange: (value: string) => void;
    onSortChange: (value: string) => void;
}
export function SearchAndSort({ searchQuery, sortBy, onSearchChange, onSortChange, }: SearchAndSortProps) {
    return (<div className="flex gap-2">
      <div className="relative flex-1">
        <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground"/>
<<<<<<< HEAD
        <Input placeholder="Поиск треков..." value={searchQuery} onChange={(e) => onSearchChange(e.target.value)} className="pl-10 pr-8"/>
=======
        <Input placeholder="Search tracks..." value={searchQuery} onChange={(e) => onSearchChange(e.target.value)} className="pl-10 pr-8"/>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
        {searchQuery && (<button type="button" className="absolute right-2 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground transition-colors cursor-pointer" onClick={() => onSearchChange("")}>
            <XCircle className="h-4 w-4"/>
          </button>)}
      </div>
      <Select value={sortBy} onValueChange={onSortChange}>
        <SelectTrigger className="w-[200px] gap-1.5">
          <ArrowUpDown className="h-4 w-4"/>
<<<<<<< HEAD
          <SelectValue placeholder="Сортировка"/>
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="default">По умолчанию</SelectItem>
          <SelectItem value="title-asc">Название (А-Я)</SelectItem>
          <SelectItem value="title-desc">Название (Я-А)</SelectItem>
          <SelectItem value="artist-asc">Исполнитель (А-Я)</SelectItem>
          <SelectItem value="artist-desc">Исполнитель (Я-А)</SelectItem>
          <SelectItem value="duration-asc">Продолжительность (меньше)</SelectItem>
          <SelectItem value="duration-desc">Продолжительность (больше)</SelectItem>
          <SelectItem value="plays-asc">Прослушивания (меньше)</SelectItem>
          <SelectItem value="plays-desc">Прослушивания (больше)</SelectItem>
          <SelectItem value="downloaded">Скачанные</SelectItem>
          <SelectItem value="not-downloaded">Не скачанные</SelectItem>
          <SelectItem value="failed">Ошибки загрузки</SelectItem>
=======
          <SelectValue placeholder="Sort by"/>
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="default">Default</SelectItem>
          <SelectItem value="title-asc">Title (A-Z)</SelectItem>
          <SelectItem value="title-desc">Title (Z-A)</SelectItem>
          <SelectItem value="artist-asc">Artist (A-Z)</SelectItem>
          <SelectItem value="artist-desc">Artist (Z-A)</SelectItem>
          <SelectItem value="duration-asc">Duration (Short)</SelectItem>
          <SelectItem value="duration-desc">Duration (Long)</SelectItem>
          <SelectItem value="plays-asc">Plays (Low)</SelectItem>
          <SelectItem value="plays-desc">Plays (High)</SelectItem>
          <SelectItem value="downloaded">Downloaded</SelectItem>
          <SelectItem value="not-downloaded">Not Downloaded</SelectItem>
          <SelectItem value="failed">Failed Downloads</SelectItem>
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
        </SelectContent>
      </Select>
    </div>);
}
