import Link from 'next/link';

const Navigation = () => {
  return (
    <nav>
      <ul>
        <li>
          <Link href="/women">
            <a>Women</a>
          </Link>
        </li>
        <li>
          <Link href="/men">
            <a>Men</a>
          </Link>
        </li>
        <li>
          <Link href="/kids">
            <a>Kids</a>
          </Link>
        </li>
      </ul>
    </nav>
  );
};

export default Navigation;
