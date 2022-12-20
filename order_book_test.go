package tradingo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderBook(t *testing.T) {
	orderBook := NewOrderBook("")

	orderBook.Update(12, 20, false)
	orderBook.Update(11, 20, false)
	orderBook.Update(13, 20, false)
	orderBook.Update(9, 20, false)
	orderBook.Update(6, 20, false)

	orderBook.Update(1, 20, true)
	orderBook.Update(4, 20, true)
	orderBook.Update(3, 20, true)
	orderBook.Update(4, 20, true)
	orderBook.Update(2, 20, true)
	orderBook.Update(5, 20, true)

	require.Equal(t, 6., orderBook.BestAsk().price)
	require.Equal(t, 20., orderBook.BestAsk().quantity)
	require.Equal(t, 5., orderBook.BestBid().price)
	require.Equal(t, 20., orderBook.BestBid().quantity)
}

func TestOrderBook1(t *testing.T) {
	orderBook := NewOrderBook("")

	orderBook.Update(12, 20, false)
	orderBook.Update(11, 20, false)
	orderBook.Update(13, 20, false)
	orderBook.Update(9, 25, false)
	orderBook.Update(6, 20, false)

	orderBook.Update(1, 20, true)
	orderBook.Update(4, 20, true)
	orderBook.Update(3, 20, true)
	orderBook.Update(7, 15, true)
	orderBook.Update(2, 20, true)
	orderBook.Update(5, 20, true)

	require.Equal(t, 9., orderBook.BestAsk().price)
	require.Equal(t, 25., orderBook.BestAsk().quantity)
	require.Equal(t, 7., orderBook.BestBid().price)
	require.Equal(t, 15., orderBook.BestBid().quantity)
}

func TestOrderBookDelete(t *testing.T) {
	orderBook := NewOrderBook("")

	orderBook.Update(12, 23, false)
	orderBook.Update(11, 21, false)
	orderBook.Update(13, 24, false)
	orderBook.Asks.deleteHead()
	orderBook.Asks.deleteTail()

	orderBook.Update(1, 10, true)
	orderBook.Update(4, 14, true)
	orderBook.Bids.deleteHead()

	require.False(t, orderBook.IsEmpty())
	require.Equal(t, 12., orderBook.BestAsk().price)
	require.Equal(t, 23., orderBook.BestAsk().quantity)
	require.Equal(t, 1., orderBook.BestBid().price)
	require.Equal(t, 10., orderBook.BestBid().quantity)
}

func TestOrderBookEmpty(t *testing.T) {
	orderBook := NewOrderBook("")

	orderBook.Update(12, 20, false)
	orderBook.Update(11, 20, false)
	orderBook.Update(13, 20, false)
	orderBook.Asks.deleteHead()
	orderBook.Asks.deleteTail()
	orderBook.Asks.deleteHead()

	orderBook.Update(1, 20, true)
	orderBook.Update(4, 20, true)
	orderBook.Bids.deleteHead()
	orderBook.Bids.deleteTail()

	require.True(t, orderBook.IsEmpty())
	require.Nil(t, orderBook.BestAsk())
	require.Nil(t, orderBook.BestBid())
	require.Nil(t, orderBook.Asks.head)
	require.Nil(t, orderBook.Asks.tail)
	require.Nil(t, orderBook.Bids.head)
	require.Nil(t, orderBook.Bids.tail)
}
