package main

import (
	"github.com/go-example/demo/geo/tool"
	"github.com/golang/geo/s2"
	"log"
)

func main() {
	//S2使用的是WGS84坐标（GPS导航坐标）,如果你获得的是百度或者高德的地理坐标，请将其转换为GPS坐标，以防止计算的位置不准确。

	ll := s2.LatLngFromDegrees(36.683, 117.1412)

	log.Println("------------- 计算cell ID --------------")
	cellId := s2.CellIDFromLatLng(ll)
	log.Println(cellId.Face(), cellId.LatLng().Lat, cellId.LatLng().Lng)

	cell := s2.CellFromCellID(cellId)
	log.Println(cell.RectBound())
	log.Println("latlng = ", ll)
	log.Println("cell level = ", cellId.Level())
	log.Printf("cell = %d\n", cellId)

	smallCell := s2.CellFromCellID(cellId.Parent(10))
	log.Printf("smallCell level = %d\n", smallCell.Level())
	log.Printf("smallCell id = %b\n", smallCell.ID())
	log.Printf("smallCell ApproxArea = %v\n", smallCell.ApproxArea())
	log.Printf("smallCell AverageArea = %v\n", smallCell.AverageArea())
	log.Printf("smallCell ExactArea = %v\n", smallCell.ExactArea())

	log.Println("------------- 计算 距离 2.718 公里 --------------")
	lon,lat := tool.GCJ02toWGS84(121.498414, 31.177459)
	log.Println(lon,lat)
	start := s2.LatLngFromDegrees(lat, lon)
	lon,lat = tool.GCJ02toWGS84(121.526996, 31.175403)
	log.Println(lon,lat)
	end := s2.LatLngFromDegrees(lat, lon)
	start = s2.LatLngFromDegrees(31.177459, 121.498414)
	end = s2.LatLngFromDegrees(31.175403, 121.526996)
	dis := end.Distance(start)
	log.Println(dis.E6())

	log.Println("------------- 构建一个围栏 --------------")
	//[[[121.498414,31.177459],[121.526996,31.175403],[121.513006,31.162771],[121.513006,31.162771]]]
	p1 := s2.LatLngFromDegrees(31.177459, 121.498414)
	p2 := s2.LatLngFromDegrees(31.175403, 121.526996)
	p3 := s2.LatLngFromDegrees(31.162771, 121.513006)
	log.Println(p1, p2.Lng, p3)

	rect := s2.RectFromLatLng(p1)
	rect.AddPoint(p2)
	rect.AddPoint(p3)
	rc := &s2.RegionCoverer{MaxLevel: 20, MaxCells: 20, MinLevel: 2}
	r := s2.Region(rect.CapBound())
	covering := rc.Covering(r)

	//验证点是否在围栏里
	cellId = s2.CellIDFromLatLng(s2.LatLngFromDegrees(31.177459, 121.498415))
	log.Println(covering.ContainsCell(s2.CellFromCellID(cellId)))
	log.Println("------------- 构建一个围栏 loop --------------")
	//loop操作
	loop := s2.LoopFromPoints([]s2.Point{
		s2.PointFromLatLng(p1),
		s2.PointFromLatLng(p2),
		s2.PointFromLatLng(p3),
	})
	log.Println(loop.ContainsPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(32.177469, 122.498415))))
	log.Println("------------- 最近和最远的点及距离 --------------")
	polylines := []s2.Polyline{
		// This is an iteration = 3 Koch snowflake centered at the
		// center of the continental US.
		s2.Polyline{
			s2.PointFromLatLng(s2.LatLngFromDegrees(47.5467, -103.6035)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.9214, -103.7320)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.1527, -105.8000)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(44.2866, -103.8538)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(42.6450, -103.9695)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.8743, -105.9314)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(42.7141, -107.8226)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.0743, -107.8377)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(40.2486, -109.6869)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(39.4333, -107.8521)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.7936, -107.8658)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(38.5849, -106.0503)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.7058, -104.2841)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(36.0638, -104.3793)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(35.3062, -106.1585)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(34.4284, -104.4703)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.8024, -104.5573)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(33.5273, -102.8163)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.6053, -101.1982)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(34.2313, -101.0361)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(34.9120, -99.2189)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(33.9382, -97.6134)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.3185, -97.8489)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.9481, -96.0510)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(31.9449, -94.5321)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(33.5521, -94.2263)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(34.1285, -92.3780)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(35.1678, -93.9070)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(36.7893, -93.5734)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.3529, -91.6381)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(36.2777, -90.1050)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.8824, -89.6824)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(38.3764, -87.7108)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(39.4869, -89.2407)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.0883, -88.7784)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(40.5829, -90.8289)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.6608, -92.4765)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(43.2777, -92.0749)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(43.7961, -89.9408)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(44.8865, -91.6533)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.4844, -91.2100)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.9512, -93.4327)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.9863, -95.2792)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.3722, -95.6237)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(44.7496, -97.7776)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.7189, -99.6629)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(47.3422, -99.4244)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.6523, -101.6056)),
		},
		s2.Polyline{
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.3185, -97.8489)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(32.9481, -96.0510)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(31.9449, -94.5321)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(33.5521, -94.2263)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(34.1285, -92.3780)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(35.1678, -93.9070)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(36.7893, -93.5734)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.3529, -91.6381)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(36.2777, -90.1050)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(37.8824, -89.6824)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(38.3764, -87.7108)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(39.4869, -89.2407)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.0883, -88.7784)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(40.5829, -90.8289)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(41.6608, -92.4765)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(43.2777, -92.0749)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(43.7961, -89.9408)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(44.8865, -91.6533)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.4844, -91.2100)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.9512, -93.4327)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.9863, -95.2792)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.3722, -95.6237)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(44.7496, -97.7776)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(45.7189, -99.6629)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(47.3422, -99.4244)),
			s2.PointFromLatLng(s2.LatLngFromDegrees(46.6523, -101.6056)),
		},
	}

	// 设置一个点，找离它最近的边
	point := s2.PointFromLatLng(s2.LatLngFromDegrees(38.7, -100.5))

	// 将形状加载到索引中
	index := s2.NewShapeIndex()
	for _, l := range polylines {
		index.Add(&l)
	}

	// 创建一个closesEdgeQuery并指定我们想要7个最接近的。
	q := s2.NewClosestEdgeQuery(index, s2.NewClosestEdgeQueryOptions().MaxResults(7))
	target := s2.NewMinDistanceToPointTarget(point)

	for _, result := range q.FindEdges(target) {
		polylineIndex := result.ShapeID()
		edgeIndex := result.EdgeID()
		distance := result.Distance()
		log.Printf("Polyline %d, Edge %d is %0.4f degrees from Point (%0.6f, %0.6f, %0.6f)\n",
			polylineIndex, edgeIndex,
			distance.Angle().Degrees(), point.X, point.Y, point.Z)
	}

	// 使用一个点，我们想找到最远的边缘。
	point = s2.PointFromLatLng(s2.LatLngFromDegrees(37.7, -122.5))
	// 将形状加载到索引中
	index = s2.NewShapeIndex()
	for _, l := range polylines {
		index.Add(&l)
	}
	// 创建一个FurthestEdgeQuery并指定我们想要最远的3个。
	q = s2.NewFurthestEdgeQuery(index, s2.NewFurthestEdgeQueryOptions().MaxResults(3))
	furthest := s2.NewMaxDistanceToPointTarget(point)

	for _, result := range q.FindEdges(furthest) {
		polylineIndex := result.ShapeID()
		edgeIndex := result.EdgeID()
		distance := result.Distance()
		log.Printf("Polyline %d, Edge %d is %0.3f degrees from Point (%0.3f, %0.3f, %0.3f)\n",
			polylineIndex, edgeIndex,
			distance.Angle().Degrees(), point.X, point.Y, point.Z)
	}

}
