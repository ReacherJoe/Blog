import { Input } from 'antd';
import { SearchOutlined } from '@ant-design/icons';

export const SearchInput = ({ keyword, setKeyword }) => {
  return (
    <div className="w-full p-2 border-[1px] rounded-md border-light-accent1 flex items-center gap-2">
      <div className="text-xl font-bold text-light-accent1">
        <SearchOutlined />
        <Input
        value={keyword}
        className="border-none outline-none dark:bg-dark-background bg-light-background dark:text-dark-text text-light-text"
        placeholder="Search..."
        onChange={(evt) => setKeyword(evt.target.value)}
      />
      </div>

      
    </div>
  );
};
