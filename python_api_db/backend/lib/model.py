from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column
from sqlalchemy import VARCHAR

class Base(DeclarativeBase):
    pass

class Movies(Base):
    __tablename__ = "movies"

    seq: Mapped[int] = mapped_column(primary_key=True, autoincrement=True, nullable=False)
    name: Mapped[str] = mapped_column(VARCHAR(200), nullable=False)
    movie_id: Mapped[str] = mapped_column(VARCHAR(200), nullable=False)
    rating: Mapped[str] = mapped_column(VARCHAR(200), nullable=False)
    timestamp: Mapped[str] = mapped_column(VARCHAR(200), nullable=False)

    def __repr__(self) -> str:
        return f"User(name={self.name}"