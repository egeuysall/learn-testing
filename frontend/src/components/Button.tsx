interface ButtonProps {
  onClick: () => void;
  children: React.ReactNode;
  disabled?: boolean;
}

export default function Button({ onClick, children, disabled = false }: ButtonProps) {
  return (
    <button onClick={onClick} disabled={disabled} className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-300">
      {children}
    </button>
  );
}
