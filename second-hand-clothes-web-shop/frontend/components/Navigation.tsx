import Link from 'next/link';

const Navigation = () => {
  return (
    <nav>
      <ul>
        <li>
          <Link href="/women">
            Women
          </Link>
        </li>
        <li>
          <Link href="/men">
            Men
          </Link>
        </li>
        <li>
          <Link href="/kids">
            Kids
          </Link>
        </li>
      </ul>
    </nav>
  );
};

export default Navigation;
