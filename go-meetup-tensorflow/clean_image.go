output = op.Div(s,
	op.Sub(s,
		// Resize
		op.ResizeBilinear(s,
			// Create a batch containing a single image
			op.ExpandDims(s,
				// Use decoded pixel values
				op.Cast(s, decode, tf.Float),
				op.Const(s.SubScope("make_batch"), int32(0))),
			op.Const(s.SubScope("size"), []int32{H, W})),
		op.Const(s.SubScope("mean"), Mean)),
	op.Const(s.SubScope("scale"), Scale))